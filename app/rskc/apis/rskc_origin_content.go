package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/rskc/task"

	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type RskcOriginContent struct {
	apis.Api
}

// GetPage 获取微众json存储列表
// @Summary 获取微众json存储列表
// @Description 获取微众json存储列表
// @Tags 微众json存储
// @Param id query int64 false "主键"
// @Param contentId query string false "uuid4"
// @Param uscId query string false "统一社会信用代码"
// @Param yearMonth query string false "数据更新年月"
// @Param content query string false "原始JSON STRING数据"
// @Param statusCode query int false "状态码"
// @Param updateBy query int64 false "更新人"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param createBy query int64 false "创建人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RskcOriginContent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rskc-origin-content [get]
// @Security Bearer
func (e RskcOriginContent) GetPage(c *gin.Context) {
	req := dto.RskcOriginContentGetPageReq{}
	s := service.RskcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRskcOriginContentFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RskcOriginContent, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRskcOriginContentFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取微众json存储
// @Summary 获取微众json存储
// @Description 获取微众json存储
// @Tags 微众json存储
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RskcOriginContent} "{"code": 200, "data": [...]}"
// @Router /api/v1/rskc-origin-content/{id} [get]
// @Security Bearer
func (e RskcOriginContent) Get(c *gin.Context) {
	req := dto.RskcOriginContentGetReq{}
	resp := dto.RskcOriginContentGetResp{}
	s := service.RskcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRskcOriginContentFail", err))
		return
	}
	var object models.RskcOriginContent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRskcOriginContentFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建微众json存储
// @Summary 创建微众json存储
// @Description 创建微众json存储
// @Tags 微众json存储
// @Accept application/json
// @Product application/json
// @Param data body dto.RskcOriginContentInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rskc-origin-content [post]
// @Security Bearer
func (e RskcOriginContent) Insert(c *gin.Context) {
	req := dto.RskcOriginContentInsertReq{}
	s := service.RskcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRskcOriginContentFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRskcOriginContentFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改微众json存储
// @Summary 修改微众json存储
// @Description 修改微众json存储
// @Tags 微众json存储
// @Accept application/json
// @Product application/json
// @Param data body dto.RskcOriginContentUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rskc-origin-content/{id} [put]
// @Security Bearer
func (e RskcOriginContent) Update(c *gin.Context) {
	req := dto.RskcOriginContentUpdateReq{}
	s := service.RskcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRskcOriginContentFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRskcOriginContentFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除微众json存储
// @Summary 删除微众json存储
// @Description 删除微众json存储
// @Tags 微众json存储
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rskc-origin-content [delete]
// @Security Bearer
func (e RskcOriginContent) Delete(c *gin.Context) {
	s := service.RskcOriginContent{}
	req := dto.RskcOriginContentDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRskcOriginContentFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRskcOriginContentFail", err))
		return
	}
	e.OK(req.GetId())
}

func (e RskcOriginContent) TaskSyncOriginContent(c *gin.Context) {
	s := service.RskcOriginContent{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "TaskSyncOriginContentFail", err))
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = task.SyncOriginJsonContent(&s, p)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "TaskSyncOriginContentFail", err))
		return
	}
	e.OK(nil)
}
