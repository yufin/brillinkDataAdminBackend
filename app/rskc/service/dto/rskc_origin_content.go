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
	common.ControlBy
}

func (s *OriginContentInsertReq) Generate(model *models.OriginContent) {
	model.ContentId = s.ContentId
	model.UscId = s.UscId
	model.YearMonth = s.YearMonth
	model.OriginJsonContent = s.OriginJsonContent
	model.CreateBy = s.CreateBy
}

func (s *OriginContentInsertReq) GetId() interface{} {
	return s.ContentId
}

type OriginContentGetPageReq struct {
	dto.Pagination `search:"-"`
	UscId          string `form:"uscId"  search:"type:exact;column:usc_id;table:rskc_origin_content" comment:"社会统一信用代码"`
	YearMonth      string `form:"yearMonth"  search:"type:exact;column:year_month;table:rskc_origin_content" comment:"数据生成年月"`
}

type OriginContentPageOrder struct {
	YearMonthOrder string `form:"yearMonthOrder" search:"type:order;column:year_month;table:rskc_origin_content"`
}

func (m *OriginContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}
