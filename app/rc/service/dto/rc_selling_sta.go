package dto

import (
	"go-admin/app/rc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type RcSellingStaGetPageReq struct {
	dto.Pagination `search:"-"`

	ContentId int64  `form:"contentId"  search:"type:exact;column:content_id;table:rc_selling_sta" comment:"foreign key from RcOriginContent"`
	Cgje      string `form:"cgje"  search:"type:exact;column:cgje;table:rc_selling_sta" comment:"cgje"`
	Jezb      string `form:"jezb"  search:"type:exact;column:jezb;table:rc_selling_sta" comment:"jezb"`
	Ssspdl    string `form:"ssspdl"  search:"type:exact;column:ssspdl;table:rc_selling_sta" comment:"ssspdl"`
	Ssspxl    string `form:"ssspxl"  search:"type:exact;column:ssspxl;table:rc_selling_sta" comment:"ssspxl"`
	Ssspzl    string `form:"ssspzl"  search:"type:exact;column:ssspzl;table:rc_selling_sta" comment:"ssspzl"`
	RcSellingStaPageOrder
}

type RcSellingStaPageOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:rc_selling_sta"`
	ContentId string `form:"contentIdOrder"  search:"type:order;column:content_id;table:rc_selling_sta"`
	Cgje      string `form:"cgjeOrder"  search:"type:order;column:cgje;table:rc_selling_sta"`
	Jezb      string `form:"jezbOrder"  search:"type:order;column:jezb;table:rc_selling_sta"`
	Ssspdl    string `form:"ssspdlOrder"  search:"type:order;column:ssspdl;table:rc_selling_sta"`
	Ssspxl    string `form:"ssspxlOrder"  search:"type:order;column:ssspxl;table:rc_selling_sta"`
	Ssspzl    string `form:"ssspzlOrder"  search:"type:order;column:ssspzl;table:rc_selling_sta"`
}

func (m *RcSellingStaGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RcSellingStaGetResp struct {
	Id        int64  `json:"id"`        // 主键
	ContentId int64  `json:"contentId"` // foreign key from RcOriginContent
	Cgje      string `json:"cgje"`      // cgje
	Jezb      string `json:"jezb"`      // jezb
	Ssspdl    string `json:"ssspdl"`    // ssspdl
	Ssspxl    string `json:"ssspxl"`    // ssspxl
	Ssspzl    string `json:"ssspzl"`    // ssspzl
	common.ControlBy
}

func (s *RcSellingStaGetResp) Generate(model *models.RcSellingSta) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ContentId = model.ContentId
	s.Cgje = model.Cgje
	s.Jezb = model.Jezb
	s.Ssspdl = model.Ssspdl
	s.Ssspxl = model.Ssspxl
	s.Ssspzl = model.Ssspzl
	s.CreateBy = model.CreateBy
}

type RcSellingStaInsertReq struct {
	Id        int64  `json:"-"`         // 主键
	ContentId int64  `json:"contentId"` // foreign key from RcOriginContent
	Cgje      string `json:"cgje"`      // cgje
	Jezb      string `json:"jezb"`      // jezb
	Ssspdl    string `json:"ssspdl"`    // ssspdl
	Ssspxl    string `json:"ssspxl"`    // ssspxl
	Ssspzl    string `json:"ssspzl"`    // ssspzl
	common.ControlBy
}

func (s *RcSellingStaInsertReq) Generate(model *models.RcSellingSta) {
	if s.Id == 0 {
		model.Model = common.Model{Id: utils.NewFlakeId()}
	}
	model.ContentId = s.ContentId
	model.Cgje = s.Cgje
	model.Jezb = s.Jezb
	model.Ssspdl = s.Ssspdl
	model.Ssspxl = s.Ssspxl
	model.Ssspzl = s.Ssspzl
	model.CreateBy = s.CreateBy
}

func (s *RcSellingStaInsertReq) GetId() interface{} {
	return s.Id
}

type RcSellingStaUpdateReq struct {
	Id        int64  `uri:"id"`         // 主键
	ContentId int64  `json:"contentId"` // foreign key from RcOriginContent
	Cgje      string `json:"cgje"`      // cgje
	Jezb      string `json:"jezb"`      // jezb
	Ssspdl    string `json:"ssspdl"`    // ssspdl
	Ssspxl    string `json:"ssspxl"`    // ssspxl
	Ssspzl    string `json:"ssspzl"`    // ssspzl
	common.ControlBy
}

func (s *RcSellingStaUpdateReq) Generate(model *models.RcSellingSta) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ContentId = s.ContentId
	model.Cgje = s.Cgje
	model.Jezb = s.Jezb
	model.Ssspdl = s.Ssspdl
	model.Ssspxl = s.Ssspxl
	model.Ssspzl = s.Ssspzl
	model.UpdateBy = s.UpdateBy
}

func (s *RcSellingStaUpdateReq) GetId() interface{} {
	return s.Id
}

// RcSellingStaGetReq 功能获取请求参数
type RcSellingStaGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RcSellingStaGetReq) GetId() interface{} {
	return s.Id
}

// RcSellingStaDeleteReq 功能删除请求参数
type RcSellingStaDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RcSellingStaDeleteReq) GetId() interface{} {
	return s.Ids
}

type RcSellingStaExport struct {
	Id        int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:主键" xlsx:"主键"`
	ContentId int64  `json:"contentId" gorm:"comment:foreign key from RcOriginContent" xlsx:"foreign key from RcOriginContent"`
	Cgje      string `json:"cgje" gorm:"comment:cgje" xlsx:"cgje"`
	Jezb      string `json:"jezb" gorm:"comment:jezb" xlsx:"jezb"`
	Ssspdl    string `json:"ssspdl" gorm:"comment:ssspdl" xlsx:"ssspdl"`
	Ssspxl    string `json:"ssspxl" gorm:"comment:ssspxl" xlsx:"ssspxl"`
	Ssspzl    string `json:"ssspzl" gorm:"comment:ssspzl" xlsx:"ssspzl"`
}
