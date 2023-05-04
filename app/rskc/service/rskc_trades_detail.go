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

type RskcTradesDetail struct {
	service.Service
}

// GetPage 获取RskcTradesDetail列表
func (e *RskcTradesDetail) GetPage(c *dto.RskcTradesDetailGetPageReq, p *actions.DataPermission, list *[]models.RskcTradesDetail, count *int64) error {
	var err error
	var data models.RskcTradesDetail

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RskcTradesDetailService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RskcTradesDetail对象
func (e *RskcTradesDetail) Get(d *dto.RskcTradesDetailGetReq, p *actions.DataPermission, model *models.RskcTradesDetail) error {
	var data models.RskcTradesDetail

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRskcTradesDetail error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RskcTradesDetail对象
func (e *RskcTradesDetail) Insert(c *dto.RskcTradesDetailInsertReq) error {
	var err error
	var data models.RskcTradesDetail
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RskcTradesDetailService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RskcTradesDetail对象
func (e *RskcTradesDetail) Update(c *dto.RskcTradesDetailUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RskcTradesDetail{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RskcTradesDetailService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RskcTradesDetail
func (e *RskcTradesDetail) Remove(d *dto.RskcTradesDetailDeleteReq, p *actions.DataPermission) error {
	var data models.RskcTradesDetail

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRskcTradesDetail error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
