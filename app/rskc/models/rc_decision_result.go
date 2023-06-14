package models

import (
	"go-admin/common/models"
)

// RcDecisionResult wezoom decision engine returend result
type RcDecisionResult struct {
	models.Model
	ParamId      int64  `json:"paramId" gorm:"comment:rc_decision_param" xlsx:"rc_decision_param"`
	TaskId       string `json:"taskId" gorm:"comment:task_id" xlsx:"task_id"`
	FinalResult  string `json:"finalResult" gorm:"comment:决策建议结果(REFUSE:拒绝，PASS:通过)" xlsx:"决策建议结果(REFUSE:拒绝，PASS:通过)"`
	AphScore     string `json:"aphScore" gorm:"comment:APH分数" xlsx:"APH分数"`
	FxSwJxccClnx string `json:"fxSwJxccClnx" gorm:"comment:经营年限" xlsx:"经营年限"`
	LhQylx       int    `json:"lhQylx" gorm:"comment:1:生产型，2:贸易型" xlsx:"1:生产型，2:贸易型"`
	Msg          string `json:"msg" gorm:"comment:返回结果描述" xlsx:"返回结果描述"`
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
