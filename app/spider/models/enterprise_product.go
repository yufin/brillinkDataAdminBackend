package models

import (
	"go-admin/common/models"
)

// EnterpriseProduct 企业产品
type EnterpriseProduct struct {
	ProdId      int64  `json:"prodId" gorm:"primaryKey;comment:主键"`
	UscId       string `json:"uscId" gorm:"comment:社会统一信用代码"`
	ProductData string `json:"productData" gorm:"comment:json格式的产品分类"`
	StatusCode  int    `json:"statusCode" gorm:"comment:状态码"`
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
