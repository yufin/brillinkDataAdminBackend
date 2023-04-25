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

type Enterprise struct {
	apis.Api
}

// GetPage 获取企业主体列表
// @Summary 获取企业主体列表
// @Description 获取企业主体列表
// @Tags 企业主体
// @Param id query int64 false "主键"
// @Param uscId query string false "统一社会信用代码"
// @Param statusCode query int false "状态标识码"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param created_at query time.Time false "创建时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.Enterprise}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise [get]
// @Security Bearer
func (e Enterprise) GetPage(c *gin.Context) {
	req := dto.EnterpriseGetPageReq{}
	s := service.Enterprise{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Enterprise, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取企业主体
// @Summary 获取企业主体
// @Description 获取企业主体
// @Tags 企业主体
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.Enterprise} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise/{id} [get]
// @Security Bearer
func (e Enterprise) Get(c *gin.Context) {
	req := dto.EnterpriseGetReq{}
	resp := dto.EnterpriseGetResp{}
	s := service.Enterprise{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseFail", err))
		return
	}
	var object models.Enterprise

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建企业主体
// @Summary 创建企业主体
// @Description 创建企业主体
// @Tags 企业主体
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise [post]
// @Security Bearer
func (e Enterprise) Insert(c *gin.Context) {
	req := dto.EnterpriseInsertReq{}
	s := service.Enterprise{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改企业主体
// @Summary 修改企业主体
// @Description 修改企业主体
// @Tags 企业主体
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise/{id} [put]
// @Security Bearer
func (e Enterprise) Update(c *gin.Context) {
	req := dto.EnterpriseUpdateReq{}
	s := service.Enterprise{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除企业主体
// @Summary 删除企业主体
// @Description 删除企业主体
// @Tags 企业主体
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise [delete]
// @Security Bearer
func (e Enterprise) Delete(c *gin.Context) {
	s := service.Enterprise{}
	req := dto.EnterpriseDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseFail", err))
		return
	}
	e.OK(req.GetId())
}
