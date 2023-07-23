package models

import "go-admin/common/models"

type RcRdmResult struct {
	models.Model
	AppType int    `json:"appType" gorm:"type:int;comment:;default:null;"`
	DepId   int64  `json:"depId" gorm:"type:bigint;comment:;default:null;"`
	Comment string `json:"comment" gorm:"type:text;comment:;default:null;"`
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
