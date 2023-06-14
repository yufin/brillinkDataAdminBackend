package dto

import (
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RcI18nContentGetPageReq struct {
	dto.Pagination `search:"-"`

	ProcessedId int64  `form:"processedId"  search:"type:exact;column:processed_id;table:rc_ixxn_content" comment:"rskc_processed_content.id"`
	Lang        string `form:"lang"  search:"type:exact;column:lang;table:rc_ixxn_content" comment:"语言类型(en,...)"`
	Content     string `form:"content"  search:"type:exact;column:content;table:rc_ixxn_content" comment:"报文json string"`
	RcI18nContentPageOrder
}

type RcI18nContentPageOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:rc_ixxn_content"`
	ProcessedId string `form:"processedIdOrder"  search:"type:order;column:processed_id;table:rc_ixxn_content"`
	Lang        string `form:"langOrder"  search:"type:order;column:lang;table:rc_ixxn_content"`
	Content     string `form:"contentOrder"  search:"type:order;column:content;table:rc_ixxn_content"`
}

func (m *RcI18nContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RcI18nContentGetResp struct {
	Id          int64  `json:"id"`          // 主键
	ProcessedId int64  `json:"processedId"` // rskc_processed_content.id
	Lang        string `json:"lang"`        // 语言类型(en,...)
	Content     string `json:"content"`     // 报文json string
	common.ControlBy
}

func (s *RcI18nContentGetResp) Generate(model *models.RcI18nContent) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ProcessedId = model.ProcessedId
	s.Lang = model.Lang
	s.Content = model.Content
	s.CreateBy = model.CreateBy
}

type RcI18nContentInsertReq struct {
	Id          int64  `json:"-"`           // 主键
	ProcessedId int64  `json:"processedId"` // rskc_processed_content.id
	Lang        string `json:"lang"`        // 语言类型(en,...)
	Content     string `json:"content"`     // 报文json string
	common.ControlBy
}

func (s *RcI18nContentInsertReq) Generate(model *models.RcI18nContent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProcessedId = s.ProcessedId
	model.Lang = s.Lang
	model.Content = s.Content
	model.CreateBy = s.CreateBy
}

func (s *RcI18nContentInsertReq) GetId() interface{} {
	return s.Id
}

type RcI18nContentUpdateReq struct {
	Id          int64  `uri:"id"`           // 主键
	ProcessedId int64  `json:"processedId"` // rskc_processed_content.id
	Lang        string `json:"lang"`        // 语言类型(en,...)
	Content     string `json:"content"`     // 报文json string
	common.ControlBy
}

func (s *RcI18nContentUpdateReq) Generate(model *models.RcI18nContent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProcessedId = s.ProcessedId
	model.Lang = s.Lang
	model.Content = s.Content
	model.UpdateBy = s.UpdateBy
}

func (s *RcI18nContentUpdateReq) GetId() interface{} {
	return s.Id
}

// RcI18nContentGetReq 功能获取请求参数
type RcI18nContentGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RcI18nContentGetReq) GetId() interface{} {
	return s.Id
}

// RcI18nContentDeleteReq 功能删除请求参数
type RcI18nContentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RcI18nContentDeleteReq) GetId() interface{} {
	return s.Ids
}

type RcI18nContentExport struct {
	Id          int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:主键" xlsx:"主键"`
	ProcessedId int64  `json:"processedId" gorm:"comment:rskc_processed_content.id" xlsx:"rskc_processed_content.id"`
	Lang        string `json:"lang" gorm:"comment:语言类型(en,...)" xlsx:"语言类型(en,...)"`
	Content     string `json:"content" gorm:"comment:报文json string" xlsx:"报文json string"`
}
