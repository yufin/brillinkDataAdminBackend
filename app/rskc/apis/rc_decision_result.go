package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/google/uuid"
	"go-admin/common"
	"os"

	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type RcDecisionResult struct {
	apis.Api
}

// GetPage 获取列表
// @Summary 获取列表
// @Description 获取列表
// @Tags
// @Param id query int64 false ""
// @Param paramId query int64 false ""
// @Param taskId query string false "任务id"
// @Param finalResult query string false "决策建议结果(REFUSE:拒绝，PASS:通过)"
// @Param aphScore query float64 false "APH分数"
// @Param fxSwJxccClnx query string false "经营年限"
// @Param lhQylx query int false "1:生产型，2:贸易型"
// @Param msg query string false "返回结果描述"
// @Param createdAt query string false ""
// @Param updatedAt query string false ""
// @Param deletedAt query string false ""
// @Param createBy query int64 false ""
// @Param updateBy query int64 false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcDecisionResult}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-decision-result [get]
// @Security Bearer
func (e RcDecisionResult) GetPage(c *gin.Context) {
	req := dto.RcDecisionResultGetPageReq{}
	s := service.RcDecisionResult{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcDecisionResultFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcDecisionResult, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcDecisionResultFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取
// @Summary 获取
// @Description 获取
// @Tags
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcDecisionResult} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-decision-result/{id} [get]
// @Security Bearer
func (e RcDecisionResult) Get(c *gin.Context) {
	req := dto.RcDecisionResultGetReq{}
	resp := dto.RcDecisionResultGetResp{}
	s := service.RcDecisionResult{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcDecisionResultFail", err))
		return
	}
	var object models.RcDecisionResult

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcDecisionResultFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建
// @Summary 创建
// @Description 创建
// @Tags
// @Accept application/json
// @Product application/json
// @Param data body dto.RcDecisionResultInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-decision-result [post]
// @Security Bearer
func (e RcDecisionResult) Insert(c *gin.Context) {
	req := dto.RcDecisionResultInsertReq{}
	s := service.RcDecisionResult{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcDecisionResultFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcDecisionResultFail", err))
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
// @Param data body dto.RcDecisionResultUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-decision-result/{id} [put]
// @Security Bearer
func (e RcDecisionResult) Update(c *gin.Context) {
	req := dto.RcDecisionResultUpdateReq{}
	s := service.RcDecisionResult{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcDecisionResultFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcDecisionResultFail", err))
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
// @Router /api/v1/rc-decision-result [delete]
// @Security Bearer
func (e RcDecisionResult) Delete(c *gin.Context) {
	s := service.RcDecisionResult{}
	req := dto.RcDecisionResultDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcDecisionResultFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcDecisionResultFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出列表
// @Summary 导出列表
// @Description 导出列表
// @Tags
// @Param id query int64 false ""
// @Param paramId query int64 false ""
// @Param taskId query string false "任务id"
// @Param finalResult query string false "决策建议结果(REFUSE:拒绝，PASS:通过)"
// @Param aphScore query float64 false "APH分数"
// @Param fxSwJxccClnx query string false "经营年限"
// @Param lhQylx query int false "1:生产型，2:贸易型"
// @Param msg query string false "返回结果描述"
// @Param createdAt query string false ""
// @Param updatedAt query string false ""
// @Param deletedAt query string false ""
// @Param createBy query int64 false ""
// @Param updateBy query int64 false ""
// @Router /api/v1/rc-decision-result/export [get]
// @Security Bearer
func (e RcDecisionResult) Export(c *gin.Context) {
	req := dto.RcDecisionResultGetPageReq{}
	s := service.RcDecisionResult{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(500, "初始化服务失败", err))
		return
	}

	list, err := s.Export(&req)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(500, "查询数据失败", err))
		return
	}

	f := common.WriteXlsx("Sheet1", list)
	filename := uuid.New().String() + ".xlsx"
	path := "temp/excel/"
	pathname := path + filename
	if !pkg.PathExist(path) {
		err := pkg.PathCreate("temp/excel/")
		if err != nil {
			panic(exception.WithMsg(500, "创建路径失败", err))
			return
		}
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(pathname); err != nil {
		panic(exception.WithMsg(500, "保存文件失败", err))
		return
	}
	e.File(pathname)
	_ = os.Remove(pathname)
}
