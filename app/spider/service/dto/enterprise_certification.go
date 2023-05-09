package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
	"time"
)

type EnterpriseCertificationGetPageReq struct {
	dto.Pagination `search:"-"`

	CertId                 int64     `form:"certId"  search:"type:exact;column:cert_id;table:enterprise_certification" comment:"主键"`
	CertificationTitle     string    `form:"certificationTitle"  search:"type:exact;column:certification_title;table:enterprise_certification" comment:"认证名称"`
	CertificationCode      string    `form:"certificationCode"  search:"type:exact;column:certification_code;table:enterprise_certification" comment:"认证编号"`
	CertificationLevel     string    `form:"certificationLevel"  search:"type:exact;column:certification_level;table:enterprise_certification" comment:"认证等级(省级,市级,国家级)"`
	CertificationType      string    `form:"certificationType"  search:"type:exact;column:certification_type;table:enterprise_certification" comment:"认证类型(荣誉,科技型企业,...)"`
	CertificationSource    string    `form:"certificationSource"  search:"type:exact;column:certification_source;table:enterprise_certification" comment:"认证来源(eg:2022年度浙江省中小企业公共服务示范平台名单)"`
	CertificationDate      time.Time `form:"certificationDate"  search:"type:exact;column:certification_date;table:enterprise_certification" comment:"发证日期"`
	CertificationTermStart time.Time `form:"certificationTermStart"  search:"type:exact;column:certification_term_start;table:enterprise_certification" comment:"有效期起"`
	CertificationTermEnd   time.Time `form:"certificationTermEnd"  search:"type:exact;column:certification_term_end;table:enterprise_certification" comment:"有效期至"`
	CertificationAuthority string    `form:"certificationAuthority"  search:"type:exact;column:certification_authority;table:enterprise_certification" comment:"发证机关"`
	UscId                  string    `form:"uscId"  search:"type:exact;column:usc_id;table:enterprise_certification" comment:"社会统一信用代码"`
	StatusCode             int64     `form:"statusCode"  search:"type:exact;column:status_code;table:enterprise_certification" comment:"状态标识码"`
	EnterpriseCertificationPageOrder
}

type EnterpriseCertificationPageOrder struct {
	CertId                 int64     `form:"certIdOrder"  search:"type:order;column:cert_id;table:enterprise_certification"`
	CertificationTitle     string    `form:"certificationTitleOrder"  search:"type:order;column:certification_title;table:enterprise_certification"`
	CertificationCode      string    `form:"certificationCodeOrder"  search:"type:order;column:certification_code;table:enterprise_certification"`
	CertificationLevel     string    `form:"certificationLevelOrder"  search:"type:order;column:certification_level;table:enterprise_certification"`
	CertificationType      string    `form:"certificationTypeOrder"  search:"type:order;column:certification_type;table:enterprise_certification"`
	CertificationSource    string    `form:"certificationSourceOrder"  search:"type:order;column:certification_source;table:enterprise_certification"`
	CertificationDate      time.Time `form:"certificationDateOrder"  search:"type:order;column:certification_date;table:enterprise_certification"`
	CertificationTermStart time.Time `form:"certificationTermStartOrder"  search:"type:order;column:certification_term_start;table:enterprise_certification"`
	CertificationTermEnd   time.Time `form:"certificationTermEndOrder"  search:"type:order;column:certification_term_end;table:enterprise_certification"`
	CertificationAuthority string    `form:"certificationAuthorityOrder"  search:"type:order;column:certification_authority;table:enterprise_certification"`
	UscId                  string    `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise_certification"`
	StatusCode             int64     `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_certification"`
}

