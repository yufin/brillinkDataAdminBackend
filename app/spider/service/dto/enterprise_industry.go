package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type EnterpriseIndustryGetPageReq struct {
	dto.Pagination `search:"-"`

	UscId        string `form:"uscId"  search:"type:exact;column:usc_id;table:enterprise_industry" comment:"社会统一信用代码"`
	IndustryData string `form:"industryData"  search:"type:exact;column:industry_data;table:enterprise_industry" comment:"json格式的行业分类"`
	StatusCode   int    `form:"statusCode"  search:"type:exact;column:status_code;table:enterprise_industry" comment:"状态标识码"`
	EnterpriseIndustryPageOrder
}

type EnterpriseIndustryPageOrder struct {
	IndId        int64  `form:"indIdOrder"  search:"type:order;column:ind_id;table:enterprise_industry"`
	UscId        string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise_industry"`
	IndustryData string `form:"industryDataOrder"  search:"type:order;column:industry_data;table:enterprise_industry"`
	StatusCode   int    `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_industry"`
}

func (m *EnterpriseIndustryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseIndustryGetResp struct {
	IndId        int64  `json:"indId"`        // 主键
	UscId        string `json:"uscId"`        // 社会统一信用代码
	IndustryData string `json:"industryData"` // json格式的行业分类
	StatusCode   int    `json:"statusCode"`   // 状态标识码
	common.ControlBy
}

func (s *EnterpriseIndustryGetResp) Generate(model *models.EnterpriseIndustry) {
	s.IndId = model.IndId
	s.UscId = model.UscId
	s.IndustryData = model.IndustryData
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type EnterpriseIndustryInsertReq struct {
	IndId        int64  `json:"-"`            // 主键
	UscId        string `json:"uscId"`        // 社会统一信用代码
	IndustryData string `json:"industryData"` // json格式的行业分类
	StatusCode   int    `json:"statusCode"`   // 状态标识码
	common.ControlBy
}

func (s *EnterpriseIndustryInsertReq) Generate(model *models.EnterpriseIndustry) {
	if s.IndId == 0 {
		s.IndId = utils.NewFlakeId()
	}
	model.IndId = s.IndId
	if s.UscId != "" {
		model.UscId = s.UscId
	}
	if s.IndustryData != "" {
		model.IndustryData = s.IndustryData
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseIndustryInsertReq) GetId() interface{} {
	return s.IndId
}

type EnterpriseIndustryUpdateReq struct {
	IndId        int64  `uri:"indId"`         // 主键
	UscId        string `json:"uscId"`        // 社会统一信用代码
	IndustryData string `json:"industryData"` // json格式的行业分类
	StatusCode   int    `json:"statusCode"`   // 状态标识码
	common.ControlBy
}

func (s *EnterpriseIndustryUpdateReq) Generate(model *models.EnterpriseIndustry) {
	model.IndId = s.IndId
	if s.UscId != "" {
		model.UscId = s.UscId
	}
	if s.IndustryData != "" {
		model.IndustryData = s.IndustryData
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseIndustryUpdateReq) GetId() interface{} {
	return s.IndId
}

// EnterpriseIndustryGetReq 功能获取请求参数
type EnterpriseIndustryGetReq struct {
	IndId int64 `uri:"indId"`
}

func (s *EnterpriseIndustryGetReq) GetId() interface{} {
	return s.IndId
}

// EnterpriseIndustryDeleteReq 功能删除请求参数
type EnterpriseIndustryDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseIndustryDeleteReq) GetId() interface{} {
	return s.Ids
}
