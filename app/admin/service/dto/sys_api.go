package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysApiGetPageReq 功能列表请求参数
type SysApiGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:sys_api" comment:"标题"`
	Path           string `form:"path"  search:"type:contains;column:path;table:sys_api" comment:"地址"`
	Method         string `form:"method"  search:"type:contains;column:method;table:sys_api" comment:"类型"`
	IsHistory      string `form:"isHistory"  search:"type:exact;column:is_history;table:sys_api" comment:"是否记录"`
	SysApiOrder
}

type SysApiOrder struct {
	NameOrder      string `search:"type:order;column:name;table:sys_api" form:"nameOrder"`
	PathOrder      string `search:"type:order;column:path;table:sys_api" form:"pathOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_api" form:"createdAtOrder"`
}

func (m *SysApiGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysApiInsertReq 功能创建请求参数
type SysApiInsertReq struct {
	Id     int    `json:"-" comment:"编码"` // 编码
	Handle string `json:"handle" comment:"handle"`
	Name   string `json:"name" comment:"标题"`
	Path   string `json:"path" comment:"地址"`
	Type   string `json:"type" comment:""`
	Method string `json:"method" comment:"类型"`
	common.ControlBy
}

func (s *SysApiInsertReq) Generate(model *models.SysApi) {
	model.Handle = s.Handle
	model.Name = s.Name
	model.Path = s.Path
	model.Type = s.Type
	model.Method = s.Method
}

func (s *SysApiInsertReq) GetId() interface{} {
	return s.Id
}

// SysApiUpdateReq 功能更新请求参数
type SysApiUpdateReq struct {
	Id     int    `uri:"id" comment:"编码"` // 编码
	Handle string `json:"handle" comment:"handle"`
	Name   string `json:"name" comment:"标题"`
	Path   string `json:"path" comment:"地址"`
	Type   string `json:"type" comment:""`
	Method string `json:"method" comment:"类型"`
	common.ControlBy
}

func (s *SysApiUpdateReq) Generate(model *models.SysApi) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Handle = s.Handle
	model.Name = s.Name
	model.Path = s.Path
	model.Type = s.Type
	model.Method = s.Method
}

func (s *SysApiUpdateReq) GetId() interface{} {
	return s.Id
}

// SysApiGetReq 功能获取请求参数
type SysApiGetReq struct {
	Id int `uri:"id"`
}

func (s *SysApiGetReq) GetId() interface{} {
	return s.Id
}

// SysApiDeleteReq 功能删除请求参数
type SysApiDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysApiDeleteReq) Generate(model *models.SysApi) {
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

func (s *SysApiDeleteReq) GetId() interface{} {
	return s.Ids
}
