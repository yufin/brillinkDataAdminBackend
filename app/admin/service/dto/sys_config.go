package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysConfigGetPageReq 列表或者搜索使用结构体
type SysConfigGetPageReq struct {
	dto.Pagination `search:"-"`
	ConfigModule   string `form:"configModule" search:"type:contains;column:config_module;table:sys_config"`
	ConfigName     string `form:"configName" search:"type:contains;column:config_name;table:sys_config"`
	ConfigKey      string `form:"configKey" search:"type:contains;column:config_key;table:sys_config"`
	ConfigType     string `form:"configType" search:"type:exact;column:config_type;table:sys_config"`
	IsFrontend     int    `form:"isFrontend" search:"type:exact;column:is_frontend;table:sys_config"`
	SysConfigOrder
}

type SysConfigOrder struct {
	IdOrder         string `search:"type:order;column:id;table:sys_config" form:"idOrder"`
	ConfigNameOrder string `search:"type:order;column:config_name;table:sys_config" form:"configNameOrder"`
	ConfigKeyOrder  string `search:"type:order;column:config_key;table:sys_config" form:"configKeyOrder"`
	ConfigTypeOrder string `search:"type:order;column:config_type;table:sys_config" form:"configTypeOrder"`
	CreatedAtOrder  string `search:"type:order;column:created_at;table:sys_config" form:"createdAtOrder"`
}

