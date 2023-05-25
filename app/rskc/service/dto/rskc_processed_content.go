package dto

import (
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type RskcProcessedContentGetPageReq struct {
	dto.Pagination `search:"-"`

	ContentId  int64  `form:"contentId"  search:"type:exact;column:content_id;table:rskc_processed_content" comment:"外键(rskc_origin_content.id)"`
	Content    string `form:"content"  search:"type:exact;column:content;table:rskc_processed_content" comment:"数据(json字符串格式)"`
	StatusCode int    `form:"statusCode"  search:"type:exact;column:status_code;table:rskc_processed_content" comment:"状态码"`
	RskcProcessedContentPageOrder
}

type RskcProcessedContentPageOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:rskc_processed_content"`
	ContentId  string `form:"contentIdOrder"  search:"type:order;column:content_id;table:rskc_processed_content"`
	Content    string `form:"contentOrder"  search:"type:order;column:content;table:rskc_processed_content"`
	StatusCode string `form:"statusCodeOrder"  search:"type:order;column:status_code;table:rskc_processed_content"`
}

func (m *RskcProcessedContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RskcProcessedContentGetResp struct {
	Id         int64  `json:"id"`         // 主键
	ContentId  int64  `json:"contentId"`  // 外键(rskc_origin_content.id)
	Content    string `json:"content"`    // 数据(json字符串格式)
	StatusCode int    `json:"statusCode"` // 状态码
	common.ControlBy
}

func (s *RskcProcessedContentGetResp) Generate(model *models.RskcProcessedContent) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ContentId = model.ContentId
	s.Content = model.Content
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type RskcProcessedContentInsertReq struct {
	Id         int64  `json:"-"`          // 主键
	ContentId  int64  `json:"contentId"`  // 外键(rskc_origin_content.id)
	Content    string `json:"content"`    // 数据(json字符串格式)
	StatusCode int    `json:"statusCode"` // 状态码
	common.ControlBy
}

func (s *RskcProcessedContentInsertReq) Generate(model *models.RskcProcessedContent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: utils.NewFlakeId()}
	}
	model.ContentId = s.ContentId
	model.Content = s.Content
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *RskcProcessedContentInsertReq) GetId() interface{} {
	return s.Id
}

type RskcProcessedContentUpdateReq struct {
	Id         int64  `uri:"id"`          // 主键
	ContentId  int64  `json:"contentId"`  // 外键(rskc_origin_content.id)
	Content    string `json:"content"`    // 数据(json字符串格式)
	StatusCode int    `json:"statusCode"` // 状态码
	common.ControlBy
}

func (s *RskcProcessedContentUpdateReq) Generate(model *models.RskcProcessedContent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ContentId = s.ContentId
	model.Content = s.Content
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *RskcProcessedContentUpdateReq) GetId() interface{} {
	return s.Id
}

// RskcProcessedContentGetReq 功能获取请求参数
type RskcProcessedContentGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RskcProcessedContentGetReq) GetId() interface{} {
	return s.Id
}

// RskcProcessedContentDeleteReq 功能删除请求参数
type RskcProcessedContentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RskcProcessedContentDeleteReq) GetId() interface{} {
	return s.Ids
}

type RskcProcessedContentExport struct {
	Id         int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:主键" xlsx:"主键"`
	ContentId  int64  `json:"contentId" gorm:"comment:外键(rskc_origin_content.id)" xlsx:"外键(rskc_origin_content.id)"`
	Content    string `json:"content" gorm:"comment:数据(json字符串格式)" xlsx:"数据(json字符串格式)"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态码" xlsx:"状态码"`
}
