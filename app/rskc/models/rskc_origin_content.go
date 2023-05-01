package models

import "go-admin/common/models"

type OriginContent struct {
	ContentId         string `json:"contentId" gorm:"primaryKey;size:36;comment:主键(uuid4)"`
	UscId             string `json:"uscId" gorm:"comment:社会统一信用代码"`
	YearMonth         string `json:"yearMonth" gorm:"size:10;comment:数据更新年月(eg:2020-08);"`
	OriginJsonContent string `json:"originJsonContent" gorm:"comment:json数据字符串"`
	models.ModelTime
	models.ControlBy
}

func (*OriginContent) TableName() string {
	return "rskc_origin_content"
}

func (e *OriginContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OriginContent) GetId() interface{} {
	return e.ContentId
}

type OriginContentInfo struct {
	ContentId string `json:"contentId" gorm:"primaryKey;size:36;comment:主键(uuid4)"`
	UscId     string `json:"uscId" gorm:"comment:社会统一信用代码"`
	YearMonth string `json:"yearMonth" gorm:"size:10;comment:数据更新年月(eg:2020-08);"`
	models.ModelTime
	models.ControlBy
}

func (*OriginContentInfo) TableName() string {
	return "rskc_origin_content"
}

func (e *OriginContentInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OriginContentInfo) GetId() interface{} {
	return e.ContentId
}
