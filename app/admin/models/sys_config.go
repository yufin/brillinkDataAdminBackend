package models

import (
	"go-admin/common/models"
)

type SysConfig struct {
	models.Model
	ConfigModule string `json:"configModule" gorm:"size:128;comment:ConfigModule"`
	ConfigName   string `json:"configName" gorm:"size:128;comment:配置项名称"`
	ConfigKey    string `json:"configKey" gorm:"size:128;unique;comment:配置key"`
	ConfigValue  string `json:"configValue" gorm:"size:255;comment:配置value"`
	ConfigType   string `json:"configType" gorm:"size:64;comment:配置类型"`
	IsFrontend   bool   `json:"isFrontend" gorm:"comment:是否前台"`
	IsSecret     bool   `json:"isSecret" gorm:"size:64;comment:是否密钥"`
	Remark       string `json:"remark" gorm:"size:128;comment:备注"`
	models.ControlBy
	models.ModelTime
}

func (SysConfig) TableName() string {
	return "sys_config"
}

func (e *SysConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysConfig) GetId() interface{} {
	return e.Id
}
