package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type RcOriginContent struct {
	service.Service
}

// GetPage 获取RcOriginContent列表
func (e *RcOriginContent) GetPage(c *dto.RcOriginContentGetPageReq, p *actions.DataPermission, list *[]models.RcOriginContent, count *int64) error {
	var err error
	var data models.RcOriginContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcOriginContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *RcOriginContent) GetPageNoContent(c *dto.RcOriginContentGetPageReq, p *actions.DataPermission, list *[]models.RcOriginContentInfo, count *int64) error {
	var err error
	var data models.RcOriginContentInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcOriginContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcOriginContent对象
func (e *RcOriginContent) Get(d *dto.RcOriginContentGetReq, p *actions.DataPermission, model *models.RcOriginContent) error {
	var data models.RcOriginContent

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcOriginContent error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcOriginContent对象
func (e *RcOriginContent) Insert(c *dto.RcOriginContentInsertReq) error {
	var err error
	var data models.RcOriginContent
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcOriginContentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcOriginContent对象
func (e *RcOriginContent) Update(c *dto.RcOriginContentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcOriginContent{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcOriginContentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcOriginContent
func (e *RcOriginContent) Remove(d *dto.RcOriginContentDeleteReq, p *actions.DataPermission) error {
	var data models.RcOriginContent

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcOriginContent error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
