package dto

import (
	"encoding/json"
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type SysAbnormalLogGetPageReq struct {
	dto.Pagination `search:"-"`
	CreatedAt      time.Time `form:"createdAt"  search:"type:exact;column:created_at;table:sys_abnormal_log" comment:"创建时间"`
	SysAbnormalLogPageOrder
}

type SysAbnormalLogPageOrder struct {
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_abnormal_log"`
}

func (m *SysAbnormalLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysAbnormalLogGetPageResp struct {
	AbId     int64  `json:"abId" comment:"编码"` // 编码
	Method   string `json:"method" comment:"请求方式"`
	Url      string `json:"url" comment:"请求地址"`
	Ip       string `json:"ip" comment:"ip"`
	AbInfo   string `json:"abInfo" comment:"异常信息"`
	UserName string `json:"userName" comment:"操作人"`
	common.ControlBy
	common.ModelTime
}

type SysAbnormalLogGetResp struct {
	AbId       int64           `json:"-" comment:"编码"` // 编码
	Method     string          `json:"method" comment:"请求方式"`
	Url        string          `json:"url" comment:"请求地址"`
	Ip         string          `json:"ip" comment:"ip"`
	AbInfo     string          `json:"abInfo" comment:"异常信息"`
	AbSource   string          `json:"abSource" comment:"异常来源"`
	AbFunc     string          `json:"abFunc" comment:"异常方法"`
	UserId     int64           `json:"userId" comment:"用户id"`
	UserName   string          `json:"userName" comment:"操作人"`
	Headers    json.RawMessage `json:"headers" comment:"请求头"`
	Body       json.RawMessage `json:"body" comment:"请求数据"`
	Resp       json.RawMessage `json:"resp" comment:"回调数据"`
	StackTrace string          `json:"stackTrace" comment:"堆栈追踪"`
	common.ControlBy
	common.ModelTime
}

func (s *SysAbnormalLogGetResp) Generate(model *models.SysAbnormalLog) {
	s.AbId = model.AbId
	s.Method = model.Method
	s.Url = model.Url
	s.Ip = model.Ip
	s.AbInfo = model.AbInfo
	s.AbSource = model.AbSource
	s.AbFunc = model.AbFunc
	s.UserId = model.UserId
	s.UserName = model.UserName
	if model.Headers == "" {
		model.Headers = "{}"
	}
	s.Headers = json.RawMessage(model.Headers)
	if model.Body == "" {
		model.Body = "{}"
	}
	s.Body = json.RawMessage(model.Body)
	if model.Resp == "" {
		model.Resp = "{}"
	}
	s.Resp = json.RawMessage(model.Resp)
	s.StackTrace = model.StackTrace
	s.CreateBy = model.CreateBy
	s.CreatedAt = model.CreatedAt
}

func (s *SysAbnormalLogGetResp) GetId() interface{} {
	return s.AbId
}

type SysAbnormalLogUpdateReq struct {
	AbId       int64  `uri:"abId" comment:"编码"` // 编码
	Method     string `json:"method" comment:"请求方式"`
	Url        string `json:"url" comment:"请求地址"`
	Ip         string `json:"ip" comment:"ip"`
	AbInfo     string `json:"abInfo" comment:"异常信息"`
	AbSource   string `json:"abSource" comment:"异常来源"`
	AbFunc     string `json:"abFunc" comment:"异常方法"`
	UserId     int64  `json:"userId" comment:"用户id"`
	UserName   string `json:"userName" comment:"操作人"`
	Headers    string `json:"headers" comment:"请求头"`
	Body       string `json:"body" comment:"请求数据"`
	Resp       string `json:"resp" comment:"回调数据"`
	StackTrace string `json:"stackTrace" comment:"堆栈追踪"`
	common.ControlBy
}

func (s *SysAbnormalLogUpdateReq) Generate(model *models.SysAbnormalLog) {
	model.AbId = s.AbId
	model.Method = s.Method
	model.Url = s.Url
	model.Ip = s.Ip
	model.AbInfo = s.AbInfo
	model.AbSource = s.AbSource
	model.AbFunc = s.AbFunc
	model.UserId = s.UserId
	model.UserName = s.UserName
	model.Headers = s.Headers
	model.Body = s.Body
	model.Resp = s.Resp
	model.StackTrace = s.StackTrace
	model.UpdateBy = s.UpdateBy
}

func (s *SysAbnormalLogUpdateReq) GetId() interface{} {
	return s.AbId
}

// SysAbnormalLogGetReq 功能获取请求参数
type SysAbnormalLogGetReq struct {
	AbId int64 `uri:"id"`
}

func (s *SysAbnormalLogGetReq) GetId() interface{} {
	return s.AbId
}

// SysAbnormalLogDeleteReq 功能删除请求参数
type SysAbnormalLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysAbnormalLogDeleteReq) GetId() interface{} {
	return s.Ids
}
