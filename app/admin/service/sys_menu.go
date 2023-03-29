package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/errors"
	"go-admin/common/middleware"
	"gorm.io/gorm"
	"sort"
	"strings"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
	cModels "go-admin/common/models"
	"go-admin/common/service"
)

type SysMenu struct {
	service.Service
}

// GetPage 获取SysMenu列表
func (e *SysMenu) GetPage(c *dto.SysMenuGetPageReq, menus *[]models.SysMenu) *SysMenu {
	var menu = new([]models.SysMenu)
	menu, err := e.getPage(c)
	if err != nil {
		err = errors.WithStack(err)
		_ = e.AddError(err)
		return e
	}
	for i := 0; i < len(*menu); i++ {
		if (*menu)[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCallp(menu, (*menu)[i])
		*menus = append(*menus, menusInfo)
	}
	return e
}

// getPage 菜单分页列表
func (e *SysMenu) getPage(r *dto.SysMenuGetPageReq) (list *[]models.SysMenu, err error) {
	list, err = service.GetList[models.SysMenu](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SysApi").Scopes(
			cDto.OrderDest("sort", false),
		)
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Get 获取SysMenu对象
func (e *SysMenu) Get(r *dto.SysMenuGetReq) (model *models.SysMenu, err error) {

	model, err = service.Get[models.SysMenu](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SysApi")
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}

	apis := make([]int, 0)
	for _, v := range model.SysApi {
		apis = append(apis, v.Id)
	}
	model.Apis = apis
	return
}

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(c *gin.Context, r *dto.SysMenuInsertReq) (err error) {
	apiList, err := service.GetList[models.SysApi](e, nil,
		func(db *gorm.DB) *gorm.DB {
			return db.Where("id in ?", r.Apis)
		})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	r.SysApi = *apiList
	model, err := service.Insert[models.SysMenu](e, r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	err = e.initPaths(e.Orm, model)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	r.MenuId = model.MenuId
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"新增",
		fmt.Sprintf("新增了菜单数据，ID：%v", model.MenuId),
		"{}",
		string(after),
	)
	return
}

func (e *SysMenu) initPaths(tx *gorm.DB, menu *models.SysMenu) error {
	var err error
	var data models.SysMenu
	parentMenu := new(models.SysMenu)
	if menu.ParentId != 0 {
		tx.Model(&data).First(parentMenu, menu.ParentId)
		if parentMenu.Paths == "" {
			err = errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return err
		}
		menu.Paths = parentMenu.Paths + "/" + pkg.IntToString(menu.MenuId)
	} else {
		menu.Paths = "/0/" + pkg.IntToString(menu.MenuId)
	}
	tx.Model(&data).Where("menu_id = ?", menu.MenuId).Update("paths", menu.Paths)
	return err
}

// Update 修改SysMenu对象
func (e *SysMenu) Update(c *gin.Context, r *dto.SysMenuUpdateReq) *SysMenu {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var alist = make([]models.SysApi, 0)
	var model = models.SysMenu{}
	tx.Preload("SysApi").First(&model, r.GetId())
	oldPath := model.Paths
	before, _ := json.Marshal(model)
	tx.Where("id in ?", r.Apis).Find(&alist)
	err = tx.Model(&model).Association("SysApi").Delete(model.SysApi)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("delete policy error:%s", err)
		_ = e.AddError(err)
		return e
	}
	r.Generate(&model)
	model.SysApi = alist
	db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		_ = e.AddError(errors.New("无权更新该数据"))
		return e
	}
	err = e.initPaths(tx, &model)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return e
	}
	if oldPath != "" {
		var menuList []models.SysMenu
		tx.Where("paths like ? and paths<>'' and length(paths)>?", oldPath+"%", len(oldPath)).Find(&menuList)
		for _, v := range menuList {
			v.Paths = strings.Replace(v.Paths, oldPath, model.Paths, 1)
			tx.Model(&v).Update("paths", v.Paths)
		}
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"修改",
		fmt.Sprintf("更新菜单，ID：%v", model.MenuId),
		string(before),
		string(after),
	)
	return e
}

// Remove 删除SysMenu
func (e *SysMenu) Remove(c *gin.Context, r *dto.SysMenuDeleteReq) *SysMenu {
	var err error
	var data []models.SysMenu

	before, _ := json.Marshal(data)
	db := e.Orm.Find(&data, r.Ids)
	db = e.Orm.Delete(&data)
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		_ = e.AddError(err)
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		_ = e.AddError(err)
	}

	middleware.SetContextOperateLog(c,
		"删除",
		fmt.Sprintf("删除菜单，ID：%v", r.Ids),
		string(before),
		"{}",
	)
	return e
}

