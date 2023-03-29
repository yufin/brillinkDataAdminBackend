package service

import (
	"github.com/pkg/errors"

	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type SysAbnormalLog struct {
	service.Service
}

// GetPage 获取SysAbnormalLog列表
func (e *SysAbnormalLog) GetPage(c *dto.SysAbnormalLogGetPageReq, p *actions.DataPermission, list *[]dto.SysAbnormalLogGetPageResp, count *int64) error {
	var err error
	var data models.SysAbnormalLog

	err = e.Orm.Table(data.TableName()).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("SysAbnormalLogService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysAbnormalLog对象
func (e *SysAbnormalLog) Get(d *dto.SysAbnormalLogGetReq, p *actions.DataPermission, model *models.SysAbnormalLog) error {
	var data models.SysAbnormalLog

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysAbnormalLog error:%s \r\n", err)
		return err
	}
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Remove 删除SysAbnormalLog
func (e *SysAbnormalLog) Remove(d *dto.SysAbnormalLogDeleteReq, p *actions.DataPermission) error {
	var data models.SysAbnormalLog

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysAbnormalLog error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
