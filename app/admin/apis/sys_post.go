package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type SysPost struct {
	apis.Api
}

// GetPage
// @Summary 岗位列表数据
// @Description 获取JSON
// @Tags 岗位
// @Param postName query string false "postName"
// @Param postCode query string false "postCode"
// @Param postId query string false "postId"
// @Param status query string false "status"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security Bearer
func (e SysPost) GetPage(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.GetPagePostFail, err))
		return
	}

	list, count, err := s.GetPage(&req)
	if err != nil {
		panic(exception.New(exception.GetPagePostFail, err))
		return
	}
	e.PageOK(list, *count, req.GetPageIndex(), req.GetPageSize())
}

// Get
// @Summary 获取岗位信息
// @Description 获取JSON
// @Tags 岗位
// @Param id path int true "编码"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{postId} [get]
// @Security Bearer
func (e SysPost) Get(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.GetPostFail, err))
		return
	}
	object, err := s.Get(&req)
	if err != nil {
		panic(exception.New(exception.GetPostFail, err))
		return
	}

	e.OK(object)
}

// Insert
// @Summary 添加岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysPostInsertReq true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/post [post]
// @Security Bearer
func (e SysPost) Insert(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.InsertPostFail, err))
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.New(exception.InsertPostFail, err))
		return
	}
	e.OK(req.GetId())
}

// Update
// @Summary 修改岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysPostUpdateReq true "body"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [put]
// @Security Bearer
func (e SysPost) Update(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.UpdatePostFail, err))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(c, &req)
	if err != nil {
		panic(exception.New(exception.UpdatePostFail, err))
		return
	}
	e.OK(req.GetId())
}

// Delete
// @Summary 删除岗位
// @Description 删除数据
// @Tags 岗位
// @Param id body dto.SysPostDeleteReq true "请求参数"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 500 {string} string	"{"code": 500, "message": "删除失败"}"
// @Router /api/v1/post [delete]
// @Security Bearer
func (e SysPost) Delete(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.DeletePostFail, err))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(c, &req)
	if err != nil {
		panic(exception.New(exception.DeletePostFail, err))
		return
	}
	e.OK(req.GetId())
}

// GetOption
// @Summary 职位下拉框数据
// @Description Get JSON
// @Tags 职位/Post
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post-option [get]
// @Security Bearer
func (e SysPost) GetOption(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.DeletePostFail, err))
		return
	}
	list, err := s.GetList(&req)
	if err != nil {
		panic(exception.New(exception.DeletePostFail, err))
		return
	}
	l := make([]dto.SysPostGetListResp, 0)
	for _, i := range *list {
		d := dto.SysPostGetListResp{}
		e.Translate(i, &d)
		l = append(l, d)
	}
	e.OK(l)
}
