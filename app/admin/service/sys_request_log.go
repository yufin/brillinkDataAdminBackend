package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	cDto "go-admin/common/dto"
	"go-admin/common/service"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysRequestLog struct {
	service.Service
}

// GetPage 获取SysRequestLog列表
func (e *SysRequestLog) GetPage(c *dto.SysRequestLogGetPageReq, list *[]dto.SysRequestLogGetPageResp, count *int64) error {
	var err error
	var data models.SysRequestLog

	err = e.Orm.Table(data.TableName()).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysRequestLogPage error:%s", err.Error())
		return err
	}
	return nil
}

// Get 获取SysRequestLog对象
func (e *SysRequestLog) Get(d *dto.SysRequestLogGetReq, model *models.SysRequestLog) error {
	var data models.SysRequestLog

	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysRequestLog error:%s", err.Error())
		return err
	}
	if err != nil {
		e.Log.Errorf("Service GetSysRequestLog error:%s", err.Error())
		return err
	}
	return nil
}

// Remove 删除SysRequestLog
func (e *SysRequestLog) Remove(d *dto.SysRequestLogDeleteReq) error {
	var err error
	var data models.SysRequestLog

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysRequestLog error:%s", err.Error())
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// GetPv 获取GetPv
func (e *SysRequestLog) GetPv(c *dto.SysRequestLogGetPvReq, list *[]dto.SysRequestLogGetPvResp) error {
	var err error
	var data models.SysRequestLog
	selectStr := "DATE_FORMAT(oper_time, '%m') x,count(1) y"
	groupStr := "x"
	if config.DatabaseConfig.Driver == "sqlite3" {
		selectStr = "strftime('%m',datetime(oper_time,'+8 hour')) x,count(1) y"
	}
	if c.Type == "month" || c.Type == "week" || (c.Type == "" && c.BeginTime != nil && common.GetDiffDays(*c.EndTime, *c.BeginTime) != 0) {
		selectStr = "DATE_FORMAT(oper_time, '%Y-%m-%d') x,count(1) y"
		if config.DatabaseConfig.Driver == "sqlite3" {
			selectStr = "strftime('%Y-%m-%d',datetime(oper_time,'+8 hour')) x,count(1) y"
		}
	} else if c.BeginTime != nil && common.GetDiffDays(*c.EndTime, *c.BeginTime) == 0 {
		selectStr = "DATE_FORMAT(oper_time, '%H') x,count(1) y"
		if config.DatabaseConfig.Driver == "sqlite3" {
			selectStr = "strftime('%H',datetime(oper_time,'+8 hour')) x,count(1) y"
		}
	} else if c.Type == "today" {
		selectStr = "DATE_FORMAT(oper_time, '%H') x,count(1) y"
		if config.DatabaseConfig.Driver == "sqlite3" {
			selectStr = "strftime('%H',datetime(oper_time,'+8 hour')) x,count(1) y"
		}
	}
	err = e.Orm.Table(data.TableName()).Select(selectStr).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Group(groupStr).
		Find(list).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysRequestLogPage error:%s", err.Error())
		return err
	}
	return nil
}

// GetUri 获取GetUri
func (e *SysRequestLog) GetUri(c *dto.SysRequestLogGetUriReq, list *[]dto.SysRequestLogGetPvResp) error {
	var err error
	var data models.SysRequestLog
	selectStr := "DATE_FORMAT(oper_time, '%m') x,count(1) y"
	groupStr := "x"
	if config.DatabaseConfig.Driver == "sqlite3" {
		selectStr = "strftime('%m',datetime(oper_time,'+8 hour')) x,count(1) y"
	}
	if c.Type == "month" || c.Type == "week" || (c.Type == "" && c.BeginTime != nil) {
		selectStr = "DATE_FORMAT(oper_time, '%Y-%m-%d') x,count(1) y"
		if config.DatabaseConfig.Driver == "sqlite3" {
			selectStr = "strftime('%Y-%m-%d',datetime(oper_time,'+8 hour')) x,count(1) y"
		}
	} else if c.Type == "today" {
		selectStr = "DATE_FORMAT(oper_time, '%H') x,count(1) y"
		if config.DatabaseConfig.Driver == "sqlite3" {
			selectStr = "strftime('%H',datetime(oper_time,'+8 hour')) x,count(1) y"
		}
	}
	err = e.Orm.Table(data.TableName()).Select(selectStr).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Group(groupStr).
		Find(list).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysRequestLogPage error:%s", err.Error())
		return err
	}
	return nil
}

// GetMethod 获取GetMethod
func (e *SysRequestLog) GetMethod(c *dto.SysRequestLogGetUriReq, list *[]dto.SysRequestLogGetPvResp) error {
	var err error
	var data models.SysRequestLog

	err = e.Orm.Table(data.TableName()).Select("request_method x,count(1) y").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Group("request_method").
		Find(list).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysRequestLogPage error:%s", err.Error())
		return err
	}
	return nil
}

func (e *SysRequestLog) GetPvWithType(c *dto.SysRequestLogGetPvReq, frist *dto.SysRequestLogGetPvWithTypeResp) error {
	var err error
	var data models.SysRequestLog
	var count int64
	err = e.Orm.Table(data.TableName()).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Count(&count).Error
	frist.Y = int(count)
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysRequestLogPage error:%s", err.Error())
		return err
	}
	return nil
}
