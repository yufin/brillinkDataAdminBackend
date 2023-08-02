package apis

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/cms/models"
	"go-admin/app/cms/service"
	"go-admin/app/cms/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type TbCmsPage struct {
	apis.Api
}

// GetPage 获取列表
// @Summary      获取列表
// @Description  获取列表
// @Tags
// @Param     id         query     int64                                                    false  "主键编码"
// @Param     title      query     string                                                   false  "页面名称"
// @Param     mark       query     string                                                   false  "页面标记"
// @Param     source     query     string                                                   false  "引用来源"
// @Param     author     query     string                                                   false  "作者"
// @Param     content    query     string                                                   false  "内容"
// @Param     pubTime    query     time.Time                                                false  "发布时间"
// @Param     createdAt  query     time.Time                                                false  "创建时间"
// @Param     updatedAt  query     time.Time                                                false  "最后更新时间"
// @Param     deletedAt  query     time.Time                                                false  "删除时间"
// @Param     createBy   query     int64                                                    false  "创建者"
// @Param     updateBy   query     int64                                                    false  "更新者"
// @Param     pageSize   query     int                                                      false  "页条数"
// @Param     pageIndex  query     int                                                      false  "页码"
// @Success   200        {object}  antd.Response{data=antd.Pages{list=[]models.TbCmsPage}}  "{"code": 200, "data": [...]}"
// @Router    /api/v1/tb-cms-page [get]
// @Security  Bearer
func (e TbCmsPage) GetPage(c *gin.Context) {
	req := dto.TbCmsPageGetPageReq{}
	s := service.TbCmsPage{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTbCmsPageFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.TbCmsPage, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTbCmsPageFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取
// @Summary      获取
// @Description  获取
// @Tags
// @Param     id   path      string                                false  "id"
// @Success   200  {object}  antd.Response{data=models.TbCmsPage}  "{"code": 200, "data": [...]}"
// @Router    /api/v1/tb-cms-page/{id} [get]
// @Security  Bearer
func (e TbCmsPage) Get(c *gin.Context) {
	req := dto.TbCmsPageGetReq{}
	resp := dto.TbCmsPageGetResp{}
	s := service.TbCmsPage{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetTbCmsPageFail", err))
		return
	}
	var object models.TbCmsPage

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetTbCmsPageFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建
// @Summary      创建
// @Description  创建
// @Tags
// @Accept    application/json
// @Product   application/json
// @Param     data  body      dto.TbCmsPageInsertReq  true  "data"
// @Success   200   {object}  antd.Response           "{"code": 200, "message": "添加成功"}"
// @Router    /api/v1/tb-cms-page [post]
// @Security  Bearer
func (e TbCmsPage) Insert(c *gin.Context) {
	req := dto.TbCmsPageInsertReq{}
	s := service.TbCmsPage{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertTbCmsPageFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertTbCmsPageFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改
// @Summary      修改
// @Description  修改
// @Tags
// @Accept    application/json
// @Product   application/json
// @Param     data  body      dto.TbCmsPageUpdateReq  true  "body"
// @Success   200   {object}  antd.Response           "{"code": 200, "message": "修改成功"}"
// @Router    /api/v1/tb-cms-page/{id} [put]
// @Security  Bearer
func (e TbCmsPage) Update(c *gin.Context) {
	req := dto.TbCmsPageUpdateReq{}
	s := service.TbCmsPage{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateTbCmsPageFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateTbCmsPageFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除
// @Summary      删除
// @Description  删除
// @Tags
// @Param     ids  body      []int          false  "ids"
// @Success   200  {object}  antd.Response  "{"code": 200, "message": "删除成功"}"
// @Router    /api/v1/tb-cms-page [delete]
// @Security  Bearer
func (e TbCmsPage) Delete(c *gin.Context) {
	s := service.TbCmsPage{}
	req := dto.TbCmsPageDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteTbCmsPageFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteTbCmsPageFail", err))
		return
	}
	e.OK(req.GetId())
}
