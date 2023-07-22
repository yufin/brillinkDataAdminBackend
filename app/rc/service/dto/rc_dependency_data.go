package dto

import (
	"go-admin/app/rc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type RcDependencyDataGetPageReq struct {
	dto.Pagination `search:"-"`

	ContentId int64  `form:"contentId"  search:"type:exact;column:content_id;table:rc_dependency_data" comment:"rc_origin_content.id"`
	UscId     string `form:"uscId"  search:"type:exact;column:usc_id;table:rc_dependency_data" comment:"统一信用社会代码"`
	LhQylx    *int   `form:"lhQylx"  search:"type:exact;column:lh_qylx;table:rc_dependency_data" comment:"企业类型"`
	LhCylwz   *int   `form:"lhCylwz"  search:"type:exact;column:lh_cylwz;table:rc_dependency_data" comment:"产业链位置"`
	LhGdct    *int   `form:"lhGdct"  search:"type:exact;column:lh_gdct;table:rc_dependency_data" comment:"股东穿透"`
	LhYhsx    *int   `form:"lhYhsx"  search:"type:exact;column:lh_yhsx;table:rc_dependency_data" comment:"银行授信"`
	LhSfsx    *int   `form:"lhSfsx"  search:"type:exact;column:lh_sfsx;table:rc_dependency_data" comment:"三方授信"`
	LhQybq    *int   `form:"lhQybq"  search:"type:exact;column:lh_qybq;table:rc_dependency_data" comment:"企业标签"`
	RcDependencyDataPageOrder
}

type RcDependencyDataPageOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:rc_dependency_data"`
	ContentId string `form:"contentIdOrder"  search:"type:order;column:content_id;table:rc_dependency_data"`
	UscId     string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:rc_dependency_data"`
	LhQylx    string `form:"lhQylxOrder"  search:"type:order;column:lh_qylx;table:rc_dependency_data"`
	LhCylwz   string `form:"lhCylwzOrder"  search:"type:order;column:lh_cylwz;table:rc_dependency_data"`
	LhGdct    string `form:"lhGdctOrder"  search:"type:order;column:lh_gdct;table:rc_dependency_data"`
	LhYhsx    string `form:"lhYhsxOrder"  search:"type:order;column:lh_yhsx;table:rc_dependency_data"`
	LhSfsx    string `form:"lhSfsxOrder"  search:"type:order;column:lh_sfsx;table:rc_dependency_data"`
	LhQybq    string `form:"lhQybqOrder"  search:"type:order;column:lh_qybq;table:rc_dependency_data"`
}

func (m *RcDependencyDataGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RcDependencyDataGetResp struct {
	Id        int64  `json:"id"`        // 主键
	ContentId int64  `json:"contentId"` // rc_origin_content.id
	UscId     string `json:"uscId"`     // 统一信用社会代码
	LhQylx    *int   `json:"lhQylx"`    // 企业类型
	LhCylwz   *int   `json:"lhCylwz"`   // 产业链位置
	LhGdct    *int   `json:"lhGdct"`    // 股东穿透
	LhYhsx    *int   `json:"lhYhsx"`    // 银行授信
	LhSfsx    *int   `json:"lhSfsx"`    // 三方授信
	LhQybq    *int   `json:"lhQybq"`    // 企业标签
	common.ControlBy
}

func (s *RcDependencyDataGetResp) Generate(model *models.RcDependencyData) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ContentId = model.ContentId
	s.UscId = model.UscId
	s.LhQylx = model.LhQylx
	s.LhCylwz = model.LhCylwz
	s.LhGdct = model.LhGdct
	s.LhYhsx = model.LhYhsx
	s.LhSfsx = model.LhSfsx
	s.LhQybq = model.LhQybq
	s.CreateBy = model.CreateBy
}

