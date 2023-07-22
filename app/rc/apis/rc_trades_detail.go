package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/rc/models"
	"go-admin/app/rc/service"
	"go-admin/app/rc/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type RcTradesDetail struct {
	apis.Api
}

func (e RcTradesDetail) GetJoin(c *gin.Context) {
	req := dto.RcTradesDetailGetPageReq{}
	s := service.RcTradesDetail{}
	err := e.MakeContext(c).Bind(&req).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(500, "GetJoinRcTradesDetailFail", err))
		return
	}
	p := actions.GetPermissionFromContext(c)
	resp := make([]dto.RcTradesDetailJoinWaitListResp, 0)
	err = s.GetJoinWaitList(req.ContentId, &resp, p)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(500, "GetJoinRcTradesDetailFail", err))
	}
	e.OK(resp)
}

// GetPage 获取客户、供应商交易细节（来自origin_content表)列表
// @Summary 获取客户、供应商交易细节（来自origin_content表)列表
// @Description 获取客户、供应商交易细节（来自origin_content表)列表
// @Tags 客户、供应商交易细节（来自origin_content表)
// @Param id query int64 false "主键"
// @Param contentId query string false "外键"
// @Param enterpriseName query string false "企业名称"
// @Param commodityRatio query string false "货物占比"
// @Param commodityName query string false "货物种类名称"
// @Param ratioAmountTax query string false "税额占比"
// @Param sumAmountTax query string false "税总额"
// @Param detailType query int false "1:customer_12,2:customer_24,3:supplier_12,4:supplier_24"
// @Param tagIndustry query string false "行业标签"
// @Param tagAuthorized query string false "认证标签"
// @Param tagProduct query string false "产品标签"
// @Param tagList query string false "榜单标签"
// @Param enterpriseInfo query string false "企业信息"
// @Param statusCode query string false "状态码"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcTradesDetail}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-trades-detail [get]
// @Security Bearer
func (e RcTradesDetail) GetPage(c *gin.Context) {
	req := dto.RcTradesDetailGetPageReq{}
	s := service.RcTradesDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcTradesDetailFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcTradesDetail, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcTradesDetailFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取客户、供应商交易细节（来自origin_content表)
// @Summary 获取客户、供应商交易细节（来自origin_content表)
// @Description 获取客户、供应商交易细节（来自origin_content表)
// @Tags 客户、供应商交易细节（来自origin_content表)
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcTradesDetail} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-trades-detail/{id} [get]
// @Security Bearer
func (e RcTradesDetail) Get(c *gin.Context) {
	req := dto.RcTradesDetailGetReq{}
	resp := dto.RcTradesDetailGetResp{}
	s := service.RcTradesDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcTradesDetailFail", err))
		return
	}
	var object models.RcTradesDetail

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcTradesDetailFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建客户、供应商交易细节（来自origin_content表)
// @Summary 创建客户、供应商交易细节（来自origin_content表)
// @Description 创建客户、供应商交易细节（来自origin_content表)
// @Tags 客户、供应商交易细节（来自origin_content表)
// @Accept application/json
// @Product application/json
// @Param data body dto.RcTradesDetailInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-trades-detail [post]
// @Security Bearer
func (e RcTradesDetail) Insert(c *gin.Context) {
	req := dto.RcTradesDetailInsertReq{}
	s := service.RcTradesDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcTradesDetailFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcTradesDetailFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改客户、供应商交易细节（来自origin_content表)
// @Summary 修改客户、供应商交易细节（来自origin_content表)
// @Description 修改客户、供应商交易细节（来自origin_content表)
// @Tags 客户、供应商交易细节（来自origin_content表)
// @Accept application/json
// @Product application/json
// @Param data body dto.RcTradesDetailUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-trades-detail/{id} [put]
// @Security Bearer
func (e RcTradesDetail) Update(c *gin.Context) {
	req := dto.RcTradesDetailUpdateReq{}
	s := service.RcTradesDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcTradesDetailFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcTradesDetailFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除客户、供应商交易细节（来自origin_content表)
// @Summary 删除客户、供应商交易细节（来自origin_content表)
// @Description 删除客户、供应商交易细节（来自origin_content表)
// @Tags 客户、供应商交易细节（来自origin_content表)
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rc-trades-detail [delete]
// @Security Bearer
func (e RcTradesDetail) Delete(c *gin.Context) {
	s := service.RcTradesDetail{}
	req := dto.RcTradesDetailDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcTradesDetailFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcTradesDetailFail", err))
		return
	}
	e.OK(req.GetId())
}

//// TaskSyncTradesDetail 同步客户、供应商交易细节
//func (e RcTradesDetail) TaskSyncTradesDetail(c *gin.Context) {
//	s := service.RcTradesDetail{}
//	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//	sContent := service.RcOriginContent{}
//	err = e.MakeContext(c).MakeOrm().MakeService(&sContent.Service).Errors
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//
//	p := actions.GetPermissionFromContext(c)
//	err = task.SyncTradesDetail(&s, &sContent, p)
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//	e.OK(nil)
//}

//// TaskSyncWaitList 同步未爬取数据至待爬取列表
//func (e RcTradesDetail) TaskSyncWaitList(c *gin.Context) {
//	sDetail := service.RcTradesDetail{}
//	err := e.MakeContext(c).MakeOrm().MakeService(&sDetail.Service).Errors
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//	sContent := service.RcOriginContent{}
//	err = e.MakeContext(c).MakeOrm().MakeService(&sContent.Service).Errors
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//	sWait := service2.EnterpriseWaitList{}
//	err = e.MakeContext(c).MakeOrm().MakeService(&sWait.Service).Errors
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//
//	p := actions.GetPermissionFromContext(c)
//	sw := task.ServicesWrap{
//		SWait:    &sWait,
//		SContent: &sContent,
//		SDetail:  &sDetail,
//		P:        p,
//	}
//	err = task.SyncWaitList(&sw)
//	if err != nil {
//		e.Logger.Error(err)
//		panic(exception.WithMsg(50000, "TaskSyncTradesDetailFail", err))
//		return
//	}
//	e.OK(nil)
//}
