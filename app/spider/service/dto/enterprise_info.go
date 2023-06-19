package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
	"time"
)

type EnterpriseInfoGetPageReq struct {
	dto.Pagination `search:"-"`

	EnterpriseTitle               string     `form:"enterpriseTitle"  search:"type:exact;column:enterprise_title;table:enterprise_info" comment:"企业名称"`
	EnterpriseTitleEn             string     `form:"enterpriseTitleEn"  search:"type:exact;column:enterprise_title_en;table:enterprise_info" comment:"企业英文名称"`
	BusinessRegistrationNumber    string     `form:"businessRegistrationNumber"  search:"type:exact;column:business_registration_number;table:enterprise_info" comment:"工商注册号"`
	EstablishedDate               *time.Time `form:"establishedDate"  search:"type:exact;column:established_date;table:enterprise_info" comment:"成立日期"`
	Region                        string     `form:"region"  search:"type:exact;column:region;table:enterprise_info" comment:"所属地区"`
	ApprovedDate                  *time.Time `form:"approvedDate"  search:"type:exact;column:approved_date;table:enterprise_info" comment:"核准日期"`
	RegisteredAddress             string     `form:"registeredAddress"  search:"type:exact;column:registered_address;table:enterprise_info" comment:"注册地址"`
	RegisteredCapital             string     `form:"registeredCapital"  search:"type:exact;column:registered_capital;table:enterprise_info" comment:"注册资本币种"`
	PaidInCapital                 string     `form:"paidInCapital"  search:"type:exact;column:paid_in_capital;table:enterprise_info" comment:"实缴资本币种"`
	EnterpriseType                string     `form:"enterpriseType"  search:"type:exact;column:enterprise_type;table:enterprise_info" comment:"企业类型"`
	StuffSize                     string     `form:"stuffSize"  search:"type:exact;column:stuff_size;table:enterprise_info" comment:"人员规模"`
	StuffInsuredNumber            int        `form:"stuffInsuredNumber"  search:"type:exact;column:stuff_insured_number;table:enterprise_info" comment:"参保人数"`
	BusinessScope                 string     `form:"businessScope"  search:"type:exact;column:business_scope;table:enterprise_info" comment:"经营范围"`
	ImportExportQualificationCode string     `form:"importExportQualificationCode"  search:"type:exact;column:import_export_qualification_code;table:enterprise_info" comment:"进出口企业代码"`
	LegalRepresentative           string     `form:"legalRepresentative"  search:"type:exact;column:legal_representative;table:enterprise_info" comment:"法定代表人"`
	RegistrationAuthority         string     `form:"registrationAuthority"  search:"type:exact;column:registration_authority;table:enterprise_info" comment:"登记机关"`
	RegistrationStatus            string     `form:"registrationStatus"  search:"type:exact;column:registration_status;table:enterprise_info" comment:"登记状态"`
	TaxpayerQualification         string     `form:"taxpayerQualification"  search:"type:exact;column:taxpayer_qualification;table:enterprise_info" comment:"纳税人资质"`
	OrganizationCode              string     `form:"organizationCode"  search:"type:exact;column:organization_code;table:enterprise_info" comment:"组织机构代码"`
	UrlQcc                        string     `form:"urlQcc"  search:"type:exact;column:url_qcc;table:enterprise_info" comment:"企查查url"`
	UrlHomepage                   string     `form:"urlHomepage"  search:"type:exact;column:url_homepage;table:enterprise_info" comment:"官网url"`
	BusinessTermStart             *time.Time `form:"businessTermStart"  search:"type:exact;column:business_term_start;table:enterprise_info" comment:"营业期限开始"`
	BusinessTermEnd               *time.Time `form:"businessTermEnd"  search:"type:exact;column:business_term_end;table:enterprise_info" comment:"营业期限结束"`
	UscId                         string     `form:"uscId"  search:"type:exact;column:usc_id;table:enterprise_info" comment:"社会统一信用代码"`
	StatusCode                    int        `form:"statusCode"  search:"type:exact;column:status_code;table:enterprise_info" comment:"状态标识码"`
	EnterpriseInfoPageOrder
}

