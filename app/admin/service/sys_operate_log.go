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

type SysOperateLog struct {
	service.Service
}

// GetPage 获取SysOperateLog列表
func (e *SysOperateLog) GetPage(c *dto.SysOperateLogGetPageReq, p *actions.DataPermission, list *[]dto.SysOperateLogGetPageResp, count *int64) error {
	var err error
	var data models.SysOperateLog

	err = e.Orm.Table(data.TableName()).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			//actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("SysOperateLogService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysOperateLog对象
func (e *SysOperateLog) Get(d *dto.SysOperateLogGetReq, p *actions.DataPermission, model *models.SysOperateLog) error {
	var data models.SysOperateLog
	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysOperateLog error:%s \r\n", err)
		return err
	}
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Remove 删除SysOperateLog
func (e *SysOperateLog) Remove(d *dto.SysOperateLogDeleteReq, p *actions.DataPermission) error {
	var data models.SysOperateLog

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysOperateLog error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
