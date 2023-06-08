package models

import (
	"go-admin/common/models"
)

// RcSellingSta parse from content sellingSta
type RcSellingSta struct {
	models.Model
	ContentId int64  `json:"contentId" gorm:"comment:foreign key from RcOriginContent" xlsx:"foreign key from RcOriginContent"`
	Cgje      string `json:"cgje" gorm:"comment:cgje" xlsx:"cgje"`
	Jezb      string `json:"jezb" gorm:"comment:jezb" xlsx:"jezb"`
	Ssspdl    string `json:"ssspdl" gorm:"comment:ssspdl" xlsx:"ssspdl"`
	Ssspxl    string `json:"ssspxl" gorm:"comment:ssspxl" xlsx:"ssspxl"`
	Ssspzl    string `json:"ssspzl" gorm:"comment:ssspzl" xlsx:"ssspzl"`
	models.ModelTime
	models.ControlBy
}

func (*RcSellingSta) TableName() string {
	return "rc_selling_sta"
}

func (e *RcSellingSta) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcSellingSta) GetId() interface{} {
	return e.Id
}
