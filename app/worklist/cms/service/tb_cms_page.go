package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/cms/models"
	"go-admin/app/cms/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type TbCmsPage struct {
	service.Service
}

// GetPage 获取TbCmsPage列表
func (e *TbCmsPage) GetPage(c *dto.TbCmsPageGetPageReq, p *actions.DataPermission, list *[]models.TbCmsPage, count *int64) error {
	var err error
	var data models.TbCmsPage

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbCmsPageService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbCmsPage对象
func (e *TbCmsPage) Get(d *dto.TbCmsPageGetReq, p *actions.DataPermission, model *models.TbCmsPage) error {
	var data models.TbCmsPage

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbCmsPage error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbCmsPage对象
func (e *TbCmsPage) Insert(c *dto.TbCmsPageInsertReq) error {
	var err error
	var data models.TbCmsPage
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbCmsPageService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbCmsPage对象
func (e *TbCmsPage) Update(c *dto.TbCmsPageUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbCmsPage{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("TbCmsPageService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbCmsPage
func (e *TbCmsPage) Remove(d *dto.TbCmsPageDeleteReq, p *actions.DataPermission) error {
	var data models.TbCmsPage

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbCmsPage error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
