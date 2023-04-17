package models

import (
	"go-admin/common/models"
)

// EnterpriseRanking 企业榜单排名数据
type EnterpriseRanking struct {
	RankId                 int64  `json:"rankId" gorm:"primaryKey;autoIncrement;comment:主键"`
	EnterpriseId           int64  `json:"enterpriseId" gorm:"comment:外键(enterprise表的id)"`
	RankingListId          int64  `json:"rankingListId" gorm:"comment:外键(enterprise_ranking_list表的id)"`
	RankingPosition        int    `json:"rankingPosition" gorm:"comment:榜内位置"`
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle" gorm:"comment:榜单中的企业名称"`
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
