package dto

import (
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type RskcTradesDetailGetPageReq struct {
	dto.Pagination `search:"-"`
	ContentId      int64  `form:"contentId"  search:"type:exact;column:content_id;table:rskc_trades_detail" comment:"外键"`
	EnterpriseName string `form:"enterpriseName"  search:"type:exact;column:enterprise_name;table:rskc_trades_detail" comment:"企业名称"`
	CommodityRatio string `form:"commodityRatio"  search:"type:exact;column:commodity_ratio;table:rskc_trades_detail" comment:"货物占比"`
	CommodityName  string `form:"commodityName"  search:"type:exact;column:commodity_name;table:rskc_trades_detail" comment:"货物种类名称"`
	RatioAmountTax string `form:"ratioAmountTax"  search:"type:exact;column:ratio_amount_tax;table:rskc_trades_detail" comment:"税额占比"`
	SumAmountTax   string `form:"sumAmountTax"  search:"type:exact;column:sum_amount_tax;table:rskc_trades_detail" comment:"税总额"`
	DetailType     int    `form:"detailType"  search:"type:exact;column:detail_type;table:rskc_trades_detail" comment:"1:customer_12,2:customer_24,3:supplier_12,4:supplier_24"`
	TagIndustry    string `form:"tagIndustry"  search:"type:exact;column:tag_industry;table:rskc_trades_detail" comment:"行业标签"`
	TagAuthorized  string `form:"tagAuthorized"  search:"type:exact;column:tag_authorized;table:rskc_trades_detail" comment:"认证标签"`
	TagProduct     string `form:"tagProduct"  search:"type:exact;column:tag_product;table:rskc_trades_detail" comment:"产品标签"`
	TagList        string `form:"tagList"  search:"type:exact;column:tag_list;table:rskc_trades_detail" comment:"榜单标签"`
	EnterpriseInfo string `form:"enterpriseInfo"  search:"type:exact;column:enterprise_info;table:rskc_trades_detail" comment:"企业信息"`
	StatusCode     int    `form:"statusCode"  search:"type:exact;column:status_code;table:rskc_trades_detail" comment:"状态码: 1.待确认企业数据已采集，2.待采集，已经同步至waitList, 3.采集完成, 4.匹配并录入完成"`
	RskcTradesDetailPageOrder
}

type RskcTradesDetailPageOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:rskc_trades_detail"`
	ContentId      int64  `form:"contentIdOrder"  search:"type:order;column:content_id;table:rskc_trades_detail"`
	EnterpriseName string `form:"enterpriseNameOrder"  search:"type:order;column:enterprise_name;table:rskc_trades_detail"`
	CommodityRatio string `form:"commodityRatioOrder"  search:"type:order;column:commodity_ratio;table:rskc_trades_detail"`
	CommodityName  string `form:"commodityNameOrder"  search:"type:order;column:commodity_name;table:rskc_trades_detail"`
	RatioAmountTax string `form:"ratioAmountTaxOrder"  search:"type:order;column:ratio_amount_tax;table:rskc_trades_detail"`
	SumAmountTax   string `form:"sumAmountTaxOrder"  search:"type:order;column:sum_amount_tax;table:rskc_trades_detail"`
	DetailType     string `form:"detailTypeOrder"  search:"type:order;column:detail_type;table:rskc_trades_detail"`
	TagIndustry    string `form:"tagIndustryOrder"  search:"type:order;column:tag_industry;table:rskc_trades_detail"`
	TagAuthorized  string `form:"tagAuthorizedOrder"  search:"type:order;column:tag_authorized;table:rskc_trades_detail"`
	TagProduct     string `form:"tagProductOrder"  search:"type:order;column:tag_product;table:rskc_trades_detail"`
	TagList        string `form:"tagListOrder"  search:"type:order;column:tag_list;table:rskc_trades_detail"`
	EnterpriseInfo string `form:"enterpriseInfoOrder"  search:"type:order;column:enterprise_info;table:rskc_trades_detail"`
	StatusCode     string `form:"statusCodeOrder"  search:"type:order;column:status_code;table:rskc_trades_detail"`
}

