package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type RcTradesDetail struct {
	service.Service
}

func (e *RcTradesDetail) GetJoinWaitList(contentId int64, resp *[]dto.RcTradesDetailJoinWaitListResp, p *actions.DataPermission) error {
	var err error
	data := models.RcTradesDetail{}

	err = e.Orm.Model(&data).
		Scopes(actions.Permission(data.TableName(), p)).
		Select(" rc_trades_detail.id as rtd_id, rc_trades_detail.status_code as rtd_status_code, rc_trades_detail.content_id as content_id,"+
			" ewl.enterprise_name as enterprise_name, ewl.status_code as ewl_status_code, ewl.id as ewl_id").
		Joins("left join enterprise_wait_list ewl on rc_trades_detail.enterprise_name = ewl.enterprise_name").
		Where("rc_trades_detail.content_id = ?", contentId).
		Scan(&resp).
		Error
	if err != nil {
		e.Log.Errorf("RcTradesDetailService GetJoinWaitList error:%s \r\n", err)
		return err
	}
	return nil
}

// GetPage 获取RcTradesDetail列表
func (e *RcTradesDetail) GetPage(c *dto.RcTradesDetailGetPageReq, p *actions.DataPermission, list *[]models.RcTradesDetail, count *int64) error {
	var err error
	var data models.RcTradesDetail

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcTradesDetailService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcTradesDetail对象
func (e *RcTradesDetail) Get(d *dto.RcTradesDetailGetReq, p *actions.DataPermission, model *models.RcTradesDetail) error {
	var data models.RcTradesDetail

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcTradesDetail error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcTradesDetail对象
func (e *RcTradesDetail) Insert(c *dto.RcTradesDetailInsertReq) error {
	var err error
	var data models.RcTradesDetail
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcTradesDetailService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcTradesDetail对象
func (e *RcTradesDetail) Update(c *dto.RcTradesDetailUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RcTradesDetail{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("RcTradesDetailService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RcTradesDetail
func (e *RcTradesDetail) Remove(d *dto.RcTradesDetailDeleteReq, p *actions.DataPermission) error {
	var data models.RcTradesDetail

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		err = errors.WithStack(db.Error)
		e.Log.Errorf("Service RemoveRcTradesDetail error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
