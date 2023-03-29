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

type SysOperateLog struct {
	apis.Api
}

// GetPage 获取操作日志列表
// @Summary 获取操作日志列表
// @Description 获取操作日志列表
// @Tags 操作日志
// @Param logId query int64 false "编码"
// @Param type query string false "操作类型"
// @Param description query string false "操作说明"
// @Param userName query string false "用户"
// @Param userId query int64 false "用户id"
// @Param updateBefore query string false "更新前"
// @Param updateAfter query string false "更新后"
// @Param createdAt query time.Time false "创建时间"
// @Param updatedAt query time.Time false "最后修改时间"
// @Param deletedAt query time.Time false "删除时间"
// @Param createBy query int64 false "创建人"
// @Param updateBy query int64 false "最后更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.SysOperateLog}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-operate-log [get]
// @Security Bearer
func (e SysOperateLog) GetPage(c *gin.Context) {
	req := dto.SysOperateLogGetPageReq{}
	s := service.SysOperateLog{}
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
	req.SysOperateLogPageOrder.CreatedAt = "desc"
	p := actions.GetPermission(c)
	list := make([]dto.SysOperateLogGetPageResp, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, fmt.Sprintf("获取操作日志 失败，失败信息 %s", err.Error()), "")
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取操作日志
// @Summary 获取操作日志
// @Description 获取操作日志
// @Tags 操作日志
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.SysOperateLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-operate-log/{id} [get]
// @Security Bearer
func (e SysOperateLog) Get(c *gin.Context) {
	req := dto.SysOperateLogGetReq{}
	resp := dto.SysOperateLogGetResp{}
	s := service.SysOperateLog{}
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
	var object models.SysOperateLog
	p := actions.GetPermission(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, fmt.Sprintf("获取操作日志失败，失败信息 %s", err.Error()), "")
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Delete 删除操作日志
// @Summary 删除操作日志
// @Description 删除操作日志
// @Tags 操作日志
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-operate-log [delete]
// @Security Bearer
func (e SysOperateLog) Delete(c *gin.Context) {
	s := service.SysOperateLog{}
	req := dto.SysOperateLogDeleteReq{}
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
		e.Error(500, fmt.Sprintf("删除操作日志失败，\r\n失败信息 %s", err.Error()), "")
		return
	}
	e.OK(req.GetId())
}
