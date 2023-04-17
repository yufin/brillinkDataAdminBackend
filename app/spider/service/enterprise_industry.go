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

type EnterpriseIndustry struct {
	service.Service
}

// GetPage 获取EnterpriseIndustry列表
func (e *EnterpriseIndustry) GetPage(c *dto.EnterpriseIndustryGetPageReq, p *actions.DataPermission, list *[]models.EnterpriseIndustry, count *int64) error {
	var err error
	var data models.EnterpriseIndustry

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnterpriseIndustryService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EnterpriseIndustry对象
func (e *EnterpriseIndustry) Get(d *dto.EnterpriseIndustryGetReq, p *actions.DataPermission, model *models.EnterpriseIndustry) error {
	var data models.EnterpriseIndustry

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnterpriseIndustry error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EnterpriseIndustry对象
func (e *EnterpriseIndustry) Insert(c *dto.EnterpriseIndustryInsertReq) error {
	var err error
	var data models.EnterpriseIndustry
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnterpriseIndustryService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EnterpriseIndustry对象
func (e *EnterpriseIndustry) Update(c *dto.EnterpriseIndustryUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EnterpriseIndustry{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("EnterpriseIndustryService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EnterpriseIndustry
func (e *EnterpriseIndustry) Remove(d *dto.EnterpriseIndustryDeleteReq, p *actions.DataPermission) error {
	var data models.EnterpriseIndustry

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveEnterpriseIndustry error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
