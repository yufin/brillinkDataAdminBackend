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

type RcSellingSta struct {
	service.Service
}

// GetPage 获取RcSellingSta列表
func (e *RcSellingSta) GetPage(c *gin.Context, r *dto.RcSellingStaGetPageReq, p *actions.DataPermission, list *[]models.RcSellingSta, count *int64) error {
	var err error
	var data models.RcSellingSta

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			cDto.Paginate(r.GetPageSize(), r.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcSellingStaService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取RcSellingSta列表
func (e *RcSellingSta) GetList(c *gin.Context, r *dto.RcSellingStaGetPageReq, p *actions.DataPermission, list *[]models.RcSellingSta) error {
	var err error
	var data models.RcSellingSta

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("RcSellingStaService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcSellingSta对象
func (e *RcSellingSta) Get(c *gin.Context, r *dto.RcSellingStaGetReq, p *actions.DataPermission, model *models.RcSellingSta) error {
	var data models.RcSellingSta

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, r.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcSellingSta error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcSellingSta对象
func (e *RcSellingSta) Insert(c *gin.Context, r *dto.RcSellingStaInsertReq) error {
	var err error
	var data models.RcSellingSta
	r.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcSellingStaService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcSellingSta对象
func (e *RcSellingSta) Update(c *gin.Context, r *dto.RcSellingStaUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcSellingSta{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, r.GetId())
	r.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcSellingStaService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcSellingSta
func (e *RcSellingSta) Remove(c *gin.Context, r *dto.RcSellingStaDeleteReq, p *actions.DataPermission) error {
	var data models.RcSellingSta

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, r.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcSellingSta error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Export 导出SysRole列表
func (e *RcSellingSta) Export(r *dto.RcSellingStaGetPageReq) (list *[]dto.RcSellingStaExport, err error) {
	modelList, err := service.GetListOutDiff[models.RcSellingSta, dto.RcSellingStaExport](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return &modelList, err
}
