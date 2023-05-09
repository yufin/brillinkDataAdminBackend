package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type EnterpriseProductGetPageReq struct {
	dto.Pagination `search:"-"`

	UscId       string `form:"uscId"  search:"type:exact;column:usc_id;table:enterprise_product" comment:"社会统一信用代码"`
	ProductData string `form:"productData"  search:"type:exact;column:product_data;table:enterprise_product" comment:"json格式的产品分类"`
	StatusCode  int    `form:"statusCode"  search:"type:exact;column:status_code;table:enterprise_product" comment:"状态码"`
	EnterpriseProductPageOrder
}

type EnterpriseProductPageOrder struct {
	ProdId      int64  `form:"prodIdOrder"  search:"type:order;column:prod_id;table:enterprise_product"`
	UscId       string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise_product"`
	ProductData string `form:"productDataOrder"  search:"type:order;column:product_data;table:enterprise_product"`
	StatusCode  int    `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_product"`
}

func (m *EnterpriseProductGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseProductGetResp struct {
	ProdId      int64  `json:"prodId"`      // 主键
	UscId       string `json:"uscId"`       // 社会统一信用代码
	ProductData string `json:"productData"` // json格式的产品分类
	StatusCode  int    `json:"statusCode"`  // 状态码
	common.ControlBy
}

func (s *EnterpriseProductGetResp) Generate(model *models.EnterpriseProduct) {
	s.ProdId = model.ProdId
	s.UscId = model.UscId
	s.ProductData = model.ProductData
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type EnterpriseProductInsertReq struct {
	ProdId      int64  `json:"-"`           // 主键
	UscId       string `json:"uscId"`       // 社会统一信用代码
	ProductData string `json:"productData"` // json格式的产品分类
	StatusCode  int    `json:"statusCode"`  // 状态码
	common.ControlBy
}

func (s *EnterpriseProductInsertReq) Generate(model *models.EnterpriseProduct) {
	if s.ProdId == 0 {
		s.ProdId = utils.NewFlakeId()
	}
	model.ProdId = s.ProdId
	model.UscId = s.UscId
	model.ProductData = s.ProductData
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseProductInsertReq) GetId() interface{} {
	return s.ProdId
}

type EnterpriseProductUpdateReq struct {
	ProdId      int64  `uri:"prodId"`       // 主键
	UscId       string `json:"uscId"`       // 社会统一信用代码
	ProductData string `json:"productData"` // json格式的产品分类
	StatusCode  int    `json:"statusCode"`  // 状态码
	common.ControlBy
}

func (s *EnterpriseProductUpdateReq) Generate(model *models.EnterpriseProduct) {
	model.ProdId = s.ProdId
	model.UscId = s.UscId
	model.ProductData = s.ProductData
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseProductUpdateReq) GetId() interface{} {
	return s.ProdId
}

// EnterpriseProductGetReq 功能获取请求参数
type EnterpriseProductGetReq struct {
	ProdId int64 `uri:"prodId"`
}

func (s *EnterpriseProductGetReq) GetId() interface{} {
	return s.ProdId
}

// EnterpriseProductDeleteReq 功能删除请求参数
type EnterpriseProductDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseProductDeleteReq) GetId() interface{} {
	return s.Ids
}
