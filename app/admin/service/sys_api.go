package service

import (
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

type SysApi struct {
	service.Service
}

// GetPage 获取SysApi列表
func (e *SysApi) GetPage(r *dto.SysApiGetPageReq, p *actions.DataPermission) (list *[]models.SysApi, count *int64, err error) {
	list, count, err = service.GetPage[models.SysApi](e, r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// GetList 获取SysApi列表
func (e *SysApi) GetList(r *dto.SysApiGetPageReq, p *actions.DataPermission) (list *[]models.SysApi, err error) {
	list, err = service.GetList[models.SysApi](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Where("is_history = ? and type = ?", false, "BUS")
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Get 获取SysApi对象with id
func (e *SysApi) Get(r *dto.SysApiGetReq, p *actions.DataPermission) (model *models.SysApi, err error) {
	model, err = service.Get[models.SysApi](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			actions.Permission(models.SysApi{}.TableName(), p),
		)
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Update 修改SysApi对象
func (e *SysApi) Update(c *gin.Context, r *dto.SysApiUpdateReq, p *actions.DataPermission) (err error) {
	var before, after []byte
	before, after, ok, err := service.GetAndUpdate[models.SysApi](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			actions.Permission(models.SysApi{}.TableName(), p),
		)
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	if ok {
		middleware.SetContextOperateLog(c,
			"修改",
			fmt.Sprintf("更新了Post数据，ID：%v", r.GetId()),
			string(before),
			string(after),
		)
	}
	return
}

// Remove 删除SysApi
func (e *SysApi) Remove(c *gin.Context, r *dto.SysApiDeleteReq, p *actions.DataPermission) (err error) {
	ok, err := service.Delete[models.SysApi](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			actions.Permission(models.SysApi{}.TableName(), p),
		)
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	if ok {
		middleware.SetContextOperateLog(c,
			"删除",
			fmt.Sprintf("删除了数据，ID：%v", r.GetId()),
			"{}",
			"{}",
		)
	}
	return
}