func (m *EnterpriseCertificationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseCertificationGetResp struct {
	CertId                 int64     `json:"certId"`                 // 主键
	CertificationTitle     string    `json:"certificationTitle"`     // 认证名称
	CertificationCode      string    `json:"certificationCode"`      // 认证编号
	CertificationLevel     string    `json:"certificationLevel"`     // 认证等级(省级,市级,国家级)
	CertificationType      string    `json:"certificationType"`      // 认证类型(荣誉,科技型企业,...)
	CertificationSource    string    `json:"certificationSource"`    // 认证来源(eg:2022年度浙江省中小企业公共服务示范平台名单)
	CertificationDate      time.Time `json:"certificationDate"`      // 发证日期
	CertificationTermStart time.Time `json:"certificationTermStart"` // 有效期起
	CertificationTermEnd   time.Time `json:"certificationTermEnd"`   // 有效期至
	CertificationAuthority string    `json:"certificationAuthority"` // 发证机关
	UscId                  string    `json:"uscId"`                  // 社会统一信用代码
	StatusCode             int64     `json:"statusCode"`             // 状态标识码
	common.ControlBy
}

func (s *EnterpriseCertificationGetResp) Generate(model *models.EnterpriseCertification) {
	s.CertId = model.CertId
	s.CertificationTitle = model.CertificationTitle
	s.CertificationCode = model.CertificationCode
	s.CertificationLevel = model.CertificationLevel
	s.CertificationType = model.CertificationType
	s.CertificationSource = model.CertificationSource
	s.CertificationDate = model.CertificationDate
	s.CertificationTermStart = model.CertificationTermStart
	s.CertificationTermEnd = model.CertificationTermEnd
	s.CertificationAuthority = model.CertificationAuthority
	s.UscId = model.UscId
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type EnterpriseCertificationInsertReq struct {
	CertId                 int64     `json:"-"`                      // 主键
	CertificationTitle     string    `json:"certificationTitle"`     // 认证名称
	CertificationCode      string    `json:"certificationCode"`      // 认证编号
	CertificationLevel     string    `json:"certificationLevel"`     // 认证等级(省级,市级,国家级)
	CertificationType      string    `json:"certificationType"`      // 认证类型(荣誉,科技型企业,...)
	CertificationSource    string    `json:"certificationSource"`    // 认证来源(eg:2022年度浙江省中小企业公共服务示范平台名单)
	CertificationDate      time.Time `json:"certificationDate"`      // 发证日期
	CertificationTermStart time.Time `json:"certificationTermStart"` // 有效期起
	CertificationTermEnd   time.Time `json:"certificationTermEnd"`   // 有效期至
	CertificationAuthority string    `json:"certificationAuthority"` // 发证机关
	UscId                  string    `json:"uscId"`                  // 社会统一信用代码
	StatusCode             int64     `json:"statusCode"`             // 状态标识码
	common.ControlBy
}

func (s *EnterpriseCertificationInsertReq) Generate(model *models.EnterpriseCertification) {
	if s.CertId == 0 {
		s.CertId = utils.NewFlakeId()
	}
	model.CertId = s.CertId
	model.CertificationTitle = s.CertificationTitle
	model.CertificationCode = s.CertificationCode
	model.CertificationLevel = s.CertificationLevel
	model.CertificationType = s.CertificationType
	model.CertificationSource = s.CertificationSource
	model.CertificationDate = s.CertificationDate
	model.CertificationTermStart = s.CertificationTermStart
	model.CertificationTermEnd = s.CertificationTermEnd
	model.CertificationAuthority = s.CertificationAuthority
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseCertificationInsertReq) GetId() interface{} {
	return s.CertId
}

type EnterpriseCertificationUpdateReq struct {
	CertId                 int64     `uri:"certId"`                  // 主键
	CertificationTitle     string    `json:"certificationTitle"`     // 认证名称
	CertificationCode      string    `json:"certificationCode"`      // 认证编号
	CertificationLevel     string    `json:"certificationLevel"`     // 认证等级(省级,市级,国家级)
	CertificationType      string    `json:"certificationType"`      // 认证类型(荣誉,科技型企业,...)
	CertificationSource    string    `json:"certificationSource"`    // 认证来源(eg:2022年度浙江省中小企业公共服务示范平台名单)
	CertificationDate      time.Time `json:"certificationDate"`      // 发证日期
	CertificationTermStart time.Time `json:"certificationTermStart"` // 有效期起
	CertificationTermEnd   time.Time `json:"certificationTermEnd"`   // 有效期至
	CertificationAuthority string    `json:"certificationAuthority"` // 发证机关
	UscId                  string    `json:"uscId"`                  // 社会统一信用代码
	StatusCode             int64     `json:"statusCode"`             // 状态标识码
	common.ControlBy
}

func (s *EnterpriseCertificationUpdateReq) Generate(model *models.EnterpriseCertification) {
	model.CertId = s.CertId
	model.CertificationTitle = s.CertificationTitle
	model.CertificationCode = s.CertificationCode
	model.CertificationLevel = s.CertificationLevel
	model.CertificationType = s.CertificationType
	model.CertificationSource = s.CertificationSource
	model.CertificationDate = s.CertificationDate
	model.CertificationTermStart = s.CertificationTermStart
	model.CertificationTermEnd = s.CertificationTermEnd
	model.CertificationAuthority = s.CertificationAuthority
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseCertificationUpdateReq) GetId() interface{} {
	return s.CertId
}

// EnterpriseCertificationGetReq 功能获取请求参数
type EnterpriseCertificationGetReq struct {
	CertId int64 `uri:"certId"`
}

func (s *EnterpriseCertificationGetReq) GetId() interface{} {
	return s.CertId
}

// EnterpriseCertificationDeleteReq 功能删除请求参数
type EnterpriseCertificationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseCertificationDeleteReq) GetId() interface{} {
	return s.Ids
}
