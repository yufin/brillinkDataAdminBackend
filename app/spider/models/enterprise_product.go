package models

import (
	"go-admin/common/models"
)

// EnterpriseProduct 企业产品
type EnterpriseProduct struct {
	ProdId       int64  `json:"prodId" gorm:"primaryKey;autoIncrement;comment:主键"`
	EnterpriseId int64  `json:"enterpriseId" gorm:"comment:外键-企业id"`
	ProductData  string `json:"productData" gorm:"comment:json格式的产品分类"`
	StatusCode   int    `json:"statusCode" gorm:"comment:状态码"`
	models.ModelTime
	models.ControlBy
}

func (*EnterpriseProduct) TableName() string {
	return "enterprise_product"
}

func (e *EnterpriseProduct) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseProduct) GetId() interface{} {
	return e.ProdId
}
