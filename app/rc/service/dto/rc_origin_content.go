package dto

import (
	"go-admin/app/rc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type RcOriginContentGetPageReq struct {
	dto.Pagination `search:"-"`
	UscId          string `form:"uscId"  search:"type:exact;column:usc_id;table:rc_origin_content" comment:"统一社会信用代码"`
	EnterpriseName string `form:"enterpriseName"  search:"type:exact;column:enterprise_name;table:rc_origin_content" comment:"企业名称"`
	YearMonth      string `form:"yearMonth"  search:"type:exact;column:year_month;table:rc_origin_content" comment:"数据更新年月"`
	StatusCode     int    `form:"statusCode"  search:"type:exact;column:status_code;table:rc_origin_content" comment:"状态码"`
	RcOriginContentPageOrder
}

type RcOriginContentPageOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:rc_origin_content"`
	UscId      string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:rc_origin_content"`
	YearMonth  string `form:"yearMonthOrder"  search:"type:order;column:year_month;table:rc_origin_content"`
	StatusCode string `form:"statusCodeOrder"  search:"type:order;column:status_code;table:rc_origin_content"`
}

func (m *RcOriginContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RcOriginContentGetResp struct {
	Id             int64  `json:"id"`             // 主键
	UscId          string `json:"uscId"`          // 统一社会信用代码
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	YearMonth      string `json:"yearMonth"`      // 数据更新年月
	Content        string `json:"content"`        // 原始JSON STRING数据
	StatusCode     int    `json:"statusCode"`     // 状态码
	common.ControlBy
}

func (s *RcOriginContentGetResp) Generate(model *models.RcOriginContent) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.UscId = model.UscId
	s.EnterpriseName = model.EnterpriseName
	s.YearMonth = model.YearMonth
	s.Content = model.Content
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type RcOriginContentInsertReq struct {
	Id             int64  `json:"-"`              // 主键
	UscId          string `json:"uscId"`          // 统一社会信用代码
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	YearMonth      string `json:"yearMonth"`      // 数据更新年月
	Content        string `json:"content"`        // 原始JSON STRING数据
	StatusCode     int    `json:"statusCode"`     // 状态码
	common.ControlBy
}

func (s *RcOriginContentInsertReq) Generate(model *models.RcOriginContent) {
	if s.Id == 0 {
		//model.Model = common.Model{Id: s.Id}
		s.Id = utils.NewFlakeId()
	}
	model.Model = common.Model{Id: s.Id}
	model.UscId = s.UscId
	model.EnterpriseName = s.EnterpriseName
	model.YearMonth = s.YearMonth
	model.Content = s.Content
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *RcOriginContentInsertReq) GetId() interface{} {
	return s.Id
}

type RcOriginContentUpdateReq struct {
	Id             int64  `uri:"id"`              // 主键
	UscId          string `json:"uscId"`          // 统一社会信用代码
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	YearMonth      string `json:"yearMonth"`      // 数据更新年月
	Content        string `json:"content"`        // 原始JSON STRING数据
	StatusCode     int    `json:"statusCode"`     // 状态码
	common.ControlBy
}

func (s *RcOriginContentUpdateReq) Generate(model *models.RcOriginContent) {
	model.Id = s.Id

	if s.UscId != "" {
		model.UscId = s.UscId
	}
	if s.EnterpriseName != "" {
		model.EnterpriseName = s.EnterpriseName
	}
	if s.YearMonth != "" {
		model.YearMonth = s.YearMonth
	}
	if s.Content != "" {
		model.Content = s.Content
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	model.UpdateBy = s.UpdateBy
}

func (s *RcOriginContentUpdateReq) GetId() interface{} {
	return s.Id
}

// RcOriginContentGetReq 功能获取请求参数
type RcOriginContentGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RcOriginContentGetReq) GetId() interface{} {
	return s.Id
}

// RcOriginContentDeleteReq 功能删除请求参数
type RcOriginContentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RcOriginContentDeleteReq) GetId() interface{} {
	return s.Ids
}
