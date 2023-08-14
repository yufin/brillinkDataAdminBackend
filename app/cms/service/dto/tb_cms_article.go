package dto

import (
	"go-admin/app/cms/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type TbCmsArticleGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string    `form:"title"  search:"type:exact;column:title;table:tb_cms_article" comment:"页面名称"`
	Mark           string    `form:"mark"  search:"type:exact;column:mark;table:tb_cms_article" comment:"页面标记"`
	Source         string    `form:"source"  search:"type:exact;column:source;table:tb_cms_article" comment:"引用来源"`
	Author         string    `form:"author"  search:"type:exact;column:author;table:tb_cms_article" comment:"作者"`
	Category       string    `form:"category"  search:"type:exact;column:category;table:tb_cms_article" comment:"分类"`
	File           string    `form:"file"  search:"type:exact;column:file;table:tb_cms_article" comment:"上传文件路径"`
	PubTime        time.Time `form:"pubTime"  search:"type:exact;column:pub_time;table:tb_cms_article" comment:"发布时间"`
	CreatedAt      time.Time `form:"createdAt"  search:"type:exact;column:created_at;table:tb_cms_article" comment:"创建时间"`
	TbCmsArticlePageOrder
}

type TbCmsArticlePageOrder struct {
	Id        int64     `form:"idOrder"  search:"type:order;column:id;table:tb_cms_article"`
	Title     string    `form:"titleOrder"  search:"type:order;column:title;table:tb_cms_article"`
	Mark      string    `form:"markOrder"  search:"type:order;column:mark;table:tb_cms_article"`
	Source    string    `form:"sourceOrder"  search:"type:order;column:source;table:tb_cms_article"`
	Author    string    `form:"authorOrder"  search:"type:order;column:author;table:tb_cms_article"`
	Category  string    `form:"categoryOrder"  search:"type:order;column:category;table:tb_cms_article"`
	Content   string    `form:"contentOrder"  search:"type:order;column:content;table:tb_cms_article"`
	File      string    `form:"fileOrder"  search:"type:order;column:file;table:tb_cms_article"`
	PubTime   time.Time `form:"pubTimeOrder"  search:"type:order;column:pub_time;table:tb_cms_article"`
	CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_cms_article"`
}

func (m *TbCmsArticleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbCmsArticleGetResp struct {
	Id       int64      `json:"id"`       // 主键编码
	Title    string     `json:"title"`    // 页面名称
	Mark     string     `json:"mark"`     // 页面标记
	Source   string     `json:"source"`   // 引用来源
	Author   string     `json:"author"`   // 作者
	Category string     `json:"category"` // 分类
	Content  string     `json:"content"`  // 内容
	File     string     `json:"file"`     // 上传文件路径
	PubTime  *time.Time `json:"pubTime"`  // 发布时间
	common.ControlBy
}

func (s *TbCmsArticleGetResp) Generate(model *models.TbCmsArticle) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.Title = model.Title
	s.Mark = model.Mark
	s.Source = model.Source
	s.Author = model.Author
	s.Category = model.Category
	s.Content = model.Content
	s.File = model.File
	s.PubTime = model.PubTime
	s.CreateBy = model.CreateBy
}

type TbCmsArticleInsertReq struct {
	Id       int64      `json:"-"`        // 主键编码
	Title    string     `json:"title"`    // 页面名称
	Mark     string     `json:"mark"`     // 页面标记
	Source   string     `json:"source"`   // 引用来源
	Author   string     `json:"author"`   // 作者
	Category string     `json:"category"` // 分类
	Content  string     `json:"content"`  // 内容
	File     string     `json:"file"`     // 上传文件路径
	PubTime  *time.Time `json:"pubTime"`  // 发布时间
	common.ControlBy
}

func (s *TbCmsArticleInsertReq) Generate(model *models.TbCmsArticle) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Mark = s.Mark
	model.Source = s.Source
	model.Author = s.Author
	model.Category = s.Category
	model.Content = s.Content
	model.File = s.File
	model.PubTime = s.PubTime
	model.CreateBy = s.CreateBy
}

func (s *TbCmsArticleInsertReq) GetId() interface{} {
	return s.Id
}

type TbCmsArticleUpdateReq struct {
	Id       int64      `uri:"id"`        // 主键编码
	Title    string     `json:"title"`    // 页面名称
	Mark     string     `json:"mark"`     // 页面标记
	Source   string     `json:"source"`   // 引用来源
	Author   string     `json:"author"`   // 作者
	Category string     `json:"category"` // 分类
	Content  string     `json:"content"`  // 内容
	File     string     `json:"file"`     // 上传文件路径
	PubTime  *time.Time `json:"pubTime"`  // 发布时间
	common.ControlBy
}

func (s *TbCmsArticleUpdateReq) Generate(model *models.TbCmsArticle) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Mark = s.Mark
	model.Source = s.Source
	model.Author = s.Author
	model.Category = s.Category
	model.Content = s.Content
	model.File = s.File
	model.PubTime = s.PubTime
	model.UpdateBy = s.UpdateBy
}

func (s *TbCmsArticleUpdateReq) GetId() interface{} {
	return s.Id
}

// TbCmsArticleGetReq 功能获取请求参数
type TbCmsArticleGetReq struct {
	Id int64 `uri:"id"`
}

func (s *TbCmsArticleGetReq) GetId() interface{} {
	return s.Id
}

// TbCmsArticleDeleteReq 功能删除请求参数
type TbCmsArticleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbCmsArticleDeleteReq) GetId() interface{} {
	return s.Ids
}