func (m *RskcTradesDetailGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RskcTradesDetailGetResp struct {
	Id             int64  `json:"id"`             // 主键
	ContentId      int64  `json:"contentId"`      // 外键
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	CommodityRatio string `json:"commodityRatio"` // 货物占比
	CommodityName  string `json:"commodityName"`  // 货物种类名称
	RatioAmountTax string `json:"ratioAmountTax"` // 税额占比
	SumAmountTax   string `json:"sumAmountTax"`   // 税总额
	DetailType     int    `json:"detailType"`     // 1:customer_12,2:customer_24,3:supplier_12,4:supplier_24
	TagIndustry    string `json:"tagIndustry"`    // 行业标签
	TagAuthorized  string `json:"tagAuthorized"`  // 认证标签
	TagProduct     string `json:"tagProduct"`     // 产品标签
	TagList        string `json:"tagList"`        // 榜单标签
	EnterpriseInfo string `json:"enterpriseInfo"` // 企业信息
	StatusCode     int    `json:"statusCode"`     // 状态码
	common.ControlBy
}

func (s *RskcTradesDetailGetResp) Generate(model *models.RskcTradesDetail) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ContentId = model.ContentId
	s.EnterpriseName = model.EnterpriseName
	s.CommodityRatio = model.CommodityRatio
	s.CommodityName = model.CommodityName
	s.RatioAmountTax = model.RatioAmountTax
	s.SumAmountTax = model.SumAmountTax
	s.DetailType = model.DetailType
	s.TagIndustry = model.TagIndustry
	s.TagAuthorized = model.TagAuthorized
	s.TagProduct = model.TagProduct
	s.TagList = model.TagList
	s.EnterpriseInfo = model.EnterpriseInfo
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type RskcTradesDetailInsertReq struct {
	Id             int64  `json:"id"`             // 主键
	ContentId      int64  `json:"contentId"`      // 外键
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	CommodityRatio string `json:"commodityRatio"` // 货物占比
	CommodityName  string `json:"commodityName"`  // 货物种类名称
	RatioAmountTax string `json:"ratioAmountTax"` // 税额占比
	SumAmountTax   string `json:"sumAmountTax"`   // 税总额
	DetailType     int    `json:"detailType"`     // 1:customer_12,2:customer_24,3:supplier_12,4:supplier_24
	TagIndustry    string `json:"tagIndustry"`    // 行业标签
	TagAuthorized  string `json:"tagAuthorized"`  // 认证标签
	TagProduct     string `json:"tagProduct"`     // 产品标签
	TagList        string `json:"tagList"`        // 榜单标签
	EnterpriseInfo string `json:"enterpriseInfo"` // 企业信息
	StatusCode     int    `json:"statusCode"`     // 状态码
	common.ControlBy
}

func (s *RskcTradesDetailInsertReq) Generate(model *models.RskcTradesDetail) {
	if s.Id == 0 {
		//model.Model = common.Model{Id: s.Id}
		s.Id = utils.NewFlakeId()
	}
	model.Model = common.Model{Id: utils.NewFlakeId()}
	model.ContentId = s.ContentId
	model.EnterpriseName = s.EnterpriseName
	model.CommodityRatio = s.CommodityRatio
	model.CommodityName = s.CommodityName
	model.RatioAmountTax = s.RatioAmountTax
	model.SumAmountTax = s.SumAmountTax
	model.DetailType = s.DetailType
	model.TagIndustry = s.TagIndustry
	model.TagAuthorized = s.TagAuthorized
	model.TagProduct = s.TagProduct
	model.TagList = s.TagList
	model.EnterpriseInfo = s.EnterpriseInfo
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *RskcTradesDetailInsertReq) GetId() interface{} {
	return s.Id
}

type RskcTradesDetailUpdateReq struct {
	Id             int64  `uri:"id"`              // 主键
	ContentId      int64  `json:"contentId"`      // 外键
	EnterpriseName string `json:"enterpriseName"` // 企业名称
	CommodityRatio string `json:"commodityRatio"` // 货物占比
	CommodityName  string `json:"commodityName"`  // 货物种类名称
	RatioAmountTax string `json:"ratioAmountTax"` // 税额占比
	SumAmountTax   string `json:"sumAmountTax"`   // 税总额
	DetailType     int    `json:"detailType"`     // 1:customer_12,2:customer_24,3:supplier_12,4:supplier_24
	TagIndustry    string `json:"tagIndustry"`    // 行业标签
	TagAuthorized  string `json:"tagAuthorized"`  // 认证标签
	TagProduct     string `json:"tagProduct"`     // 产品标签
	TagList        string `json:"tagList"`        // 榜单标签
	EnterpriseInfo string `json:"enterpriseInfo"` // 企业信息
	StatusCode     int    `json:"statusCode"`     // 状态码: 1.待确认企业数据已采集，2.待采集，已经同步至waitList, 3.采集完成, 4.匹配并录入完成
	common.ControlBy
}

func (s *RskcTradesDetailUpdateReq) Generate(model *models.RskcTradesDetail) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.ContentId != 0 {
		model.ContentId = s.ContentId
	}
	if s.EnterpriseName != "" {
		model.EnterpriseName = s.EnterpriseName
	}
	if s.CommodityRatio != "" {
		model.CommodityRatio = s.CommodityRatio
	}
	if s.CommodityName != "" {
		model.CommodityName = s.CommodityName
	}
	if s.RatioAmountTax != "" {
		model.RatioAmountTax = s.RatioAmountTax
	}
	if s.SumAmountTax != "" {
		model.SumAmountTax = s.SumAmountTax
	}
	if s.DetailType != 0 {
		model.DetailType = s.DetailType
	}
	if s.TagIndustry != "" {
		model.TagIndustry = s.TagIndustry
	}
	if s.TagAuthorized != "" {
		model.TagAuthorized = s.TagAuthorized
	}
	if s.TagProduct != "" {
		model.TagProduct = s.TagProduct
	}
	if s.TagList != "" {
		model.TagList = s.TagList
	}
	if s.EnterpriseInfo != "" {
		model.EnterpriseInfo = s.EnterpriseInfo
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	model.UpdateBy = s.UpdateBy
}

func (s *RskcTradesDetailUpdateReq) GetId() interface{} {
	return s.Id
}

// RskcTradesDetailGetReq 功能获取请求参数
type RskcTradesDetailGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RskcTradesDetailGetReq) GetId() interface{} {
	return s.Id
}

// RskcTradesDetailDeleteReq 功能删除请求参数
type RskcTradesDetailDeleteReq struct {
	Ids []int64 `json:"ids"`
}

func (s *RskcTradesDetailDeleteReq) GetId() interface{} {
	return s.Ids
}
