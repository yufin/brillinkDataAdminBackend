package models

import (
	"go-admin/common/models"
)

// RcDecisionResult
type RcDecisionResult struct {
	models.Model
	ParamId      int64   `json:"paramId" gorm:"comment:ParamId" xlsx:"ParamId"`
	TaskId       string  `json:"taskId" gorm:"comment:任务id" xlsx:"任务id"`
	FinalResult  string  `json:"final_result" gorm:"comment:决策建议结果(REFUSE:拒绝，PASS:通过)" xlsx:"决策建议结果(REFUSE:拒绝，PASS:通过)"`
	AphScore     float64 `json:"AHP_SCORE" gorm:"comment:APH分数" xlsx:"APH分数"`
	FxSwJxccClnx string  `json:"fx_sw_jxcc_clnx" gorm:"comment:经营年限" xlsx:"经营年限"`
	LhQylx       int     `json:"lh_qylx" gorm:"comment:1:生产型，2:贸易型" xlsx:"1:生产型，2:贸易型"`
	Msg          string  `json:"msg" gorm:"comment:返回结果描述" xlsx:"返回结果描述"`
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
