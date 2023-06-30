package models

import (
	"time"
)

type SysRequestLog struct {
	Model
	RequestMethod string    `json:"requestMethod" gorm:"size:128;comment:请求方式"`
	OperName      string    `json:"operName" gorm:"size:128;comment:操作者"`
	OperUrl       string    `json:"operUrl" gorm:"size:255;comment:访问地址"`
	OperIp        string    `json:"operIp" gorm:"size:128;comment:客户端ip"`
	OperLocation  string    `json:"operLocation" gorm:"size:128;comment:访问位置"`
	OperParam     string    `json:"operParam" gorm:"type:json;comment:请求参数"`
	OperHeaders   string    `json:"operHeaders" gorm:"type:json;comment:请求头"`
	OperTime      time.Time `json:"operTime" gorm:"type:timestamp;comment:操作时间"`
	JsonResult    string    `json:"jsonResult" gorm:"type:json;comment:返回数据"`
	LatencyTime   string    `json:"latencyTime" gorm:"size:128;comment:耗时"`
	UserAgent     string    `json:"userAgent" gorm:"size:255;comment:ua"`
	ModelTime
	ControlBy
}

func (SysRequestLog) TableName() string {
	return "sys_request_log"
}
