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

type RcSellingSta struct {
	apis.Api
}

// GetPage 获取parse from content sellingSta列表
// @Summary 获取parse from content sellingSta列表
// @Description 获取parse from content sellingSta列表
// @Tags parse from content sellingSta
// @Param id query int64 false "主键"
// @Param contentId query int64 false "foreign key from RcOriginContent"
// @Param cgje query string false "cgje"
// @Param jezb query string false "jezb"
// @Param ssspdl query string false "ssspdl"
// @Param ssspxl query string false "ssspxl"
// @Param ssspzl query string false "ssspzl"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcSellingSta}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-selling-sta [get]
// @Security Bearer
func (e RcSellingSta) GetPage(c *gin.Context) {
	req := dto.RcSellingStaGetPageReq{}
	s := service.RcSellingSta{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcSellingStaFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcSellingSta, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcSellingStaFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取parse from content sellingSta
// @Summary 获取parse from content sellingSta
// @Description 获取parse from content sellingSta
// @Tags parse from content sellingSta
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcSellingSta} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-selling-sta/{id} [get]
// @Security Bearer
func (e RcSellingSta) Get(c *gin.Context) {
	req := dto.RcSellingStaGetReq{}
	resp := dto.RcSellingStaGetResp{}
	s := service.RcSellingSta{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcSellingStaFail", err))
		return
	}
	var object models.RcSellingSta

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcSellingStaFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建parse from content sellingSta
// @Summary 创建parse from content sellingSta
// @Description 创建parse from content sellingSta
// @Tags parse from content sellingSta
// @Accept application/json
// @Product application/json
// @Param data body dto.RcSellingStaInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-selling-sta [post]
// @Security Bearer
func (e RcSellingSta) Insert(c *gin.Context) {
	req := dto.RcSellingStaInsertReq{}
	s := service.RcSellingSta{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcSellingStaFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcSellingStaFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改parse from content sellingSta
// @Summary 修改parse from content sellingSta
// @Description 修改parse from content sellingSta
// @Tags parse from content sellingSta
// @Accept application/json
// @Product application/json
// @Param data body dto.RcSellingStaUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-selling-sta/{id} [put]
// @Security Bearer
func (e RcSellingSta) Update(c *gin.Context) {
	req := dto.RcSellingStaUpdateReq{}
	s := service.RcSellingSta{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcSellingStaFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcSellingStaFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除parse from content sellingSta
// @Summary 删除parse from content sellingSta
// @Description 删除parse from content sellingSta
// @Tags parse from content sellingSta
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rc-selling-sta [delete]
// @Security Bearer
func (e RcSellingSta) Delete(c *gin.Context) {
	s := service.RcSellingSta{}
	req := dto.RcSellingStaDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcSellingStaFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcSellingStaFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出parse from content sellingSta列表
// @Summary 导出parse from content sellingSta列表
// @Description 导出parse from content sellingSta列表
// @Tags parse from content sellingSta
// @Param id query int64 false "主键"
// @Param contentId query int64 false "foreign key from RcOriginContent"
// @Param cgje query string false "cgje"
// @Param jezb query string false "jezb"
// @Param ssspdl query string false "ssspdl"
// @Param ssspxl query string false "ssspxl"
// @Param ssspzl query string false "ssspzl"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Router /api/v1/rc-selling-sta/export [get]
// @Security Bearer
func (e RcSellingSta) Export(c *gin.Context) {
	req := dto.RcSellingStaGetPageReq{}
	s := service.RcSellingSta{}
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
