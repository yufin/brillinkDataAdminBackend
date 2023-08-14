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

type TbCmsFriendlink struct {
	service.Service
}

// GetPage 获取TbCmsFriendlink列表
func (e *TbCmsFriendlink) GetPage(c *dto.TbCmsFriendlinkGetPageReq, p *actions.DataPermission, list *[]models.TbCmsFriendlink, count *int64) error {
	var err error
	var data models.TbCmsFriendlink

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbCmsFriendlinkService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbCmsFriendlink对象
func (e *TbCmsFriendlink) Get(d *dto.TbCmsFriendlinkGetReq, p *actions.DataPermission, model *models.TbCmsFriendlink) error {
	var data models.TbCmsFriendlink

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbCmsFriendlink error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbCmsFriendlink对象
func (e *TbCmsFriendlink) Insert(c *dto.TbCmsFriendlinkInsertReq) error {
	var err error
	var data models.TbCmsFriendlink
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbCmsFriendlinkService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbCmsFriendlink对象
func (e *TbCmsFriendlink) Update(c *dto.TbCmsFriendlinkUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbCmsFriendlink{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("TbCmsFriendlinkService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbCmsFriendlink
func (e *TbCmsFriendlink) Remove(d *dto.TbCmsFriendlinkDeleteReq, p *actions.DataPermission) error {
	var data models.TbCmsFriendlink

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbCmsFriendlink error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