type EnterpriseInfoPageOrder struct {
	InfoId                        int64      `form:"infoIdOrder"  search:"type:order;column:info_id;table:enterprise_info"`
	EnterpriseTitle               string     `form:"enterpriseTitleOrder"  search:"type:order;column:enterprise_title;table:enterprise_info"`
	EnterpriseTitleEn             string     `form:"enterpriseTitleEnOrder"  search:"type:order;column:enterprise_title_en;table:enterprise_info"`
	BusinessRegistrationNumber    string     `form:"businessRegistrationNumberOrder"  search:"type:order;column:business_registration_number;table:enterprise_info"`
	EstablishedDate               *time.Time `form:"establishedDateOrder"  search:"type:order;column:established_date;table:enterprise_info"`
	Region                        string     `form:"regionOrder"  search:"type:order;column:region;table:enterprise_info"`
	ApprovedDate                  *time.Time `form:"approvedDateOrder"  search:"type:order;column:approved_date;table:enterprise_info"`
	RegisteredAddress             string     `form:"registeredAddressOrder"  search:"type:order;column:registered_address;table:enterprise_info"`
	RegisteredCapital             string     `form:"registeredCapitalOrder"  search:"type:order;column:registered_capital;table:enterprise_info"`
	PaidInCapital                 string     `form:"paidInCapitalOrder"  search:"type:order;column:paid_in_capital;table:enterprise_info"`
	EnterpriseType                string     `form:"enterpriseTypeOrder"  search:"type:order;column:enterprise_type;table:enterprise_info"`
	StuffSize                     string     `form:"stuffSizeOrder"  search:"type:order;column:stuff_size;table:enterprise_info"`
	StuffInsuredNumber            int        `form:"stuffInsuredNumberOrder"  search:"type:order;column:stuff_insured_number;table:enterprise_info"`
	BusinessScope                 string     `form:"businessScopeOrder"  search:"type:order;column:business_scope;table:enterprise_info"`
	ImportExportQualificationCode string     `form:"importExportQualificationCodeOrder"  search:"type:order;column:import_export_qualification_code;table:enterprise_info"`
	LegalRepresentative           string     `form:"legalRepresentativeOrder"  search:"type:order;column:legal_representative;table:enterprise_info"`
	RegistrationAuthority         string     `form:"registrationAuthorityOrder"  search:"type:order;column:registration_authority;table:enterprise_info"`
	RegistrationStatus            string     `form:"registrationStatusOrder"  search:"type:order;column:registration_status;table:enterprise_info"`
	TaxpayerQualification         string     `form:"taxpayerQualificationOrder"  search:"type:order;column:taxpayer_qualification;table:enterprise_info"`
	OrganizationCode              string     `form:"organizationCodeOrder"  search:"type:order;column:organization_code;table:enterprise_info"`
	UrlQcc                        string     `form:"urlQccOrder"  search:"type:order;column:url_qcc;table:enterprise_info"`
	UrlHomepage                   string     `form:"urlHomepageOrder"  search:"type:order;column:url_homepage;table:enterprise_info"`
	BusinessTermStart             *time.Time `form:"businessTermStartOrder"  search:"type:order;column:business_term_start;table:enterprise_info"`
	BusinessTermEnd               *time.Time `form:"businessTermEndOrder"  search:"type:order;column:business_term_end;table:enterprise_info"`
	UscId                         string     `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise_info"`
	StatusCode                    int        `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_info"`
}

