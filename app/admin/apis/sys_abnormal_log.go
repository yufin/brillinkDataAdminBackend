package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	_ "go-admin/common/response/antd"
)

type SysAbnormalLog struct {
	apis.Api
}

// GetPage 获取异常日志列表
// @Summary      获取异常日志列表
// @Description  获取异常日志列表
// @Tags         异常日志
// @Param        abId        query     int64                                                         false  "编码"
// @Param        method      query     string                                                        false  "请求方式"
// @Param        url         query     string                                                        false  "请求地址"
// @Param        ip          query     string                                                        false  "ip"
// @Param        abInfo      query     string                                                        false  "异常信息"
// @Param        abSource    query     string                                                        false  "异常来源"
// @Param        abFunc      query     string                                                        false  "异常方法"
// @Param        userId      query     int64                                                         false  "用户id"
// @Param        userName    query     string                                                        false  "操作人"
// @Param        headers     query     string                                                        false  "请求头"
// @Param        body        query     string                                                        false  "请求数据"
// @Param        resp        query     string                                                        false  "回调数据"
// @Param        stackTrace  query     string                                                        false  "堆栈追踪"
// @Param        pageSize    query     int                                                           false  "页条数"
// @Param        pageIndex   query     int                                                           false  "页码"
// @Success      200         {object}  antd.Response{data=antd.Pages{list=[]models.SysAbnormalLog}}  "{"code": 200, "data": [...]}"
// @Router       /api/v1/sys-abnormal-log [get]
// @Security     Bearer
func (e SysAbnormalLog) GetPage(c *gin.Context) {
	req := dto.SysAbnormalLogGetPageReq{}
	s := service.SysAbnormalLog{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err.Error(), "")
		return
	}

	p := actions.GetPermission(c)
	list := make([]dto.SysAbnormalLogGetPageResp, 0)
	var count int64
	req.SysAbnormalLogPageOrder.CreatedAt = "desc"
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, fmt.Sprintf("获取异常日志 失败，失败信息 %s", err.Error()), "")
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取异常日志
// @Summary      获取异常日志
// @Description  获取异常日志
// @Tags         异常日志
// @Param        id   path      string                                     false  "id"
// @Success      200  {object}  antd.Response{data=models.SysAbnormalLog}  "{"code": 200, "data": [...]}"
// @Router       /api/v1/sys-abnormal-log/{id} [get]
// @Security     Bearer
func (e SysAbnormalLog) Get(c *gin.Context) {
	req := dto.SysAbnormalLogGetReq{}
	resp := dto.SysAbnormalLogGetResp{}
	s := service.SysAbnormalLog{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err.Error(), "")
		return
	}
	var object models.SysAbnormalLog

	p := actions.GetPermission(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, fmt.Sprintf("获取异常日志失败，失败信息 %s", err.Error()), "")
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Delete 删除异常日志
// @Summary      删除异常日志
// @Description  删除异常日志
// @Tags         异常日志
// @Param        ids  body      []int          false  "ids"
// @Success      200  {object}  antd.Response  "{"code": 200, "message": "删除成功"}"
// @Router       /api/v1/sys-abnormal-log [delete]
// @Security     Bearer
func (e SysAbnormalLog) Delete(c *gin.Context) {
	s := service.SysAbnormalLog{}
	req := dto.SysAbnormalLogDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err.Error(), "")
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermission(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, fmt.Sprintf("删除异常日志失败，\r\n失败信息 %s", err.Error()), "")
		return
	}
	e.OK(req.GetId())
}
