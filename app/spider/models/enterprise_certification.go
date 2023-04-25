package models

import (
	"go-admin/common/models"
	"time"
)

// EnterpriseCertification
type EnterpriseCertification struct {
	CertId                 int64     `json:"certId" gorm:"primaryKey;comment:主键"`
	CertificationTitle     string    `json:"certificationTitle" gorm:"comment:认证名称"`
	CertificationCode      string    `json:"certificationCode" gorm:"comment:认证编号"`
	CertificationLevel     string    `json:"certificationLevel" gorm:"comment:认证等级(省级,市级,国家级)"`
	CertificationType      string    `json:"certificationType" gorm:"comment:认证类型(荣誉,科技型企业,...)"`
	CertificationSource    string    `json:"certificationSource" gorm:"comment:认证来源(eg:2022年度浙江省中小企业公共服务示范平台名单)"`
	CertificationDate      time.Time `json:"certificationDate" gorm:"comment:发证日期"`
	CertificationTermStart time.Time `json:"certificationTermStart" gorm:"comment:有效期起"`
	CertificationTermEnd   time.Time `json:"certificationTermEnd" gorm:"comment:有效期至"`
	CertificationAuthority string    `json:"certificationAuthority" gorm:"comment:发证机关"`
	UscId                  string    `json:"uscId" gorm:"comment:社会统一信用代码"`
	StatusCode             int64     `json:"statusCode" gorm:"comment:状态标识码"`
	models.ModelTime
	models.ControlBy
}

func (*EnterpriseCertification) TableName() string {
	return "enterprise_certification"
}

func (e *EnterpriseCertification) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseCertification) GetId() interface{} {
	return e.CertId
}
