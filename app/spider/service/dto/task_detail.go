package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type TaskDetailGetPageReq struct {
	dto.Pagination `search:"-"`

	UscId      string `form:"uscId"  search:"type:exact;column:usc_id;table:task_detail" comment:"企业统一信用代码"`
	StatusCode int    `form:"statusCode"  search:"type:exact;column:status_code;table:task_detail" comment:"状态码(1.未完成, 2.完成)"`
	Topic      string `form:"topic"  search:"type:exact;column:topic;table:task_detail" comment:"主题"`
	Priority   int    `form:"priority"  search:"type:exact;column:priority;table:task_detail" comment:"优先级"`
	Comment    string `form:"comment"  search:"type:exact;column:comment;table:task_detail" comment:"备注"`
	TaskDetailPageOrder
}

type TaskDetailPageOrder struct {
	Id         int64  `form:"idOrder"  search:"type:order;column:id;table:task_detail"`
	UscId      string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:task_detail"`
	StatusCode int    `form:"statusCodeOrder"  search:"type:order;column:status_code;table:task_detail"`
	Topic      string `form:"topicOrder"  search:"type:order;column:topic;table:task_detail"`
	Priority   int    `form:"priorityOrder"  search:"type:order;column:priority;table:task_detail"`
	Comment    string `form:"commentOrder"  search:"type:order;column:comment;table:task_detail"`
}

func (m *TaskDetailGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TaskDetailGetResp struct {
	Id         int64  `json:"id"`         // 主键
	UscId      string `json:"uscId"`      // 企业统一信用代码
	StatusCode int    `json:"statusCode"` // 状态码(1.未完成, 2.完成)
	Topic      string `json:"topic"`      // 主题
	Priority   int    `json:"priority"`   // 优先级
	Comment    string `json:"comment"`    // 备注
	common.ControlBy
}

func (s *TaskDetailGetResp) Generate(model *models.TaskDetail) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.UscId = model.UscId
	s.StatusCode = model.StatusCode
	s.Topic = model.Topic
	s.Priority = model.Priority
	s.Comment = model.Comment
	s.CreateBy = model.CreateBy
}

type TaskDetailInsertReq struct {
	Id         int64  `json:"-"`          // 主键
	UscId      string `json:"uscId"`      // 企业统一信用代码
	StatusCode int    `json:"statusCode"` // 状态码(1.未完成, 2.完成)
	Topic      string `json:"topic"`      // 主题
	Priority   int    `json:"priority"`   // 优先级
	Comment    string `json:"comment"`    // 备注
	common.ControlBy
}

func (s *TaskDetailInsertReq) Generate(model *models.TaskDetail) {
	if s.Id == 0 {
		model.Model = common.Model{Id: utils.NewFlakeId()}
	}
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.Topic = s.Topic
	model.Priority = s.Priority
	model.Comment = s.Comment
	model.CreateBy = s.CreateBy
}

func (s *TaskDetailInsertReq) GetId() interface{} {
	return s.Id
}

type TaskDetailUpdateReq struct {
	Id         int64  `uri:"id"`          // 主键
	UscId      string `json:"uscId"`      // 企业统一信用代码
	StatusCode int    `json:"statusCode"` // 状态码(1.未完成, 2.完成)
	Topic      string `json:"topic"`      // 主题
	Priority   int    `json:"priority"`   // 优先级
	Comment    string `json:"comment"`    // 备注
	common.ControlBy
}

func (s *TaskDetailUpdateReq) Generate(model *models.TaskDetail) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.Topic = s.Topic
	model.Priority = s.Priority
	model.Comment = s.Comment
	model.UpdateBy = s.UpdateBy
}

func (s *TaskDetailUpdateReq) GetId() interface{} {
	return s.Id
}

// TaskDetailGetReq 功能获取请求参数
type TaskDetailGetReq struct {
	Id int64 `uri:"id"`
}

func (s *TaskDetailGetReq) GetId() interface{} {
	return s.Id
}

// TaskDetailDeleteReq 功能删除请求参数
type TaskDetailDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TaskDetailDeleteReq) GetId() interface{} {
	return s.Ids
}
