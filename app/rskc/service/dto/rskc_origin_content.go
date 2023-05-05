package dto

import (
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OriginContentInsertReq struct {
	ContentId         string `json:"contentId"`
	UscId             string `json:"uscId"`
	YearMonth         string `json:"yearMonth"`
	OriginJsonContent string `json:"originJsonContent"`
	StatusCode        int    `json:"statusCode"`
	common.ControlBy
}

func (s *OriginContentInsertReq) Generate(model *models.OriginContent) {
	model.ContentId = s.ContentId
	model.UscId = s.UscId
	model.YearMonth = s.YearMonth
	model.OriginJsonContent = s.OriginJsonContent
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *OriginContentInsertReq) GetId() interface{} {
	return s.ContentId
}

type OriginContentGetPageReq struct {
	dto.Pagination `search:"-"`
	UscId          string `form:"uscId"  search:"type:exact;column:usc_id;table:rskc_origin_content" comment:"社会统一信用代码"`
	YearMonth      string `form:"yearMonth"  search:"type:exact;column:year_month;table:rskc_origin_content" comment:"数据生成年月"`
	StatusCode     int    `form:"statusCode" search:"type:exact;column:status_code;table:rskc_origin_content" comment:"状态码状态码：1.等待解析, 2.上下游企业已同步至trades_detail表，等待采集/确认 3.数据采集匹配录入完成"`
}

type OriginContentPageOrder struct {
	YearMonthOrder string `form:"yearMonthOrder" search:"type:order;column:year_month;table:rskc_origin_content"`
}

func (m *OriginContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OriginContentUpdateReq struct {
	ContentId         string `json:"contentId"`
	UscId             string `json:"uscId"`
	YearMonth         string `json:"yearMonth"`
	OriginJsonContent string `json:"originJsonContent"`
	StatusCode        int    `json:"statusCode"`
	common.ControlBy
}

func (s *OriginContentUpdateReq) Generate(model *models.OriginContent) {
	model.ContentId = s.ContentId
	model.UscId = s.UscId
	model.YearMonth = s.YearMonth
	model.OriginJsonContent = s.OriginJsonContent
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *OriginContentUpdateReq) GetId() interface{} {
	return s.ContentId
}

type OriginContentDeleteReq struct {
	Ids []string `json:"ids"`
}

func (s *OriginContentDeleteReq) GetIds() interface{} {
	return s.Ids
}
