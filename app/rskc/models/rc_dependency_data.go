package models

import (
	"go-admin/common/models"
)

// RcDependencyData 依赖数据
type RcDependencyData struct {
	models.Model
	ContentId int64  `json:"contentId" gorm:"comment:rskc_origin_content.id" xlsx:"rskc_origin_content.id"`
	UscId     string `json:"uscId" gorm:"comment:统一信用社会代码" xlsx:"统一信用社会代码"`
	LhQylx    *int   `json:"lhQylx" gorm:"comment:企业类型" xlsx:"企业类型"`
	LhCylwz   *int   `json:"lhCylwz" gorm:"comment:产业链位置" xlsx:"产业链位置"`
	LhGdct    *int   `json:"lhGdct" gorm:"comment:股东穿透" xlsx:"股东穿透"`
	LhYhsx    *int   `json:"lhYhsx" gorm:"comment:银行授信" xlsx:"银行授信"`
	LhSfsx    *int   `json:"lhSfsx" gorm:"comment:三方授信" xlsx:"三方授信"`
	LhQybq    *int   `json:"lhQybq" gorm:"comment:企业标签" xlsx:"企业标签"`
	models.ModelTime
	models.ControlBy
}

func (*RcDependencyData) TableName() string {
	return "rc_dependency_data"
}

func (e *RcDependencyData) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcDependencyData) GetId() interface{} {
	return e.Id
}
