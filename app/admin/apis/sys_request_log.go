package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/apis"
	_ "go-admin/common/response/antd"
)

type SysRequestLog struct {
	apis.Api
}

// GetPage 操作日志列表
// @Summary 操作日志列表
// @Description 获取JSON
// @Tags 操作日志
// @Param title query string false "title"
// @Param method query string false "method"
// @Param requestMethod  query string false "requestMethod"
// @Param operUrl query string false "operUrl"
// @Param operIp query string false "operIp"
// @Param status query string false "status"
// @Param beginTime query string false "beginTime"
// @Param endTime query string false "endTime"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-request-log [get]
// @Security Bearer
func (e SysRequestLog) GetPage(c *gin.Context) {
	s := service.SysRequestLog{}
	req := new(dto.SysRequestLogGetPageReq)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req.OperTimeOrder = "desc"
	list := make([]dto.SysRequestLogGetPageResp, 0)
	var count int64

	err = s.GetPage(req, &list, &count)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 操作日志通过id获取
// @Summary 操作日志通过id获取
// @Description 获取JSON
// @Tags 操作日志
// @Param id path string false "id"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-request-log/{id} [get]
// @Security Bearer
func (e SysRequestLog) Get(c *gin.Context) {
	s := new(service.SysRequestLog)
	req := dto.SysRequestLogGetReq{}
	resp := dto.SysRequestLogGetResp{}
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
	var object models.SysRequestLog
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}

	resp.Generate(&object)
	e.OK(object)
}

// Delete 操作日志删除
// DeleteSysMenu 操作日志删除
// @Summary 删除操作日志
// @Description 删除数据
// @Tags 操作日志
// @Param data body dto.SysRequestLogDeleteReq true "body"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-request-log [delete]
// @Security Bearer
func (e SysRequestLog) Delete(c *gin.Context) {
	s := new(service.SysRequestLog)
	req := dto.SysRequestLogDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, fmt.Sprintf("删除失败！错误详情：%s", err.Error()), "")
		return
	}
	e.OK(req.GetId())
}
