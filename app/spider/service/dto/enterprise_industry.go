package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EnterpriseIndustryGetPageReq struct {
	dto.Pagination `search:"-"`

	EnterpriseId int64  `form:"enterpriseId"  search:"type:;column:enterprise_id;table:enterprise_industry" comment:"外键-企业id"`
	IndustryData string `form:"industryData"  search:"type:;column:industry_data;table:enterprise_industry" comment:"json格式的行业分类"`
	StatusCode   int    `form:"statusCode"  search:"type:;column:status_code;table:enterprise_industry" comment:"状态标识码"`
	EnterpriseIndustryPageOrder
}

type EnterpriseIndustryPageOrder struct {
	IndId        int64  `form:"indIdOrder"  search:"type:order;column:ind_id;table:enterprise_industry"`
	EnterpriseId int64  `form:"enterpriseIdOrder"  search:"type:order;column:enterprise_id;table:enterprise_industry"`
	IndustryData string `form:"industryDataOrder"  search:"type:order;column:industry_data;table:enterprise_industry"`
	StatusCode   int    `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_industry"`
}

func (m *EnterpriseIndustryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseIndustryGetResp struct {
	IndId        int64  `json:"indId"`        // 主键
	EnterpriseId int64  `json:"enterpriseId"` // 外键-企业id
	IndustryData string `json:"industryData"` // json格式的行业分类
	StatusCode   int    `json:"statusCode"`   // 状态标识码
	common.ControlBy
}

func (s *EnterpriseIndustryGetResp) Generate(model *models.EnterpriseIndustry) {
	s.IndId = model.IndId
	s.EnterpriseId = model.EnterpriseId
	s.IndustryData = model.IndustryData
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type EnterpriseIndustryInsertReq struct {
	IndId        int64  `json:"-"`            // 主键
	EnterpriseId int64  `json:"enterpriseId"` // 外键-企业id
	IndustryData string `json:"industryData"` // json格式的行业分类
	StatusCode   int    `json:"statusCode"`   // 状态标识码
	common.ControlBy
}

func (s *EnterpriseIndustryInsertReq) Generate(model *models.EnterpriseIndustry) {
	model.IndId = s.IndId
	model.EnterpriseId = s.EnterpriseId
	model.IndustryData = s.IndustryData
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseIndustryInsertReq) GetId() interface{} {
	return s.IndId
}

type EnterpriseIndustryUpdateReq struct {
	IndId        int64  `uri:"indId"`         // 主键
	EnterpriseId int64  `json:"enterpriseId"` // 外键-企业id
	IndustryData string `json:"industryData"` // json格式的行业分类
	StatusCode   int    `json:"statusCode"`   // 状态标识码
	common.ControlBy
}

func (s *EnterpriseIndustryUpdateReq) Generate(model *models.EnterpriseIndustry) {
	model.IndId = s.IndId
	model.EnterpriseId = s.EnterpriseId
	model.IndustryData = s.IndustryData
	model.StatusCode = s.StatusCode
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
