package dto

import (
	"go-admin/app/cms/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type TbCmsPageGetPageReq struct {
	dto.Pagination `search:"-"`

	Id        int64     `form:"id"  search:"type:exact;column:id;table:tb_cms_page" comment:"主键编码"`
	Title     string    `form:"title"  search:"type:exact;column:title;table:tb_cms_page" comment:"页面名称"`
	Mark      string    `form:"mark"  search:"type:exact;column:mark;table:tb_cms_page" comment:"页面标记"`
	Source    string    `form:"source"  search:"type:exact;column:source;table:tb_cms_page" comment:"引用来源"`
	Author    string    `form:"author"  search:"type:exact;column:author;table:tb_cms_page" comment:"作者"`
	Content   string    `form:"content"  search:"type:exact;column:content;table:tb_cms_page" comment:"内容"`
	PubTime   time.Time `form:"pubTime"  search:"type:exact;column:pub_time;table:tb_cms_page" comment:"发布时间"`
	StartTime string    `form:"startTime" search:"type:gte;column:created_at;table:tb_cms_page" comment:"创建时间"`
	EndTime   string    `form:"endTime" search:"type:lte;column:created_at;table:tb_cms_page" comment:"创建时间"`
	TbCmsPagePageOrder
}

type TbCmsPagePageOrder struct {
	Id      int64     `form:"idOrder"  search:"type:order;column:id;table:tb_cms_page"`
	Title   string    `form:"titleOrder"  search:"type:order;column:title;table:tb_cms_page"`
	Mark    string    `form:"markOrder"  search:"type:order;column:mark;table:tb_cms_page"`
	Source  string    `form:"sourceOrder"  search:"type:order;column:source;table:tb_cms_page"`
	Author  string    `form:"authorOrder"  search:"type:order;column:author;table:tb_cms_page"`
	Content string    `form:"contentOrder"  search:"type:order;column:content;table:tb_cms_page"`
	PubTime time.Time `form:"pubTimeOrder"  search:"type:order;column:pub_time;table:tb_cms_page"`
}

func (m *TbCmsPageGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbCmsPageGetResp struct {
	Id      int64     `json:"id"`      // 主键编码
	Title   string    `json:"title"`   // 页面名称
	Mark    string    `json:"mark"`    // 页面标记
	Source  string    `json:"source"`  // 引用来源
	Author  string    `json:"author"`  // 作者
	Content string    `json:"content"` // 内容
	PubTime time.Time `json:"pubTime"` // 发布时间
	common.ControlBy
}

func (s *TbCmsPageGetResp) Generate(model *models.TbCmsPage) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.Title = model.Title
	s.Mark = model.Mark
	s.Source = model.Source
	s.Author = model.Author
	s.Content = model.Content
	s.PubTime = model.PubTime
	s.CreateBy = model.CreateBy
}

type TbCmsPageInsertReq struct {
	Id      int64     `json:"-"`       // 主键编码
	Title   string    `json:"title"`   // 页面名称
	Mark    string    `json:"mark"`    // 页面标记
	Source  string    `json:"source"`  // 引用来源
	Author  string    `json:"author"`  // 作者
	Content string    `json:"content"` // 内容
	PubTime time.Time `json:"pubTime"` // 发布时间
	common.ControlBy
}

func (s *TbCmsPageInsertReq) Generate(model *models.TbCmsPage) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Mark = s.Mark
	model.Source = s.Source
	model.Author = s.Author
	model.Content = s.Content
	model.PubTime = s.PubTime
	model.CreateBy = s.CreateBy
}

func (s *TbCmsPageInsertReq) GetId() interface{} {
	return s.Id
}

type TbCmsPageUpdateReq struct {
	Id      int64     `uri:"id"`       // 主键编码
	Title   string    `json:"title"`   // 页面名称
	Mark    string    `json:"mark"`    // 页面标记
	Source  string    `json:"source"`  // 引用来源
	Author  string    `json:"author"`  // 作者
	Content string    `json:"content"` // 内容
	PubTime time.Time `json:"pubTime"` // 发布时间
	common.ControlBy
}

func (s *TbCmsPageUpdateReq) Generate(model *models.TbCmsPage) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Mark = s.Mark
	model.Source = s.Source
	model.Author = s.Author
	model.Content = s.Content
	model.PubTime = s.PubTime
	model.UpdateBy = s.UpdateBy
}

func (s *TbCmsPageUpdateReq) GetId() interface{} {
	return s.Id
}

// TbCmsPageGetReq 功能获取请求参数
type TbCmsPageGetReq struct {
	Id int64 `uri:"id"`
}

func (s *TbCmsPageGetReq) GetId() interface{} {
	return s.Id
}

// TbCmsPageDeleteReq 功能删除请求参数
type TbCmsPageDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbCmsPageDeleteReq) GetId() interface{} {
	return s.Ids
}
