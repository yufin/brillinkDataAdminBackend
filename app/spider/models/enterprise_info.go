package models

import (
	"go-admin/common/models"
	"time"
)

// EnterpriseInfo 企业主体信息
type EnterpriseInfo struct {
	InfoId                        int64      `json:"infoId" gorm:"primaryKey;comment:主键"`
	EnterpriseTitle               string     `json:"enterpriseTitle" gorm:"comment:企业名称"`
	EnterpriseTitleEn             string     `json:"enterpriseTitleEn" gorm:"comment:企业英文名称"`
	BusinessRegistrationNumber    string     `json:"businessRegistrationNumber" gorm:"comment:工商注册号"`
	EstablishedDate               *time.Time `json:"establishedDate" gorm:"comment:成立日期"`
	Region                        string     `json:"region" gorm:"comment:所属地区"`
	ApprovedDate                  *time.Time `json:"approvedDate" gorm:"comment:核准日期"`
	RegisteredAddress             string     `json:"registeredAddress" gorm:"comment:注册地址"`
	RegisteredCapital             float64    `json:"registeredCapital" gorm:"comment:注册资本"`
	RegisteredCapitalCurrency     string     `json:"registeredCapitalCurrency" gorm:"comment:注册资本币种"`
	PaidInCapital                 float64    `json:"paidInCapital" gorm:"comment:实缴资本"`
	PaidInCapitalCurrency         string     `json:"paidInCapitalCurrency" gorm:"comment:实缴资本币种"`
	EnterpriseType                string     `json:"enterpriseType" gorm:"comment:企业类型"`
	StuffSize                     string     `json:"stuffSize" gorm:"comment:人员规模"`
	StuffInsuredNumber            int        `json:"stuffInsuredNumber" gorm:"comment:参保人数"`
	BusinessScope                 string     `json:"businessScope" gorm:"comment:经营范围"`
	ImportExportQualificationCode string     `json:"importExportQualificationCode" gorm:"comment:进出口企业代码"`
	LegalRepresentative           string     `json:"legalRepresentative" gorm:"comment:法定代表人"`
	RegistrationAuthority         string     `json:"registrationAuthority" gorm:"comment:登记机关"`
	RegistrationStatus            string     `json:"registrationStatus" gorm:"comment:登记状态"`
	TaxpayerQualification         string     `json:"taxpayerQualification" gorm:"comment:纳税人资质"`
	OrganizationCode              string     `json:"organizationCode" gorm:"comment:组织机构代码"`
	UrlQcc                        string     `json:"urlQcc" gorm:"comment:企查查url"`
	UrlHomepage                   string     `json:"urlHomepage" gorm:"comment:官网url"`
	BusinessTermStart             *time.Time `json:"businessTermStart" gorm:"comment:营业期限开始"`
	BusinessTermEnd               *time.Time `json:"businessTermEnd" gorm:"comment:营业期限结束"`
	UscId                         string     `json:"uscId" gorm:"comment:社会统一信用代码"`
	StatusCode                    int        `json:"statusCode" gorm:"comment:状态标识码"`
	models.ModelTime
	models.ControlBy
}

func (*EnterpriseInfo) TableName() string {
	return "enterprise_info"
}

func (e *EnterpriseInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseInfo) GetId() interface{} {
	return e.InfoId
}
