package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EnterpriseGetPageReq struct {
	dto.Pagination `search:"-"`

	UscId      string `form:"uscId"  search:"type:;column:usc_id;table:enterprise" comment:"统一社会信用代码"`
	StatusCode int    `form:"statusCode"  search:"type:;column:status_code;table:enterprise" comment:"状态标识码"`
	EnterprisePageOrder
}

type EnterprisePageOrder struct {
	Id         int64  `form:"idOrder"  search:"type:order;column:id;table:enterprise"`
	UscId      string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise"`
	StatusCode int    `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise"`
}

func (m *EnterpriseGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseGetResp struct {
	Id         int64  `json:"id"`         // 主键
	UscId      string `json:"uscId"`      // 统一社会信用代码
	StatusCode int    `json:"statusCode"` // 状态标识码
	common.ControlBy
}

func (s *EnterpriseGetResp) Generate(model *models.Enterprise) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.UscId = model.UscId
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type EnterpriseInsertReq struct {
	Id         int64  `json:"-"`          // 主键
	UscId      string `json:"uscId"`      // 统一社会信用代码
	StatusCode int    `json:"statusCode"` // 状态标识码
	common.ControlBy
}

func (s *EnterpriseInsertReq) Generate(model *models.Enterprise) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseInsertReq) GetId() interface{} {
	return s.Id
}

type EnterpriseUpdateReq struct {
	Id         int64  `uri:"id"`          // 主键
	UscId      string `json:"uscId"`      // 统一社会信用代码
	StatusCode int    `json:"statusCode"` // 状态标识码
	common.ControlBy
}

func (s *EnterpriseUpdateReq) Generate(model *models.Enterprise) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseUpdateReq) GetId() interface{} {
	return s.Id
}

// EnterpriseGetReq 功能获取请求参数
type EnterpriseGetReq struct {
	Id int64 `uri:"id"`
}

func (s *EnterpriseGetReq) GetId() interface{} {
	return s.Id
}

// EnterpriseDeleteReq 功能删除请求参数
type EnterpriseDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseDeleteReq) GetId() interface{} {
	return s.Ids
}
