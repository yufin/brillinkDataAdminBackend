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

type RcDecisionParam struct {
	apis.Api
}

// GetPage 获取列表
// @Summary 获取列表
// @Description 获取列表
// @Tags
// @Param id query int64 false "id"
// @Param swSdsnbCyrs query int64 false ""
// @Param gsGdct query int64 false ""
// @Param gsGdwdx query int64 false ""
// @Param gsFrwdx query int64 false ""
// @Param lhCylwz query int64 false ""
// @Param lhMdPpjzl query int64 false ""
// @Param mdQybq query int64 false ""
// @Param swCwbbYyzjzzts query int64 false ""
// @Param sfFhSsqkQy query string false "纳税信用评级"
// @Param swJcxxNsrxypj query string false ""
// @Param zxYhsxqk query int64 false ""
// @Param zxDsfsxqk query int64 false ""
// @Param lhQylx query int64 false ""
// @Param orderNo query string false "订单号"
// @Param nsrsbh query string false ""
// @Param applyTime query string false ""
// @Param swSbNszeZzsqysds12m query int64 false ""
// @Param swSbNszezzlZzsqysds12mA query int64 false ""
// @Param swSdsnbGzxjzzjezzl query int64 false ""
// @Param swSbzsSflhypld12m query int64 false ""
// @Param swSdsnbYjfy query int64 false ""
// @Param fpJxLxfy12m query int64 false ""
// @Param swCwbbSszb query int64 false ""
// @Param fpJySychjeZb12mLh query int64 false ""
// @Param fpJxZyjyjezb12mLh query int64 false ""
// @Param fpXxXychjeZb12mLh query int64 false ""
// @Param fpXxZyjyjezb12mLh query int64 false ""
// @Param swSbQbxse12m query int64 false ""
// @Param swSbQbxsezzl12m query int64 false ""
// @Param swSbLsxs12m query int64 false ""
// @Param swCwbbChzztsCb query int64 false ""
// @Param swCwbbZcfzl query int64 false ""
// @Param swCwbbMlrzzlv query int64 false ""
// @Param swCwbbJlrzzlv query int64 false ""
// @Param swCwbbJzcszlv query int64 false ""
// @Param swJcxxClnx query int64 false ""
// @Param statusCode query int64 false ""
// @Param createdAt query int64 false ""
// @Param updatedAt query int64 false ""
// @Param deletedAt query int64 false ""
// @Param updatedBy query int64 false ""
// @Param createdBy query int64 false ""
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcDecisionParam}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-decision-param [get]
// @Security Bearer
func (e RcDecisionParam) GetPage(c *gin.Context) {
	req := dto.RcDecisionParamGetPageReq{}
	s := service.RcDecisionParam{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcDecisionParamFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcDecisionParam, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcDecisionParamFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取
// @Summary 获取
// @Description 获取
// @Tags
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcDecisionParam} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-decision-param/{id} [get]
// @Security Bearer
func (e RcDecisionParam) Get(c *gin.Context) {
	req := dto.RcDecisionParamGetReq{}
	resp := dto.RcDecisionParamGetResp{}
	s := service.RcDecisionParam{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcDecisionParamFail", err))
		return
	}
	var object models.RcDecisionParam

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcDecisionParamFail", err))
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
// @Param data body dto.RcDecisionParamInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-decision-param [post]
// @Security Bearer
func (e RcDecisionParam) Insert(c *gin.Context) {
	req := dto.RcDecisionParamInsertReq{}
	s := service.RcDecisionParam{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcDecisionParamFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcDecisionParamFail", err))
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
// @Param data body dto.RcDecisionParamUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-decision-param/{id} [put]
// @Security Bearer
func (e RcDecisionParam) Update(c *gin.Context) {
	req := dto.RcDecisionParamUpdateReq{}
	s := service.RcDecisionParam{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcDecisionParamFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcDecisionParamFail", err))
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
// @Router /api/v1/rc-decision-param [delete]
// @Security Bearer
func (e RcDecisionParam) Delete(c *gin.Context) {
	s := service.RcDecisionParam{}
	req := dto.RcDecisionParamDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcDecisionParamFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcDecisionParamFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出列表
// @Summary 导出列表
// @Description 导出列表
// @Tags
// @Param id query int64 false "id"
// @Param swSdsnbCyrs query int64 false ""
// @Param gsGdct query int64 false ""
// @Param gsGdwdx query int64 false ""
// @Param gsFrwdx query int64 false ""
// @Param lhCylwz query int64 false ""
// @Param lhMdPpjzl query int64 false ""
// @Param mdQybq query int64 false ""
// @Param swCwbbYyzjzzts query int64 false ""
// @Param sfFhSsqkQy query string false "纳税信用评级"
// @Param swJcxxNsrxypj query string false ""
// @Param zxYhsxqk query int64 false ""
// @Param zxDsfsxqk query int64 false ""
// @Param lhQylx query int64 false ""
// @Param orderNo query string false "订单号"
// @Param nsrsbh query string false ""
// @Param applyTime query string false ""
// @Param swSbNszeZzsqysds12m query int64 false ""
// @Param swSbNszezzlZzsqysds12mA query int64 false ""
// @Param swSdsnbGzxjzzjezzl query int64 false ""
// @Param swSbzsSflhypld12m query int64 false ""
// @Param swSdsnbYjfy query int64 false ""
// @Param fpJxLxfy12m query int64 false ""
// @Param swCwbbSszb query int64 false ""
// @Param fpJySychjeZb12mLh query int64 false ""
// @Param fpJxZyjyjezb12mLh query int64 false ""
// @Param fpXxXychjeZb12mLh query int64 false ""
// @Param fpXxZyjyjezb12mLh query int64 false ""
// @Param swSbQbxse12m query int64 false ""
// @Param swSbQbxsezzl12m query int64 false ""
// @Param swSbLsxs12m query int64 false ""
// @Param swCwbbChzztsCb query int64 false ""
// @Param swCwbbZcfzl query int64 false ""
// @Param swCwbbMlrzzlv query int64 false ""
// @Param swCwbbJlrzzlv query int64 false ""
// @Param swCwbbJzcszlv query int64 false ""
// @Param swJcxxClnx query int64 false ""
// @Param statusCode query int64 false ""
// @Param createdAt query int64 false ""
// @Param updatedAt query int64 false ""
// @Param deletedAt query int64 false ""
// @Param updatedBy query int64 false ""
// @Param createdBy query int64 false ""
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Router /api/v1/rc-decision-param/export [get]
// @Security Bearer
func (e RcDecisionParam) Export(c *gin.Context) {
	req := dto.RcDecisionParamGetPageReq{}
	s := service.RcDecisionParam{}
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
