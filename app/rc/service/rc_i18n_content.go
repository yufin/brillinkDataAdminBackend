package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type RcI18nContent struct {
	service.Service
}

// GetPage 获取RcI18nContent列表
func (e *RcI18nContent) GetPage(c *gin.Context, r *dto.RcI18nContentGetPageReq, p *actions.DataPermission, list *[]models.RcI18nContent, count *int64) error {
	var err error
	var data models.RcI18nContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcI18nContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RcI18nContent列表
func (e *RcI18nContent) GetList(c *gin.Context, r *dto.RcI18nContentGetPageReq, p *actions.DataPermission, list *[]models.RcI18nContent) error {
	var err error
	var data models.RcI18nContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RcI18nContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcI18nContent对象
func (e *RcI18nContent) Get(c *gin.Context, r *dto.RcI18nContentGetReq, p *actions.DataPermission, model *models.RcI18nContent) error {
	var data models.RcI18nContent

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcI18nContent error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcI18nContent对象
func (e *RcI18nContent) Insert(c *gin.Context, r *dto.RcI18nContentInsertReq) error {
	var err error
	var data models.RcI18nContent
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcI18nContentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcI18nContent对象
func (e *RcI18nContent) Update(c *gin.Context, r *dto.RcI18nContentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcI18nContent{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcI18nContentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcI18nContent
func (e *RcI18nContent) Remove(c *gin.Context, r *dto.RcI18nContentDeleteReq, p *actions.DataPermission) error {
	var data models.RcI18nContent

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcI18nContent error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RcI18nContent) Export(r *dto.RcI18nContentGetPageReq) (list *[]dto.RcI18nContentExport, err error) {
	modelList, err := service.GetListOutDiff[models.RcI18nContent, dto.RcI18nContentExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
