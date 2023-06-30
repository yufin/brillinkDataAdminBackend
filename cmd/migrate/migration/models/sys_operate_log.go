package models

import (
	"go-admin/common/models"
)

// SysOperateLog 操作日志
type SysOperateLog struct {
	LogId        int64  `json:"logId" gorm:"primaryKey;autoIncrement;comment:编码"`
	Type         string `json:"type" gorm:"size:128;comment:操作类型"`
	Description  string `json:"description" gorm:"size:128;comment:操作说明"`
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
