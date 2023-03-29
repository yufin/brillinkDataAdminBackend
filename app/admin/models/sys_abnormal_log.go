package models

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/storage"
	"github.com/pkg/errors"
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

// SaveAbnormalLog 从队列中获取操作日志
func SaveAbnormalLog(message storage.Messager) (err error) {
	//准备db
	db := sdk.Runtime.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var l SysAbnormalLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	if l.Body == "" {
		l.Body = "[]"
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	return nil
}
