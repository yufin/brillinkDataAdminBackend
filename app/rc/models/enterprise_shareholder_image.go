package models

import "go-admin/common/models"

//create table enterprise_shareholder_image
//(
//id         bigint      not null
//primary key,
//usc_id     varchar(18) not null comment '企业社会统一信用代码',
//image_time datetime(3) not null,
//created_at datetime(3) null,
//updated_at datetime(3) null,
//deleted_at datetime(3) null,
//create_by  bigint      null,
//update_by  bigint      null
//);

type EnterpriseShareholderImage struct {
	models.Model
	UscId     string `json:"usc_id" gorm:"comment:企业社会统一信用代码;default:NULL;size:18;"`
	ImageTime string `json:"image_time" gorm:"comment:;default:NULL;"`
	models.ControlBy
	models.ModelTime
}

func (*EnterpriseShareholderImage) TableName() string {
	return "enterprise_shareholder_image"
}

func (e *EnterpriseShareholderImage) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseShareholderImage) GetId() interface{} {
	return e.Id
}
