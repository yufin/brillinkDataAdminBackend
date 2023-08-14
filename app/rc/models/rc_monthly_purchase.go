package models

import "go-admin/common/models"

type RcMonthlyPurchase struct {
	models.Model
	ContentId       int64  `json:"contentId" gorm:"column:content_id;comment:;type:bigint;size:19;"`
	AttributedMonth string `json:"attributedMonth" gorm:"column:attributed_month;comment:;type:varchar(122);size:122;"`
	GncgM           string `json:"gncgM" gorm:"column:gncg_m;comment:国内采购（上游发票统计）;type:varchar(122);size:122;"`
	JkcgM           string `json:"jkcgM" gorm:"column:jkcg_m;comment:进口采购（海关增值税进口缴款书统计）;type:varchar(122);size:122;"`
	HjM             string `json:"hjM" gorm:"column:hj_m;comment:采购总额;type:varchar(122);size:122;"`
	models.ModelTime
	models.ControlBy
}

func (e *RcMonthlyPurchase) TableName() string {
	return "rc_monthly_purchase"
}

func (e *RcMonthlyPurchase) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcMonthlyPurchase) GetId() interface{} {
	return e.Id
}
