package models

import "go-admin/common/models"

type RcMonthlySales struct {
	models.Model
	ContentId       int64  `json:"contentId" gorm:"column:content_id;comment:;type:bigint;size:19;"`
	AttributedMonth string `json:"attributedMonth" gorm:"column:attributed_month;comment:;type:varchar(122);size:122;"`
	Nxsr            string `json:"nxsr" gorm:"column:nxsr;comment:内销收入;type:varchar(122);size:122;"`
	Ckxssr          string `json:"ckxssr" gorm:"column:ckxssr;comment:出口销售收入;type:varchar(122);size:122;"`
	Sbkjzsr         string `json:"sbkjzsr" gorm:"column:sbkjzsr;comment:;type:varchar(122);size:122;"`
	Fpkjsr          string `json:"fpkjsr" gorm:"column:fpkjsr;comment:发票口径收入;type:varchar(122);size:122;"`
	models.ModelTime
	models.ControlBy
}

func (e *RcMonthlySales) TableName() string {
	return "rc_monthly_sales"
}

func (e *RcMonthlySales) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcMonthlySales) GetId() interface{} {
	return e.Id
}
