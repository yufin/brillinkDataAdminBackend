package service

import (
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type OriginContent struct {
	service.Service
}

func (e *OriginContent) Insert(c *dto.OriginContentInsertReq) error {
	var err error
	var data models.OriginContent
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RskcOriginContentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *OriginContent) GetPage(c *dto.OriginContentGetPageReq, p *actions.DataPermission, list *[]models.OriginContent, count *int64) error {
	var err error
	var data models.OriginContent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RskcOriginContent GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *OriginContent) GetPageWithoutContent(c *dto.OriginContentGetPageReq, p *actions.DataPermission, list *[]models.OriginContentInfo, count *int64) error {
	var err error
	var data models.OriginContentInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RskcOriginContent GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *OriginContent) CountByInfo(c *dto.OriginContentGetPageReq, count *int64) error {
	var err error
	var data models.OriginContentInfo
	err = e.Orm.Model(&data).Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
	).Count(count).Error
	if err != nil {
		e.Log.Errorf("RskOriginContent CountByInfo error:%s \r\n", err)
		return err
	}
	return nil
}