func (m *SysConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysConfigGetToSysAppReq struct {
	IsFrontend int `form:"isFrontend" search:"type:exact;column:is_frontend;table:sys_config"`
}

func (m *SysConfigGetToSysAppReq) GetNeedSearch() interface{} {
	return *m
}

type SysConfigGetPageResp struct {
	Id           int64  `json:"id" comment:"编码"` // 编码
	ConfigModule string `json:"configModule"`
	ConfigName   string `json:"configName" comment:""`
	ConfigKey    string `json:"configKey" comment:""`
	ConfigValue  string `json:"configValue" comment:""`
	ConfigType   string `json:"configType" comment:""`
	IsFrontend   bool   `json:"isFrontend"`
	IsSecret     bool   `json:"isSecret"`
	Remark       string `json:"remark" comment:""`
	common.ModelTime
	common.ControlBy
}

func (s *SysConfigGetPageResp) Generate(modelList *[]models.SysConfig) (list *[]SysConfigGetPageResp) {
	l := make([]SysConfigGetPageResp, 0)
	for _, model := range *modelList {
		e := SysConfigGetPageResp{}
		e.Id = model.Id
		e.ConfigModule = model.ConfigModule
		e.ConfigName = model.ConfigName
		e.ConfigKey = model.ConfigKey
		e.ConfigValue = model.ConfigValue
		e.ConfigType = model.ConfigType
		if model.IsSecret {
			e.ConfigValue = "*********"
		}
		e.IsFrontend = model.IsFrontend
		e.IsSecret = model.IsSecret
		e.Remark = model.Remark
		l = append(l, e)
	}
	return &l
}

// SysConfigGetResp 查询
type SysConfigGetResp struct {
	Id           int64  `json:"id" comment:"编码"` // 编码
	ConfigModule string `json:"configModule"`
	ConfigName   string `json:"configName" comment:""`
	ConfigKey    string `json:"configKey" comment:""`
	ConfigValue  string `json:"configValue" comment:""`
	ConfigType   string `json:"configType" comment:""`
	IsFrontend   bool   `json:"isFrontend"`
	IsSecret     bool   `json:"isSecret"`
	Remark       string `json:"remark" comment:""`
	common.ControlBy
}

// Generate 结构体数据转化 从 SysConfigControl 至 system.SysConfig 对应的模型
func (s *SysConfigGetResp) Generate(model *models.SysConfig) {
	s.Id = model.Id
	s.ConfigModule = model.ConfigModule
	s.ConfigName = model.ConfigName
	s.ConfigKey = model.ConfigKey
	s.ConfigValue = model.ConfigValue
	s.ConfigType = model.ConfigType
	if model.IsSecret {
		s.ConfigValue = "*********"
	}
	s.IsFrontend = model.IsFrontend
	s.IsSecret = model.IsSecret
	s.Remark = model.Remark
}

// SysConfigControl 增、改使用的结构体
type SysConfigControl struct {
	Id           int64  `uri:"Id" comment:"编码"` // 编码
	ConfigModule string `json:"configModule"`
	ConfigName   string `json:"configName" comment:""`
	ConfigKey    string `uri:"configKey" json:"configKey" comment:""`
	ConfigValue  string `json:"configValue" comment:""`
	ConfigType   string `json:"configType" comment:""`
	IsFrontend   bool   `json:"isFrontend"`
	IsSecret     bool   `json:"isSecret"`
	Remark       string `json:"remark" comment:""`
	common.ControlBy
}

// Generate 结构体数据转化 从 SysConfigControl 至 system.SysConfig 对应的模型
func (s *SysConfigControl) Generate(model *models.SysConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ConfigModule = s.ConfigModule
	model.ConfigName = s.ConfigName
	model.ConfigKey = s.ConfigKey
	model.ConfigValue = s.ConfigValue
	model.ConfigType = s.ConfigType
	model.IsFrontend = s.IsFrontend
	model.IsSecret = s.IsSecret
	model.Remark = s.Remark

}

// GetId 获取数据对应的ID
func (s *SysConfigControl) GetId() interface{} {
	return s.Id
}

type GetSetSysConfigCustomResp struct {
	Id          int64  `json:"id,omitempty"`
	ConfigName  string `json:"configName" comment:""`
	ConfigKey   string `json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
	IsFrontend  bool   `json:"isFrontend"`
	IsSecret    int    `json:"isSecret,omitempty"`
}

// GetSetSysConfigReq 增、改使用的结构体
type GetSetSysConfigReq struct {
	ConfigKey   string `json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
	IsSecret    int    `json:"-"`
}

// Generate 结构体数据转化 从 SysConfigControl 至 system.SysConfig 对应的模型
func (s *GetSetSysConfigReq) Generate(model *models.SysConfig) {
	model.ConfigValue = s.ConfigValue
}

type UpdateSetSysConfigReq struct {
	ConfigKey   string      `json:"configKey" comment:""`
	ConfigValue interface{} `json:"configValue" comment:""`
	IsSecret    int         `json:"-"`
	Custom      string      `json:"custom"`
}

//type UpdateSetSysConfigReq map[string]string

type UpdateSetSysConfigCustomReq struct {
	Id          int64  `json:"-"`
	ConfigName  string `json:"configName" comment:""`
	ConfigKey   string `json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
	IsFrontend  bool   `json:"isFrontend,omitempty"`
	IsSecret    bool   `json:"isSecret,omitempty"`
}

func (s *UpdateSetSysConfigCustomReq) Generate(model *models.SysConfig) {
	model.ConfigModule = "custom"
	model.ConfigName = s.ConfigName
	model.ConfigKey = s.ConfigKey
	model.ConfigValue = s.ConfigValue
	model.ConfigType = "N"
	model.IsFrontend = s.IsFrontend
	model.IsSecret = s.IsSecret

}

// SysConfigByKeyReq 根据Key获取配置
type SysConfigByKeyReq struct {
	ConfigKey string `uri:"configKey" search:"type:contains;column:config_key;table:sys_config"`
}

func (m *SysConfigByKeyReq) GetNeedSearch() interface{} {
	return *m
}

type GetSysConfigByKEYForServiceResp struct {
	ConfigKey   string `json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
}

type SysConfigGetReq struct {
	Id int `uri:"id"`
}

func (s *SysConfigGetReq) GetId() interface{} {
	return s.Id
}

type SysConfigDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
