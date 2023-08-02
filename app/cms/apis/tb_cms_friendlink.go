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

type TbCmsFriendlink struct {
	apis.Api
}

// GetPage 获取列表
// @Summary      获取列表
// @Description  获取列表
// @Tags
// @Param     id         query     int64                                                          false  "主键编码"
// @Param     name       query     string                                                         false  "链接名称"
// @Param     link       query     string                                                         false  "链接地址"
// @Param     createdAt  query     time.Time                                                      false  "创建时间"
// @Param     updatedAt  query     time.Time                                                      false  "最后更新时间"
// @Param     deletedAt  query     time.Time                                                      false  "删除时间"
// @Param     createBy   query     int64                                                          false  "创建者"
// @Param     updateBy   query     int64                                                          false  "更新者"
// @Param     pageSize   query     int                                                            false  "页条数"
// @Param     pageIndex  query     int                                                            false  "页码"
// @Success   200        {object}  antd.Response{data=antd.Pages{list=[]models.TbCmsFriendlink}}  "{"code": 200, "data": [...]}"
// @Router    /api/v1/tb-cms-friendlink [get]
// @Security  Bearer
func (e TbCmsFriendlink) GetPage(c *gin.Context) {
	req := dto.TbCmsFriendlinkGetPageReq{}
	s := service.TbCmsFriendlink{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTbCmsFriendlinkFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.TbCmsFriendlink, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTbCmsFriendlinkFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取
// @Summary      获取
// @Description  获取
// @Tags
// @Param     id   path      string                                      false  "id"
// @Success   200  {object}  antd.Response{data=models.TbCmsFriendlink}  "{"code": 200, "data": [...]}"
// @Router    /api/v1/tb-cms-friendlink/{id} [get]
// @Security  Bearer
func (e TbCmsFriendlink) Get(c *gin.Context) {
	req := dto.TbCmsFriendlinkGetReq{}
	resp := dto.TbCmsFriendlinkGetResp{}
	s := service.TbCmsFriendlink{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetTbCmsFriendlinkFail", err))
		return
	}
	var object models.TbCmsFriendlink

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetTbCmsFriendlinkFail", err))
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
// @Param     data  body      dto.TbCmsFriendlinkInsertReq  true  "data"
// @Success   200   {object}  antd.Response                 "{"code": 200, "message": "添加成功"}"
// @Router    /api/v1/tb-cms-friendlink [post]
// @Security  Bearer
func (e TbCmsFriendlink) Insert(c *gin.Context) {
	req := dto.TbCmsFriendlinkInsertReq{}
	s := service.TbCmsFriendlink{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertTbCmsFriendlinkFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertTbCmsFriendlinkFail", err))
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
// @Param     data  body      dto.TbCmsFriendlinkUpdateReq  true  "body"
// @Success   200   {object}  antd.Response                 "{"code": 200, "message": "修改成功"}"
// @Router    /api/v1/tb-cms-friendlink/{id} [put]
// @Security  Bearer
func (e TbCmsFriendlink) Update(c *gin.Context) {
	req := dto.TbCmsFriendlinkUpdateReq{}
	s := service.TbCmsFriendlink{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateTbCmsFriendlinkFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateTbCmsFriendlinkFail", err))
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
// @Router    /api/v1/tb-cms-friendlink [delete]
// @Security  Bearer
func (e TbCmsFriendlink) Delete(c *gin.Context) {
	s := service.TbCmsFriendlink{}
	req := dto.TbCmsFriendlinkDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteTbCmsFriendlinkFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteTbCmsFriendlinkFail", err))
		return
	}
	e.OK(req.GetId())
}
