package models

import (
	"go-admin/common/models"
)

// EnterpriseRanking 企业榜单排名数据
type EnterpriseRanking struct {
	RankId                 int64  `json:"rankId" gorm:"primaryKey;comment:主键"`
	UscId                  string `json:"uscId" gorm:"comment:社会统一信用代码"`
	ListId                 int64  `json:"listId" gorm:"comment:外键(enterprise_ranking_list表的id)"`
	RankingPosition        int    `json:"rankingPosition" gorm:"comment:榜内位置"`
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle" gorm:"comment:榜单中的企业名称"`
	StatusCode             int    `json:"statusCode" gorm:"comment:状态码"` //TODD:加字段
	models.ModelTime
	models.ControlBy
}

func (*EnterpriseRanking) TableName() string {
	return "enterprise_ranking"
}

func (e *EnterpriseRanking) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnterpriseRanking) GetId() interface{} {
	return e.RankId
}

type EnterpriseRankingList struct {
	RankingList
	RankId                 int64  `json:"rankId" gorm:"comment:主键"`
	RankingPosition        int    `json:"rankingPosition" gorm:"comment:榜内位置"`
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle" gorm:"comment:榜单中的企业名称"`
	UscId                  string `json:"uscId" gorm:"comment:社会统一信用代码"`
}
