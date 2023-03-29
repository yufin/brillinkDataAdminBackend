package models

import (
	"encoding/json"
	"github.com/pkg/errors"
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/storage"

	"go-admin/common/models"
)

type SysRequestLog struct {
	models.Model
	RequestMethod string    `json:"requestMethod" gorm:"size:128;comment:请求方式"`
	OperName      string    `json:"operName" gorm:"size:128;comment:操作者"`
	OperUrl       string    `json:"operUrl" gorm:"size:255;comment:访问地址"`
	OperIp        string    `json:"operIp" gorm:"size:128;comment:客户端ip"`
	OperLocation  string    `json:"operLocation" gorm:"size:128;comment:访问位置"`
	OperParam     string    `json:"operParam" gorm:"type:json;comment:请求参数"`
	OperHeaders   string    `json:"operHeaders" gorm:"type:json;comment:请求头"`
	OperTime      time.Time `json:"operTime" gorm:"comment:操作时间"`
	JsonResult    string    `json:"jsonResult" gorm:"type:json;comment:返回数据"`
	LatencyTime   string    `json:"latencyTime" gorm:"size:128;comment:耗时"`
	UserAgent     string    `json:"userAgent" gorm:"size:255;comment:ua"`
	models.ModelTime
	models.ControlBy
}

func (SysRequestLog) TableName() string {
	return "sys_request_log"
}

func (e *SysRequestLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRequestLog) GetId() interface{} {
	return e.Id
}

// SaveRequestLog 从队列中获取操作日志
func SaveRequestLog(message storage.Messager) (err error) {
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
	var l SysRequestLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	return nil
}
