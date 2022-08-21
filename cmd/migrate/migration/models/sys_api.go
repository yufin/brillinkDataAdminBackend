package models

type SysApi struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle    string `json:"handle" gorm:"size:128;comment:handle"`
	Name      string `json:"name" gorm:"size:128;comment:标题"`
	Path      string `json:"path" gorm:"size:128;comment:地址"`
	Type      string `json:"type" gorm:"size:16;comment:接口类型"`
	Method    string `json:"method" gorm:"size:16;comment:请求类型"`
	Project   string `json:"project" gorm:"size:128;comment:项目"`
	Bus       string `json:"bus" gorm:"size:128;comment:业务模块"`
	Func      string `json:"func" gorm:"size:128;comment:func"`
	IsHistory bool   `json:"isHistory" gorm:"comment:是否历史接口"`
	ModelTime
	ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}
