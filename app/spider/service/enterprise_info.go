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

type EnterpriseInfo struct {
	service.Service
}

// GetPage 获取EnterpriseInfo列表
func (e *EnterpriseInfo) GetPage(c *dto.EnterpriseInfoGetPageReq, p *actions.DataPermission, list *[]models.EnterpriseInfo, count *int64) error {
	var err error
	var data models.EnterpriseInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnterpriseInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EnterpriseInfo对象
func (e *EnterpriseInfo) Get(d *dto.EnterpriseInfoGetReq, p *actions.DataPermission, model *models.EnterpriseInfo) error {
	var data models.EnterpriseInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnterpriseInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EnterpriseInfo对象
func (e *EnterpriseInfo) Insert(c *dto.EnterpriseInfoInsertReq) error {
	var err error
	var data models.EnterpriseInfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnterpriseInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EnterpriseInfo对象
func (e *EnterpriseInfo) Update(c *dto.EnterpriseInfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EnterpriseInfo{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("EnterpriseInfoService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EnterpriseInfo
func (e *EnterpriseInfo) Remove(d *dto.EnterpriseInfoDeleteReq, p *actions.DataPermission) error {
	var data models.EnterpriseInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveEnterpriseInfo error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
