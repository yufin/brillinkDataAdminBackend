package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	_ "go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type SysApi struct {
	apis.Api
}

// GetPage 获取接口管理列表
// @Summary      获取接口管理列表
// @Description  获取接口管理列表
// @Tags         接口管理
// @Param        name       query     string                                                false  "名称"
// @Param        title      query     string                                                false  "标题"
// @Param        path       query     string                                                false  "地址"
// @Param        action     query     string                                                false  "类型"
// @Param        pageSize   query     int                                                   false  "页条数"
// @Param        pageIndex  query     int                                                   false  "页码"
// @Success      200        {object}  antd.Response{data=antd.Pages{list=[]models.SysApi}}  "{"code": 200, "data": [...]}"
// @Router       /api/v1/sys-api [get]
// @Security     Bearer
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	//数据权限检查
	p := actions.GetPermission(c)
	list, count, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}
	e.PageOK(list, *count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取接口管理
// @Summary      获取接口管理
// @Description  获取接口管理
// @Tags         接口管理
// @Param        id   path      string                             false  "id"
// @Success      200  {object}  antd.Response{data=models.SysApi}  "{"code": 200, "data": [...]}"
// @Router       /api/v1/sys-api/{id} [get]
// @Security     Bearer
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	p := actions.GetPermission(c)
	object, err := s.Get(&req, p)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}
	e.OK(object)
}

// Update 修改接口管理
// @Summary      修改接口管理
// @Description  修改接口管理
// @Tags         接口管理
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.SysApiUpdateReq  true  "body"
// @Success      200   {object}  antd.Response        "{"code": 200, "message": "修改成功"}"
// @Router       /api/v1/sys-api/{id} [put]
// @Security     Bearer
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermission(c)
	err = s.Update(c, &req, p)
	if err != nil {
		e.Error(500, "更新失败", "")
		return
	}
	e.OK(req.GetId())
}

// Delete 删除接口管理
// @Summary      删除接口管理
// @Description  删除接口管理
// @Tags         接口管理
// @Param        data  body      dto.SysApiDeleteReq  true  "body"
// @Success      200   {object}  antd.Response        "{"code": 200, "message": "删除成功"}"
// @Router       /api/v1/sys-api [delete]
// @Security     Bearer
func (e SysApi) Delete(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermission(c)
	err = s.Remove(c, &req, p)
	if err != nil {
		e.Error(500, "删除失败", "")
		return
	}
	e.OK(req.GetId())
}

// GetList 获取接口列表配置菜单接口使用
// @Summary      获取接口列表配置菜单接口使用
// @Description  获取接口列表配置菜单接口使用
// @Tags         接口管理
// @Param        name       query     string                                                false  "名称"
// @Param        title      query     string                                                false  "标题"
// @Param        path       query     string                                                false  "地址"
// @Param        action     query     string                                                false  "类型"
// @Param        pageSize   query     int                                                   false  "页条数"
// @Param        pageIndex  query     int                                                   false  "页码"
// @Success      200        {object}  antd.Response{data=antd.Pages{list=[]models.SysApi}}  "{"code": 200, "data": [...]}"
// @Router       /api/v1/sys-api/list [get]
// @Security     Bearer
func (e SysApi) GetList(c *gin.Context) {
	req := dto.SysApiGetPageReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	//数据权限检查
	p := actions.GetPermission(c)
	//var count int64
	list, err := s.GetList(&req, p)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}
	e.ListOK(list)
}
