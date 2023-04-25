package models

import (
	"go-admin/common/models"
)

// EnterpriseWaitList 待爬取列表
type EnterpriseWaitList struct {
	models.Model
	EnterpriseName string `json:"enterpriseName" gorm:"comment:企业名称"`
	UscId          string `json:"uscId" gorm:"comment:社会统一信用代码"`
	Priority       int    `json:"priority" gorm:"comment:优先级"`
	QccUrl         string `json:"qccUrl" gorm:"comment:qcc主体网址"`
	StatusCode     int    `json:"statusCode" gorm:"comment:数据爬取状态码"`
	Source         string `json:"source" gorm:"comment:来源备注"`
	models.ModelTime
	models.ControlBy
}

func (*EnterpriseWaitList) TableName() string {
	return "enterprise_wait_list"
}

func (e *EnterpriseWaitList) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseWaitList) GetId() interface{} {
	return e.Id
}
