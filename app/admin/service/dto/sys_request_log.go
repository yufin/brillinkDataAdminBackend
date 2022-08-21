package dto

import (
	"encoding/json"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysRequestLogGetPageReq struct {
	dto.Pagination `search:"-"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:oper_time;table:sys_request_log" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:oper_time;table:sys_request_log" comment:"创建时间"`
	SysRequestLogOrder
}

type SysRequestLogOrder struct {
	OperTimeOrder string `search:"type:order;column:oper_time;table:sys_request_log" form:"operTimeOrder"`
}

func (m *SysRequestLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysRequestLogGetPageResp struct {
	ID            int       `uri:"Id" json:"id" comment:"编码"`     // 编码
	RequestMethod string    `json:"requestMethod" comment:"请求方式"` // 请求方式
	OperName      string    `json:"operName" comment:"操作者"`       // 操作者
	OperUrl       string    `json:"operUrl" comment:"访问地址"`       // 访问地址
	OperIp        string    `json:"operIp" comment:"客户端ip"`       // 客户端ip
	OperLocation  string    `json:"operLocation" comment:"访问位置"`  // 访问位置
	OperTime      time.Time `json:"operTime" comment:"操作时间"`      // 操作时间
	LatencyTime   string    `json:"latencyTime" comment:"耗时"`     // 耗时
	UserAgent     string    `json:"userAgent" comment:"ua"`       // ua
	common.ModelTime
	common.ControlBy
}

type SysRequestLogGetUriReq struct {
	SysRequestLogGetPvReq
}

type SysRequestLogGetPvReq struct {
	Type      string     `form:"type" search:"-"`
	BeginTime *time.Time `form:"beginTime" search:"type:gte;column:oper_time;table:sys_request_log" comment:"创建时间"`
	EndTime   *time.Time `form:"endTime" search:"type:lte;column:oper_time;table:sys_request_log" comment:"创建时间"`
	SysRequestLogOrder
}

func (m *SysRequestLogGetPvReq) GetNeedSearch() interface{} {
	return *m
}

type SysRequestLogGetMethodResp struct {
	RequestMethod string `json:"requestMethod,omitempty"`
	Y             int    `json:"y"`
}

type SysRequestLogGetPvResp struct {
	X string `json:"x,omitempty"`
	Y int    `json:"y"`
}

type SysRequestLogGetPvWithTypeResp struct {
	Y int `json:"y"`
}

type SysRequestLogGetResp struct {
	ID            int64           `uri:"Id" comment:"编码"`                  // 编码
	RequestMethod string          `json:"requestMethod" comment:"请求方式"`    // 请求方式
	OperName      string          `json:"operName" comment:"操作者"`          // 操作者
	OperUrl       string          `json:"operUrl" comment:"访问地址"`          // 访问地址
	OperIp        string          `json:"operIp" comment:"客户端ip"`          // 客户端ip
	OperLocation  string          `json:"operLocation" comment:"访问位置"`     // 访问位置
	OperParam     json.RawMessage `json:"operParam" comment:"请求参数"`        // 请求参数
	OperHeaders   json.RawMessage `json:"operHeaders" comment:"请求Headers"` // 请求Headers
	OperTime      time.Time       `json:"operTime" comment:"操作时间"`         // 操作时间
	JsonResult    json.RawMessage `json:"jsonResult" comment:"返回数据"`       // 返回数据
	LatencyTime   string          `json:"latencyTime" comment:"耗时"`        // 耗时
	UserAgent     string          `json:"userAgent" comment:"ua"`          // ua
	common.ModelTime
	common.ControlBy
}

func (s *SysRequestLogGetResp) Generate(model *models.SysRequestLog) {
	s.ID = model.Id
	s.RequestMethod = model.RequestMethod
	s.OperName = model.OperName
	s.OperUrl = model.OperUrl
	s.OperIp = model.OperIp
	s.OperLocation = model.OperLocation
	if model.OperParam == "" {
		model.OperParam = "[]"
	}
	s.OperParam = json.RawMessage(model.OperParam)
	if model.OperHeaders == "" {
		model.OperHeaders = "[]"
	}
	s.OperHeaders = json.RawMessage(model.OperHeaders)
	s.OperTime = model.OperTime
	if model.JsonResult == "" {
		model.JsonResult = "[]"
	}
	s.JsonResult = json.RawMessage(model.JsonResult)
	s.LatencyTime = model.LatencyTime
	s.UserAgent = model.UserAgent
}

type SysRequestLogControl struct {
	ID            int64           `uri:"Id" comment:"编码"` // 编码
	RequestMethod string          `json:"requestMethod" comment:"请求方式"`
	OperName      string          `json:"operName" comment:"操作者"`
	OperUrl       string          `json:"operUrl" comment:"访问地址"`
	OperIp        string          `json:"operIp" comment:"客户端ip"`
	OperLocation  string          `json:"operLocation" comment:"访问位置"`
	OperParam     json.RawMessage `json:"operParam" comment:"请求参数"`
	OperHeaders   json.RawMessage `json:"operHeaders" comment:"请求Headers"` // 请求Headers
	OperTime      time.Time       `json:"operTime" comment:"操作时间"`
	JsonResult    json.RawMessage `json:"jsonResult" comment:"返回数据"`
	LatencyTime   string          `json:"latencyTime" comment:"耗时"`
	UserAgent     string          `json:"userAgent" comment:"ua"`
}

func (s *SysRequestLogControl) Generate() (*models.SysRequestLog, error) {
	return &models.SysRequestLog{
		Model:         common.Model{Id: s.ID},
		RequestMethod: s.RequestMethod,
		OperName:      s.OperName,
		OperUrl:       s.OperUrl,
		OperIp:        s.OperIp,
		OperLocation:  s.OperLocation,
		OperParam:     string(s.OperParam),
		OperHeaders:   string(s.OperHeaders),
		OperTime:      s.OperTime,
		JsonResult:    string(s.JsonResult),
		LatencyTime:   s.LatencyTime,
		UserAgent:     s.UserAgent,
	}, nil
}

func (s *SysRequestLogControl) GetId() interface{} {
	return s.ID
}

type SysRequestLogGetReq struct {
	Id int `uri:"id"`
}

func (s *SysRequestLogGetReq) GetId() interface{} {
	return s.Id
}

// SysRequestLogDeleteReq 功能删除请求参数
type SysRequestLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysRequestLogDeleteReq) GetId() interface{} {
	return s.Ids
}
