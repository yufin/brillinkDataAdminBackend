package models

import "go-admin/common/models"

type RcRiskIndex struct {
	models.Model
	ContentId  int64  `json:"content_id"`
	RiskDec    string `json:"riskDec"`
	IndexDec   string `json:"indexDec"`
	IndexValue string `json:"indexValue"`
	IndexFlag  string `json:"indexFlag"`
	models.ModelTime
	models.ControlBy
}

func (*RcRiskIndex) TableName() string {
	return "rc_risk_index"
}

func (e *RcRiskIndex) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcRiskIndex) GetId() interface{} {
	return e.Id
}

//create table rc_risk_index
//(
//id         bigint       not null
//primary key,
//content_id bigint       null,
//risk_desc  text         null,
//`index`    varchar(255) null,
//value      varchar(255) null,
//flag       varchar(255) null,
//created_at datetime(3)  null,
//updated_at datetime(3)  null,
//deleted_at datetime(3)  null,
//create_by  bigint       null,
//update_by  int          null
//);
