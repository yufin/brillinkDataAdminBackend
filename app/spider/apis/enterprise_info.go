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

type EnterpriseInfo struct {
	apis.Api
}

// GetPage 获取企业主体信息列表
// @Summary 获取企业主体信息列表
// @Description 获取企业主体信息列表
// @Tags 企业主体信息
// @Param infoId query int64 false "主键"
// @Param enterpriseTitle query string false "企业名称"
// @Param enterpriseTitleEn query string false "企业英文名称"
// @Param businessRegistrationNumber query string false "工商注册号"
// @Param establishedDate query time.Time false "成立日期"
// @Param region query string false "所属地区"
// @Param approvedDate query time.Time false "核准日期"
// @Param registeredAddress query string false "注册地址"
// @Param registeredCapital query float64 false "注册资本"
// @Param registeredCapitalCurrency query string false "注册资本币种"
// @Param paidInCapital query float64 false "实缴资本"
// @Param paidInCapitalCurrency query string false "实缴资本币种"
// @Param enterpriseType query string false "企业类型"
// @Param stuffSize query string false "人员规模"
// @Param stuffInsuredNumber query int false "参保人数"
// @Param businessScope query string false "经营范围"
// @Param importExportQualificationCode query string false "进出口企业代码"
// @Param legalRepresentative query string false "法定代表人"
// @Param registrationAuthority query string false "登记机关"
// @Param registrationStatus query string false "登记状态"
// @Param taxpayerQualification query string false "纳税人资质"
// @Param organizationCode query string false "组织机构代码"
// @Param urlQcc query string false "企查查url"
// @Param urlHomepage query string false "官网url"
// @Param businessTermStart query time.Time false "营业期限开始"
// @Param businessTermEnd query time.Time false "营业期限结束"
// @Param uscId query string false "社会统一信用代码"
// @Param statusCode query int false "状态标识码"
// @Param createdAt query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param createBy query int64 false "创建人"
// @Param updateBy query int64 false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-info [get]
// @Security Bearer
func (e EnterpriseInfo) GetPage(c *gin.Context) {
	req := dto.EnterpriseInfoGetPageReq{}
	s := service.EnterpriseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseInfoFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseInfoFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取企业主体信息
// @Summary 获取企业主体信息
// @Description 获取企业主体信息
// @Tags 企业主体信息
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-info/{id} [get]
// @Security Bearer
func (e EnterpriseInfo) Get(c *gin.Context) {
	req := dto.EnterpriseInfoGetReq{}
	resp := dto.EnterpriseInfoGetResp{}
	s := service.EnterpriseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseInfoFail", err))
		return
	}
	var object models.EnterpriseInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseInfoFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建企业主体信息
// @Summary 创建企业主体信息
// @Description 创建企业主体信息
// @Tags 企业主体信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseInfoInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-info [post]
// @Security Bearer
func (e EnterpriseInfo) Insert(c *gin.Context) {
	req := dto.EnterpriseInfoInsertReq{}
	s := service.EnterpriseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseInfoFail", err))
		return
	}
	req.StatusCode = 1
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseInfoFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改企业主体信息
// @Summary 修改企业主体信息
// @Description 修改企业主体信息
// @Tags 企业主体信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseInfoUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-info/{id} [put]
// @Security Bearer
func (e EnterpriseInfo) Update(c *gin.Context) {
	req := dto.EnterpriseInfoUpdateReq{}
	s := service.EnterpriseInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseInfoFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseInfoFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除企业主体信息
// @Summary 删除企业主体信息
// @Description 删除企业主体信息
// @Tags 企业主体信息
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-info [delete]
// @Security Bearer
func (e EnterpriseInfo) Delete(c *gin.Context) {
	s := service.EnterpriseInfo{}
	req := dto.EnterpriseInfoDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseInfoFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseInfoFail", err))
		return
	}
	e.OK(req.GetId())
}

func (e EnterpriseInfo) GetpageUnion
