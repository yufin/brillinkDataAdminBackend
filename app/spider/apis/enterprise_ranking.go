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

type EnterpriseRanking struct {
	apis.Api
}

// GetPage 获取企业榜单排名数据列表
// @Summary 获取企业榜单排名数据列表
// @Description 获取企业榜单排名数据列表
// @Tags 企业榜单排名数据
// @Param rankId query int64 false "主键"
// @Param enterpriseId query int64 false "外键(enterprise表的id)"
// @Param rankingListId query int64 false "外键(enterprise_ranking_list表的id)"
// @Param rankingPosition query int false "榜内位置"
// @Param rankingEnterpriseTitle query string false "榜单中的企业名称"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseRanking}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-ranking [get]
// @Security Bearer
func (e EnterpriseRanking) GetPage(c *gin.Context) {
	req := dto.EnterpriseRankingGetPageReq{}
	s := service.EnterpriseRanking{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseRankingFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseRanking, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseRankingFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取企业榜单排名数据
// @Summary 获取企业榜单排名数据
// @Description 获取企业榜单排名数据
// @Tags 企业榜单排名数据
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseRanking} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-ranking/{id} [get]
// @Security Bearer
func (e EnterpriseRanking) Get(c *gin.Context) {
	req := dto.EnterpriseRankingGetReq{}
	resp := dto.EnterpriseRankingGetResp{}
	s := service.EnterpriseRanking{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseRankingFail", err))
		return
	}
	var object models.EnterpriseRanking

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseRankingFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建企业榜单排名数据
// @Summary 创建企业榜单排名数据
// @Description 创建企业榜单排名数据
// @Tags 企业榜单排名数据
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseRankingInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-ranking [post]
// @Security Bearer
func (e EnterpriseRanking) Insert(c *gin.Context) {
	req := dto.EnterpriseRankingInsertReq{}
	s := service.EnterpriseRanking{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseRankingFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseRankingFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改企业榜单排名数据
// @Summary 修改企业榜单排名数据
// @Description 修改企业榜单排名数据
// @Tags 企业榜单排名数据
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseRankingUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-ranking/{id} [put]
// @Security Bearer
func (e EnterpriseRanking) Update(c *gin.Context) {
	req := dto.EnterpriseRankingUpdateReq{}
	s := service.EnterpriseRanking{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseRankingFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseRankingFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除企业榜单排名数据
// @Summary 删除企业榜单排名数据
// @Description 删除企业榜单排名数据
// @Tags 企业榜单排名数据
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-ranking [delete]
// @Security Bearer
func (e EnterpriseRanking) Delete(c *gin.Context) {
	s := service.EnterpriseRanking{}
	req := dto.EnterpriseRankingDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseRankingFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseRankingFail", err))
		return
	}
	e.OK(req.GetId())
}
