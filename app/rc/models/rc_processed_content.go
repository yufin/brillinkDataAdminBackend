package models

import (
	"go-admin/common/models"
)

// RcProcessedContent 报文处理
type RcProcessedContent struct {
	models.Model
	ContentId  int64  `json:"contentId" gorm:"comment:外键(rc_origin_content.id)" xlsx:"外键(rc_origin_content.id)"`
	Content    string `json:"content" gorm:"comment:数据(json字符串格式)" xlsx:"数据(json字符串格式)"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态码" xlsx:"状态码"`
	models.ModelTime
	models.ControlBy
}

func (*RcProcessedContent) TableName() string {
	return "rc_processed_content"
}

func (e *RcProcessedContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcProcessedContent) GetId() interface{} {
	return e.Id
}
