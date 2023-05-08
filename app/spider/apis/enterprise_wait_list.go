package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go-admin/app/spider/models"
	"go-admin/app/spider/service"
	"go-admin/app/spider/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	dtoCommon "go-admin/common/dto"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
	"math"
	"net/url"
)

type EnterpriseWaitList struct {
	apis.Api
}

// UpdateStatusCode 更新状态码
// @Summary 更新状态码, 3检查info表信息是否健全,4.检查certification,industry,product,ranking表信息是否健全.
func (e EnterpriseWaitList) UpdateStatusCode(c *gin.Context) {

}

// UpdateQccUrls 插入匹配的url
// @Summary 通过id与名称插入匹配的url,
// @Description 修改待爬取列表
// @Tags 待爬取列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseWaitListUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-wait-list/qccUrl/{id} [put]
// @Security Bearer
func (e EnterpriseWaitList) UpdateQccUrls(c *gin.Context) {
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
	u, err := url.Parse(req.QccUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		e.Logger.Error(err)
		panic(exception.WithMsg(500, "NotValidUrl", err))
		return
	}

	var object models.EnterpriseWaitList
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&dto.EnterpriseWaitListGetReq{Id: req.Id}, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseWaitListFailWhileUpdateQccUrl", err))
		return
	}

	req.StatusCode = int(math.Max(float64(object.StatusCode), float64(2)))
	req.SetUpdateBy(int64(user.GetUserId(c)))
	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	e.OK(req.GetId())
}

// GetEnterprisePageWaitingForMatch 获取企业等待匹配列表
// @Summary 获取企业等待匹配url的列表: 条件: qccUrl为空字符串, statusCode=0
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Param statusCode query int 状态码: 1.等待匹配url 2.等待爬取主体信息(enterprise_info), 3.等待爬取其他信息(tag,industry...)，4.完成爬取
// @Router /api/v1/enterprise-wait-lit/waiting [get]
func (e EnterpriseWaitList) GetEnterprisePageWaitingForMatch(c *gin.Context) {

	paginationReq := dtoCommon.Pagination{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().Bind(&paginationReq).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}
	//statusCodeParam, err := strconv.Atoi(c.Query("statusCode"))
	//if err != nil {
	//	e.Logger.Error(err)
	//	panic(exception.WithMsg(500, "QueryParamStatusCodeParseFail", err))
	//	return
	//}
	req := dto.EnterpriseWaitListGetPageReq{
		Pagination: paginationReq,
		StatusCode: 1,
		//StatusCode: statusCodeParam,
		//QccUrl: "-",
	}
	req.EnterpriseWaitListPageOrder.PriorityOrder = "desc"

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseWaitList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}

	respList := make([]dto.EnterpriseWaitListWaitingGetPageResp, 0)
	for _, v := range list {
		resp := dto.EnterpriseWaitListWaitingGetPageResp{}
		resp.Generate(&v)
		respList = append(respList, resp)
	}

	e.PageOK(respList, count, paginationReq.GetPageIndex(), paginationReq.GetPageSize())
}

// GetPage 获取待爬取列表列表
// @Summary 获取待爬取列表列表
// @Description 获取待爬取列表列表
// @Tags 待爬取列表
// @Param id query int64 false "主键id"
// @Param enterpriseName query string false "企业名称"
// @Param uscId query string false "社会统一信用代码"
// @Param priority query int false "优先级"
// @Param qccUrl query string false "qcc主体网址"
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

	if req.Priority == 0 {
		req.Priority = 1
	}

	u, err := url.Parse(req.QccUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		req.StatusCode = 1 // not valid url
	} else {
		req.StatusCode = 2
	}
	req.SetCreateBy(int64(user.GetUserId(c)))

	p := actions.GetPermissionFromContext(c)
	err = s.Insert(&req)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			// column not unique
			list := make([]models.EnterpriseWaitList, 0)
			var count int64
			err = s.GetPage(&dto.EnterpriseWaitListGetPageReq{
				EnterpriseName: req.EnterpriseName, UscId: req.UscId}, p, &list, &count)
			if err != nil {
				e.Logger.Error(err)
				panic(exception.WithMsg(50000, "InsertEnterpriseWaitListFail", err))
				return
			}
			if len(list) > 0 {
				e.OK(list[0])
				return
			}
		} else {
			e.Logger.Error(err)
			panic(exception.WithMsg(500, "InsertEnterpriseWaitListFail", err))
			return
		}
	}
	e.OK(req)
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
