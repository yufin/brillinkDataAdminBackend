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

type RcDependencyData struct {
	service.Service
}

// GetPage 获取RcDependencyData列表
func (e *RcDependencyData) GetPage(c *gin.Context, r *dto.RcDependencyDataGetPageReq, p *actions.DataPermission, list *[]models.RcDependencyData, count *int64) error {
	var err error
	var data models.RcDependencyData

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcDependencyDataService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RcDependencyData列表
func (e *RcDependencyData) GetList(c *gin.Context, r *dto.RcDependencyDataGetPageReq, p *actions.DataPermission, list *[]models.RcDependencyData) error {
	var err error
	var data models.RcDependencyData

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RcDependencyDataService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcDependencyData对象
func (e *RcDependencyData) Get(c *gin.Context, r *dto.RcDependencyDataGetReq, p *actions.DataPermission, model *models.RcDependencyData) error {
	var data models.RcDependencyData

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcDependencyData error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcDependencyData对象
func (e *RcDependencyData) Insert(c *gin.Context, r *dto.RcDependencyDataInsertReq) error {
	var err error
	var data models.RcDependencyData
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcDependencyDataService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcDependencyData对象
func (e *RcDependencyData) Update(c *gin.Context, r *dto.RcDependencyDataUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcDependencyData{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcDependencyDataService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcDependencyData
func (e *RcDependencyData) Remove(c *gin.Context, r *dto.RcDependencyDataDeleteReq, p *actions.DataPermission) error {
	var data models.RcDependencyData

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcDependencyData error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RcDependencyData) Export(r *dto.RcDependencyDataGetPageReq) (list *[]dto.RcDependencyDataExport, err error) {
	modelList, err := service.GetListOutDiff[models.RcDependencyData, dto.RcDependencyDataExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
