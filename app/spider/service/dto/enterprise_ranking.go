package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type EnterpriseRankingGetPageReq struct {
	dto.Pagination `search:"-"`

	UscId                  string `form:"uscId"  search:"type:exact;column:usc_id;table:enterprise_ranking" comment:"社会统一信用代码"`
	ListId                 int64  `form:"listId"  search:"type:exact;column:list_id;table:enterprise_ranking" comment:"外键(enterprise_ranking_list表的id)"`
	RankingPosition        int    `form:"rankingPosition"  search:"type:exact;column:ranking_position;table:enterprise_ranking" comment:"榜内位置"`
	RankingEnterpriseTitle string `form:"rankingEnterpriseTitle"  search:"type:exact;column:ranking_enterprise_title;table:enterprise_ranking" comment:"榜单中的企业名称"`
	EnterpriseRankingPageOrder
}

type EnterpriseRankingPageOrder struct {
	RankId                 int64  `form:"rankIdOrder"  search:"type:order;column:rank_id;table:enterprise_ranking"`
	UscId                  string `form:"uscIdOrder"  search:"type:order;column:usc_id;table:enterprise_ranking"`
	ListId                 int64  `form:"listIdOrder"  search:"type:order;column:list_id;table:enterprise_ranking"`
	RankingPosition        int    `form:"rankingPositionOrder"  search:"type:order;column:ranking_position;table:enterprise_ranking"`
	RankingEnterpriseTitle string `form:"rankingEnterpriseTitleOrder"  search:"type:order;column:ranking_enterprise_title;table:enterprise_ranking"`
}

func (m *EnterpriseRankingGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseRankingGetResp struct {
	RankId                 int64  `json:"rankId"`                 // 主键
	UscId                  string `json:"uscId"`                  // 社会统一信用代码
	ListId                 int64  `json:"listId"`                 // 外键(enterprise_ranking_list表的id)
	RankingPosition        int    `json:"rankingPosition"`        // 榜内位置
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"` // 榜单中的企业名称
	common.ControlBy
}

func (s *EnterpriseRankingGetResp) Generate(model *models.EnterpriseRanking) {
	s.RankId = model.RankId
	s.UscId = model.UscId
	s.ListId = model.ListId
	s.RankingPosition = model.RankingPosition
	s.RankingEnterpriseTitle = model.RankingEnterpriseTitle
	s.CreateBy = model.CreateBy
}

type EnterpriseRankingInsertReq struct {
	RankId                 int64  `json:"-"`                      // 主键
	UscId                  string `json:"uscId"`                  // 社会统一信用代码
	ListId                 int64  `json:"listId"`                 // 外键(enterprise_ranking_list表的id)
	RankingPosition        int    `json:"rankingPosition"`        // 榜内位置
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"` // 榜单中的企业名称
	common.ControlBy
}

func (s *EnterpriseRankingInsertReq) Generate(model *models.EnterpriseRanking) {
	if s.RankId == 0 {
		s.RankId = utils.NewFlakeId()
	}
	model.RankId = s.RankId
	model.UscId = s.UscId
	model.ListId = s.ListId
	model.RankingPosition = s.RankingPosition
	model.RankingEnterpriseTitle = s.RankingEnterpriseTitle
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseRankingInsertReq) GetId() interface{} {
	return s.RankId
}

type EnterpriseRankingUpdateReq struct {
	RankId                 int64  `uri:"rankId"`                  // 主键
	UscId                  string `json:"uscId"`                  // 社会统一信用代码
	ListId                 int64  `json:"listId"`                 // 外键(enterprise_ranking_list表的id)
	RankingPosition        int    `json:"rankingPosition"`        // 榜内位置
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"` // 榜单中的企业名称
	common.ControlBy
}

func (s *EnterpriseRankingUpdateReq) Generate(model *models.EnterpriseRanking) {
	model.RankId = s.RankId
	model.UscId = s.UscId
	model.ListId = s.ListId
	model.RankingPosition = s.RankingPosition
	model.RankingEnterpriseTitle = s.RankingEnterpriseTitle
	model.UpdateBy = s.UpdateBy
}

func (s *EnterpriseRankingUpdateReq) GetId() interface{} {
	return s.RankId
}

// EnterpriseRankingGetReq 功能获取请求参数
type EnterpriseRankingGetReq struct {
	RankId int64 `uri:"rankId"`
}

func (s *EnterpriseRankingGetReq) GetId() interface{} {
	return s.RankId
}

// EnterpriseRankingDeleteReq 功能删除请求参数
type EnterpriseRankingDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnterpriseRankingDeleteReq) GetId() interface{} {
	return s.Ids
}
