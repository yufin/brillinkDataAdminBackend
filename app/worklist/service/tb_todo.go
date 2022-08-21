package service

import (
	"go-admin/app/worklist/models"
	"go-admin/app/worklist/service/dto"
	cDto "go-admin/common/dto"
	"go-admin/common/service"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TbTodo struct {
	service.Service
}

// GetPage 获取TbTodo列表
func (e *TbTodo) GetPage(c *dto.TbTodoGetPageReq, list *[]models.TbTodo, count *int64) error {
	var err error
	var data models.TbTodo
	c.TbTodoOrder.EndAtOrder = "desc"
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysApiPage error:%s", err)
		return err
	}
	return nil
}

// GetTotal 获取统计信息
func (e *TbTodo) GetTotal(c *dto.TbTodoGetPageReq, normal *int64, active *int64, exception *int64, success *int64) error {
	var err error
	var data models.TbTodo
	c.TbTodoOrder.EndAtOrder = "desc"
	err = e.Orm.Model(&data).
		Where("status = 'success'").
		Count(success).
		Error
	if err != nil {
		e.Log.Errorf("Service GetTotal error:%s", err)
		return err
	}
	err = e.Orm.Model(&data).
		Where("status = 'normal'").
		Count(normal).
		Error
	err = e.Orm.Model(&data).
		Where("status = 'exception'").
		Count(exception).
		Error
	if err != nil {
		e.Log.Errorf("Service GetTotal error:%s", err)
		return err
	}
	err = e.Orm.Model(&data).
		Where("status = 'active'").
		Count(active).
		Error
	if err != nil {
		e.Log.Errorf("Service GetTotal error:%s", err)
		return err
	}
	return nil
}

func (e *TbTodo) GetList(list *[]models.TbTodo) error {
	var err error
	var data models.TbTodo

	err = e.Orm.Model(&data).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("Service GetSysApiPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取 TbTodo 对象with id
func (e *TbTodo) Get(d *dto.TbTodoGetReq, model *models.TbTodo) *TbTodo {
	var data models.TbTodo
	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysApi error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	return e
}

// Insert 创建SysDept对象
func (e *TbTodo) Insert(c *dto.TbTodoInsertReq) error {
	var err error
	var data models.TbTodo
	c.Generate(&data)
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	data.Logo = "https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png"
	data.Status = "active"
	data.Percent = 1
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改
func (e *TbTodo) Update(c *dto.TbTodoUpdateReq) error {
	var model = models.TbTodo{}
	db := e.Orm.First(&model, c.GetId())
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	db = e.Orm.Save(&model)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysApi error:%s", err)
		return err
	}

	return nil
}

// Remove 删除
func (e *TbTodo) Remove(d *dto.TbTodoDeleteReq) error {
	var data models.TbTodo

	db := e.Orm.Model(&data).
		Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
