package models

import (
	"go-admin/common/models"
)

// Enterprise 企业主体
type Enterprise struct {
	models.Model
	UscId      string `json:"uscId" gorm:"comment:统一社会信用代码"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态标识码"`
	models.ModelTime
	models.ControlBy
}

func (*Enterprise) TableName() string {
	return "enterprise"
}

func (e *Enterprise) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Enterprise) GetId() interface{} {
	return e.Id
}
