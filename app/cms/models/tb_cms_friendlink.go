package models

import (
	"go-admin/common/models"
)

// TbCmsFriendlink
type TbCmsFriendlink struct {
	models.Model
	Name string `json:"name" gorm:"size:128;comment:链接名称"`
	Link string `json:"link" gorm:"size:128;comment:链接地址"`
	models.ModelTime
	models.ControlBy
}

func (TbCmsFriendlink) TableName() string {
	return "tb_cms_friendlink"
}

func (e *TbCmsFriendlink) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbCmsFriendlink) GetId() interface{} {
	return e.Id
}