func (m *EnterpriseInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseInfoGetResp struct {
	InfoId                        int64      `json:"infoId"`                        // 主键
	EnterpriseTitle               string     `json:"enterpriseTitle"`               // 企业名称
	EnterpriseTitleEn             string     `json:"enterpriseTitleEn"`             // 企业英文名称
	BusinessRegistrationNumber    string     `json:"businessRegistrationNumber"`    // 工商注册号
	EstablishedDate               *time.Time `json:"establishedDate"`               // 成立日期
	Region                        string     `json:"region"`                        // 所属地区
	ApprovedDate                  *time.Time `json:"approvedDate"`                  // 核准日期
	RegisteredAddress             string     `json:"registeredAddress"`             // 注册地址
	RegisteredCapital             string     `json:"registeredCapital"`             // 注册资本币种
	PaidInCapital                 string     `json:"paidInCapital"`                 // 实缴资本币种
	EnterpriseType                string     `json:"enterpriseType"`                // 企业类型
	StuffSize                     string     `json:"stuffSize"`                     // 人员规模
	StuffInsuredNumber            int        `json:"stuffInsuredNumber"`            // 参保人数
	BusinessScope                 string     `json:"businessScope"`                 // 经营范围
	ImportExportQualificationCode string     `json:"importExportQualificationCode"` // 进出口企业代码
	LegalRepresentative           string     `json:"legalRepresentative"`           // 法定代表人
	RegistrationAuthority         string     `json:"registrationAuthority"`         // 登记机关
	RegistrationStatus            string     `json:"registrationStatus"`            // 登记状态
	TaxpayerQualification         string     `json:"taxpayerQualification"`         // 纳税人资质
	OrganizationCode              string     `json:"organizationCode"`              // 组织机构代码
	UrlQcc                        string     `json:"urlQcc"`                        // 企查查url
	UrlHomepage                   string     `json:"urlHomepage"`                   // 官网url
	BusinessTermStart             *time.Time `json:"businessTermStart"`             // 营业期限开始
	BusinessTermEnd               *time.Time `json:"businessTermEnd"`               // 营业期限结束
	UscId                         string     `json:"uscId"`                         // 社会统一信用代码
	StatusCode                    int        `json:"statusCode"`                    // 状态标识码
	common.ControlBy
}

func (s *EnterpriseInfoGetResp) Generate(model *models.EnterpriseInfo) {
	s.InfoId = model.InfoId
	s.EnterpriseTitle = model.EnterpriseTitle
	s.EnterpriseTitleEn = model.EnterpriseTitleEn
	s.BusinessRegistrationNumber = model.BusinessRegistrationNumber
	s.EstablishedDate = model.EstablishedDate
	s.Region = model.Region
	s.ApprovedDate = model.ApprovedDate
	s.RegisteredAddress = model.RegisteredAddress
	s.RegisteredCapital = model.RegisteredCapital
	s.PaidInCapital = model.PaidInCapital
	s.EnterpriseType = model.EnterpriseType
	s.StuffSize = model.StuffSize
	s.StuffInsuredNumber = model.StuffInsuredNumber
	s.BusinessScope = model.BusinessScope
	s.ImportExportQualificationCode = model.ImportExportQualificationCode
	s.LegalRepresentative = model.LegalRepresentative
	s.RegistrationAuthority = model.RegistrationAuthority
	s.RegistrationStatus = model.RegistrationStatus
	s.TaxpayerQualification = model.TaxpayerQualification
	s.OrganizationCode = model.OrganizationCode
	s.UrlQcc = model.UrlQcc
	s.UrlHomepage = model.UrlHomepage
	s.BusinessTermStart = model.BusinessTermStart
	s.BusinessTermEnd = model.BusinessTermEnd
	s.UscId = model.UscId
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type EnterpriseInfoInsertReq struct {
	InfoId                        int64      `json:"-"`                             // 主键
	EnterpriseTitle               string     `json:"enterpriseTitle"`               // 企业名称
	EnterpriseTitleEn             string     `json:"enterpriseTitleEn"`             // 企业英文名称
	BusinessRegistrationNumber    string     `json:"businessRegistrationNumber"`    // 工商注册号
	EstablishedDate               *time.Time `json:"establishedDate"`               // 成立日期
	Region                        string     `json:"region"`                        // 所属地区
	ApprovedDate                  *time.Time `json:"approvedDate"`                  // 核准日期
	RegisteredAddress             string     `json:"registeredAddress"`             // 注册地址
	RegisteredCapital             string     `json:"registeredCapital"`             // 注册资本币种
	PaidInCapital                 string     `json:"paidInCapital"`                 // 实缴资本币种
	EnterpriseType                string     `json:"enterpriseType"`                // 企业类型
	StuffSize                     string     `json:"stuffSize"`                     // 人员规模
	StuffInsuredNumber            int        `json:"stuffInsuredNumber"`            // 参保人数
	BusinessScope                 string     `json:"businessScope"`                 // 经营范围
	ImportExportQualificationCode string     `json:"importExportQualificationCode"` // 进出口企业代码
	LegalRepresentative           string     `json:"legalRepresentative"`           // 法定代表人
	RegistrationAuthority         string     `json:"registrationAuthority"`         // 登记机关
	RegistrationStatus            string     `json:"registrationStatus"`            // 登记状态
	TaxpayerQualification         string     `json:"taxpayerQualification"`         // 纳税人资质
	OrganizationCode              string     `json:"organizationCode"`              // 组织机构代码
	UrlQcc                        string     `json:"urlQcc"`                        // 企查查url
	UrlHomepage                   string     `json:"urlHomepage"`                   // 官网url
	BusinessTermStart             *time.Time `json:"businessTermStart"`             // 营业期限开始
	BusinessTermEnd               *time.Time `json:"businessTermEnd"`               // 营业期限结束
	UscId                         string     `json:"uscId"`                         // 社会统一信用代码
	StatusCode                    int        `json:"statusCode"`                    // 状态标识码
	common.ControlBy
}

func (s *EnterpriseInfoInsertReq) Generate(model *models.EnterpriseInfo) {
	minTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if s.InfoId == 0 {
		s.InfoId = utils.NewFlakeId()
	}
	model.InfoId = s.InfoId
	if s.EnterpriseTitle != "" {
		model.EnterpriseTitle = s.EnterpriseTitle
	}
	if s.EnterpriseTitleEn != "" {
		model.EnterpriseTitleEn = s.EnterpriseTitleEn
	}
	if s.BusinessRegistrationNumber != "" {
		model.BusinessRegistrationNumber = s.BusinessRegistrationNumber
	}
	if s.EstablishedDate.After(minTime) {
		model.EstablishedDate = new(time.Time)
		model.EstablishedDate = s.EstablishedDate
	}
	if s.Region != "" {
		model.Region = s.Region
	}
	if s.ApprovedDate.After(minTime) {
		model.ApprovedDate = new(time.Time)
		model.ApprovedDate = s.ApprovedDate
	}
	if s.RegisteredAddress != "" {
		model.RegisteredAddress = s.RegisteredAddress
	}
	if s.RegisteredCapital != "" {
		model.RegisteredCapital = s.RegisteredCapital
	}
	if s.BusinessRegistrationNumber != "" {
		model.BusinessRegistrationNumber = s.BusinessRegistrationNumber
	}
	if s.ApprovedDate.After(minTime) {
		model.ApprovedDate = new(time.Time)
		model.ApprovedDate = s.ApprovedDate
	}
	if s.RegisteredAddress != "" {
		model.RegisteredAddress = s.RegisteredAddress
	}
	if s.PaidInCapital != "" {
		model.PaidInCapital = s.PaidInCapital
	}
	if s.EnterpriseType != "" {
		model.EnterpriseType = s.EnterpriseType
	}
	if s.StuffSize != "" {
		model.StuffSize = s.StuffSize
	}
	if s.StuffInsuredNumber != 0 {
		model.StuffInsuredNumber = s.StuffInsuredNumber
	}
	if s.BusinessScope != "" {
		model.BusinessScope = s.BusinessScope
	}
	if s.ImportExportQualificationCode != "" {
		model.ImportExportQualificationCode = s.ImportExportQualificationCode
	}
	if s.LegalRepresentative != "" {
		model.LegalRepresentative = s.LegalRepresentative
	}
	if s.RegistrationAuthority != "" {
		model.RegistrationAuthority = s.RegistrationAuthority
	}
	if s.RegistrationStatus != "" {
		model.RegistrationStatus = s.RegistrationStatus
	}
	if s.TaxpayerQualification != "" {
		model.TaxpayerQualification = s.TaxpayerQualification
	}
	if s.OrganizationCode != "" {
		model.OrganizationCode = s.OrganizationCode
	}
	if s.UrlQcc != "" {
		model.UrlQcc = s.UrlQcc
	}
	if s.UrlHomepage != "" {
		model.UrlHomepage = s.UrlHomepage
	}
	if s.BusinessTermStart.After(minTime) {
		model.BusinessTermStart = new(time.Time)
		model.BusinessTermStart = s.BusinessTermStart
	}
	if s.BusinessTermEnd.After(minTime) {
		model.BusinessTermEnd = new(time.Time)
		model.BusinessTermEnd = s.BusinessTermEnd
	}
	if s.UscId != "" {
		model.UscId = s.UscId
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseInfoInsertReq) GetId() interface{} {
	return s.InfoId
}

type EnterpriseInfoUpdateReq struct {
	InfoId                        int64      `uri:"infoId"`                         // 主键
	EnterpriseTitle               string     `json:"enterpriseTitle"`               // 企业名称
	EnterpriseTitleEn             string     `json:"enterpriseTitleEn"`             // 企业英文名称
	BusinessRegistrationNumber    string     `json:"businessRegistrationNumber"`    // 工商注册号
	EstablishedDate               *time.Time `json:"establishedDate"`               // 成立日期
	Region                        string     `json:"region"`                        // 所属地区
	ApprovedDate                  *time.Time `json:"approvedDate"`                  // 核准日期
	RegisteredAddress             string     `json:"registeredAddress"`             // 注册地址
	RegisteredCapital             string     `json:"registeredCapital"`             // 注册资本币种
	PaidInCapital                 string     `json:"paidInCapital"`                 // 实缴资本币种
	EnterpriseType                string     `json:"enterpriseType"`                // 企业类型
	StuffSize                     string     `json:"stuffSize"`                     // 人员规模
	StuffInsuredNumber            int        `json:"stuffInsuredNumber"`            // 参保人数
	BusinessScope                 string     `json:"businessScope"`                 // 经营范围
	ImportExportQualificationCode string     `json:"importExportQualificationCode"` // 进出口企业代码
	LegalRepresentative           string     `json:"legalRepresentative"`           // 法定代表人
	RegistrationAuthority         string     `json:"registrationAuthority"`         // 登记机关
	RegistrationStatus            string     `json:"registrationStatus"`            // 登记状态
	TaxpayerQualification         string     `json:"taxpayerQualification"`         // 纳税人资质
	OrganizationCode              string     `json:"organizationCode"`              // 组织机构代码
	UrlQcc                        string     `json:"urlQcc"`                        // 企查查url
	UrlHomepage                   string     `json:"urlHomepage"`                   // 官网url
	BusinessTermStart             *time.Time `json:"businessTermStart"`             // 营业期限开始
	BusinessTermEnd               *time.Time `json:"businessTermEnd"`               // 营业期限结束
	UscId                         string     `json:"uscId"`                         // 社会统一信用代码
	StatusCode                    int        `json:"statusCode"`                    // 状态标识码
	common.ControlBy
}

func (s *EnterpriseInfoUpdateReq) Generate(model *models.EnterpriseInfo) {
	minTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	model.InfoId = s.InfoId
	if s.EnterpriseTitle != "" {
		model.EnterpriseTitle = s.EnterpriseTitle
	}
	if s.EnterpriseTitleEn != "" {
		model.EnterpriseTitleEn = s.EnterpriseTitleEn
	}
	if s.BusinessRegistrationNumber != "" {
		model.BusinessRegistrationNumber = s.BusinessRegistrationNumber
	}
	if s.EstablishedDate.After(minTime) {
		model.EstablishedDate = new(time.Time)
		model.EstablishedDate = s.EstablishedDate
	}
	if s.Region != "" {
		model.Region = s.Region
	}
	if s.ApprovedDate.After(minTime) {
		model.ApprovedDate = new(time.Time)
		model.ApprovedDate = s.ApprovedDate
	}
	if s.RegisteredAddress != "" {
		model.RegisteredAddress = s.RegisteredAddress
	}
	if s.BusinessRegistrationNumber != "" {
		model.BusinessRegistrationNumber = s.BusinessRegistrationNumber
	}
	if s.ApprovedDate.After(minTime) {
		model.ApprovedDate = new(time.Time)
		model.ApprovedDate = s.ApprovedDate
	}
	if s.RegisteredAddress != "" {
		model.RegisteredAddress = s.RegisteredAddress
	}
	if s.RegisteredCapital != "" {
		model.RegisteredCapital = s.RegisteredCapital
	}
	if s.PaidInCapital != "" {
		model.PaidInCapital = s.PaidInCapital
	}
	if s.EnterpriseType != "" {
		model.EnterpriseType = s.EnterpriseType
	}
	if s.StuffSize != "" {
		model.StuffSize = s.StuffSize
	}
	if s.StuffInsuredNumber != 0 {
		model.StuffInsuredNumber = s.StuffInsuredNumber
	}
	if s.BusinessScope != "" {
		model.BusinessScope = s.BusinessScope
	}
	if s.ImportExportQualificationCode != "" {
		model.ImportExportQualificationCode = s.ImportExportQualificationCode
	}
	if s.LegalRepresentative != "" {
		model.LegalRepresentative = s.LegalRepresentative
	}
	if s.RegistrationAuthority != "" {
		model.RegistrationAuthority = s.RegistrationAuthority
	}
	if s.RegistrationStatus != "" {
		model.RegistrationStatus = s.RegistrationStatus
	}
	if s.TaxpayerQualification != "" {
		model.TaxpayerQualification = s.TaxpayerQualification
	}
	if s.OrganizationCode != "" {
		model.OrganizationCode = s.OrganizationCode
	}
	if s.UrlQcc != "" {
		model.UrlQcc = s.UrlQcc
	}
	if s.UrlHomepage != "" {
		model.UrlHomepage = s.UrlHomepage
	}
	if s.BusinessTermStart.After(minTime) {
		model.BusinessTermStart = new(time.Time)
		model.BusinessTermStart = s.BusinessTermStart
	}
	if s.BusinessTermEnd.After(minTime) {
		model.BusinessTermEnd = new(time.Time)
		model.BusinessTermEnd = s.BusinessTermEnd
	}
	if s.UscId != "" {
		model.UscId = s.UscId
	}
	if s.StatusCode != 0 {
		model.StatusCode = s.StatusCode
	}
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseInfoUpdateReq) GetId() interface{} {
	return s.InfoId
}

// EnterpriseInfoGetReq 功能获取请求参数
type EnterpriseInfoGetReq struct {
	InfoId int64 `uri:"infoId"`
}

func (s *EnterpriseInfoGetReq) GetId() interface{} {
	return s.InfoId
}

// EnterpriseInfoDeleteReq 功能删除请求参数
type EnterpriseInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
