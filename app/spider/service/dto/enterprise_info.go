package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type EnterpriseInfoGetPageReq struct {
	dto.Pagination `search:"-"`

	EnterpriseTitle               string    `form:"enterpriseTitle"  search:"type:;column:enterprise_title;table:enterprise_info" comment:"企业名称"`
	EnterpriseTitleEn             string    `form:"enterpriseTitleEn"  search:"type:;column:enterprise_title_en;table:enterprise_info" comment:"企业英文名称"`
	BusinessRegistrationNumber    string    `form:"businessRegistrationNumber"  search:"type:;column:business_registration_number;table:enterprise_info" comment:"工商注册号"`
	EstablishedDate               time.Time `form:"establishedDate"  search:"type:;column:established_date;table:enterprise_info" comment:"成立日期"`
	Region                        string    `form:"region"  search:"type:;column:region;table:enterprise_info" comment:"所属地区"`
	ApprovedDate                  time.Time `form:"approvedDate"  search:"type:;column:approved_date;table:enterprise_info" comment:"核准日期"`
	RegisteredAddress             string    `form:"registeredAddress"  search:"type:;column:registered_address;table:enterprise_info" comment:"注册地址"`
	RegisteredCapital             float64   `form:"registeredCapital"  search:"type:;column:registered_capital;table:enterprise_info" comment:"注册资本"`
	RegisteredCapitalCurrency     string    `form:"registeredCapitalCurrency"  search:"type:;column:registered_capital_currency;table:enterprise_info" comment:"注册资本币种"`
	PaidInCapital                 float64   `form:"paidInCapital"  search:"type:;column:paid_in_capital;table:enterprise_info" comment:"实缴资本"`
	PaidInCapitalCurrency         string    `form:"paidInCapitalCurrency"  search:"type:;column:paid_in_capital_currency;table:enterprise_info" comment:"实缴资本币种"`
	EnterpriseType                string    `form:"enterpriseType"  search:"type:;column:enterprise_type;table:enterprise_info" comment:"企业类型"`
	StuffSize                     string    `form:"stuffSize"  search:"type:;column:stuff_size;table:enterprise_info" comment:"人员规模"`
	StuffInsuredNumber            int       `form:"stuffInsuredNumber"  search:"type:;column:stuff_insured_number;table:enterprise_info" comment:"参保人数"`
	BusinessScope                 string    `form:"businessScope"  search:"type:;column:business_scope;table:enterprise_info" comment:"经营范围"`
	ImportExportQualificationCode string    `form:"importExportQualificationCode"  search:"type:;column:import_export_qualification_code;table:enterprise_info" comment:"进出口企业代码"`
	LegalRepresentative           string    `form:"legalRepresentative"  search:"type:;column:legal_representative;table:enterprise_info" comment:"法定代表人"`
	RegistrationAuthority         string    `form:"registrationAuthority"  search:"type:;column:registration_authority;table:enterprise_info" comment:"登记机关"`
	RegistrationStatus            string    `form:"registrationStatus"  search:"type:;column:registration_status;table:enterprise_info" comment:"登记状态"`
	TaxpayerQualification         string    `form:"taxpayerQualification"  search:"type:;column:taxpayer_qualification;table:enterprise_info" comment:"纳税人资质"`
	OrganizationCode              string    `form:"organizationCode"  search:"type:;column:organization_code;table:enterprise_info" comment:"组织机构代码"`
	UrlQcc                        string    `form:"urlQcc"  search:"type:;column:url_qcc;table:enterprise_info" comment:"企查查url"`
	UrlHomepage                   string    `form:"urlHomepage"  search:"type:;column:url_homepage;table:enterprise_info" comment:"官网url"`
	BusinessTermStart             time.Time `form:"businessTermStart"  search:"type:;column:business_term_start;table:enterprise_info" comment:"营业期限开始"`
	BusinessTermEnd               time.Time `form:"businessTermEnd"  search:"type:;column:business_term_end;table:enterprise_info" comment:"营业期限结束"`
	UscId                         string    `form:"uscId"  search:"type:;column:usc_id;table:enterprise_info" comment:"社会统一信用代码"`
	StatusCode                    int       `form:"statusCode"  search:"type:;column:status_code;table:enterprise_info" comment:"状态标识码"`
	EnterpriseInfoPageOrder
}

