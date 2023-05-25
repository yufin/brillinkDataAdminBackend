package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type RskcProcessedContent struct {
	service.Service
}

// GetPage 获取RskcProcessedContent列表
func (e *RskcProcessedContent) GetPage(c *gin.Context, r *dto.RskcProcessedContentGetPageReq, p *actions.DataPermission, list *[]models.RskcProcessedContent, count *int64) error {
	var err error
	var data models.RskcProcessedContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RskcProcessedContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RskcProcessedContent列表
func (e *RskcProcessedContent) GetList(c *gin.Context, r *dto.RskcProcessedContentGetPageReq, p *actions.DataPermission, list *[]models.RskcProcessedContent) error {
	var err error
	var data models.RskcProcessedContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RskcProcessedContentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RskcProcessedContent对象
func (e *RskcProcessedContent) Get(c *gin.Context, r *dto.RskcProcessedContentGetReq, p *actions.DataPermission, model *models.RskcProcessedContent) error {
	var data models.RskcProcessedContent

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRskcProcessedContent error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RskcProcessedContent对象
func (e *RskcProcessedContent) Insert(c *gin.Context, r *dto.RskcProcessedContentInsertReq) error {
	var err error
	var data models.RskcProcessedContent
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RskcProcessedContentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RskcProcessedContent对象
func (e *RskcProcessedContent) Update(c *gin.Context, r *dto.RskcProcessedContentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RskcProcessedContent{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RskcProcessedContentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RskcProcessedContent
func (e *RskcProcessedContent) Remove(c *gin.Context, r *dto.RskcProcessedContentDeleteReq, p *actions.DataPermission) error {
	var data models.RskcProcessedContent

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRskcProcessedContent error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RskcProcessedContent) Export(r *dto.RskcProcessedContentGetPageReq) (list *[]dto.RskcProcessedContentExport, err error) {
	modelList, err := service.GetListOutDiff[models.RskcProcessedContent, dto.RskcProcessedContentExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
