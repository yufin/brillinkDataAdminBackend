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

type EnterpriseWaitList struct {
	apis.Api
}

// GetPage 获取待爬取列表列表
// @Summary 获取待爬取列表列表
// @Description 获取待爬取列表列表
// @Tags 待爬取列表
// @Param id query int64 false "主键id"
// @Param enterpriseName query string false "企业名称"
// @Param unifiedSocialCreditCode query string false "纳税人识别号"
// @Param priority query int false "优先级"
// @Param qccUrl query string false "qcc主体网址"
// @Param enterpriseId query string false "企业uuid4"
// @Param statusCode query int false "数据爬取状态码"
// @Param source query string false "来源备注"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param createdAt query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseWaitList}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-wait-list [get]
// @Security Bearer
func (e EnterpriseWaitList) GetPage(c *gin.Context) {
	req := dto.EnterpriseWaitListGetPageReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseWaitList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取待爬取列表
// @Summary 获取待爬取列表
// @Description 获取待爬取列表
// @Tags 待爬取列表
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseWaitList} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-wait-list/{id} [get]
// @Security Bearer
func (e EnterpriseWaitList) Get(c *gin.Context) {
	req := dto.EnterpriseWaitListGetReq{}
	resp := dto.EnterpriseWaitListGetResp{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseWaitListFail", err))
		return
	}
	var object models.EnterpriseWaitList

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseWaitListFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建待爬取列表
// @Summary 创建待爬取列表
// @Description 创建待爬取列表
// @Tags 待爬取列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseWaitListInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-wait-list [post]
// @Security Bearer
func (e EnterpriseWaitList) Insert(c *gin.Context) {
	req := dto.EnterpriseWaitListInsertReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseWaitListFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseWaitListFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改待爬取列表
// @Summary 修改待爬取列表
// @Description 修改待爬取列表
// @Tags 待爬取列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseWaitListUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-wait-list/{id} [put]
// @Security Bearer
func (e EnterpriseWaitList) Update(c *gin.Context) {
	req := dto.EnterpriseWaitListUpdateReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除待爬取列表
// @Summary 删除待爬取列表
// @Description 删除待爬取列表
// @Tags 待爬取列表
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-wait-list [delete]
// @Security Bearer
func (e EnterpriseWaitList) Delete(c *gin.Context) {
	s := service.EnterpriseWaitList{}
	req := dto.EnterpriseWaitListDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseWaitListFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseWaitListFail", err))
		return
	}
	e.OK(req.GetId())
}
