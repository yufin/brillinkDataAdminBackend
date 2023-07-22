package models

import "go-admin/common/models"

type RcReportOss struct {
	models.Model
	DepId   int64 `json:"depId" gorm:"comment:rc_report_oss.dep_id" xlsx:"rc_report_oss.dep_id"`
	OssId   int64 `json:"ossId" gorm:"comment:oss_id" xlsx:"oss_id"`
	Version int   `json:"version" gorm:"comment:version of report" xlsx:"version"`
	models.ModelTime
	models.ControlBy
}

func (*RcReportOss) TableName() string {
	return "rc_report_oss"
}

func (e *RcReportOss) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcReportOss) GetId() interface{} {
	return e.Id
}
