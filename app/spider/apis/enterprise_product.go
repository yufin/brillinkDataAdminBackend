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

type EnterpriseProduct struct {
	apis.Api
}

// GetPage 获取企业产品列表
// @Summary 获取企业产品列表
// @Description 获取企业产品列表
// @Tags 企业产品
// @Param prodId query int64 false "主键"
// @Param enterpriseId query int64 false "外键-企业id"
// @Param productData query string false "json格式的产品分类"
// @Param statusCode query int false "状态码"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseProduct}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-product [get]
// @Security Bearer
func (e EnterpriseProduct) GetPage(c *gin.Context) {
	req := dto.EnterpriseProductGetPageReq{}
	s := service.EnterpriseProduct{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseProductFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseProduct, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseProductFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取企业产品
// @Summary 获取企业产品
// @Description 获取企业产品
// @Tags 企业产品
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseProduct} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-product/{id} [get]
// @Security Bearer
func (e EnterpriseProduct) Get(c *gin.Context) {
	req := dto.EnterpriseProductGetReq{}
	resp := dto.EnterpriseProductGetResp{}
	s := service.EnterpriseProduct{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseProductFail", err))
		return
	}
	var object models.EnterpriseProduct

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseProductFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建企业产品
// @Summary 创建企业产品
// @Description 创建企业产品
// @Tags 企业产品
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseProductInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-product [post]
// @Security Bearer
func (e EnterpriseProduct) Insert(c *gin.Context) {
	req := dto.EnterpriseProductInsertReq{}
	s := service.EnterpriseProduct{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseProductFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseProductFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改企业产品
// @Summary 修改企业产品
// @Description 修改企业产品
// @Tags 企业产品
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseProductUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-product/{id} [put]
// @Security Bearer
func (e EnterpriseProduct) Update(c *gin.Context) {
	req := dto.EnterpriseProductUpdateReq{}
	s := service.EnterpriseProduct{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseProductFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseProductFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除企业产品
// @Summary 删除企业产品
// @Description 删除企业产品
// @Tags 企业产品
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-product [delete]
// @Security Bearer
func (e EnterpriseProduct) Delete(c *gin.Context) {
	s := service.EnterpriseProduct{}
	req := dto.EnterpriseProductDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseProductFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseProductFail", err))
		return
	}
	e.OK(req.GetId())
}
