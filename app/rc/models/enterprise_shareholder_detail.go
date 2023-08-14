package models

import "go-admin/common/models"

//create table enterprise_shareholder_detail
//(
//id                bigint          not null
//primary key,
//image_id          bigint          null,
//shareholder_type  varchar(122)    null,
//credit_no         varchar(52)     null,
//enterprise_name   varchar(256)    null,
//shareholding_rate decimal(18, 10) null comment '持股占比: (例如持股50%: 50) ',
//capital_paid_in   decimal(28, 5)  null comment '认缴出资额（万元）',
//capital_paid_date datetime        null comment '认缴出资日期',
//created_at        datetime(3)     null,
//updated_at        datetime(3)     null,
//deleted_at        datetime(3)     null,
//create_by         bigint          null,
//update_by         bigint          null
//);

type EnterpriseShareholderDetail struct {
	models.Model
	ImageId          *int64  `json:"image_id" gorm:"comment:;default:NULL;"`
	ShareholderType  string  `json:"shareholder_type" gorm:"comment:;default:NULL;size:122;"`
	CreditNo         string  `json:"credit_no" gorm:"comment:;default:NULL;size:52;"`
	EnterpriseName   string  `json:"enterprise_name" gorm:"comment:;default:NULL;size:256;"`
	ShareholdingRate float64 `json:"shareholding_rate" gorm:"comment:持股占比: (例如持股50%: 50);default:NULL;type:decimal(18,10);"`
	CapitalPaidIn    float64 `json:"capital_paid_in" gorm:"comment:认缴出资额（万元）;default:NULL;type:decimal(28,5);"`
	CapitalPaidDate  string  `json:"capital_paid_date" gorm:"comment:认缴出资日期;default:NULL;"`
	models.ControlBy
	models.ModelTime
}

func (*EnterpriseShareholderDetail) TableName() string {
	return "enterprise_shareholder_detail"
}

func (e *EnterpriseShareholderDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseShareholderDetail) GetId() interface{} {
	return e.Id
}
