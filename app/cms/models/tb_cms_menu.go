package models

import (
	"go-admin/common/models"
)

// TbCmsMenu
type TbCmsMenu struct {
	models.Model
	Name   string `json:"name" gorm:"size:128;comment:页面名称"`
	Type   string `json:"type" gorm:"size:128;comment:菜单类型，list列表，page详情页"`
	Link   string `json:"link" gorm:"size:128;comment:引用表的id，或者mark"`
	Parent string `json:"parent" gorm:"size:128;comment:父节点"`
	models.ModelTime
	models.ControlBy
}

func (TbCmsMenu) TableName() string {
	return "tb_cms_menu"
}

func (e *TbCmsMenu) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbCmsMenu) GetId() interface{} {
	return e.Id
}
