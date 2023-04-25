package models

import (
	"go-admin/common/models"
)

// RankingList 排名榜单表
type RankingList struct {
	models.Model
	ListTitle             string `json:"listTitle" gorm:"comment:榜单名称"`
	ListType              string `json:"listType" gorm:"comment:榜单类型(品牌产品榜,企业榜,...)"`
	ListSource            string `json:"listSource" gorm:"comment:排名来源(JsonArray格式;eg:["德本咨询", "eNet研究院", "互联网周刊"]"`
	ListParticipantsTotal int64  `json:"listParticipantsTotal" gorm:"comment:参与排名企业数"`
	ListPublishedDate     string `json:"listPublishedDate" gorm:"comment:排名发布日期"`
	ListUrlQcc            string `json:"listUrlQcc" gorm:"comment:排名url(qcc)"`
	ListUrlOrigin         string `json:"listUrlOrigin" gorm:"comment:排名url(来源)"`
	StatusCode            int64  `json:"statusCode" gorm:"comment:状态标识码"`
	models.ModelTime
	models.ControlBy
}

func (*RankingList) TableName() string {
	return "ranking_list"
}

func (e *RankingList) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RankingList) GetId() interface{} {
	return e.Id
}
