package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/middleware"
	"go-admin/common/service"

	"gorm.io/gorm"
)

type SysPersonal struct {
	service.Service
}

func (e *SysPersonal) GetAccess(roleId int) ([]string, error) {
	var err error
	model := new([]models.SysMenu)
	if roleId == 1 {
		model, err = service.GetList[models.SysMenu](e, nil)
		if err != nil {
			err = errors.WithStack(err)
			e.GetLog().Errorf("database operation failed:%s \r", err)
			return nil, err
		}
	} else {
		modelRole, err := service.GetById[models.SysRole](e, roleId, func(db *gorm.DB) *gorm.DB {
			return db.Preload("SysMenu")
		})
		if err != nil {
			err = errors.WithStack(err)
			e.GetLog().Errorf("database operation failed:%s \r", err)
			return nil, err
		}
		model = modelRole.SysMenu
	}
	m := make([]string, 0)
	for i := 0; i < len(*model); i++ {
		if (*model)[i].Permission != "" {
			m = append(m, (*model)[i].Permission)
		}
	}
	return m, nil
}

func (e *SysPersonal) GetInfo(userId int64) (*models.SysUser, error) {
	var err error
	model := new(models.SysUser)
	r := dto.SysUserById{}
	r.Id = userId
	model, err = service.Get[models.SysUser](e, &r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return nil, err
	}

	return model, nil
}

func (e *SysPersonal) Update(c *gin.Context, r *dto.UpdatePersonalReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, r.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service SysPersonal error: %s", err)
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
		"个人信息",
	)
	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysPersonal) UpdateAvatar(c *gin.Context, r *dto.UpdatePersonalAvatarReq, p *actions.DataPermission) error {
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
	err = e.Orm.Model(&model).Omit("password").UpdateColumns(&model).Error
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
		"个人信息",
	)
	return nil
}
