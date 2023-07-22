package models

import "go-admin/common/models"

type RcRdmResult struct {
	models.Model
	DepId int64   `json:"depId" gorm:"type:bigint;comment:;default:null;"`
	Field string  `json:"field" gorm:"type:varchar(512);comment:;default:null;"`
	Level int     `json:"level" gorm:"type:int;comment:;default:null;"`
	Score float64 `json:"score" gorm:"type:decimal(32, 10);comment:;default:null;"`
	models.ControlBy
	models.ModelTime
}

func (*RcRdmResult) TableName() string {
	return "rc_rdm_result"
}

func (e *RcRdmResult) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcRdmResult) GetId() interface{} {
	return e.Id
}
