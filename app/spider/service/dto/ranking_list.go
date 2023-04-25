package dto

import (
	"go-admin/app/spider/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RankingListGetPageReq struct {
	dto.Pagination `search:"-"`

	Id                    int64  `form:"id"  search:"type:;column:id;table:ranking_list" comment:"主键id"`
	ListTitle             string `form:"listTitle"  search:"type:;column:list_title;table:ranking_list" comment:"榜单名称"`
	ListType              string `form:"listType"  search:"type:;column:list_type;table:ranking_list" comment:"榜单类型(品牌产品榜,企业榜,...)"`
	ListSource            string `form:"listSource"  search:"type:;column:list_source;table:ranking_list" comment:"排名来源(JsonArray格式;eg:["德本咨询", "eNet研究院", "互联网周刊"]"`
	ListParticipantsTotal int64  `form:"listParticipantsTotal"  search:"type:;column:list_participants_total;table:ranking_list" comment:"参与排名企业数"`
	ListPublishedDate     string `form:"listPublishedDate"  search:"type:;column:list_published_date;table:ranking_list" comment:"排名发布日期"`
	ListUrlQcc            string `form:"listUrlQcc"  search:"type:;column:list_url_qcc;table:ranking_list" comment:"排名url(qcc)"`
	ListUrlOrigin         string `form:"listUrlOrigin"  search:"type:;column:list_url_origin;table:ranking_list" comment:"排名url(来源)"`
	StatusCode            int64  `form:"statusCode"  search:"type:;column:status_code;table:ranking_list" comment:"状态标识码"`
	StartTime             string `form:"startTime" search:"type:gte;column:created_at;table:ranking_list" comment:"创建时间"`
	EndTime               string `form:"endTime" search:"type:lte;column:created_at;table:ranking_list" comment:"创建时间"`
	RankingListPageOrder
}

type RankingListPageOrder struct {
	Id                    int64  `form:"idOrder"  search:"type:order;column:id;table:ranking_list"`
	ListTitle             string `form:"listTitleOrder"  search:"type:order;column:list_title;table:ranking_list"`
	ListType              string `form:"listTypeOrder"  search:"type:order;column:list_type;table:ranking_list"`
	ListSource            string `form:"listSourceOrder"  search:"type:order;column:list_source;table:ranking_list"`
	ListParticipantsTotal int64  `form:"listParticipantsTotalOrder"  search:"type:order;column:list_participants_total;table:ranking_list"`
	ListPublishedDate     string `form:"listPublishedDateOrder"  search:"type:order;column:list_published_date;table:ranking_list"`
	ListUrlQcc            string `form:"listUrlQccOrder"  search:"type:order;column:list_url_qcc;table:ranking_list"`
	ListUrlOrigin         string `form:"listUrlOriginOrder"  search:"type:order;column:list_url_origin;table:ranking_list"`
	StatusCode            int64  `form:"statusCodeOrder"  search:"type:order;column:status_code;table:ranking_list"`
}

func (m *RankingListGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RankingListGetResp struct {
	Id                    int64  `json:"id"`                    // 主键id
	ListTitle             string `json:"listTitle"`             // 榜单名称
	ListType              string `json:"listType"`              // 榜单类型(品牌产品榜,企业榜,...)
	ListSource            string `json:"listSource"`            // 排名来源(JsonArray格式;eg:["德本咨询", "eNet研究院", "互联网周刊"]
	ListParticipantsTotal int64  `json:"listParticipantsTotal"` // 参与排名企业数
	ListPublishedDate     string `json:"listPublishedDate"`     // 排名发布日期
	ListUrlQcc            string `json:"listUrlQcc"`            // 排名url(qcc)
	ListUrlOrigin         string `json:"listUrlOrigin"`         // 排名url(来源)
	StatusCode            int64  `json:"statusCode"`            // 状态标识码
	common.ControlBy
}

func (s *RankingListGetResp) Generate(model *models.RankingList) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.ListTitle = model.ListTitle
	s.ListType = model.ListType
	s.ListSource = model.ListSource
	s.ListParticipantsTotal = model.ListParticipantsTotal
	s.ListPublishedDate = model.ListPublishedDate
	s.ListUrlQcc = model.ListUrlQcc
	s.ListUrlOrigin = model.ListUrlOrigin
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type RankingListInsertReq struct {
	Id                    int64  `json:"-"`                     // 主键id
	ListTitle             string `json:"listTitle"`             // 榜单名称
	ListType              string `json:"listType"`              // 榜单类型(品牌产品榜,企业榜,...)
	ListSource            string `json:"listSource"`            // 排名来源(JsonArray格式;eg:["德本咨询", "eNet研究院", "互联网周刊"]
	ListParticipantsTotal int64  `json:"listParticipantsTotal"` // 参与排名企业数
	ListPublishedDate     string `json:"listPublishedDate"`     // 排名发布日期
	ListUrlQcc            string `json:"listUrlQcc"`            // 排名url(qcc)
	ListUrlOrigin         string `json:"listUrlOrigin"`         // 排名url(来源)
	StatusCode            int64  `json:"statusCode"`            // 状态标识码
	common.ControlBy
}

func (s *RankingListInsertReq) Generate(model *models.RankingList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ListTitle = s.ListTitle
	model.ListType = s.ListType
	model.ListSource = s.ListSource
	model.ListParticipantsTotal = s.ListParticipantsTotal
	model.ListPublishedDate = s.ListPublishedDate
	model.ListUrlQcc = s.ListUrlQcc
	model.ListUrlOrigin = s.ListUrlOrigin
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *RankingListInsertReq) GetId() interface{} {
	return s.Id
}

type RankingListUpdateReq struct {
	Id                    int64  `uri:"id"`                     // 主键id
	ListTitle             string `json:"listTitle"`             // 榜单名称
	ListType              string `json:"listType"`              // 榜单类型(品牌产品榜,企业榜,...)
	ListSource            string `json:"listSource"`            // 排名来源(JsonArray格式;eg:["德本咨询", "eNet研究院", "互联网周刊"]
	ListParticipantsTotal int64  `json:"listParticipantsTotal"` // 参与排名企业数
	ListPublishedDate     string `json:"listPublishedDate"`     // 排名发布日期
	ListUrlQcc            string `json:"listUrlQcc"`            // 排名url(qcc)
	ListUrlOrigin         string `json:"listUrlOrigin"`         // 排名url(来源)
	StatusCode            int64  `json:"statusCode"`            // 状态标识码
	common.ControlBy
}

func (s *RankingListUpdateReq) Generate(model *models.RankingList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ListTitle = s.ListTitle
	model.ListType = s.ListType
	model.ListSource = s.ListSource
	model.ListParticipantsTotal = s.ListParticipantsTotal
	model.ListPublishedDate = s.ListPublishedDate
	model.ListUrlQcc = s.ListUrlQcc
	model.ListUrlOrigin = s.ListUrlOrigin
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *RankingListUpdateReq) GetId() interface{} {
	return s.Id
}

// RankingListGetReq 功能获取请求参数
type RankingListGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RankingListGetReq) GetId() interface{} {
	return s.Id
}

// RankingListDeleteReq 功能删除请求参数
type RankingListDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RankingListDeleteReq) GetId() interface{} {
	return s.Ids
}
