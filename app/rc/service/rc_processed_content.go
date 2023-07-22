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

type RcProcessedContent struct {
	service.Service
}

// GetPage 获取RcProcessedContent列表
func (e *RcProcessedContent) GetPage(c *gin.Context, r *dto.RcProcessedContentGetPageReq, p *actions.DataPermission, list *[]models.RcProcessedContent, count *int64) error {
	var err error
	var data models.RcProcessedContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcProcessedContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RcProcessedContent列表
func (e *RcProcessedContent) GetList(c *gin.Context, r *dto.RcProcessedContentGetPageReq, p *actions.DataPermission, list *[]models.RcProcessedContent) error {
	var err error
	var data models.RcProcessedContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RcProcessedContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcProcessedContent对象
func (e *RcProcessedContent) Get(c *gin.Context, r *dto.RcProcessedContentGetReq, p *actions.DataPermission, model *models.RcProcessedContent) error {
	var data models.RcProcessedContent

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcProcessedContent error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcProcessedContent对象
func (e *RcProcessedContent) Insert(c *gin.Context, r *dto.RcProcessedContentInsertReq) error {
	var err error
	var data models.RcProcessedContent
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcProcessedContentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcProcessedContent对象
func (e *RcProcessedContent) Update(c *gin.Context, r *dto.RcProcessedContentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcProcessedContent{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcProcessedContentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcProcessedContent
func (e *RcProcessedContent) Remove(c *gin.Context, r *dto.RcProcessedContentDeleteReq, p *actions.DataPermission) error {
	var data models.RcProcessedContent

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcProcessedContent error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RcProcessedContent) Export(r *dto.RcProcessedContentGetPageReq) (list *[]dto.RcProcessedContentExport, err error) {
	modelList, err := service.GetListOutDiff[models.RcProcessedContent, dto.RcProcessedContentExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
