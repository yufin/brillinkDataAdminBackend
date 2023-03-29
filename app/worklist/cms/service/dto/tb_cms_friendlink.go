package dto

import (
	"go-admin/app/cms/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type TbCmsFriendlinkGetPageReq struct {
	dto.Pagination `search:"-"`

	Id        int64     `form:"id"  search:"type:exact;column:id;table:tb_cms_friendlink" comment:"主键编码"`
	Name      string    `form:"name"  search:"type:exact;column:name;table:tb_cms_friendlink" comment:"链接名称"`
	Link      string    `form:"link"  search:"type:exact;column:link;table:tb_cms_friendlink" comment:"链接地址"`
	CreatedAt time.Time `form:"createdAt"  search:"type:exact;column:created_at;table:tb_cms_friendlink" comment:"创建时间"`
	UpdatedAt time.Time `form:"updatedAt"  search:"type:exact;column:updated_at;table:tb_cms_friendlink" comment:"最后更新时间"`
	DeletedAt time.Time `form:"deletedAt"  search:"type:exact;column:deleted_at;table:tb_cms_friendlink" comment:"删除时间"`
	CreateBy  int64     `form:"createBy"  search:"type:exact;column:create_by;table:tb_cms_friendlink" comment:"创建者"`
	UpdateBy  int64     `form:"updateBy"  search:"type:exact;column:update_by;table:tb_cms_friendlink" comment:"更新者"`
	TbCmsFriendlinkPageOrder
}

type TbCmsFriendlinkPageOrder struct {
	Id        int64     `form:"idOrder"  search:"type:order;column:id;table:tb_cms_friendlink"`
	Name      string    `form:"nameOrder"  search:"type:order;column:name;table:tb_cms_friendlink"`
	Link      string    `form:"linkOrder"  search:"type:order;column:link;table:tb_cms_friendlink"`
	CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_cms_friendlink"`
	UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_cms_friendlink"`
	DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_cms_friendlink"`
	CreateBy  int64     `form:"createByOrder"  search:"type:order;column:create_by;table:tb_cms_friendlink"`
	UpdateBy  int64     `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_cms_friendlink"`
}

func (m *TbCmsFriendlinkGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbCmsFriendlinkGetResp struct {
	Id   int64  `json:"id"`   // 主键编码
	Name string `json:"name"` // 链接名称
	Link string `json:"link"` // 链接地址
	common.ControlBy
}

func (s *TbCmsFriendlinkGetResp) Generate(model *models.TbCmsFriendlink) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.Name = model.Name
	s.Link = model.Link
	s.CreateBy = model.CreateBy
}

type TbCmsFriendlinkInsertReq struct {
	Id   int64  `json:"-"`    // 主键编码
	Name string `json:"name"` // 链接名称
	Link string `json:"link"` // 链接地址
	common.ControlBy
}

func (s *TbCmsFriendlinkInsertReq) Generate(model *models.TbCmsFriendlink) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Link = s.Link
	model.CreateBy = s.CreateBy
}

func (s *TbCmsFriendlinkInsertReq) GetId() interface{} {
	return s.Id
}

type TbCmsFriendlinkUpdateReq struct {
	Id   int64  `uri:"id"`    // 主键编码
	Name string `json:"name"` // 链接名称
	Link string `json:"link"` // 链接地址
	common.ControlBy
}

func (s *TbCmsFriendlinkUpdateReq) Generate(model *models.TbCmsFriendlink) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Link = s.Link
	model.UpdateBy = s.UpdateBy
}

func (s *TbCmsFriendlinkUpdateReq) GetId() interface{} {
	return s.Id
}

// TbCmsFriendlinkGetReq 功能获取请求参数
type TbCmsFriendlinkGetReq struct {
	Id int64 `uri:"id"`
}

func (s *TbCmsFriendlinkGetReq) GetId() interface{} {
	return s.Id
}

// TbCmsFriendlinkDeleteReq 功能删除请求参数
type TbCmsFriendlinkDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbCmsFriendlinkDeleteReq) GetId() interface{} {
	return s.Ids
}
