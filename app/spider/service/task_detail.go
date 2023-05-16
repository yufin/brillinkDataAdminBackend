package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/spider/models"
	"go-admin/app/spider/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type TaskDetail struct {
	service.Service
}

// GetPage 获取TaskDetail列表
func (e *TaskDetail) GetPage(c *gin.Context, r *dto.TaskDetailGetPageReq, p *actions.DataPermission, list *[]models.TaskDetail, count *int64) error {
	var err error
	var data models.TaskDetail

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TaskDetailService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取TaskDetail列表
func (e *TaskDetail) GetList(c *gin.Context, r *dto.TaskDetailGetPageReq, p *actions.DataPermission, list *[]models.TaskDetail) error {
	var err error
	var data models.TaskDetail

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("TaskDetailService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TaskDetail对象
func (e *TaskDetail) Get(c *gin.Context, r *dto.TaskDetailGetReq, p *actions.DataPermission, model *models.TaskDetail) error {
	var data models.TaskDetail

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTaskDetail error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TaskDetail对象
func (e *TaskDetail) Insert(c *gin.Context, r *dto.TaskDetailInsertReq) error {
	var err error
	var data models.TaskDetail
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TaskDetailService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TaskDetail对象
func (e *TaskDetail) Update(c *gin.Context, r *dto.TaskDetailUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TaskDetail{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("TaskDetailService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TaskDetail
func (e *TaskDetail) Remove(c *gin.Context, r *dto.TaskDetailDeleteReq, p *actions.DataPermission) error {
	var data models.TaskDetail

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveTaskDetail error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
