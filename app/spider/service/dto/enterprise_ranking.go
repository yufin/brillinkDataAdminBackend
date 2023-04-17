package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EnterpriseRankingGetPageReq struct {
	dto.Pagination `search:"-"`

	EnterpriseId           int64  `form:"enterpriseId"  search:"type:;column:enterprise_id;table:enterprise_ranking" comment:"外键(enterprise表的id)"`
	RankingListId          int64  `form:"rankingListId"  search:"type:;column:ranking_list_id;table:enterprise_ranking" comment:"外键(enterprise_ranking_list表的id)"`
	RankingPosition        int    `form:"rankingPosition"  search:"type:;column:ranking_position;table:enterprise_ranking" comment:"榜内位置"`
	RankingEnterpriseTitle string `form:"rankingEnterpriseTitle"  search:"type:;column:ranking_enterprise_title;table:enterprise_ranking" comment:"榜单中的企业名称"`
	EnterpriseRankingPageOrder
}

type EnterpriseRankingPageOrder struct {
	RankId                 int64  `form:"rankIdOrder"  search:"type:order;column:rank_id;table:enterprise_ranking"`
	EnterpriseId           int64  `form:"enterpriseIdOrder"  search:"type:order;column:enterprise_id;table:enterprise_ranking"`
	RankingListId          int64  `form:"rankingListIdOrder"  search:"type:order;column:ranking_list_id;table:enterprise_ranking"`
	RankingPosition        int    `form:"rankingPositionOrder"  search:"type:order;column:ranking_position;table:enterprise_ranking"`
	RankingEnterpriseTitle string `form:"rankingEnterpriseTitleOrder"  search:"type:order;column:ranking_enterprise_title;table:enterprise_ranking"`
}

func (m *EnterpriseRankingGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnterpriseRankingGetResp struct {
	RankId                 int64  `json:"rankId"`                 // 主键
	EnterpriseId           int64  `json:"enterpriseId"`           // 外键(enterprise表的id)
	RankingListId          int64  `json:"rankingListId"`          // 外键(enterprise_ranking_list表的id)
	RankingPosition        int    `json:"rankingPosition"`        // 榜内位置
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"` // 榜单中的企业名称
	common.ControlBy
}

func (s *EnterpriseRankingGetResp) Generate(model *models.EnterpriseRanking) {
	s.RankId = model.RankId
	s.EnterpriseId = model.EnterpriseId
	s.RankingListId = model.RankingListId
	s.RankingPosition = model.RankingPosition
	s.RankingEnterpriseTitle = model.RankingEnterpriseTitle
	s.CreateBy = model.CreateBy
}

type EnterpriseRankingInsertReq struct {
	RankId                 int64  `json:"-"`                      // 主键
	EnterpriseId           int64  `json:"enterpriseId"`           // 外键(enterprise表的id)
	RankingListId          int64  `json:"rankingListId"`          // 外键(enterprise_ranking_list表的id)
	RankingPosition        int    `json:"rankingPosition"`        // 榜内位置
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"` // 榜单中的企业名称
	common.ControlBy
}

func (s *EnterpriseRankingInsertReq) Generate(model *models.EnterpriseRanking) {
	model.RankId = s.RankId
	model.EnterpriseId = s.EnterpriseId
	model.RankingListId = s.RankingListId
	model.RankingPosition = s.RankingPosition
	model.RankingEnterpriseTitle = s.RankingEnterpriseTitle
	model.CreateBy = s.CreateBy
}

func (s *EnterpriseRankingInsertReq) GetId() interface{} {
	return s.RankId
}

type EnterpriseRankingUpdateReq struct {
	RankId                 int64  `uri:"rankId"`                  // 主键
	EnterpriseId           int64  `json:"enterpriseId"`           // 外键(enterprise表的id)
	RankingListId          int64  `json:"rankingListId"`          // 外键(enterprise_ranking_list表的id)
	RankingPosition        int    `json:"rankingPosition"`        // 榜内位置
	RankingEnterpriseTitle string `json:"rankingEnterpriseTitle"` // 榜单中的企业名称
	common.ControlBy
}

func (s *EnterpriseRankingUpdateReq) Generate(model *models.EnterpriseRanking) {
	model.RankId = s.RankId
	model.EnterpriseId = s.EnterpriseId
	model.RankingListId = s.RankingListId
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
