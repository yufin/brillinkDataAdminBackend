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

type EnterpriseRanking struct {
	service.Service
}

// GetPage 获取EnterpriseRanking列表
func (e *EnterpriseRanking) GetPage(c *dto.EnterpriseRankingGetPageReq, p *actions.DataPermission, list *[]models.EnterpriseRanking, count *int64) error {
	var err error
	var data models.EnterpriseRanking

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnterpriseRankingService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EnterpriseRanking对象
func (e *EnterpriseRanking) Get(d *dto.EnterpriseRankingGetReq, p *actions.DataPermission, model *models.EnterpriseRanking) error {
	var data models.EnterpriseRanking

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnterpriseRanking error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EnterpriseRanking对象
func (e *EnterpriseRanking) Insert(c *dto.EnterpriseRankingInsertReq) error {
	var err error
	var data models.EnterpriseRanking
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnterpriseRankingService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EnterpriseRanking对象
func (e *EnterpriseRanking) Update(c *dto.EnterpriseRankingUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EnterpriseRanking{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("EnterpriseRankingService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EnterpriseRanking
func (e *EnterpriseRanking) Remove(d *dto.EnterpriseRankingDeleteReq, p *actions.DataPermission) error {
	var data models.EnterpriseRanking

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveEnterpriseRanking error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
