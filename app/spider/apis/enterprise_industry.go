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

type EnterpriseIndustry struct {
	apis.Api
}

// GetPage 获取企业行业分类表列表
// @Summary 获取企业行业分类表列表
// @Description 获取企业行业分类表列表
// @Tags 企业行业分类表
// @Param indId query int64 false "主键"
// @Param uscId query string false "社会统一信用代码"
// @Param industryData query string false "json格式的行业分类"
// @Param statusCode query int false "状态标识码"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseIndustry}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-industry [get]
// @Security Bearer
func (e EnterpriseIndustry) GetPage(c *gin.Context) {
	req := dto.EnterpriseIndustryGetPageReq{}
	s := service.EnterpriseIndustry{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseIndustryFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseIndustry, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseIndustryFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取企业行业分类表
// @Summary 获取企业行业分类表
// @Description 获取企业行业分类表
// @Tags 企业行业分类表
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseIndustry} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-industry/{id} [get]
// @Security Bearer
func (e EnterpriseIndustry) Get(c *gin.Context) {
	req := dto.EnterpriseIndustryGetReq{}
	resp := dto.EnterpriseIndustryGetResp{}
	s := service.EnterpriseIndustry{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseIndustryFail", err))
		return
	}
	var object models.EnterpriseIndustry

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseIndustryFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建企业行业分类表
// @Summary 创建企业行业分类表
// @Description 创建企业行业分类表
// @Tags 企业行业分类表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseIndustryInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-industry [post]
// @Security Bearer
func (e EnterpriseIndustry) Insert(c *gin.Context) {
	req := dto.EnterpriseIndustryInsertReq{}
	s := service.EnterpriseIndustry{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseIndustryFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseIndustryFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改企业行业分类表
// @Summary 修改企业行业分类表
// @Description 修改企业行业分类表
// @Tags 企业行业分类表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseIndustryUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-industry/{id} [put]
// @Security Bearer
func (e EnterpriseIndustry) Update(c *gin.Context) {
	req := dto.EnterpriseIndustryUpdateReq{}
	s := service.EnterpriseIndustry{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseIndustryFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseIndustryFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除企业行业分类表
// @Summary 删除企业行业分类表
// @Description 删除企业行业分类表
// @Tags 企业行业分类表
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-industry [delete]
// @Security Bearer
func (e EnterpriseIndustry) Delete(c *gin.Context) {
	s := service.EnterpriseIndustry{}
	req := dto.EnterpriseIndustryDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseIndustryFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseIndustryFail", err))
		return
	}
	e.OK(req.GetId())
}
