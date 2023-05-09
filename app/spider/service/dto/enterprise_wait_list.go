package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type EnterpriseWaitListGetPageReq struct {
	dto.Pagination `search:"-"`

	Id             int64  `form:"id"  search:"type:exact;column:id;table:enterprise_wait_list" comment:"主键id"`
	EnterpriseName string `form:"enterpriseName"  search:"type:exact;column:enterprise_name;table:enterprise_wait_list" comment:"企业名称"`
	UscId          string `form:"uscId"  search:"type:exact;column:usc_id;table:enterprise_wait_list" comment:"社会统一信用代码"`
	Priority       int    `form:"priority"  search:"type:exact;column:priority;table:enterprise_wait_list" comment:"优先级"`
	QccUrl         string `form:"qccUrl"  search:"type:exact;column:qcc_url;table:enterprise_wait_list" comment:"qcc主体网址"`
	StatusCode     int    `form:"statusCode"  search:"type:exact;column:status_code;table:enterprise_wait_list" comment:"数据爬取状态码"`
	Source         string `form:"source"  search:"type:exact;column:source;table:enterprise_wait_list" comment:"来源备注"`
	EnterpriseWaitListPageOrder
}

type EnterpriseWaitListPageOrder struct {
	IdOrder         string `form:"idOrder"  search:"type:order;column:id;table:enterprise_wait_list"`
	PriorityOrder   string `form:"priorityOrder"  search:"type:order;column:priority;table:enterprise_wait_list"`
	StatusCodeOrder string `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_wait_list"`
}

func (m *EnterpriseWaitListGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseWaitListGetResp struct {
	Id             int64  `json:"id"`             // 主键id
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	UscId          string `json:"uscId"`          // 社会统一信用代码
	Priority       int    `json:"priority"`       // 优先级
	QccUrl         string `json:"qccUrl"`         // qcc主体网址
	StatusCode     int    `json:"statusCode"`     // 数据爬取状态码
	Source         string `json:"source"`         // 来源备注
	common.ControlBy
}

type EnterpriseWaitListWaitingGetPageResp struct {
	Id             int64  `json:"id"`             // 主键id
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	UscId          string `json:"uscId"`          // 社会统一信用代码
	Priority       int    `json:"priority"`       // 优先级
	Source         string `json:"source"`
	StatusCode     int    `json:"statusCode"`
	QccUrl         string `json:"qccUrl"`
}

func (s *EnterpriseWaitListWaitingGetPageResp) Generate(model *models.EnterpriseWaitList) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.EnterpriseName = model.EnterpriseName
	s.UscId = model.UscId
	s.Priority = model.Priority
	s.Source = model.Source
	s.StatusCode = model.StatusCode
	s.QccUrl = model.QccUrl
}

func (s *EnterpriseWaitListGetResp) Generate(model *models.EnterpriseWaitList) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.EnterpriseName = model.EnterpriseName
	s.UscId = model.UscId
	s.Priority = model.Priority
	s.QccUrl = model.QccUrl
	s.StatusCode = model.StatusCode
	s.Source = model.Source
	s.CreateBy = model.CreateBy
}

type EnterpriseWaitListInsertReq struct {
	Id             int64  `json:"-"`              // 主键id
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	UscId          string `json:"uscId"`          // 社会统一信用代码
	Priority       int    `json:"priority"`       // 优先级
	QccUrl         string `json:"qccUrl"`         // qcc主体网址
	StatusCode     int    `json:"statusCode"`     // 数据爬取状态码
	Source         string `json:"source"`         // 来源备注
	common.ControlBy
}

func (s *EnterpriseWaitListInsertReq) Generate(model *models.EnterpriseWaitList) {
	if s.Id == 0 {
		s.Id = utils.NewFlakeId()
	}
	model.Model = common.Model{Id: s.Id}
	model.EnterpriseName = s.EnterpriseName
	model.UscId = s.UscId
	model.Priority = s.Priority
	model.QccUrl = s.QccUrl
	model.StatusCode = s.StatusCode
	model.Source = s.Source
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseWaitListInsertReq) GetId() interface{} {
	return s.Id
}

type EnterpriseWaitListUpdateReq struct {
	Id             int64  `uri:"id"`              // 主键id
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	UscId          string `json:"uscId"`          // 社会统一信用代码
	Priority       int    `json:"priority"`       // 优先级
	QccUrl         string `json:"qccUrl"`         // qcc主体网址
	StatusCode     int    `json:"statusCode"`     // 数据爬取状态码
	Source         string `json:"source"`         // 来源备注
	common.ControlBy
}

func (s *EnterpriseWaitListUpdateReq) Generate(model *models.EnterpriseWaitList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.EnterpriseName != "" {
		model.EnterpriseName = s.EnterpriseName
	}
	if s.UscId != "" {
		model.UscId = s.UscId
	}
	if s.Priority != 0 {
		model.Priority = s.Priority
	}
	if s.QccUrl != "" {
		model.QccUrl = s.QccUrl
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	if s.Source != "" {
		model.Source = s.Source
	}
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseWaitListUpdateReq) GetId() interface{} {
	return s.Id
}

// EnterpriseWaitListGetReq 功能获取请求参数
type EnterpriseWaitListGetReq struct {
	Id int64 `uri:"id"`
}

func (s *EnterpriseWaitListGetReq) GetId() interface{} {
	return s.Id
}

// EnterpriseWaitListDeleteReq 功能删除请求参数
type EnterpriseWaitListDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseWaitListDeleteReq) GetId() interface{} {
	return s.Ids
}
