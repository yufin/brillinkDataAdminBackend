package models

import (
	"github.com/shopspring/decimal"
	"go-admin/common/models"
	"time"
)

//create table rc_revenue_detail
//(
//id           bigint         not null
//primary key,
//content_id   bigint         null,
//seq          int            null comment '用于json序列化的排序',
//field        varchar(122)   null,
//val          decimal(32, 5) null,
//period_start date           null,
//period_end   date           null,
//created_at   datetime(3)    null,
//updated_at   datetime(3)    null,
//deleted_at   datetime(3)    null,
//update_by    bigint         null,
//create_by    bigint         null
//)
//comment 'parse from lrbDetail';

type RcRevenueDetail struct {
	models.Model
	ContentId   int64               `json:"contentId" gorm:"column:content_id;comment:content_id"`
	Seq         *int                `json:"seq" gorm:"column:seq;comment:用于json序列化的排序"`
	Field       string              `json:"field" gorm:"column:field;comment:字段名"`
	Val         decimal.NullDecimal `json:"val" gorm:"column:val;comment:字段值"`
	PeriodStart time.Time           `json:"periodStart;type:date;column:period_start;comment:期间开始"`
	PeriodEnd   time.Time           `json:"periodEnd;type:date;column:period_end;comment:期间结束"`
	models.ControlBy
	models.ModelTime
}

func (*RcRevenueDetail) TableName() string {
	return "rc_revenue_detail"
}

func (e *RcRevenueDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcRevenueDetail) GetId() interface{} {
	return e.Id
}
