package models

import (
	"go-admin/common/models"
)

// RcI18nContent i18n content
type RcI18nContent struct {
	models.Model
	ProcessedId int64  `json:"processedId" gorm:"comment:rskc_processed_content.id" xlsx:"rskc_processed_content.id"`
	Lang        string `json:"lang" gorm:"comment:语言类型(en,...)" xlsx:"语言类型(en,...)"`
	Content     string `json:"content" gorm:"comment:报文json string" xlsx:"报文json string"`
	models.ModelTime
	models.ControlBy
}

func (*RcI18nContent) TableName() string {
	return "rc_i18n_content"
}

func (e *RcI18nContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcI18nContent) GetId() interface{} {
	return e.Id
}
