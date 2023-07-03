package models

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/storage"
	"github.com/pkg/errors"
	"go-admin/common/models"
)

// SysOperateLog 操作日志
type SysOperateLog struct {
	LogId        int64  `json:"logId" gorm:"primaryKey;autoIncrement;comment:编码"`
	Type         string `json:"type" gorm:"size:128;comment:操作类型"`
	Description  string `json:"description" gorm:"size:128;comment:操作说明"`
	Project      string `json:"project" gorm:"size:128;comment:项目"`
	UserName     string `json:"userName" gorm:"size:128;comment:用户"`
	UserId       int64  `json:"userId" gorm:"size:11;comment:用户id"`
	UpdateBefore string `json:"updateBefore" gorm:"type:json;comment:更新前"`
	UpdateAfter  string `json:"updateAfter" gorm:"type:json;comment:更新后"`
	models.ModelTime
	models.ControlBy
}

func (SysOperateLog) TableName() string {
	return "sys_operate_log"
}

func (e *SysOperateLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysOperateLog) GetId() interface{} {
	return e.LogId
}

// SaveOperateLog 从队列中获取操作日志
func SaveOperateLog(message storage.Messager) (err error) {
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
	var l SysOperateLog
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
