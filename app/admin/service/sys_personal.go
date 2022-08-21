package service

import (
	"github.com/pkg/errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
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
