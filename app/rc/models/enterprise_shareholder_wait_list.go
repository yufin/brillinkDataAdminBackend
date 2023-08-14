package models

import (
	"go-admin/common/models"
	"time"
)

//create table enterprise_shareholder_wait_list
//(
//id               bigint      not null
//primary key,
//image_id         bigint      null,
//usc_id           char(18)    null,
//valid_start_time datetime(3) null,
//valid_end_time   datetime(3) null,
//created_at       datetime(3) null,
//updated_at       datetime(3) null,
//deleted_at       datetime(3) null,
//create_by        bigint      null,
//update_by        bigint      null
//);

type EnterpriseShareholderWaitList struct {
	models.Model
	ImageId        *int64    `json:"image_id" gorm:"comment:;default:NULL;"`
	UscId          string    `json:"usc_id" gorm:"comment:;default:NULL;size:18;"`
	ValidStartTime time.Time `json:"valid_start_time" gorm:"comment:;default:NULL;"`
	ValidEndTime   time.Time `json:"valid_end_time" gorm:"comment:;default:NULL;"`
	models.ControlBy
	models.ModelTime
}

func (*EnterpriseShareholderWaitList) TableName() string {
	return "enterprise_shareholder_wait_list"
}

func (e *EnterpriseShareholderWaitList) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseShareholderWaitList) GetId() interface{} {
	return e.Id
}
