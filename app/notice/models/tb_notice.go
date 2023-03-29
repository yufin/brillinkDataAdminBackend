package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/storage"
	"github.com/google/uuid"
	"go-admin/common/models"
	"time"
)

type NoticeType string

type EventStatus string

const (
	// Event 待办
	Event NoticeType = "event"
	// Message 消息
	Message NoticeType = "message"
	// Notification 通知
	Notification NoticeType = "notification"
)

// event status 枚举值
const (
	// Processing 进行中
	Processing EventStatus = "processing"
	// Todo 未开始
	Todo EventStatus = "todo"
	// Doing 耗时xx天 等同 Processing
	Doing EventStatus = "doing"
	// Urgent 马上到期
	Urgent EventStatus = "urgent"
)

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
	models.ModelTime
	models.ControlBy
}

func (TbNotice) TableName() string {
	return "tb_notice"
}

func (e *TbNotice) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbNotice) GetId() interface{} {
	return e.Id
}

func SaveTbNotice(message storage.Messager) (err error) {
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		err := fmt.Errorf("json Marshal error, %s", err.Error())
		return err
	}
	var notice TbNotice
	err = json.Unmarshal(rb, &notice)
	if err != nil {
		err := fmt.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}
	db := sdk.Runtime.GetDbByKey("*")

	err = db.Where(TbNotice{Title: notice.Title, Datetime: notice.Datetime, Id: uuid.New().String(), Type: notice.Type, Avatar: notice.Avatar}).
		FirstOrCreate(&TbNotice{}).
		//Update("handle", v.Handler).
		Error
	if err != nil {
		err := fmt.Errorf("Models SaveTbNotice error: %s \r\n ", err.Error())
		return err
	}

	return nil
}
