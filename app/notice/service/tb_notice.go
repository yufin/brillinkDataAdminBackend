package service

import (
	"github.com/pkg/errors"
	"go-admin/app/notice/models"
	"go-admin/app/notice/service/dto"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type TbNotice struct {
	service.Service
}

// GetList 获取TbNotice列表
func (e *TbNotice) GetList(c *dto.TbNoticeGetPageReq, list *[]models.TbNotice, count *int64) error {
	var err error
	var data models.TbNotice
	//c.TbNoticeOrder.EndAtOrder = "desc"
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Count(count).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysApiPage error:%s", err)
		return err
	}
	return nil
}
