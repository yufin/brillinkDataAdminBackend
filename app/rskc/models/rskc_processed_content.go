package models

import (
	"go-admin/common/models"
)

// RskcProcessedContent 报文处理
type RskcProcessedContent struct {
	models.Model
	ContentId  int64  `json:"contentId" gorm:"comment:外键(rskc_origin_content.id)" xlsx:"外键(rskc_origin_content.id)"`
	Content    string `json:"content" gorm:"comment:数据(json字符串格式)" xlsx:"数据(json字符串格式)"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态码" xlsx:"状态码"`
	models.ModelTime
	models.ControlBy
}

func (*RskcProcessedContent) TableName() string {
	return "rskc_processed_content"
}

func (e *RskcProcessedContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RskcProcessedContent) GetId() interface{} {
	return e.Id
}