type EnterpriseInfoPageOrder struct {
	InfoId                        int64     `form:"infoIdOrder"  search:"type:order;column:info_id;table:enterprise_info"`
	EnterpriseTitle               string    `form:"enterpriseTitleOrder"  search:"type:order;column:enterprise_title;table:enterprise_info"`
	EnterpriseTitleEn             string    `form:"enterpriseTitleEnOrder"  search:"type:order;column:enterprise_title_en;table:enterprise_info"`
	BusinessRegistrationNumber    string    `form:"businessRegistrationNumberOrder"  search:"type:order;column:business_registration_number;table:enterprise_info"`
	EstablishedDate               time.Time `form:"establishedDateOrder"  search:"type:order;column:established_date;table:enterprise_info"`
	Region                        string    `form:"regionOrder"  search:"type:order;column:region;table:enterprise_info"`
	ApprovedDate                  time.Time `form:"approvedDateOrder"  search:"type:order;column:approved_date;table:enterprise_info"`
	RegisteredAddress             string    `form:"registeredAddressOrder"  search:"type:order;column:registered_address;table:enterprise_info"`
	RegisteredCapital             float64   `form:"registeredCapitalOrder"  search:"type:order;column:registered_capital;table:enterprise_info"`
	RegisteredCapitalCurrency     string    `form:"registeredCapitalCurrencyOrder"  search:"type:order;column:registered_capital_currency;table:enterprise_info"`
	PaidInCapital                 float64   `form:"paidInCapitalOrder"  search:"type:order;column:paid_in_capital;table:enterprise_info"`
	PaidInCapitalCurrency         string    `form:"paidInCapitalCurrencyOrder"  search:"type:order;column:paid_in_capital_currency;table:enterprise_info"`
	EnterpriseType                string    `form:"enterpriseTypeOrder"  search:"type:order;column:enterprise_type;table:enterprise_info"`
	StuffSize                     string    `form:"stuffSizeOrder"  search:"type:order;column:stuff_size;table:enterprise_info"`
	StuffInsuredNumber            int       `form:"stuffInsuredNumberOrder"  search:"type:order;column:stuff_insured_number;table:enterprise_info"`
	BusinessScope                 string    `form:"businessScopeOrder"  search:"type:order;column:business_scope;table:enterprise_info"`
	ImportExportQualificationCode string    `form:"importExportQualificationCodeOrder"  search:"type:order;column:import_export_qualification_code;table:enterprise_info"`
	LegalRepresentative           string    `form:"legalRepresentativeOrder"  search:"type:order;column:legal_representative;table:enterprise_info"`
	RegistrationAuthority         string    `form:"registrationAuthorityOrder"  search:"type:order;column:registration_authority;table:enterprise_info"`
	RegistrationStatus            string    `form:"registrationStatusOrder"  search:"type:order;column:registration_status;table:enterprise_info"`
	TaxpayerQualification         string    `form:"taxpayerQualificationOrder"  search:"type:order;column:taxpayer_qualification;table:enterprise_info"`
	OrganizationCode              string    `form:"organizationCodeOrder"  search:"type:order;column:organization_code;table:enterprise_info"`
	UrlQcc                        string    `form:"urlQccOrder"  search:"type:order;column:url_qcc;table:enterprise_info"`
	UrlHomepage                   string    `form:"urlHomepageOrder"  search:"type:order;column:url_homepage;table:enterprise_info"`
	BusinessTermStart             time.Time `form:"businessTermStartOrder"  search:"type:order;column:business_term_start;table:enterprise_info"`
	BusinessTermEnd               time.Time `form:"businessTermEndOrder"  search:"type:order;column:business_term_end;table:enterprise_info"`
	UscId                         string    `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise_info"`
	StatusCode                    int       `form:"statusCodeOrder"  search:"type:order;column:status_code;table:enterprise_info"`
}

