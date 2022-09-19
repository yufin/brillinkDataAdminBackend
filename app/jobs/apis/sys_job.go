package apis

import (
	"go-admin/app/jobs/models"
	"go-admin/app/jobs/service"
	"go-admin/app/jobs/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
)

type SysJob struct {
	apis.Api
}

// GetPage 获取列表
// @Summary 获取列表
// @Description 获取列表
// @Tags
// @Param jobName query string false "jobName"
// @Param jobGroup query string false "jobGroup"
// @Param jobType query string false "jobType"
// @Param cronExpression query string false "cronExpression"
// @Param invokeTarget query string false "invokeTarget"
// @Param args query string false "args"
// @Param misfirePolicy query int64 false "misfirePolicy"
// @Param concurrent query string false "concurrent"
// @Param status query string false "status"
// @Param entryId query int64 false "entryId"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.SysJob}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-job [get]
// @Security Bearer
func (e SysJob) GetPage(c *gin.Context) {
	req := dto.SysJobGetPageReq{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysJob, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取
// @Summary 获取
// @Description 获取
// @Tags
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.SysJob} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-job/{id} [get]
// @Security Bearer
func (e SysJob) Get(c *gin.Context) {
	req := dto.SysJobGetReq{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	var object models.SysJob

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(err)
		return
	}

	e.OK(object)
}

// Insert 创建
// @Summary 创建
// @Description 创建
// @Tags
// @Accept application/json
// @Product application/json
// @Param data body dto.SysJobInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-job [post]
// @Security Bearer
func (e SysJob) Insert(c *gin.Context) {
	req := dto.SysJobInsertReq{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		panic(err)
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
// @Param data body dto.SysJobUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-job/{id} [put]
// @Security Bearer
func (e SysJob) Update(c *gin.Context) {
	req := dto.SysJobUpdateReq{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(err)
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
// @Router /api/v1/sys-job/{id} [delete]
// @Security Bearer
func (e SysJob) Delete(c *gin.Context) {
	s := service.SysJob{}
	req := dto.SysJobDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(err)
		return
	}
	e.OK(req.GetId())
}

// RemoveJobForService 调用service实现
func (e SysJob) RemoveJobForService(c *gin.Context) {
	req := dto.SysJobDeleteReq{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.RemoveJob(&req)
	if err != nil {
		e.Logger.Errorf("RemoveJob error, %s", err.Error())
		return
	}
	e.OK(nil)
}

// StartJobForService 启动job service实现
func (e SysJob) StartJobForService(c *gin.Context) {
	req := dto.SysJobGetReq{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.StartJob(&req)
	if err != nil {
		e.Logger.Errorf("GetCrontabKey error, %s", err.Error())
		return
	}
	e.OK(nil)
}