type RcDependencyDataInsertReq struct {
	Id        int64  `json:"-"`         // 主键
	ContentId int64  `json:"contentId"` // rc_origin_content.id
	UscId     string `json:"uscId"`     // 统一信用社会代码
	LhQylx    *int   `json:"lhQylx"`    // 企业类型
	LhCylwz   *int   `json:"lhCylwz"`   // 产业链位置
	LhGdct    *int   `json:"lhGdct"`    // 股东穿透
	LhYhsx    *int   `json:"lhYhsx"`    // 银行授信
	LhSfsx    *int   `json:"lhSfsx"`    // 三方授信
	LhQybq    *int   `json:"lhQybq"`    // 企业标签
	common.ControlBy
}

func (s *RcDependencyDataInsertReq) Generate(model *models.RcDependencyData) {
	if s.Id == 0 {
		model.Model = common.Model{Id: utils.NewFlakeId()}
	}
	model.ContentId = s.ContentId
	model.UscId = s.UscId
	model.LhQylx = s.LhQylx
	model.LhCylwz = s.LhCylwz
	model.LhGdct = s.LhGdct
	model.LhYhsx = s.LhYhsx
	model.LhSfsx = s.LhSfsx
	model.LhQybq = s.LhQybq
	model.CreateBy = s.CreateBy

}

func (s *RcDependencyDataInsertReq) GetId() interface{} {
	return s.Id
}

type RcDependencyDataUpdateReq struct {
	Id        int64  `uri:"id"`         // 主键
	ContentId int64  `json:"contentId"` // rc_origin_content.id
	UscId     string `json:"uscId"`     // 统一信用社会代码
	LhQylx    *int   `json:"lhQylx"`    // 企业类型
	LhCylwz   *int   `json:"lhCylwz"`   // 产业链位置
	LhGdct    *int   `json:"lhGdct"`    // 股东穿透
	LhYhsx    *int   `json:"lhYhsx"`    // 银行授信
	LhSfsx    *int   `json:"lhSfsx"`    // 三方授信
	LhQybq    *int   `json:"lhQybq"`    // 企业标签
	common.ControlBy
}

func (s *RcDependencyDataUpdateReq) Generate(model *models.RcDependencyData) {
	if s.Id != 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ContentId = s.ContentId
	model.UscId = s.UscId
	model.LhQylx = s.LhQylx
	model.LhCylwz = s.LhCylwz
	model.LhGdct = s.LhGdct
	model.LhYhsx = s.LhYhsx
	model.LhSfsx = s.LhSfsx
	model.LhQybq = s.LhQybq
	model.UpdateBy = s.UpdateBy
}

func (s *RcDependencyDataUpdateReq) GetId() interface{} {
	return s.Id
}

// RcDependencyDataGetReq 功能获取请求参数
type RcDependencyDataGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RcDependencyDataGetReq) GetId() interface{} {
	return s.Id
}

// RcDependencyDataDeleteReq 功能删除请求参数
type RcDependencyDataDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RcDependencyDataDeleteReq) GetId() interface{} {
	return s.Ids
}

type RcDependencyDataExport struct {
	Id        int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:主键" xlsx:"主键"`
	ContentId int64  `json:"contentId" gorm:"comment:rc_origin_content.id" xlsx:"rc_origin_content.id"`
	UscId     string `json:"uscId" gorm:"comment:统一信用社会代码" xlsx:"统一信用社会代码"`
	LhQylx    *int   `json:"lhQylx" gorm:"comment:企业类型" xlsx:"企业类型"`
	LhCylwz   *int   `json:"lhCylwz" gorm:"comment:产业链位置" xlsx:"产业链位置"`
	LhGdct    *int   `json:"lhGdct" gorm:"comment:股东穿透" xlsx:"股东穿透"`
	LhYhsx    *int   `json:"lhYhsx" gorm:"comment:银行授信" xlsx:"银行授信"`
	LhSfsx    *int   `json:"lhSfsx" gorm:"comment:三方授信" xlsx:"三方授信"`
	LhQybq    *int   `json:"lhQybq" gorm:"comment:企业标签" xlsx:"企业标签"`
}
