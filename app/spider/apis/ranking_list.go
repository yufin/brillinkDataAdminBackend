package apis

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/models"
	"go-admin/app/spider/service"
	"go-admin/app/spider/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type RankingList struct {
	apis.Api
}

// GetPage 获取排名榜单表列表
// @Summary 获取排名榜单表列表
// @Description 获取排名榜单表列表
// @Tags 排名榜单表
// @Param id query int64 false "主键id"
// @Param listTitle query string false "榜单名称"
// @Param listType query string false "榜单类型(品牌产品榜,企业榜,...)"
// @Param listSource query string false "排名来源(JsonArray格式;eg:["德本咨询", "eNet研究院", "互联网周刊"]"
// @Param listParticipantsTotal query int64 false "参与排名企业数"
// @Param listPublishedDate query string false "排名发布日期"
// @Param listUrlQcc query string false "排名url(qcc)"
// @Param listUrlOrigin query string false "排名url(来源)"
// @Param statusCode query int64 false "状态标识码"
// @Param createdAt query string false "创建时间"
// @Param updatedAt query string false "更新时间"
// @Param deletedAt query string false "删除时间"
// @Param createBy query int64 false "创建人"
// @Param updateBy query int64 false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RankingList}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ranking-list [get]
// @Security Bearer
func (e RankingList) GetPage(c *gin.Context) {
	req := dto.RankingListGetPageReq{}
	s := service.RankingList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRankingListFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RankingList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRankingListFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取排名榜单表
// @Summary 获取排名榜单表
// @Description 获取排名榜单表
// @Tags 排名榜单表
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RankingList} "{"code": 200, "data": [...]}"
// @Router /api/v1/ranking-list/{id} [get]
// @Security Bearer
func (e RankingList) Get(c *gin.Context) {
	req := dto.RankingListGetReq{}
	resp := dto.RankingListGetResp{}
	s := service.RankingList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRankingListFail", err))
		return
	}
	var object models.RankingList

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRankingListFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建排名榜单表
// @Summary 创建排名榜单表
// @Description 创建排名榜单表
// @Tags 排名榜单表
// @Accept application/json
// @Product application/json
// @Param data body dto.RankingListInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ranking-list [post]
// @Security Bearer
func (e RankingList) Insert(c *gin.Context) {
	req := dto.RankingListInsertReq{}
	s := service.RankingList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRankingListFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRankingListFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改排名榜单表
// @Summary 修改排名榜单表
// @Description 修改排名榜单表
// @Tags 排名榜单表
// @Accept application/json
// @Product application/json
// @Param data body dto.RankingListUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ranking-list/{id} [put]
// @Security Bearer
func (e RankingList) Update(c *gin.Context) {
	req := dto.RankingListUpdateReq{}
	s := service.RankingList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRankingListFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRankingListFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除排名榜单表
// @Summary 删除排名榜单表
// @Description 删除排名榜单表
// @Tags 排名榜单表
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ranking-list [delete]
// @Security Bearer
func (e RankingList) Delete(c *gin.Context) {
	s := service.RankingList{}
	req := dto.RankingListDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRankingListFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRankingListFail", err))
		return
	}
	e.OK(req.GetId())
}
