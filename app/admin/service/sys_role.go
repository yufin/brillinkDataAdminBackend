package service

import (
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/middleware"
	"go-admin/common/service"
)

type SysRole struct {
	service.Service
}

// GetPage 获取SysRole列表
func (e *SysRole) GetPage(r *dto.SysRoleGetPageReq) (list *[]models.SysRole, count *int64, err error) {
	list, count, err = service.GetPage[models.SysRole](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SysMenu")
	})
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// GetList 获取SysRole列表
func (e *SysRole) GetList(r *dto.SysRoleGetPageReq) (list *[]models.SysRole, err error) {
	list, err = service.GetList[models.SysRole](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SysMenu")
	})
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Get 获取SysRole对象
func (e *SysRole) Get(r *dto.SysRoleGetReq) (model *models.SysRole, err error) {

	model, err = service.Get[models.SysRole](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SysMenu").Preload("SysDept")
	})
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	m := *model.SysMenu
	for i := 0; i < len(m); i++ {
		model.MenuIds = append(model.MenuIds, m[i].MenuId)
	}
	dept := *model.SysDept
	for i := 0; i < len(dept); i++ {
		model.DeptIds = append(model.DeptIds, dept[i].DeptId)
	}
	return
}

// Insert 创建新角色
func (e *SysRole) Insert(c *gin.Context, r *dto.SysRoleInsertReq, cb *casbin.SyncedEnforcer) (err error) {
	var data models.SysRole
	dataMenu, err := service.GetList[models.SysMenu](e, nil, func(db *gorm.DB) *gorm.DB {
		return db.Where("menu_id in ?", r.MenuIds)
	})
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}

	r.SysMenu = *dataMenu
	r.Generate(&data)
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = tx.Create(&data).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	tx.Commit()
	var polices = make([][]string, 0)
	mp := make(map[string]int, 0)
	for _, menu := range *dataMenu {
		for _, api := range *menu.SysApi {
			if _, ok := mp[data.RoleKey+api.Path+api.Method]; !ok {
				mp[data.RoleKey+api.Path+api.Method] = 1
				polices = append(polices, []string{data.RoleKey, api.Path, api.Method})
			}
		}
	}
	_, err = cb.AddNamedPolicies("p", polices)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("add named policy error:%s", err)
	}
	after, _ := json.Marshal(data)
	middleware.SetContextOperateLog(c,
		"新增",
		fmt.Sprintf("创建新角色，ID：%v", data.RoleId),
		"[]",
		string(after),
	)
	return nil
}

// Update 修改SysRole对象
func (e *SysRole) Update(c *gin.Context, r *dto.SysRoleUpdateReq, cb *casbin.SyncedEnforcer) error {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	var model = models.SysRole{}
	var mList = make([]models.SysMenu, 0)
	var dList = make([]models.SysDept, 0)
	tx.Preload("SysMenu").Preload("SysDept").First(&model, r.GetId())
	before, _ := json.Marshal(model)
	tx.Where("menu_id in ?", r.MenuIds).Preload("SysApi").Find(&mList)
	tx.Where("dept_id in ?", r.DeptIds).Find(&dList)
	err = tx.Model(&model).Association("SysMenu").Delete(model.SysMenu)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("delete policy error:%s", err)
		return err
	}
	err = tx.Model(&model).Association("SysDept").Delete(model.SysDept)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("delete policy error:%s", err)
		return err
	}
	r.Generate(&model)
	model.SysMenu = &mList
	model.SysDept = &dList
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	var policys = make([][]string, 0)
	mp := make(map[string]int, 0)
	for _, menu := range mList {
		for _, api := range *menu.SysApi {
			if _, ok := mp[model.RoleKey+api.Path+api.Method]; !ok {
				mp[model.RoleKey+api.Path+api.Method] = 1
				policys = append(policys, []string{model.RoleKey, api.Path, api.Method})
			}
		}
	}

	tx.Commit()

	_, err = cb.RemoveFilteredPolicy(0, model.RoleKey)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("delete policy error:%s", err)
		return err
	}
	_, err = cb.AddNamedPolicies("p", policys)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("add named policy error:%s", err)
	}

	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"修改",
		fmt.Sprintf("更新角色，ID：%v", model.RoleId),
		string(before),
		string(after),
	)
	return nil
}

// Remove 删除SysRole
func (e *SysRole) Remove(c *gin.Context, r *dto.SysRoleDeleteReq) error {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = make([]models.SysRole, 0)
	tx.Preload("SysMenu").Preload("SysDept").
		Where("role_id in ?", r.GetId()).
		Find(&model)
	before, _ := json.Marshal(model)
	db := tx.Select(clause.Associations).Delete(&model)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	middleware.SetContextOperateLog(c,
		"删除",
		fmt.Sprintf("删除角色，ID：%v", r.GetId()),
		string(before),
		"{}",
	)
	return nil
}

// GetRoleMenuId 获取角色对应的菜单ids
func (e *SysRole) GetRoleMenuId(roleId int) ([]int, error) {
	menuIds := make([]int, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].MenuId)
	}
	return menuIds, nil
}

func (e *SysRole) UpdateDataScope(c *gin.Context, r *dto.RoleDataScopeReq) *SysRole {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var dlist = make([]models.SysDept, 0)
	var model = models.SysRole{}
	tx.Preload("SysDept").First(&model, r.RoleId)
	before, _ := json.Marshal(model)
	tx.Where("id in ?", r.DeptIds).Find(&dlist)
	err = tx.Model(&model).Association("SysDept").Delete(model.SysDept)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("delete SysDept error:%s", err)
		_ = e.AddError(err)
		return e
	}
	r.Generate(&model)
	model.SysDept = &dlist
	db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		_ = e.AddError(errors.New("无权更新该数据"))
		return e
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"删除",
		fmt.Sprintf("删除角色，ID：%v", model.RoleId),
		string(before),
		string(after),
	)
	return e
}

// UpdateStatus 修改SysRole对象status
func (e *SysRole) UpdateStatus(c *gin.Context, r *dto.UpdateStatusReq) error {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	tx.First(&model, r.GetId())
	before, _ := json.Marshal(model)
	r.Generate(&model)
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"删除",
		fmt.Sprintf("删除角色，ID：%v", model.RoleId),
		string(before),
		string(after),
	)
	return nil
}

// GetWithName 获取SysRole对象
func (e *SysRole) GetWithName(d *dto.SysRoleByName, model *models.SysRole) *SysRole {
	var err error
	db := e.Orm.Where("role_name = ?", d.RoleName).First(model)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(err, "查看对象不存在或无权查看")
		return e
	}
	if err != nil {
		err = errors.WithStack(err)
		return e
	}
	model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	if err != nil {
		err = errors.WithStack(err)
		return e
	}
	return e
}

// GetById 获取SysRole对象
func (e *SysRole) GetById(roleId int) ([]string, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return permissions, nil
}
