package models

import "go-admin/common/models"

type RcScorecard struct {
	models.Model
	Identifier string `json:"identifier" gorm:"type:varchar(512);comment:;default:null;"`
	Alias      string `json:"alias" gorm:"type:varchar(512);comment:;default:null;"`
	Desc       string `json:"desc" gorm:"type:text;comment:;default:null;"`
	models.ModelTime
	models.ControlBy
}

func (*RcScorecard) TableName() string {
	return "rc_scorecard"
}

func (e *RcScorecard) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcScorecard) GetId() interface{} {
	return e.Id
}
