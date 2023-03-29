package models

import (
	"go-admin/common/models"
	"time"
)

// TbCmsArticle
type TbCmsArticle struct {
	models.Model
	Title    string     `json:"title" gorm:"size:128;comment:页面名称"`
	Mark     string     `json:"mark" gorm:"size:128;comment:页面标记"`
	Source   string     `json:"source" gorm:"size:128;comment:引用来源"`
	Author   string     `json:"author" gorm:"size:128;comment:作者"`
	Category string     `json:"category" gorm:"size:128;comment:分类"`
	Content  string     `json:"content" gorm:"type:text;comment:内容"`
	File     string     `json:"file" gorm:"size:256;comment:上传文件路径"`
	PubTime  *time.Time `json:"pubTime" gorm:"type:datetime(3);comment:发布时间"`
	models.ModelTime
	models.ControlBy
}

func (TbCmsArticle) TableName() string {
	return "tb_cms_article"
}

func (e *TbCmsArticle) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbCmsArticle) GetId() interface{} {
	return e.Id
}
