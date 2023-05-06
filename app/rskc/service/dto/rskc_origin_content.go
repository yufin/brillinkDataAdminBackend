package dto

import (
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RskcOriginContentGetPageReq struct {
	dto.Pagination `search:"-"`

	ContentId  string `form:"contentId"  search:"type:exact;column:content_id;table:rskc_origin_content" comment:"uuid4"`
	UscId      string `form:"uscId"  search:"type:exact;column:usc_id;table:rskc_origin_content" comment:"统一社会信用代码"`
	YearMonth  string `form:"yearMonth"  search:"type:exact;column:year_month;table:rskc_origin_content" comment:"数据更新年月"`
	StatusCode int    `form:"statusCode"  search:"type:exact;column:status_code;table:rskc_origin_content" comment:"状态码"`
	RskcOriginContentPageOrder
}

type RskcOriginContentPageOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:rskc_origin_content"`
	ContentId  string `form:"contentIdOrder"  search:"type:order;column:content_id;table:rskc_origin_content"`
	UscId      string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:rskc_origin_content"`
	YearMonth  string `form:"yearMonthOrder"  search:"type:order;column:year_month;table:rskc_origin_content"`
	StatusCode string `form:"statusCodeOrder"  search:"type:order;column:status_code;table:rskc_origin_content"`
}

func (m *RskcOriginContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RskcOriginContentGetResp struct {
	Id         int64  `json:"id"`         // 主键
	ContentId  string `json:"contentId"`  // uuid4
	UscId      string `json:"uscId"`      // 统一社会信用代码
	YearMonth  string `json:"yearMonth"`  // 数据更新年月
	Content    string `json:"content"`    // 原始JSON STRING数据
	StatusCode int    `json:"statusCode"` // 状态码
	common.ControlBy
}

func (s *RskcOriginContentGetResp) Generate(model *models.RskcOriginContent) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ContentId = model.ContentId
	s.UscId = model.UscId
	s.YearMonth = model.YearMonth
	s.Content = model.Content
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type RskcOriginContentInsertReq struct {
	Id         int64  `json:"-"`          // 主键
	ContentId  string `json:"contentId"`  // uuid4
	UscId      string `json:"uscId"`      // 统一社会信用代码
	YearMonth  string `json:"yearMonth"`  // 数据更新年月
	Content    string `json:"content"`    // 原始JSON STRING数据
	StatusCode int    `json:"statusCode"` // 状态码
	common.ControlBy
}

func (s *RskcOriginContentInsertReq) Generate(model *models.RskcOriginContent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ContentId = s.ContentId
	model.UscId = s.UscId
	model.YearMonth = s.YearMonth
	model.Content = s.Content
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *RskcOriginContentInsertReq) GetId() interface{} {
	return s.Id
}

type RskcOriginContentUpdateReq struct {
	Id         int64  `uri:"id"`          // 主键
	ContentId  string `json:"contentId"`  // uuid4
	UscId      string `json:"uscId"`      // 统一社会信用代码
	YearMonth  string `json:"yearMonth"`  // 数据更新年月
	Content    string `json:"content"`    // 原始JSON STRING数据
	StatusCode int    `json:"statusCode"` // 状态码
	common.ControlBy
}

func (s *RskcOriginContentUpdateReq) Generate(model *models.RskcOriginContent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.ContentId != "" {
		model.ContentId = s.ContentId
	}
	if s.UscId != "" {
		model.UscId = s.UscId
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

func (s *RskcOriginContentUpdateReq) GetId() interface{} {
	return s.Id
}

// RskcOriginContentGetReq 功能获取请求参数
type RskcOriginContentGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RskcOriginContentGetReq) GetId() interface{} {
	return s.Id
}

// RskcOriginContentDeleteReq 功能删除请求参数
type RskcOriginContentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RskcOriginContentDeleteReq) GetId() interface{} {
	return s.Ids
}
