package models

import "time"

// TbNotice 通知
type TbNotice struct {
	// Avatar 用户头像地址 适用于：event、message、notification
	Avatar string `json:"avatar,omitempty" gorm:"size:512;comment:用户头像地址 适用于：event、message、notification"`
	// ClickClose 点击 适用于：message
	ClickClose bool `json:"clickClose,omitempty" gorm:"size:512;comment:点击 适用于：message"`
	// Datetime 时间 适用于：event、message、notification
	Datetime time.Time `json:"datetime,omitempty" gorm:"size:512;comment:时间 适用于：event、message、notification"`
	// Description 描述信息 适用于：event、message
	Description string `json:"description,omitempty" gorm:"size:512;comment:描述信息 适用于：event、message"`
	// Id 编码 适用于：event、message、notification
	Id string `json:"id,omitempty" gorm:"size:512;comment:编码 适用于：event、message、notification"`
	// Title 标题 适用于：event、message、notification
	Title string `json:"title,omitempty" gorm:"size:512;comment:标题 适用于：event、message、notification"`
	// Type 类型 适用于：event、message、notification
	Type string `json:"type,omitempty" gorm:"size:512;comment:类型 适用于：event、message、notification"`
	// Extra 扩展信息 适用于：event
	Extra string `json:"extra,omitempty" gorm:"size:512;comment:扩展信息 适用于：event"`
	// Status 事件状态 适用于：event
	Status string `json:"status,omitempty" gorm:"size:512;comment:事件状态 适用于：event"`
	// Target 目标
	Target string `json:"-" gorm:"size:128;comment:目标"`
	// IsRead 是否已读
	IsRead int `json:"-" gorm:"size:4;comment:是否已读"`
	ControlBy
	ModelTime
}

func (TbNotice) TableName() string {
	return "tb_notice"
}
