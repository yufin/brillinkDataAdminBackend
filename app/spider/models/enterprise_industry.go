package models

import (
	"go-admin/common/models"
)

// EnterpriseIndustry 企业行业分类表
type EnterpriseIndustry struct {
	IndId        int64  `json:"indId" gorm:"primaryKey;autoIncrement;comment:主键"`
	EnterpriseId int64  `json:"enterpriseId" gorm:"comment:外键-企业id"`
	IndustryData string `json:"industryData" gorm:"comment:json格式的行业分类"`
	StatusCode   int    `json:"statusCode" gorm:"comment:状态标识码"`
	models.ModelTime
	models.ControlBy
}

func (*EnterpriseIndustry) TableName() string {
	return "enterprise_industry"
}

func (e *EnterpriseIndustry) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseIndustry) GetId() interface{} {
	return e.IndId
}
