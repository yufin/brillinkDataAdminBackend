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

type Enterprise struct {
	service.Service
}

// GetPage 获取Enterprise列表
func (e *Enterprise) GetPage(c *dto.EnterpriseGetPageReq, p *actions.DataPermission, list *[]models.Enterprise, count *int64) error {
	var err error
	var data models.Enterprise

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnterpriseService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Enterprise对象
func (e *Enterprise) Get(d *dto.EnterpriseGetReq, p *actions.DataPermission, model *models.Enterprise) error {
	var data models.Enterprise

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnterprise error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Enterprise对象
func (e *Enterprise) Insert(c *dto.EnterpriseInsertReq) error {
	var err error
	var data models.Enterprise
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnterpriseService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Enterprise对象
func (e *Enterprise) Update(c *dto.EnterpriseUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Enterprise{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("EnterpriseService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Enterprise
func (e *Enterprise) Remove(d *dto.EnterpriseDeleteReq, p *actions.DataPermission) error {
	var data models.Enterprise

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveEnterprise error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
