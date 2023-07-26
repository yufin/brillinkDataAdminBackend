package models

import (
	"go-admin/common/models"
)

// RcOriginContent 微众json存储
type RcOriginContent struct {
	models.Model
	//Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UscId          string `json:"uscId" gorm:"comment:统一社会信用代码"`
	EnterpriseName string `json:"enterpriseName" gorm:"comment:企业名称"`
	YearMonth      string `json:"yearMonth" gorm:"comment:数据更新年月"`
	Content        string `json:"content" gorm:"comment:原始JSON STRING数据"`
	StatusCode     int    `json:"statusCode" gorm:"comment:状态码: 1.待解析录入其他表,2.解析并录入完成,3.数据匹配并录入完成"`
	models.ModelTime
	models.ControlBy
}

func (*RcOriginContent) TableName() string {
	return "rc_origin_content"
}

func (e *RcOriginContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcOriginContent) GetId() interface{} {
	return e.Id
}

type RcOriginContentInfo struct {
	models.Model
	//ContentId      string `json:"contentId" gorm:"comment:uuid4"`
	UscId          string `json:"uscId" gorm:"comment:统一社会信用代码"`
	EnterpriseName string `json:"enterpriseName" gorm:"comment:企业名称"`
	YearMonth      string `json:"yearMonth" gorm:"comment:数据更新年月"`
	StatusCode     int    `json:"statusCode" gorm:"comment:状态码"`
	models.ModelTime
	models.ControlBy
}

func (*RcOriginContentInfo) TableName() string {
	return "rc_origin_content"
}

func (e *RcOriginContentInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcOriginContentInfo) GetId() interface{} {
	return e.Id
}
