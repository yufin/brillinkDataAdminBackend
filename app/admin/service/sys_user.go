package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/exception"
	"go-admin/common/middleware"
	"gorm.io/gorm"

	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(r *dto.SysUserGetPageReq, p *actions.DataPermission, list *[]dto.SysUserGetPageResp, count *int64) error {
	var err error
	var data models.SysUser
	err = e.Orm.Table(data.TableName()).Preload("Dept").
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, p *actions.DataPermission, model *models.SysUser) error {
	var data models.SysUser

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(err, "查看对象不存在或无权查看")
		return err
	}
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

// Insert 创建新用户
func (e *SysUser) Insert(c *gin.Context, r *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", r.Username).Count(&i).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	if i > 0 {
		panic(exception.WithMsg(5000, "用户名已存在!", errors.New("用户名已存在!")))
	}
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	after, _ := json.Marshal(data)
	middleware.SetContextOperateLog(c,
		"新增",
		fmt.Sprintf("创建新用户，ID：%v", data.UserId),
		"{}",
		string(after),
		"用户",
	)
	return nil
}

// Update 更新用户信息
func (e *SysUser) Update(c *gin.Context, r *dto.SysUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, r.GetId())
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("对象不存在或无权查看")
		return err
	}
	if err = db.Error; err != nil {
		err = errors.WithStack(err)
		return err
	}
	before, _ := json.Marshal(model)
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	r.Generate(&model)
	err = e.Orm.Model(&model).Omit("password").UpdateColumns(&model).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"修改",
		fmt.Sprintf("更新用户信息，ID：%v", model.UserId),
		string(before),
		string(after),
		"用户",
	)
	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *gin.Context, r *dto.UpdateSysUserAvatarReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, r.GetId())
	if err = db.Error; err != nil {
		err = errors.WithStack(err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	before, _ := json.Marshal(model)
	r.Generate(&model)
	err = e.Orm.Save(&model).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"修改",
		fmt.Sprintf("更新用户头像，ID：%v", model.UserId),
		string(before),
		string(after),
		"用户",
	)
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *gin.Context, r *dto.UpdateSysUserStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, r.GetId())
	if err = db.Error; err != nil {
		err = errors.WithStack(err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	before, _ := json.Marshal(model)
	r.Generate(&model)
	err = e.Orm.Save(&model).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"修改",
		fmt.Sprintf("更新用户状态，ID：%v", model.UserId),
		string(before),
		string(after),
		"用户",
	)
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *gin.Context, r *dto.ResetSysUserPwdReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	before, after, ok, err := service.GetAndUpdate[models.SysUser](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return err
	}
	if ok {
		middleware.SetContextOperateLog(c,
			"修改",
			fmt.Sprintf("重置用户密码，ID：%v", model.UserId),
			string(before),
			string(after),
			"用户",
		)
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(c *gin.Context, r *dto.SysUserById, p *actions.DataPermission) error {
	var err error
	var data []models.SysUser
	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(models.SysUser{}.TableName(), p),
		).First(&data, r.GetId())
	if err = db.Error; err != nil {
		err = errors.WithStack(err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	before, _ := json.Marshal(data)
	db = e.Orm.Delete(&data)
	if err = db.Error; err != nil {
		err = errors.WithStack(err)
		return err
	}
	after, _ := json.Marshal(data)
	middleware.SetContextOperateLog(c,
		"删除",
		fmt.Sprintf("删除了User数据，ID：%v", r.GetId()),
		string(before),
		string(after),
		"用户",
	)
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(c *gin.Context, r *dto.UpdateSysUserPwdReq, p *actions.DataPermission) (err error) {
	if r.Password == "" {
		return errors.New("新密码不能为空")
	}
	model := new(models.SysUser)

	model, err = service.Get[models.SysUser](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			actions.Permission(models.SysUser{}.TableName(), p),
		).Select("UserId", "Password", "Salt")
	})

	var ok bool
	ok, err = pkg.CompareHashAndPassword(model.Password, r.OldPassword)
	if err != nil {
		return errors.WithStack(err)
	}
	if !ok {
		err = errors.Wrap(err, "账号密码不正确")
		return
	}

	before, after, ok, err := service.GetAndUpdate[models.SysUser](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	if ok {
		middleware.SetContextOperateLog(c,
			"修改",
			fmt.Sprintf("更新了用户密码数据，ID：%v", model.UserId),
			string(before),
			string(after),
			"用户",
		)
	}
	return
}

func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	err := e.Orm.Preload("Dept").First(user, c.GetId()).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	err = e.Orm.Find(roles, user.RoleId).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	err = e.Orm.Find(posts, user.PostIds).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}
