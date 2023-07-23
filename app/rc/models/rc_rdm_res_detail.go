package models

import "go-admin/common/models"

type RcRdmResDetail struct {
	models.Model
	ResId int64   `json:"ResId" gorm:"type:bigint;comment:;default:null;"`
	Field string  `json:"field" gorm:"type:varchar(512);comment:;default:null;"`
	Level int     `json:"level" gorm:"type:int;comment:;default:null;"`
	Score float64 `json:"score" gorm:"type:decimal(32, 10);comment:;default:null;"`
	models.ControlBy
	models.ModelTime
}

func (*RcRdmResDetail) TableName() string {
	return "rc_rdm_res_detail"
}

func (e *RcRdmResDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcRdmResDetail) GetId() interface{} {
	return e.Id
}
