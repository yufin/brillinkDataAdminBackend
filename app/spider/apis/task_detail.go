package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/google/uuid"
	"go-admin/common"
	"os"

	"go-admin/app/spider/models"
	"go-admin/app/spider/service"
	"go-admin/app/spider/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type TaskDetail struct {
	apis.Api
}

// GetPage 获取任务列表列表
// @Summary 获取任务列表列表
// @Description 获取任务列表列表
// @Tags 任务列表
// @Param id query int64 false "主键"
// @Param uscId query string false "企业统一信用代码"
// @Param statusCode query int false "状态码(1.未完成, 2.完成)"
// @Param topic query string false "主题"
// @Param priority query int false "优先级"
// @Param comment query string false "备注"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.TaskDetail}} "{"code": 200, "data": [...]}"
// @Router /api/v1/task-detail [get]
// @Security Bearer
func (e TaskDetail) GetPage(c *gin.Context) {
	req := dto.TaskDetailGetPageReq{}
	s := service.TaskDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTaskDetailFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.TaskDetail, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTaskDetailFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取任务列表
// @Summary 获取任务列表
// @Description 获取任务列表
// @Tags 任务列表
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.TaskDetail} "{"code": 200, "data": [...]}"
// @Router /api/v1/task-detail/{id} [get]
// @Security Bearer
func (e TaskDetail) Get(c *gin.Context) {
	req := dto.TaskDetailGetReq{}
	resp := dto.TaskDetailGetResp{}
	s := service.TaskDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetTaskDetailFail", err))
		return
	}
	var object models.TaskDetail

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetTaskDetailFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建任务列表
// @Summary 创建任务列表
// @Description 创建任务列表
// @Tags 任务列表
// @Accept application/json
// @Product application/json
// @Param data body dto.TaskDetailInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/task-detail [post]
// @Security Bearer
func (e TaskDetail) Insert(c *gin.Context) {
	req := dto.TaskDetailInsertReq{}
	s := service.TaskDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertTaskDetailFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertTaskDetailFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改任务列表
// @Summary 修改任务列表
// @Description 修改任务列表
// @Tags 任务列表
// @Accept application/json
// @Product application/json
// @Param data body dto.TaskDetailUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/task-detail/{id} [put]
// @Security Bearer
func (e TaskDetail) Update(c *gin.Context) {
	req := dto.TaskDetailUpdateReq{}
	s := service.TaskDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateTaskDetailFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateTaskDetailFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除任务列表
// @Summary 删除任务列表
// @Description 删除任务列表
// @Tags 任务列表
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/task-detail [delete]
// @Security Bearer
func (e TaskDetail) Delete(c *gin.Context) {
	s := service.TaskDetail{}
	req := dto.TaskDetailDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteTaskDetailFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteTaskDetailFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出任务列表列表
// @Summary 导出任务列表列表
// @Description 导出任务列表列表
// @Tags 任务列表
// @Param id query int64 false "主键"
// @Param uscId query string false "企业统一信用代码"
// @Param statusCode query int false "状态码(1.未完成, 2.完成)"
// @Param topic query string false "主题"
// @Param priority query int false "优先级"
// @Param comment query string false "备注"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Router /api/v1/task-detail/export [get]
// @Security Bearer
func (e TaskDetail) Export(c *gin.Context) {
	req := dto.TaskDetailGetPageReq{}
	s := service.TaskDetail{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTaskDetailFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.TaskDetail, 0)

	err = s.GetList(c, &req, p, &list)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageTaskDetailFail", err))
		return
	}

	f := common.WriteXlsx("Sheet1", list)
	filename := uuid.New().String() + ".xlsx"
	path := "temp/excel/"
	pathname := path + filename
	if !pkg.PathExist(path) {
		pkg.PathCreate(path)
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(pathname); err != nil {
		println(err.Error())
	}
	e.File(pathname)
	_ = os.Remove(pathname)
}
