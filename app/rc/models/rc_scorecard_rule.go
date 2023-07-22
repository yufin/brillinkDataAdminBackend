package models

import "go-admin/common/models"

type RcScorecardRule struct {
	models.Model
	ScorecardId int64   `json:"scorecardId" gorm:"type:bigint;comment:;default:null;"`
	Field       string  `json:"field" gorm:"type:varchar(255);comment:;default:null;"`
	AggField    string  `json:"aggField" gorm:"type:varchar(255);comment:;default:null;"`
	Pattern     string  `json:"pattern" gorm:"type:varchar(255);comment:;default:null;"`
	Expr        string  `json:"expr" gorm:"type:varchar(512);comment:;default:null;"`
	Score       float64 `json:"score" gorm:"type:decimal(32, 10);comment:;default:null;"`
	Level       int     `json:"level" gorm:"type:int;comment:;default:null;"`
	Comment     string  `json:"comment" gorm:"type:text;comment:;default:null;"`
	models.ModelTime
	models.ControlBy
}

func (*RcScorecardRule) TableName() string {
	return "rc_scorecard_rule"
}

func (e *RcScorecardRule) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcScorecardRule) GetId() interface{} {
	return e.Id
}
