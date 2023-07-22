package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/google/uuid"
	"go-admin/common"
	"os"

	"go-admin/app/rc/models"
	"go-admin/app/rc/service"
	"go-admin/app/rc/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type RcDecisionResult struct {
	apis.Api
}

// GetPage 获取wezoom decision engine returend result列表
// @Summary 获取wezoom decision engine returend result列表
// @Description 获取wezoom decision engine returend result列表
// @Tags wezoom decision engine returend result
// @Param id query int64 false "主键"
// @Param depId query int64 false "rc_decision_param"
// @Param taskId query string false "task_id"
// @Param finalResult query string false "决策建议结果(REFUSE:拒绝，PASS:通过)"
// @Param aphScore query string false "APH分数"
// @Param fxSwJxccClnx query string false "经营年限"
// @Param lhQylx query int false "1:生产型，2:贸易型"
// @Param msg query string false "返回结果描述"
// @Param update_by query string false "更新人"
// @Param create_by query string false "创建人"
// @Param deleted_at query time.Time false "删除时间"
// @Param updated_at query time.Time false "更新时间"
// @Param created_at query time.Time false "创建时间"
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

// Get 获取wezoom decision engine returend result
// @Summary 获取wezoom decision engine returend result
// @Description 获取wezoom decision engine returend result
// @Tags wezoom decision engine returend result
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

// Insert 创建wezoom decision engine returend result
// @Summary 创建wezoom decision engine returend result
// @Description 创建wezoom decision engine returend result
// @Tags wezoom decision engine returend result
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

// Update 修改wezoom decision engine returend result
// @Summary 修改wezoom decision engine returend result
// @Description 修改wezoom decision engine returend result
// @Tags wezoom decision engine returend result
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

// Delete 删除wezoom decision engine returend result
// @Summary 删除wezoom decision engine returend result
// @Description 删除wezoom decision engine returend result
// @Tags wezoom decision engine returend result
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

// Export 导出wezoom decision engine returend result列表
// @Summary 导出wezoom decision engine returend result列表
// @Description 导出wezoom decision engine returend result列表
// @Tags wezoom decision engine returend result
// @Param id query int64 false "主键"
// @Param depId query int64 false "rc_decision_param"
// @Param taskId query string false "task_id"
// @Param finalResult query string false "决策建议结果(REFUSE:拒绝，PASS:通过)"
// @Param aphScore query string false "APH分数"
// @Param fxSwJxccClnx query string false "经营年限"
// @Param lhQylx query int false "1:生产型，2:贸易型"
// @Param msg query string false "返回结果描述"
// @Param update_by query string false "更新人"
// @Param create_by query string false "创建人"
// @Param deleted_at query time.Time false "删除时间"
// @Param updated_at query time.Time false "更新时间"
// @Param created_at query time.Time false "创建时间"
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
