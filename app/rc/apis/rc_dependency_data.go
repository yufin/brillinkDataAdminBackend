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

type RcDependencyData struct {
	apis.Api
}

// GetPage 获取依赖数据列表
// @Summary 获取依赖数据列表
// @Description 获取依赖数据列表
// @Tags 依赖数据
// @Param id query int64 false "主键"
// @Param contentId query int64 false "rc_origin_content.id"
// @Param uscId query string false "统一信用社会代码"
// @Param lhQylx query int false "企业类型"
// @Param lhCylwz query int false "产业链位置"
// @Param lhGdct query int false "股东穿透"
// @Param lhYhsx query int false "银行授信"
// @Param lhSfsx query int false "三方授信"
// @Param update_by query string false "更新人"
// @Param create_by query string false "创建人"
// @Param deleted_at query time.Time false "删除时间"
// @Param updated_at query time.Time false "更新时间"
// @Param created_at query time.Time false "创建时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcDependencyData}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-dependency-data [get]
// @Security Bearer
func (e RcDependencyData) GetPage(c *gin.Context) {
	req := dto.RcDependencyDataGetPageReq{}
	s := service.RcDependencyData{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcDependencyDataFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcDependencyData, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcDependencyDataFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取依赖数据
// @Summary 获取依赖数据
// @Description 获取依赖数据
// @Tags 依赖数据
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcDependencyData} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-dependency-data/{id} [get]
// @Security Bearer
func (e RcDependencyData) Get(c *gin.Context) {
	req := dto.RcDependencyDataGetReq{}
	resp := dto.RcDependencyDataGetResp{}
	s := service.RcDependencyData{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcDependencyDataFail", err))
		return
	}
	var object models.RcDependencyData

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcDependencyDataFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建依赖数据
// @Summary 创建依赖数据
// @Description 创建依赖数据
// @Tags 依赖数据
// @Accept application/json
// @Product application/json
// @Param data body dto.RcDependencyDataInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-dependency-data [post]
// @Security Bearer
func (e RcDependencyData) Insert(c *gin.Context) {
	req := dto.RcDependencyDataInsertReq{}
	s := service.RcDependencyData{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcDependencyDataFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcDependencyDataFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改依赖数据
// @Summary 修改依赖数据
// @Description 修改依赖数据
// @Tags 依赖数据
// @Accept application/json
// @Product application/json
// @Param data body dto.RcDependencyDataUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-dependency-data/{id} [put]
// @Security Bearer
func (e RcDependencyData) Update(c *gin.Context) {
	req := dto.RcDependencyDataUpdateReq{}
	s := service.RcDependencyData{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcDependencyDataFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcDependencyDataFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除依赖数据
// @Summary 删除依赖数据
// @Description 删除依赖数据
// @Tags 依赖数据
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rc-dependency-data [delete]
// @Security Bearer
func (e RcDependencyData) Delete(c *gin.Context) {
	s := service.RcDependencyData{}
	req := dto.RcDependencyDataDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcDependencyDataFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcDependencyDataFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出依赖数据列表
// @Summary 导出依赖数据列表
// @Description 导出依赖数据列表
// @Tags 依赖数据
// @Param id query int64 false "主键"
// @Param contentId query int64 false "rc_origin_content.id"
// @Param uscId query string false "统一信用社会代码"
// @Param lhQylx query int false "企业类型"
// @Param lhCylwz query int false "产业链位置"
// @Param lhGdct query int false "股东穿透"
// @Param lhYhsx query int false "银行授信"
// @Param lhSfsx query int false "三方授信"
// @Param update_by query string false "更新人"
// @Param create_by query string false "创建人"
// @Param deleted_at query time.Time false "删除时间"
// @Param updated_at query time.Time false "更新时间"
// @Param created_at query time.Time false "创建时间"
// @Router /api/v1/rc-dependency-data/export [get]
// @Security Bearer
func (e RcDependencyData) Export(c *gin.Context) {
	req := dto.RcDependencyDataGetPageReq{}
	s := service.RcDependencyData{}
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
