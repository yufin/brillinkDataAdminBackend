package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/spider/models"
	"go-admin/app/spider/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type EnterpriseWaitList struct {
	service.Service
}

// GetPage 获取EnterpriseWaitList列表
func (e *EnterpriseWaitList) GetPage(c *dto.EnterpriseWaitListGetPageReq, p *actions.DataPermission, list *[]models.EnterpriseWaitList, count *int64) error {
	var (
		err  error
		data models.EnterpriseWaitList
	)

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnterpriseWaitListService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EnterpriseWaitList对象
func (e *EnterpriseWaitList) Get(d *dto.EnterpriseWaitListGetReq, p *actions.DataPermission, model *models.EnterpriseWaitList) error {
	var data models.EnterpriseWaitList

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnterpriseWaitList error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EnterpriseWaitList对象
func (e *EnterpriseWaitList) Insert(c *dto.EnterpriseWaitListInsertReq) error {
	var err error
	var data models.EnterpriseWaitList
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnterpriseWaitListService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EnterpriseWaitList对象
func (e *EnterpriseWaitList) Update(c *dto.EnterpriseWaitListUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EnterpriseWaitList{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("EnterpriseWaitListService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EnterpriseWaitList
func (e *EnterpriseWaitList) Remove(d *dto.EnterpriseWaitListDeleteReq, p *actions.DataPermission) error {
	var data models.EnterpriseWaitList

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveEnterpriseWaitList error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