func (m *EnterpriseInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseInfoGetResp struct {
	InfoId                        int64     `json:"infoId"`                        // 主键
	EnterpriseTitle               string    `json:"enterpriseTitle"`               // 企业名称
	EnterpriseTitleEn             string    `json:"enterpriseTitleEn"`             // 企业英文名称
	BusinessRegistrationNumber    string    `json:"businessRegistrationNumber"`    // 工商注册号
	EstablishedDate               time.Time `json:"establishedDate"`               // 成立日期
	Region                        string    `json:"region"`                        // 所属地区
	ApprovedDate                  time.Time `json:"approvedDate"`                  // 核准日期
	RegisteredAddress             string    `json:"registeredAddress"`             // 注册地址
	RegisteredCapital             float64   `json:"registeredCapital"`             // 注册资本
	RegisteredCapitalCurrency     string    `json:"registeredCapitalCurrency"`     // 注册资本币种
	PaidInCapital                 float64   `json:"paidInCapital"`                 // 实缴资本
	PaidInCapitalCurrency         string    `json:"paidInCapitalCurrency"`         // 实缴资本币种
	EnterpriseType                string    `json:"enterpriseType"`                // 企业类型
	StuffSize                     string    `json:"stuffSize"`                     // 人员规模
	StuffInsuredNumber            int       `json:"stuffInsuredNumber"`            // 参保人数
	BusinessScope                 string    `json:"businessScope"`                 // 经营范围
	ImportExportQualificationCode string    `json:"importExportQualificationCode"` // 进出口企业代码
	LegalRepresentative           string    `json:"legalRepresentative"`           // 法定代表人
	RegistrationAuthority         string    `json:"registrationAuthority"`         // 登记机关
	RegistrationStatus            string    `json:"registrationStatus"`            // 登记状态
	TaxpayerQualification         string    `json:"taxpayerQualification"`         // 纳税人资质
	OrganizationCode              string    `json:"organizationCode"`              // 组织机构代码
	UrlQcc                        string    `json:"urlQcc"`                        // 企查查url
	UrlHomepage                   string    `json:"urlHomepage"`                   // 官网url
	BusinessTermStart             time.Time `json:"businessTermStart"`             // 营业期限开始
	BusinessTermEnd               time.Time `json:"businessTermEnd"`               // 营业期限结束
	UscId                         string    `json:"uscId"`                         // 社会统一信用代码
	StatusCode                    int       `json:"statusCode"`                    // 状态标识码
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
	s.RegisteredCapitalCurrency = model.RegisteredCapitalCurrency
	s.PaidInCapital = model.PaidInCapital
	s.PaidInCapitalCurrency = model.PaidInCapitalCurrency
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
	InfoId                        int64     `json:"-"`                             // 主键
	EnterpriseTitle               string    `json:"enterpriseTitle"`               // 企业名称
	EnterpriseTitleEn             string    `json:"enterpriseTitleEn"`             // 企业英文名称
	BusinessRegistrationNumber    string    `json:"businessRegistrationNumber"`    // 工商注册号
	EstablishedDate               time.Time `json:"establishedDate"`               // 成立日期
	Region                        string    `json:"region"`                        // 所属地区
	ApprovedDate                  time.Time `json:"approvedDate"`                  // 核准日期
	RegisteredAddress             string    `json:"registeredAddress"`             // 注册地址
	RegisteredCapital             float64   `json:"registeredCapital"`             // 注册资本
	RegisteredCapitalCurrency     string    `json:"registeredCapitalCurrency"`     // 注册资本币种
	PaidInCapital                 float64   `json:"paidInCapital"`                 // 实缴资本
	PaidInCapitalCurrency         string    `json:"paidInCapitalCurrency"`         // 实缴资本币种
	EnterpriseType                string    `json:"enterpriseType"`                // 企业类型
	StuffSize                     string    `json:"stuffSize"`                     // 人员规模
	StuffInsuredNumber            int       `json:"stuffInsuredNumber"`            // 参保人数
	BusinessScope                 string    `json:"businessScope"`                 // 经营范围
	ImportExportQualificationCode string    `json:"importExportQualificationCode"` // 进出口企业代码
	LegalRepresentative           string    `json:"legalRepresentative"`           // 法定代表人
	RegistrationAuthority         string    `json:"registrationAuthority"`         // 登记机关
	RegistrationStatus            string    `json:"registrationStatus"`            // 登记状态
	TaxpayerQualification         string    `json:"taxpayerQualification"`         // 纳税人资质
	OrganizationCode              string    `json:"organizationCode"`              // 组织机构代码
	UrlQcc                        string    `json:"urlQcc"`                        // 企查查url
	UrlHomepage                   string    `json:"urlHomepage"`                   // 官网url
	BusinessTermStart             time.Time `json:"businessTermStart"`             // 营业期限开始
	BusinessTermEnd               time.Time `json:"businessTermEnd"`               // 营业期限结束
	UscId                         string    `json:"uscId"`                         // 社会统一信用代码
	StatusCode                    int       `json:"statusCode"`                    // 状态标识码
	common.ControlBy
}

func (s *EnterpriseInfoInsertReq) Generate(model *models.EnterpriseInfo) {
	model.InfoId = s.InfoId
	model.EnterpriseTitle = s.EnterpriseTitle
	model.EnterpriseTitleEn = s.EnterpriseTitleEn
	model.BusinessRegistrationNumber = s.BusinessRegistrationNumber
	model.EstablishedDate = s.EstablishedDate
	model.Region = s.Region
	model.ApprovedDate = s.ApprovedDate
	model.RegisteredAddress = s.RegisteredAddress
	model.RegisteredCapital = s.RegisteredCapital
	model.RegisteredCapitalCurrency = s.RegisteredCapitalCurrency
	model.PaidInCapital = s.PaidInCapital
	model.PaidInCapitalCurrency = s.PaidInCapitalCurrency
	model.EnterpriseType = s.EnterpriseType
	model.StuffSize = s.StuffSize
	model.StuffInsuredNumber = s.StuffInsuredNumber
	model.BusinessScope = s.BusinessScope
	model.ImportExportQualificationCode = s.ImportExportQualificationCode
	model.LegalRepresentative = s.LegalRepresentative
	model.RegistrationAuthority = s.RegistrationAuthority
	model.RegistrationStatus = s.RegistrationStatus
	model.TaxpayerQualification = s.TaxpayerQualification
	model.OrganizationCode = s.OrganizationCode
	model.UrlQcc = s.UrlQcc
	model.UrlHomepage = s.UrlHomepage
	model.BusinessTermStart = s.BusinessTermStart
	model.BusinessTermEnd = s.BusinessTermEnd
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseInfoInsertReq) GetId() interface{} {
	return s.InfoId
}

type EnterpriseInfoUpdateReq struct {
	InfoId                        int64     `uri:"infoId"`                         // 主键
	EnterpriseTitle               string    `json:"enterpriseTitle"`               // 企业名称
	EnterpriseTitleEn             string    `json:"enterpriseTitleEn"`             // 企业英文名称
	BusinessRegistrationNumber    string    `json:"businessRegistrationNumber"`    // 工商注册号
	EstablishedDate               time.Time `json:"establishedDate"`               // 成立日期
	Region                        string    `json:"region"`                        // 所属地区
	ApprovedDate                  time.Time `json:"approvedDate"`                  // 核准日期
	RegisteredAddress             string    `json:"registeredAddress"`             // 注册地址
	RegisteredCapital             float64   `json:"registeredCapital"`             // 注册资本
	RegisteredCapitalCurrency     string    `json:"registeredCapitalCurrency"`     // 注册资本币种
	PaidInCapital                 float64   `json:"paidInCapital"`                 // 实缴资本
	PaidInCapitalCurrency         string    `json:"paidInCapitalCurrency"`         // 实缴资本币种
	EnterpriseType                string    `json:"enterpriseType"`                // 企业类型
	StuffSize                     string    `json:"stuffSize"`                     // 人员规模
	StuffInsuredNumber            int       `json:"stuffInsuredNumber"`            // 参保人数
	BusinessScope                 string    `json:"businessScope"`                 // 经营范围
	ImportExportQualificationCode string    `json:"importExportQualificationCode"` // 进出口企业代码
	LegalRepresentative           string    `json:"legalRepresentative"`           // 法定代表人
	RegistrationAuthority         string    `json:"registrationAuthority"`         // 登记机关
	RegistrationStatus            string    `json:"registrationStatus"`            // 登记状态
	TaxpayerQualification         string    `json:"taxpayerQualification"`         // 纳税人资质
	OrganizationCode              string    `json:"organizationCode"`              // 组织机构代码
	UrlQcc                        string    `json:"urlQcc"`                        // 企查查url
	UrlHomepage                   string    `json:"urlHomepage"`                   // 官网url
	BusinessTermStart             time.Time `json:"businessTermStart"`             // 营业期限开始
	BusinessTermEnd               time.Time `json:"businessTermEnd"`               // 营业期限结束
	UscId                         string    `json:"uscId"`                         // 社会统一信用代码
	StatusCode                    int       `json:"statusCode"`                    // 状态标识码
	common.ControlBy
}

func (s *EnterpriseInfoUpdateReq) Generate(model *models.EnterpriseInfo) {
	model.InfoId = s.InfoId
	model.EnterpriseTitle = s.EnterpriseTitle
	model.EnterpriseTitleEn = s.EnterpriseTitleEn
	model.BusinessRegistrationNumber = s.BusinessRegistrationNumber
	model.EstablishedDate = s.EstablishedDate
	model.Region = s.Region
	model.ApprovedDate = s.ApprovedDate
	model.RegisteredAddress = s.RegisteredAddress
	model.RegisteredCapital = s.RegisteredCapital
	model.RegisteredCapitalCurrency = s.RegisteredCapitalCurrency
	model.PaidInCapital = s.PaidInCapital
	model.PaidInCapitalCurrency = s.PaidInCapitalCurrency
	model.EnterpriseType = s.EnterpriseType
	model.StuffSize = s.StuffSize
	model.StuffInsuredNumber = s.StuffInsuredNumber
	model.BusinessScope = s.BusinessScope
	model.ImportExportQualificationCode = s.ImportExportQualificationCode
	model.LegalRepresentative = s.LegalRepresentative
	model.RegistrationAuthority = s.RegistrationAuthority
	model.RegistrationStatus = s.RegistrationStatus
	model.TaxpayerQualification = s.TaxpayerQualification
	model.OrganizationCode = s.OrganizationCode
	model.UrlQcc = s.UrlQcc
	model.UrlHomepage = s.UrlHomepage
	model.BusinessTermStart = s.BusinessTermStart
	model.BusinessTermEnd = s.BusinessTermEnd
	model.UscId = s.UscId
	model.StatusCode = s.StatusCode
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
