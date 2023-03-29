package dto

import (
	"go-admin/app/cms/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type TbCmsMenuGetPageReq struct {
	dto.Pagination `search:"-"`

	Id        int64     `form:"id"  search:"type:exact;column:id;table:tb_cms_menu" comment:"主键编码"`
	Name      string    `form:"name"  search:"type:exact;column:name;table:tb_cms_menu" comment:"页面名称"`
	Type      string    `form:"type"  search:"type:exact;column:type;table:tb_cms_menu" comment:"菜单类型，list列表，page详情页"`
	Link      string    `form:"link"  search:"type:exact;column:link;table:tb_cms_menu" comment:"引用表的id，或者mark"`
	Parent    string    `form:"parent"  search:"type:exact;column:parent;table:tb_cms_menu" comment:"父节点"`
	CreatedAt time.Time `form:"createdAt"  search:"type:exact;column:created_at;table:tb_cms_menu" comment:"创建时间"`
	UpdatedAt time.Time `form:"updatedAt"  search:"type:exact;column:updated_at;table:tb_cms_menu" comment:"最后更新时间"`
	DeletedAt time.Time `form:"deletedAt"  search:"type:exact;column:deleted_at;table:tb_cms_menu" comment:"删除时间"`
	CreateBy  int64     `form:"createBy"  search:"type:exact;column:create_by;table:tb_cms_menu" comment:"创建者"`
	UpdateBy  int64     `form:"updateBy"  search:"type:exact;column:update_by;table:tb_cms_menu" comment:"更新者"`
	TbCmsMenuPageOrder
}

type TbCmsMenuPageOrder struct {
	Id        int64     `form:"idOrder"  search:"type:order;column:id;table:tb_cms_menu"`
	Name      string    `form:"nameOrder"  search:"type:order;column:name;table:tb_cms_menu"`
	Type      string    `form:"typeOrder"  search:"type:order;column:type;table:tb_cms_menu"`
	Link      string    `form:"linkOrder"  search:"type:order;column:link;table:tb_cms_menu"`
	Parent    string    `form:"parentOrder"  search:"type:order;column:parent;table:tb_cms_menu"`
	CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_cms_menu"`
	UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_cms_menu"`
	DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_cms_menu"`
	CreateBy  int64     `form:"createByOrder"  search:"type:order;column:create_by;table:tb_cms_menu"`
	UpdateBy  int64     `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_cms_menu"`
}

func (m *TbCmsMenuGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbCmsMenuGetResp struct {
	Id     int64  `json:"id"`     // 主键编码
	Name   string `json:"name"`   // 页面名称
	Type   string `json:"type"`   // 菜单类型，list列表，page详情页
	Link   string `json:"link"`   // 引用表的id，或者mark
	Parent string `json:"parent"` // 父节点
	common.ControlBy
}

func (s *TbCmsMenuGetResp) Generate(model *models.TbCmsMenu) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.Name = model.Name
	s.Type = model.Type
	s.Link = model.Link
	s.Parent = model.Parent
	s.CreateBy = model.CreateBy
}

type TbCmsMenuInsertReq struct {
	Id     int64  `json:"-"`      // 主键编码
	Name   string `json:"name"`   // 页面名称
	Type   string `json:"type"`   // 菜单类型，list列表，page详情页
	Link   string `json:"link"`   // 引用表的id，或者mark
	Parent string `json:"parent"` // 父节点
	common.ControlBy
}

func (s *TbCmsMenuInsertReq) Generate(model *models.TbCmsMenu) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Type = s.Type
	model.Link = s.Link
	model.Parent = s.Parent
	model.CreateBy = s.CreateBy
}

func (s *TbCmsMenuInsertReq) GetId() interface{} {
	return s.Id
}

type TbCmsMenuUpdateReq struct {
	Id     int64  `uri:"id"`      // 主键编码
	Name   string `json:"name"`   // 页面名称
	Type   string `json:"type"`   // 菜单类型，list列表，page详情页
	Link   string `json:"link"`   // 引用表的id，或者mark
	Parent string `json:"parent"` // 父节点
	common.ControlBy
}

func (s *TbCmsMenuUpdateReq) Generate(model *models.TbCmsMenu) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Type = s.Type
	model.Link = s.Link
	model.Parent = s.Parent
	model.UpdateBy = s.UpdateBy
}

func (s *TbCmsMenuUpdateReq) GetId() interface{} {
	return s.Id
}

// TbCmsMenuGetReq 功能获取请求参数
type TbCmsMenuGetReq struct {
	Id int64 `uri:"id"`
}

func (s *TbCmsMenuGetReq) GetId() interface{} {
	return s.Id
}

// TbCmsMenuDeleteReq 功能删除请求参数
type TbCmsMenuDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbCmsMenuDeleteReq) GetId() interface{} {
	return s.Ids
}
