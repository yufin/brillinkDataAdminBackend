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

type EnterpriseCertification struct {
	apis.Api
}

// GetPage 获取列表
// @Summary 获取列表
// @Description 获取列表
// @Tags
// @Param certId query int64 false "主键"
// @Param certificationTitle query string false "认证名称"
// @Param certificationCode query string false "认证编号"
// @Param certificationLevel query string false "认证等级(省级,市级,国家级)"
// @Param certificationType query string false "认证类型(荣誉,科技型企业,...)"
// @Param certificationSource query string false "认证来源(eg:2022年度浙江省中小企业公共服务示范平台名单)"
// @Param certificationDate query time.Time false "发证日期"
// @Param certificationTermStart query time.Time false "有效期起"
// @Param certificationTermEnd query time.Time false "有效期至"
// @Param certificationAuthority query string false "发证机关"
// @Param enterpriseId query int64 false "外键(enterprise表的id)"
// @Param statusCode query int64 false "状态标识码"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param createBy query int64 false "创建人"
// @Param updateBy query int64 false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseCertification}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-certification [get]
// @Security Bearer
func (e EnterpriseCertification) GetPage(c *gin.Context) {
	req := dto.EnterpriseCertificationGetPageReq{}
	s := service.EnterpriseCertification{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseCertificationFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseCertification, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseCertificationFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取
// @Summary 获取
// @Description 获取
// @Tags
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseCertification} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-certification/{id} [get]
// @Security Bearer
func (e EnterpriseCertification) Get(c *gin.Context) {
	req := dto.EnterpriseCertificationGetReq{}
	resp := dto.EnterpriseCertificationGetResp{}
	s := service.EnterpriseCertification{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseCertificationFail", err))
		return
	}
	var object models.EnterpriseCertification

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseCertificationFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建
// @Summary 创建
// @Description 创建
// @Tags
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseCertificationInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-certification [post]
// @Security Bearer
func (e EnterpriseCertification) Insert(c *gin.Context) {
	req := dto.EnterpriseCertificationInsertReq{}
	s := service.EnterpriseCertification{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseCertificationFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertEnterpriseCertificationFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改
// @Summary 修改
// @Description 修改
// @Tags
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseCertificationUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-certification/{id} [put]
// @Security Bearer
func (e EnterpriseCertification) Update(c *gin.Context) {
	req := dto.EnterpriseCertificationUpdateReq{}
	s := service.EnterpriseCertification{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseCertificationFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseCertificationFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除
// @Summary 删除
// @Description 删除
// @Tags
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-certification [delete]
// @Security Bearer
func (e EnterpriseCertification) Delete(c *gin.Context) {
	s := service.EnterpriseCertification{}
	req := dto.EnterpriseCertificationDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseCertificationFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseCertificationFail", err))
		return
	}
	e.OK(req.GetId())
}
