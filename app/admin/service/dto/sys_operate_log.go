package dto

import (
	"encoding/json"
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type SysOperateLogGetPageReq struct {
	dto.Pagination `search:"-"`
	Type           string    `form:"type"  search:"type:exact;column:type;table:sys_operate_log" comment:"操作类型"`
	Description    string    `form:"description"  search:"type:exact;column:description;table:sys_operate_log" comment:"操作说明"`
	UserName       string    `form:"userName"  search:"type:exact;column:user_name;table:sys_operate_log" comment:"用户"`
	UserId         int64     `form:"userId"  search:"type:exact;column:user_id;table:sys_operate_log" comment:"用户id"`
	CreatedAt      time.Time `form:"createdAt"  search:"type:exact;column:created_at;table:sys_operate_log" comment:"创建时间"`
	SysOperateLogPageOrder
}

type SysOperateLogPageOrder struct {
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_operate_log"`
}

func (m *SysOperateLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysOperateLogGetPageResp struct {
	LogId       int64  `json:"logId" comment:"编码"` // 编码
	Type        string `json:"type" comment:"操作类型"`
	Description string `json:"description" comment:"操作说明"`
	Project     string `json:"project" comment:"项目"`
	UserName    string `json:"userName" comment:"用户"`
	UserId      int64  `json:"userId" comment:"用户id"`
	common.ControlBy
	common.ModelTime
}

type SysOperateLogGetResp struct {
	LogId        int64           `json:"-" comment:"编码"` // 编码
	Type         string          `json:"type" comment:"操作类型"`
	Description  string          `json:"description" comment:"操作说明"`
	UserName     string          `json:"userName" comment:"用户"`
	UserId       int64           `json:"userId" comment:"用户id"`
	UpdateBefore json.RawMessage `json:"updateBefore" comment:"更新前"`
	UpdateAfter  json.RawMessage `json:"updateAfter" comment:"更新后"`
	common.ControlBy
	common.ModelTime
}

func (s *SysOperateLogGetResp) Generate(model *models.SysOperateLog) {
	s.LogId = model.LogId
	s.Type = model.Type
	s.Description = model.Description
	s.UserName = model.UserName
	s.UserId = model.UserId
	s.UpdateBefore = json.RawMessage(model.UpdateBefore)
	s.UpdateAfter = json.RawMessage(model.UpdateAfter)
	s.CreateBy = model.CreateBy
	s.CreatedAt = model.CreatedAt
}

// SysOperateLogGetReq 功能获取请求参数
type SysOperateLogGetReq struct {
	LogId int64 `uri:"id"`
}

func (s *SysOperateLogGetReq) GetId() interface{} {
	return s.LogId
}

// SysOperateLogDeleteReq 功能删除请求参数
type SysOperateLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysOperateLogDeleteReq) GetId() interface{} {
	return s.Ids
}
