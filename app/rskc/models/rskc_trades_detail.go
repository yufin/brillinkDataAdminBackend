package models

import (
	"go-admin/common/models"
)

// RskcTradesDetail 客户、供应商交易细节（来自origin_content表)
type RskcTradesDetail struct {
	models.Model
	ContentId      string `json:"contentId" gorm:"comment:外键"`
	EnterpriseName string `json:"enterpriseName" gorm:"comment:企业名称"`
	CommodityRatio string `json:"commodityRatio" gorm:"default:'-';comment:货物占比"`
	CommodityName  string `json:"commodityName" gorm:"default:'-';comment:货物种类名称"`
	RatioAmountTax string `json:"ratioAmountTax" gorm:"default:'-';comment:税额占比"`
	SumAmountTax   string `json:"sumAmountTax" gorm:"default:'-';comment:税总额"`
	DetailType     int    `json:"detailType" gorm:"comment:1:customer_12,2:customer_24,3:supplier_12,4:supplier_24"`
	TagIndustry    string `json:"tagIndustry" gorm:"default:'-';comment:行业标签"`
	TagAuthorized  string `json:"tagAuthorized" gorm:"default:'-';comment:认证标签"`
	TagProduct     string `json:"tagProduct" gorm:"default:'-';comment:产品标签"`
	TagList        string `json:"tagList" gorm:"default:'-';comment:榜单标签"`
	EnterpriseInfo string `json:"enterpriseInfo" gorm:"default:'-';comment:企业信息"`
	StatusCode     int    `json:"statusCode" gorm:"comment:状态码: 1.待确认企业数据已采集，2.待采集，已经同步至waitList, 3.采集完成, 4.匹配并录入完成"`
	models.ModelTime
	models.ControlBy
}

func (*RskcTradesDetail) TableName() string {
	return "rskc_trades_detail"
}

func (e *RskcTradesDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RskcTradesDetail) GetId() interface{} {
	return e.Id
}
