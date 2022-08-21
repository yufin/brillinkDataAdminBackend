package dto

import (
	"go-admin/app/worklist/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type TbNoticeGetPageReq struct {
	dto.Pagination `search:"-"`
	//Status         string `form:"status"  search:"type:contains;column:status;table:tb_todo"`
	//Title          string `form:"title"  search:"type:contains;column:title;table:tb_todo"`
	//TbNoticeOrder
}

type TbNoticeOrder struct {
	EndAtOrder string `search:"type:order;column:end_at;table:tb_todo" form:"createdAtOrder"`
}

func (m *TbNoticeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// TbNoticeInsertReq 功能创建请求参数
type TbNoticeInsertReq struct {
	Id             int    `json:"-"  ` // 编码
	Owner          string `json:"owner" `
	SubDescription string `json:"subDescription" `
	Title          string `json:"title"`
	Percent        int    `json:"percent"`
	EndAt          string `json:"endAt"`
	Status         string `json:"status"`
	common.ControlBy
}

func (s *TbNoticeInsertReq) Generate(model *models.TbTodo) {
	model.Owner = s.Owner
	model.Title = s.Title
	model.SubDescription = s.SubDescription
	model.EndAt, _ = time.Parse("2006-01-02", s.EndAt)
	model.Status = s.Status
}

func (s *TbNoticeInsertReq) GetId() interface{} {
	return s.Id
}

type TbNoticeUpdateReq struct {
	Id             int    `uri:"id"` // 编码
	Owner          string `json:"owner" `
	SubDescription string `json:"subDescription" `
	Title          string `json:"title"`
	Percent        int    `json:"percent"`
	EndAt          string `json:"endAt"`
	Status         string `json:"status"`
	common.ControlBy
}

func (s *TbNoticeUpdateReq) Generate(model *models.TbTodo) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Owner = s.Owner
	model.SubDescription = s.SubDescription
	model.Title = s.Title
	model.Percent = s.Percent
	model.EndAt, _ = time.Parse("2006-01-02", s.EndAt)
	model.Status = s.Status
}

func (s *TbNoticeUpdateReq) GetId() interface{} {
	return s.Id
}

type TbNoticeGetReq struct {
	Id string `uri:"id"`
}

func (s *TbNoticeGetReq) GetId() interface{} {
	return s.Id
}

type TbNoticeDeleteReq struct {
	Ids string `json:"id"`
	common.ControlBy
}

func (s *TbNoticeDeleteReq) GetId() interface{} {
	return s.Ids
}