// GetList 获取菜单数据
func (e *SysMenu) GetList(r *dto.SysMenuGetPageReq) (list *[]models.SysMenu, err error) {
	list, err = service.GetList[models.SysMenu](e, r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// TODO: 需要修改名字

// SetLabel 修改角色中 设置菜单基础数据
func (e *SysMenu) SetLabel() (m []dto.MenuLabel, err error) {
	var list = new([]models.SysMenu)
	list, err = e.GetList(&dto.SysMenuGetPageReq{})

	m = make([]dto.MenuLabel, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = (*list)[i].MenuId
		e.Label = (*list)[i].Title
		deptsInfo := menuLabelCall(list, e)

		m = append(m, deptsInfo)
	}
	return
}

// GetSysMenuByRoleName 左侧菜单
func (e *SysMenu) GetSysMenuByRoleName(roleName ...string) ([]models.Menu, error) {
	var MenuList []models.Menu
	var role models.SysRole
	var err error
	admin := false
	for _, s := range roleName {
		if s == "admin" {
			admin = true
		}
	}

	if len(roleName) > 0 && admin {
		var data []models.Menu
		err = e.Orm.Where(" menu_type in ('M','C')").
			Order("sort").
			Find(&data).
			Error
		MenuList = data
	} else {
		err = e.Orm.Model(&role).Preload("SysMenu", func(db *gorm.DB) *gorm.DB {
			return db.Where(" menu_type in ('M','C')").Order("sort")
		}).Where("role_name in ?", roleName).Find(&role).
			Error
		for _, e := range *role.SysMenu {
			MenuList = append(MenuList, models.Menu{Title: e.Title, Path: e.Path})
		}
	}

	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("db error:%s", err)
	}
	return MenuList, err
}

// menuLabelCall 递归构造组织数据
func menuLabelCall(eList *[]models.SysMenu, dept dto.MenuLabel) dto.MenuLabel {
	list := *eList

	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].MenuId
		mi.Label = list[j].Title
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType != "F" {
			ms := menuLabelCall(eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

func menuCallp(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	list := *menuList

	min := make([]models.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.CreatedAt = list[j].CreatedAt
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.Children = []models.SysMenu{}

		if mi.MenuType != cModels.Button {
			ms := menuCallp(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

// menuCall 构建菜单树
func menuCall(menuList *[]models.Menu, menu models.Menu) models.Menu {
	list := *menuList

	min := make([]models.Menu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := models.Menu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.Children = []models.Menu{}

		if j == 0 && mi.MenuType == cModels.Directory {
			mDefault := models.Menu{}
			mDefault.Path = menu.Path
			mDefault.Redirect = list[j].Path
			min = append(min, mDefault)
		}
		if mi.MenuType != cModels.Button {
			ms := menuCall(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

func menuDistinct(menuList []models.Menu) (result []models.Menu) {
	distinctMap := make(map[int]struct{}, len(menuList))
	for _, menu := range menuList {
		if _, ok := distinctMap[menu.MenuId]; !ok {
			distinctMap[menu.MenuId] = struct{}{}
			result = append(result, menu)
		}
	}
	return result
}

func recursiveSetMenu(orm *gorm.DB, mIds []int, menus *[]models.Menu) error {
	if len(mIds) == 0 || menus == nil {
		return nil
	}
	var subMenus []models.Menu
	err := orm.Where(fmt.Sprintf(" menu_type in ('%s', '%s', '%s') and menu_id in ?",
		cModels.Directory, cModels.Menu, cModels.Button), mIds).Order("sort").Find(&subMenus).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	subIds := make([]int, 0)
	for _, menu := range subMenus {
		if menu.ParentId != 0 {
			subIds = append(subIds, menu.ParentId)
		}
		if menu.MenuType != cModels.Button {
			*menus = append(*menus, menu)
		}
	}
	return recursiveSetMenu(orm, subIds, menus)
}

// SetMenuRole 获取左侧菜单树使用
func (e *SysMenu) SetMenuRole(roleName string) (m []models.Menu, err error) {
	menus, err := e.getByRoleName(roleName)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	m = make([]models.Menu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return
}
func (e *SysMenu) getByRoleName(roleName string) ([]models.Menu, error) {
	var role models.SysRole
	var err error
	data := make([]models.Menu, 0)

	if roleName == "admin" {
		err = e.Orm.Where(" menu_type in ('M','C') and deleted_at is null").
			Order("sort").
			Find(&data).
			Error
		err = errors.WithStack(err)
	} else {
		role.RoleKey = roleName
		err = e.Orm.Model(&role).Where("role_key = ? ", roleName).Preload("SysMenu").First(&role).Error

		if role.SysMenu != nil {
			mIds := make([]int, 0)
			for _, menu := range *role.SysMenu {
				mIds = append(mIds, menu.MenuId)
			}
			if err := recursiveSetMenu(e.Orm, mIds, &data); err != nil {
				return nil, err
			}

			data = menuDistinct(data)
		}
	}

	sort.Sort(models.SysMenuSlice(data))
	return data, err
}
