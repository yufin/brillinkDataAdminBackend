package models

import (
	"go-admin/common/models"
)

// SysAbnormalLog 异常日志
type SysAbnormalLog struct {
	AbId       int64  `json:"abId" gorm:"primaryKey;autoIncrement;comment:编码"`
	Method     string `json:"method" gorm:"size:128;comment:请求方式"`
	Url        string `json:"url" gorm:"size:128;comment:请求地址"`
	Ip         string `json:"ip" gorm:"size:128;comment:ip"`
	AbInfo     string `json:"abInfo" gorm:"size:256;comment:异常信息"`
	AbSource   string `json:"abSource" gorm:"size:256;comment:异常来源"`
	AbFunc     string `json:"abFunc" gorm:"size:256;comment:异常方法"`
	UserId     int64  `json:"userId" gorm:"size:10;comment:用户id"`
	UserName   string `json:"userName" gorm:"size:128;comment:操作人"`
	Headers    string `json:"headers" gorm:"type:json;comment:请求头"`
	Body       string `json:"body" gorm:"type:json;comment:请求数据"`
	Resp       string `json:"resp" gorm:"type:json;comment:回调数据"`
	StackTrace string `json:"stackTrace" gorm:"type:text;comment:堆栈追踪"`
	models.ModelTime
	models.ControlBy
}

func (SysAbnormalLog) TableName() string {
	return "sys_abnormal_log"
}

func (e *SysAbnormalLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysAbnormalLog) GetId() interface{} {
	return e.AbId
}
