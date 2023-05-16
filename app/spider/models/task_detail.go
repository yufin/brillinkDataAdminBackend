package models

import (
	"go-admin/common/models"
)

// TaskDetail 任务列表
type TaskDetail struct {
	models.Model
	UscId      string `json:"uscId" gorm:"comment:企业统一信用代码" xlsx:"企业统一信用代码"`
	StatusCode int    `json:"statusCode" gorm:"comment:状态码(1.未完成, 2.完成)" xlsx:"状态码(1.未完成, 2.完成)"`
	Topic      string `json:"topic" gorm:"comment:主题" xlsx:"主题"`
	Priority   int    `json:"priority" gorm:"comment:优先级" xlsx:"优先级"`
	Comment    string `json:"comment" gorm:"comment:备注" xlsx:"备注"`
	models.ModelTime
	models.ControlBy
}

func (*TaskDetail) TableName() string {
	return "task_detail"
}

func (e *TaskDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TaskDetail) GetId() interface{} {
	return e.Id
}
