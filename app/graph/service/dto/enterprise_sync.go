package dto

import (
	modelsSp "go-admin/app/spider/models"
	"go-admin/utils"
	"time"
)

type EnterpriseInfoSyncReq struct {
	Id                            string     `json:"id"`
	InfoId                        int64      `json:"_infoId"`
	EnterpriseTitle               string     `json:"title"`
	EnterpriseTitleEn             string     `json:"enterpriseTitleEn"`
	BusinessRegistrationNumber    string     `json:"businessRegistrationNumber"`
	EstablishedDate               *time.Time `json:"establishedDate"`
	Region                        string     `json:"region"`
	ApprovedDate                  *time.Time `json:"approvedDate"`
	RegisteredAddress             string     `json:"registeredAddress"`
	RegisteredCapital             string     `json:"registeredCapital"`
	PaidInCapital                 string     `json:"paidInCapital"`
	EnterpriseType                string     `json:"enterpriseType"`
	StuffSize                     string     `json:"stuffSize"`
	StuffInsuredNumber            int        `json:"stuffInsuredNumber"`
	BusinessScope                 string     `json:"businessScope"`
	ImportExportQualificationCode string     `json:"importExportQualificationCode"`
	LegalRepresentative           string     `json:"legalRepresentative"`
	RegistrationAuthority         string     `json:"registrationAuthority"`
	RegistrationStatus            string     `json:"registrationStatus"`
	TaxpayerQualification         string     `json:"taxpayerQualification"`
	OrganizationCode              string     `json:"organizationCode"`
	UrlQcc                        string     `json:"urlQcc"`
	UrlHomepage                   string     `json:"urlHomepage"`
	BusinessTermStart             *time.Time `json:"businessTermStart"`
	BusinessTermEnd               *time.Time `json:"businessTermEnd"`
	UscId                         string     `json:"uscId"`
	CreatedAt                     time.Time  `json:"createdAt"`
}

func (s *EnterpriseInfoSyncReq) Assignment(model *modelsSp.EnterpriseInfo) {
	s.Id = utils.NewRandomUUID()
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
	s.CreatedAt = model.CreatedAt
}

type EnterpriseCertificationSyncReq struct {
	Id                     string     `json:"id"`
	CertId                 int64      `json:"_certId"`
	UscId                  string     `json:"uscId"`
	CertificationTitle     string     `json:"-"`
	CertificationCode      string     `json:"certificationCode"`
	CertificationLevel     string     `json:"certificationLevel"`
	CertificationType      string     `json:"certificationType"`
	CertificationSource    string     `json:"certificationSource"`
	CertificationDate      *time.Time `json:"certificationDate"`
	CertificationTermStart *time.Time `json:"certificationTermStart"`
	CertificationTermEnd   *time.Time `json:"certificationTermEnd"`
	CertificationAuthority string     `json:"certificationAuthority"`
	CreatedAt              time.Time  `json:"createdAt"`
}

func (s *EnterpriseCertificationSyncReq) Assignment(model *modelsSp.EnterpriseCertification) {
	s.Id = utils.NewRandomUUID()
	s.CertId = model.CertId
	s.UscId = model.UscId
	s.CertificationTitle = model.CertificationTitle
	s.CertificationCode = model.CertificationCode
	s.CertificationLevel = model.CertificationLevel
	s.CertificationType = model.CertificationType
	s.CertificationSource = model.CertificationSource
	s.CertificationDate = model.CertificationDate
	s.CertificationTermStart = model.CertificationTermStart
	s.CertificationTermEnd = model.CertificationTermEnd
	s.CertificationAuthority = model.CertificationAuthority
	s.CreatedAt = model.CreatedAt
}

type RankingListSyncReq struct {
	Id                    string    `json:"id"`
	ListId                int64     `json:"_listId"`
	ListTitle             string    `json:"title"`
	ListType              string    `json:"listType"`
	ListSource            string    `json:"listSource"`
	ListParticipantsTotal int64     `json:"listParticipantsTotal"`
	ListPublishedDate     string    `json:"listPublishedDate"`
	ListUrlQcc            string    `json:"listUrlQcc"`
	ListUrlOrigin         string    `json:"listUrlOrigin"`
	CreatedAt             time.Time `json:"createdAt"`
}

func (s *RankingListSyncReq) Assignment(model *modelsSp.EnterpriseRankingList) {
	s.Id = utils.NewRandomUUID()
	s.ListId = model.Id
	s.ListTitle = model.ListTitle
	s.ListType = model.ListType
	s.ListSource = model.ListSource
	s.ListParticipantsTotal = model.ListParticipantsTotal
	s.ListPublishedDate = model.ListPublishedDate
	s.ListUrlQcc = model.ListUrlQcc
	s.ListUrlOrigin = model.ListUrlOrigin
	s.CreatedAt = model.CreatedAt
}

type EnterpriseRankingSyncReq struct {
	Id                     string `json:"id"`
	RankId                 int64  `json:"_rankId"`
	UscId                  string `json:"-"`
	RankingPosition        int    `json:"rankingPosition"`
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"`
}

func (s *EnterpriseRankingSyncReq) Assignment(model *modelsSp.EnterpriseRankingList) {
	s.Id = utils.NewRandomUUID()
	s.RankId = model.RankId
	s.UscId = model.UscId
	s.RankingPosition = model.RankingPosition
	s.RankingEnterpriseTitle = model.RankingEnterpriseTitle
}

type EnterpriseProductSyncReq struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type EnterpriseIndustrySyncReq struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
