package models

import (
	"github.com/shopspring/decimal"
	"go-admin/common/models"
)

// RcDecisionResult wezoom decision engine returend result
type RcDecisionResult struct {
	models.Model
	DepId        int64               `json:"depId" gorm:"comment:rc_decision_param" xlsx:"rc_decision_param"`
	TaskId       string              `json:"taskId" gorm:"comment:task_id" xlsx:"task_id"`
	FinalResult  string              `json:"finalResult" gorm:"comment:决策建议结果(REFUSE:拒绝，PASS:通过)" xlsx:"决策建议结果(REFUSE:拒绝，PASS:通过)"`
	AhpScore     decimal.NullDecimal `json:"ahpScore" gorm:"comment:AHP分数" xlsx:"AHP分数"`
	FxSwJxccClnx string              `json:"fxSwJxccClnx" gorm:"comment:经营年限" xlsx:"经营年限"`
	LhQylx       int                 `json:"lhQylx" gorm:"comment:1:生产型，2:贸易型" xlsx:"1:生产型，2:贸易型"`
	Msg          string              `json:"msg" gorm:"comment:返回结果描述" xlsx:"返回结果描述"`
	models.ModelTime
	models.ControlBy
}

func (*RcDecisionResult) TableName() string {
	return "rc_decision_result"
}

func (e *RcDecisionResult) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcDecisionResult) GetId() interface{} {
	return e.Id
}
