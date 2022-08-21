package models

import (
	"go-admin/common/models"
	"time"
)

// TbCmsPage 单页面展示
type TbCmsPage struct {
	models.Model
	Title   string    `json:"title" gorm:"size:128;comment:页面名称"`
	Mark    string    `json:"mark" gorm:"size:128;comment:页面标记"`
	Source  string    `json:"source" gorm:"size:128;comment:引用来源"`
	Author  string    `json:"author" gorm:"size:128;comment:作者"`
	Content string    `json:"content" gorm:"type:text;comment:内容"`
	PubTime time.Time `json:"pubTime" gorm:"type:datetime(3);comment:发布时间"`
	models.ModelTime
	models.ControlBy
}

func (TbCmsPage) TableName() string {
	return "tb_cms_page"
}

func (e *TbCmsPage) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbCmsPage) GetId() interface{} {
	return e.Id
}
