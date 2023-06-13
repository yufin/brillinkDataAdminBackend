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

type RcDecisionParam struct {
	service.Service
}

// GetPage 获取RcDecisionParam列表
func (e *RcDecisionParam) GetPage(c *gin.Context, r *dto.RcDecisionParamGetPageReq, p *actions.DataPermission, list *[]models.RcDecisionParam, count *int64) error {
	var err error
	var data models.RcDecisionParam

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcDecisionParamService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RcDecisionParam列表
func (e *RcDecisionParam) GetList(c *gin.Context, r *dto.RcDecisionParamGetPageReq, p *actions.DataPermission, list *[]models.RcDecisionParam) error {
	var err error
	var data models.RcDecisionParam

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RcDecisionParamService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcDecisionParam对象
func (e *RcDecisionParam) Get(c *gin.Context, r *dto.RcDecisionParamGetReq, p *actions.DataPermission, model *models.RcDecisionParam) error {
	var data models.RcDecisionParam

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcDecisionParam error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcDecisionParam对象
func (e *RcDecisionParam) Insert(c *gin.Context, r *dto.RcDecisionParamInsertReq) error {
	var err error
	var data models.RcDecisionParam
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcDecisionParamService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcDecisionParam对象
func (e *RcDecisionParam) Update(c *gin.Context, r *dto.RcDecisionParamUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcDecisionParam{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcDecisionParamService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcDecisionParam
func (e *RcDecisionParam) Remove(c *gin.Context, r *dto.RcDecisionParamDeleteReq, p *actions.DataPermission) error {
	var data models.RcDecisionParam

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcDecisionParam error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RcDecisionParam) Export(r *dto.RcDecisionParamGetPageReq) (list *[]dto.RcDecisionParamExport, err error) {
	modelList, err := service.GetListOutDiff[models.RcDecisionParam, dto.RcDecisionParamExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
