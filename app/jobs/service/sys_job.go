package service

import (
	"github.com/pkg/errors"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
	"time"

	"github.com/robfig/cron/v3"

	"go-admin/app/jobs"
	"go-admin/app/jobs/models"
	"go-admin/app/jobs/service/dto"
	"go-admin/common/service"
)

type SysJob struct {
	service.Service
	Cron *cron.Cron
}

// GetPage 获取SysJob列表
func (e *SysJob) GetPage(c *dto.SysJobGetPageReq, p *actions.DataPermission, list *[]models.SysJob, count *int64) error {
	var err error
	var data models.SysJob

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysJobService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysJob对象
func (e *SysJob) Get(d *dto.SysJobGetReq, p *actions.DataPermission, model *models.SysJob) error {
	var data models.SysJob

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysJob error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysJob对象
func (e *SysJob) Insert(c *dto.SysJobInsertReq) error {
	var err error
	var data models.SysJob
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysJobService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysJob对象
func (e *SysJob) Update(c *dto.SysJobUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SysJob{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SysJobService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysJob
func (e *SysJob) Remove(d *dto.SysJobDeleteReq, p *actions.DataPermission) error {
	var data models.SysJob

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysJob error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// RemoveJob 删除job
func (e *SysJob) RemoveJob(c *dto.SysJobDeleteReq) error {
	var err error
	var data models.SysJob
	err = e.Orm.Table(data.TableName()).Where("job_id = ?", c.GetId()).First(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	cn := jobs.Remove(e.Cron, int(data.EntryId))

	select {
	case res := <-cn:
		if res {
			err = e.Orm.Table(data.TableName()).Where("entry_id = ?", data.EntryId).Update("entry_id", 0).Error
			if err != nil {
				e.Log.Errorf("db error: %s", err)
			}
			return err
		}
	case <-time.After(time.Second * 1):
		e.Msg = "操作超时！"
		return nil
	}
	return nil
}

// StartJob 启动任务
func (e *SysJob) StartJob(c *dto.SysJobGetReq) error {
	var data models.SysJob
	var err error
	err = e.Orm.Table(data.TableName()).First(&data, c.JobId).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if data.JobType == 1 {
		var j = &jobs.HttpJob{}
		j.InvokeTarget = data.InvokeTarget
		j.CronExpression = data.CronExpression
		j.JobId = data.JobId
		j.Name = data.JobName
		data.EntryId, err = jobs.AddJob(e.Cron, j)
		if err != nil {
			e.Log.Errorf("jobs AddJob[HttpJob] error: %s", err)
		}
	} else {
		var j = &jobs.ExecJob{}
		j.InvokeTarget = data.InvokeTarget
		j.CronExpression = data.CronExpression
		j.JobId = data.JobId
		j.Name = data.JobName
		j.Args = data.Args
		data.EntryId, err = jobs.AddJob(e.Cron, j)
		if err != nil {
			err = errors.WithStack(err)
			e.Log.Errorf("jobs AddJob[ExecJob] error: %s", err)
		}
	}
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	err = e.Orm.Table(data.TableName()).Where("job_id = ?", c.JobId).Updates(&data).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("db error: %s", err)
	}
	return err
}
