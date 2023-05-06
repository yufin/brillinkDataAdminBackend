package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type RskcOriginContent struct {
	service.Service
}

// GetPage 获取RskcOriginContent列表
func (e *RskcOriginContent) GetPage(c *dto.RskcOriginContentGetPageReq, p *actions.DataPermission, list *[]models.RskcOriginContent, count *int64) error {
	var err error
	var data models.RskcOriginContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RskcOriginContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *RskcOriginContent) GetPageNoContent(c *dto.RskcOriginContentGetPageReq, p *actions.DataPermission, list *[]models.RskcOriginContentInfo, count *int64) error {
	var err error
	var data models.RskcOriginContentInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RskcOriginContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RskcOriginContent对象
func (e *RskcOriginContent) Get(d *dto.RskcOriginContentGetReq, p *actions.DataPermission, model *models.RskcOriginContent) error {
	var data models.RskcOriginContent

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRskcOriginContent error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RskcOriginContent对象
func (e *RskcOriginContent) Insert(c *dto.RskcOriginContentInsertReq) error {
	var err error
	var data models.RskcOriginContent
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RskcOriginContentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RskcOriginContent对象
func (e *RskcOriginContent) Update(c *dto.RskcOriginContentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RskcOriginContent{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RskcOriginContentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RskcOriginContent
func (e *RskcOriginContent) Remove(d *dto.RskcOriginContentDeleteReq, p *actions.DataPermission) error {
	var data models.RskcOriginContent

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRskcOriginContent error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
