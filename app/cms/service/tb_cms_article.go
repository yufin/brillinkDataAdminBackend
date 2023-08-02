package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/app/cms/models"
	"go-admin/app/cms/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
)

type TbCmsArticle struct {
	service.Service
}

// GetPage 获取TbCmsArticle列表
func (e *TbCmsArticle) GetPage(c *dto.TbCmsArticleGetPageReq, p *actions.DataPermission, list *[]models.TbCmsArticle, count *int64) error {
	var err error
	var data models.TbCmsArticle

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbCmsArticleService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbCmsArticle对象
func (e *TbCmsArticle) Get(d *dto.TbCmsArticleGetReq, p *actions.DataPermission, model *models.TbCmsArticle) error {
	var data models.TbCmsArticle

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbCmsArticle error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbCmsArticle对象
func (e *TbCmsArticle) Insert(c *dto.TbCmsArticleInsertReq) error {
	var err error
	var data models.TbCmsArticle
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbCmsArticleService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbCmsArticle对象
func (e *TbCmsArticle) Update(c *dto.TbCmsArticleUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbCmsArticle{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("TbCmsArticleService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbCmsArticle
func (e *TbCmsArticle) Remove(d *dto.TbCmsArticleDeleteReq, p *actions.DataPermission) error {
	var data models.TbCmsArticle

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbCmsArticle error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
