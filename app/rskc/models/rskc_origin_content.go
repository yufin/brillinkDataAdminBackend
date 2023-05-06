package models

import (
	"go-admin/common/models"
)

// RskcOriginContent 微众json存储
type RskcOriginContent struct {
	models.Model
	ContentId  string `json:"contentId" gorm:"comment:uuid4"`
	UscId      string `json:"uscId" gorm:"comment:统一社会信用代码"`
	YearMonth  string `json:"yearMonth" gorm:"comment:数据更新年月"`
	Content    string `json:"content" gorm:"comment:原始JSON STRING数据"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态码"`
	models.ModelTime
	models.ControlBy
}

func (*RskcOriginContent) TableName() string {
	return "rskc_origin_content"
}

func (e *RskcOriginContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RskcOriginContent) GetId() interface{} {
	return e.Id
}

type RskcOriginContentInfo struct {
	models.Model
	ContentId  string `json:"contentId" gorm:"comment:uuid4"`
	UscId      string `json:"uscId" gorm:"comment:统一社会信用代码"`
	YearMonth  string `json:"yearMonth" gorm:"comment:数据更新年月"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态码"`
	models.ModelTime
	models.ControlBy
}

func (*RskcOriginContentInfo) TableName() string {
	return "rskc_origin_content"
}

func (e *RskcOriginContentInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RskcOriginContentInfo) GetId() interface{} {
	return e.Id
}
