package dto

import (
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RcDecisionResultGetPageReq struct {
	dto.Pagination `search:"-"`

	Id           int64   `form:"id"  search:"type:exact;column:id;table:rc_decision_result" comment:""`
	ParamId      int64   `form:"paramId"  search:"type:exact;column:param_id;table:rc_decision_result" comment:""`
	TaskId       string  `form:"taskId"  search:"type:exact;column:task_id;table:rc_decision_result" comment:"任务id"`
	FinalResult  string  `form:"finalResult"  search:"type:exact;column:final_result;table:rc_decision_result" comment:"决策建议结果(REFUSE:拒绝，PASS:通过)"`
	AphScore     float64 `form:"aphScore"  search:"type:exact;column:aph_score;table:rc_decision_result" comment:"APH分数"`
	FxSwJxccClnx string  `form:"fxSwJxccClnx"  search:"type:exact;column:fx_sw_jxcc_clnx;table:rc_decision_result" comment:"经营年限"`
	LhQylx       int     `form:"lhQylx"  search:"type:exact;column:lh_qylx;table:rc_decision_result" comment:"1:生产型，2:贸易型"`
	Msg          string  `form:"msg"  search:"type:exact;column:msg;table:rc_decision_result" comment:"返回结果描述"`
	StartTime    string  `form:"startTime" search:"type:gte;column:created_at;table:rc_decision_result" comment:"创建时间"`
	EndTime      string  `form:"endTime" search:"type:lte;column:created_at;table:rc_decision_result" comment:"创建时间"`
	RcDecisionResultPageOrder
}

type RcDecisionResultPageOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:rc_decision_result"`
	ParamId      string `form:"paramIdOrder"  search:"type:order;column:param_id;table:rc_decision_result"`
	TaskId       string `form:"taskIdOrder"  search:"type:order;column:task_id;table:rc_decision_result"`
	FinalResult  string `form:"finalResultOrder"  search:"type:order;column:final_result;table:rc_decision_result"`
	AphScore     string `form:"aphScoreOrder"  search:"type:order;column:aph_score;table:rc_decision_result"`
	FxSwJxccClnx string `form:"fxSwJxccClnxOrder"  search:"type:order;column:fx_sw_jxcc_clnx;table:rc_decision_result"`
	LhQylx       string `form:"lhQylxOrder"  search:"type:order;column:lh_qylx;table:rc_decision_result"`
	Msg          string `form:"msgOrder"  search:"type:order;column:msg;table:rc_decision_result"`
}

func (m *RcDecisionResultGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RcDecisionResultGetResp struct {
	Id           int64   `json:"id"`           //
	ParamId      int64   `json:"paramId"`      //
	TaskId       string  `json:"taskId"`       // 任务id
	FinalResult  string  `json:"finalResult"`  // 决策建议结果(REFUSE:拒绝，PASS:通过)
	AphScore     float64 `json:"aphScore"`     // APH分数
	FxSwJxccClnx string  `json:"fxSwJxccClnx"` // 经营年限
	LhQylx       int     `json:"lhQylx"`       // 1:生产型，2:贸易型
	Msg          string  `json:"msg"`          // 返回结果描述
	common.ControlBy
}

func (s *RcDecisionResultGetResp) Generate(model *models.RcDecisionResult) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ParamId = model.ParamId
	s.TaskId = model.TaskId
	s.FinalResult = model.FinalResult
	s.AphScore = model.AphScore
	s.FxSwJxccClnx = model.FxSwJxccClnx
	s.LhQylx = model.LhQylx
	s.Msg = model.Msg
	s.CreateBy = model.CreateBy
}

type RcDecisionResultInsertReq struct {
	Id           int64   `json:"-"`            //
	ParamId      int64   `json:"paramId"`      //
	TaskId       string  `json:"taskId"`       // 任务id
	FinalResult  string  `json:"finalResult"`  // 决策建议结果(REFUSE:拒绝，PASS:通过)
	AphScore     float64 `json:"aphScore"`     // APH分数
	FxSwJxccClnx string  `json:"fxSwJxccClnx"` // 经营年限
	LhQylx       int     `json:"lhQylx"`       // 1:生产型，2:贸易型
	Msg          string  `json:"msg"`          // 返回结果描述
	common.ControlBy
}

func (s *RcDecisionResultInsertReq) Generate(model *models.RcDecisionResult) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ParamId = s.ParamId
	model.TaskId = s.TaskId
	model.FinalResult = s.FinalResult
	model.AphScore = s.AphScore
	model.FxSwJxccClnx = s.FxSwJxccClnx
	model.LhQylx = s.LhQylx
	model.Msg = s.Msg
	model.CreateBy = s.CreateBy
}

func (s *RcDecisionResultInsertReq) GetId() interface{} {
	return s.Id
}

type RcDecisionResultUpdateReq struct {
	Id           int64   `uri:"id"`            //
	ParamId      int64   `json:"paramId"`      //
	TaskId       string  `json:"taskId"`       // 任务id
	FinalResult  string  `json:"finalResult"`  // 决策建议结果(REFUSE:拒绝，PASS:通过)
	AphScore     float64 `json:"aphScore"`     // APH分数
	FxSwJxccClnx string  `json:"fxSwJxccClnx"` // 经营年限
	LhQylx       int     `json:"lhQylx"`       // 1:生产型，2:贸易型
	Msg          string  `json:"msg"`          // 返回结果描述
	common.ControlBy
}

func (s *RcDecisionResultUpdateReq) Generate(model *models.RcDecisionResult) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ParamId = s.ParamId
	model.TaskId = s.TaskId
	model.FinalResult = s.FinalResult
	model.AphScore = s.AphScore
	model.FxSwJxccClnx = s.FxSwJxccClnx
	model.LhQylx = s.LhQylx
	model.Msg = s.Msg
	model.UpdateBy = s.UpdateBy
}

func (s *RcDecisionResultUpdateReq) GetId() interface{} {
	return s.Id
}

// RcDecisionResultGetReq 功能获取请求参数
type RcDecisionResultGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RcDecisionResultGetReq) GetId() interface{} {
	return s.Id
}

// RcDecisionResultDeleteReq 功能删除请求参数
type RcDecisionResultDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RcDecisionResultDeleteReq) GetId() interface{} {
	return s.Ids
}

type RcDecisionResultExport struct {
	Id           int64   `json:"id" gorm:"primaryKey;comment:Id" xlsx:"Id"`
	ParamId      int64   `json:"paramId" gorm:"comment:ParamId" xlsx:"ParamId"`
	TaskId       string  `json:"taskId" gorm:"comment:任务id" xlsx:"任务id"`
	FinalResult  string  `json:"finalResult" gorm:"comment:决策建议结果(REFUSE:拒绝，PASS:通过)" xlsx:"决策建议结果(REFUSE:拒绝，PASS:通过)"`
	AphScore     float64 `json:"aphScore" gorm:"comment:APH分数" xlsx:"APH分数"`
	FxSwJxccClnx string  `json:"fxSwJxccClnx" gorm:"comment:经营年限" xlsx:"经营年限"`
	LhQylx       int     `json:"lhQylx" gorm:"comment:1:生产型，2:贸易型" xlsx:"1:生产型，2:贸易型"`
	Msg          string  `json:"msg" gorm:"comment:返回结果描述" xlsx:"返回结果描述"`
}
