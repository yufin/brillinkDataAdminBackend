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

type RcDecisionResult struct {
	service.Service
}

// GetPage 获取RcDecisionResult列表
func (e *RcDecisionResult) GetPage(c *gin.Context, r *dto.RcDecisionResultGetPageReq, p *actions.DataPermission, list *[]models.RcDecisionResult, count *int64) error {
	var err error
	var data models.RcDecisionResult

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcDecisionResultService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RcDecisionResult列表
func (e *RcDecisionResult) GetList(c *gin.Context, r *dto.RcDecisionResultGetPageReq, p *actions.DataPermission, list *[]models.RcDecisionResult) error {
	var err error
	var data models.RcDecisionResult

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RcDecisionResultService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcDecisionResult对象
func (e *RcDecisionResult) Get(c *gin.Context, r *dto.RcDecisionResultGetReq, p *actions.DataPermission, model *models.RcDecisionResult) error {
	var data models.RcDecisionResult

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcDecisionResult error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcDecisionResult对象
func (e *RcDecisionResult) Insert(c *gin.Context, r *dto.RcDecisionResultInsertReq) error {
	var err error
	var data models.RcDecisionResult
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcDecisionResultService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcDecisionResult对象
func (e *RcDecisionResult) Update(c *gin.Context, r *dto.RcDecisionResultUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcDecisionResult{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcDecisionResultService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcDecisionResult
func (e *RcDecisionResult) Remove(c *gin.Context, r *dto.RcDecisionResultDeleteReq, p *actions.DataPermission) error {
	var data models.RcDecisionResult

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcDecisionResult error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RcDecisionResult) Export(r *dto.RcDecisionResultGetPageReq) (list *[]dto.RcDecisionResultExport, err error) {
	modelList, err := service.GetListOutDiff[models.RcDecisionResult, dto.RcDecisionResultExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
