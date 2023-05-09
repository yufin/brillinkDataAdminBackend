package models

import (
	"go-admin/common/models"
)

// EnterpriseWaitList 待爬取列表
type EnterpriseWaitList struct {
	models.Model
	EnterpriseName string `json:"enterpriseName" gorm:"default:null;comment:企业名称"`
	UscId          string `json:"uscId" gorm:"default:null;comment:社会统一信用代码"`
	Priority       int    `json:"priority" gorm:"default:1;comment:优先级"`
	QccUrl         string `json:"qccUrl" gorm:"default:'-';comment:qcc主体网址"`
	Source         string `json:"source" gorm:"default:'-';comment:来源备注"`
	StatusCode     int    `json:"statusCode" gorm:"default:1;comment:数据爬取状态码,1.待匹配qccUrl&uscId,2.待爬取,3.爬取完成,-1.爬取失败,9非法公司(自动忽略)"`
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
