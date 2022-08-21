package dto

import (
	"go-admin/app/worklist/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type TbTodoGetPageReq struct {
	dto.Pagination `search:"-"`
	Status         string `form:"status"  search:"type:contains;column:status;table:tb_todo"`
	Title          string `form:"title"  search:"type:contains;column:title;table:tb_todo"`
	TbTodoOrder
}

type TbTodoOrder struct {
	EndAtOrder string `search:"type:order;column:end_at;table:tb_todo" form:"createdAtOrder"`
}

func (m *TbTodoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// TbTodoInsertReq 功能创建请求参数
type TbTodoInsertReq struct {
	Id             int    `json:"-"  ` // 编码
	Owner          string `json:"owner" `
	SubDescription string `json:"subDescription" `
	Title          string `json:"title"`
	Percent        int    `json:"percent"`
	Content        string `json:"content"`
	EndAt          string `json:"endAt"`
	Priority       string `json:"priority"`
	Status         string `json:"status"`
	common.ControlBy
}

func (s *TbTodoInsertReq) Generate(model *models.TbTodo) {
	model.Owner = s.Owner
	model.Title = s.Title
	model.SubDescription = s.SubDescription
	model.Content = s.Content
	model.EndAt, _ = time.Parse("2006-01-02", s.EndAt)
	model.Status = s.Status
	model.Priority = s.Priority
	model.Percent = s.Percent
}

func (s *TbTodoInsertReq) GetId() interface{} {
	return s.Id
}

type TbTodoUpdateReq struct {
	Id             int    `uri:"id"` // 编码
	Owner          string `json:"owner" `
	SubDescription string `json:"subDescription" `
	Title          string `json:"title"`
	Percent        int    `json:"percent"`
	Content        string `json:"content"`
	EndAt          string `json:"endAt"`
	Priority       string `json:"priority"`
	Status         string `json:"status"`
	common.ControlBy
}

func (s *TbTodoUpdateReq) Generate(model *models.TbTodo) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Owner = s.Owner
	model.SubDescription = s.SubDescription
	model.Title = s.Title
	model.Percent = s.Percent
	model.Content = s.Content
	model.Priority = s.Priority
	model.EndAt, _ = time.Parse("2006-01-02", s.EndAt)
	model.Status = s.Status
}

func (s *TbTodoUpdateReq) GetId() interface{} {
	return s.Id
}

type TbTodoGetReq struct {
	Id int `uri:"id"`
}

func (s *TbTodoGetReq) GetId() interface{} {
	return s.Id
}

type TbTodoDeleteReq struct {
	Ids int `json:"id"`
	common.ControlBy
}

func (s *TbTodoDeleteReq) GetId() interface{} {
	return s.Ids
}
