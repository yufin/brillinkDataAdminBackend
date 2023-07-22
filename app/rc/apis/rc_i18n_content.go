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

type RcI18nContent struct {
	apis.Api
}

// GetPage 获取i18n content列表
// @Summary 获取i18n content列表
// @Description 获取i18n content列表
// @Tags i18n content
// @Param id query int64 false "主键"
// @Param processedId query int64 false "rc_processed_content.id"
// @Param lang query string false "语言类型(en,...)"
// @Param content query string false "报文json string"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcI18nContent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-ixxn-content [get]
// @Security Bearer
func (e RcI18nContent) GetPage(c *gin.Context) {
	req := dto.RcI18nContentGetPageReq{}
	s := service.RcI18nContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcI18nContentFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcI18nContent, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcI18nContentFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取i18n content
// @Summary 获取i18n content
// @Description 获取i18n content
// @Tags i18n content
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcI18nContent} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-ixxn-content/{id} [get]
// @Security Bearer
func (e RcI18nContent) Get(c *gin.Context) {
	req := dto.RcI18nContentGetReq{}
	resp := dto.RcI18nContentGetResp{}
	s := service.RcI18nContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcI18nContentFail", err))
		return
	}
	var object models.RcI18nContent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcI18nContentFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建i18n content
// @Summary 创建i18n content
// @Description 创建i18n content
// @Tags i18n content
// @Accept application/json
// @Product application/json
// @Param data body dto.RcI18nContentInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-ixxn-content [post]
// @Security Bearer
func (e RcI18nContent) Insert(c *gin.Context) {
	req := dto.RcI18nContentInsertReq{}
	s := service.RcI18nContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcI18nContentFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcI18nContentFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改i18n content
// @Summary 修改i18n content
// @Description 修改i18n content
// @Tags i18n content
// @Accept application/json
// @Product application/json
// @Param data body dto.RcI18nContentUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-ixxn-content/{id} [put]
// @Security Bearer
func (e RcI18nContent) Update(c *gin.Context) {
	req := dto.RcI18nContentUpdateReq{}
	s := service.RcI18nContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcI18nContentFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcI18nContentFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除i18n content
// @Summary 删除i18n content
// @Description 删除i18n content
// @Tags i18n content
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rc-ixxn-content [delete]
// @Security Bearer
func (e RcI18nContent) Delete(c *gin.Context) {
	s := service.RcI18nContent{}
	req := dto.RcI18nContentDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcI18nContentFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcI18nContentFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出i18n content列表
// @Summary 导出i18n content列表
// @Description 导出i18n content列表
// @Tags i18n content
// @Param id query int64 false "主键"
// @Param processedId query int64 false "rc_processed_content.id"
// @Param lang query string false "语言类型(en,...)"
// @Param content query string false "报文json string"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Router /api/v1/rc-ixxn-content/export [get]
// @Security Bearer
func (e RcI18nContent) Export(c *gin.Context) {
	req := dto.RcI18nContentGetPageReq{}
	s := service.RcI18nContent{}
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
