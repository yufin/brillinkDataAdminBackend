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

type RskcProcessedContent struct {
	apis.Api
}

// GetPage 获取报文处理列表
// @Summary 获取报文处理列表
// @Description 获取报文处理列表
// @Tags 报文处理
// @Param id query int64 false "主键"
// @Param contentId query int64 false "外键(rskc_origin_content.id)"
// @Param content query string false "数据(json字符串格式)"
// @Param statusCode query int false "状态码"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RskcProcessedContent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rskc-processed-content [get]
// @Security Bearer
func (e RskcProcessedContent) GetPage(c *gin.Context) {
	req := dto.RskcProcessedContentGetPageReq{}
	s := service.RskcProcessedContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRskcProcessedContentFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RskcProcessedContent, 0)
	var count int64

	err = s.GetPage(c, &req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRskcProcessedContentFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取报文处理
// @Summary 获取报文处理
// @Description 获取报文处理
// @Tags 报文处理
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RskcProcessedContent} "{"code": 200, "data": [...]}"
// @Router /api/v1/rskc-processed-content/{id} [get]
// @Security Bearer
func (e RskcProcessedContent) Get(c *gin.Context) {
	req := dto.RskcProcessedContentGetReq{}
	resp := dto.RskcProcessedContentGetResp{}
	s := service.RskcProcessedContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRskcProcessedContentFail", err))
		return
	}
	var object models.RskcProcessedContent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(c, &req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRskcProcessedContentFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建报文处理
// @Summary 创建报文处理
// @Description 创建报文处理
// @Tags 报文处理
// @Accept application/json
// @Product application/json
// @Param data body dto.RskcProcessedContentInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rskc-processed-content [post]
// @Security Bearer
func (e RskcProcessedContent) Insert(c *gin.Context) {
	req := dto.RskcProcessedContentInsertReq{}
	s := service.RskcProcessedContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRskcProcessedContentFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRskcProcessedContentFail", err))
		return
	}

	e.OK(req.GetId())
}

// Update 修改报文处理
// @Summary 修改报文处理
// @Description 修改报文处理
// @Tags 报文处理
// @Accept application/json
// @Product application/json
// @Param data body dto.RskcProcessedContentUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rskc-processed-content/{id} [put]
// @Security Bearer
func (e RskcProcessedContent) Update(c *gin.Context) {
	req := dto.RskcProcessedContentUpdateReq{}
	s := service.RskcProcessedContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRskcProcessedContentFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRskcProcessedContentFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除报文处理
// @Summary 删除报文处理
// @Description 删除报文处理
// @Tags 报文处理
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rskc-processed-content [delete]
// @Security Bearer
func (e RskcProcessedContent) Delete(c *gin.Context) {
	s := service.RskcProcessedContent{}
	req := dto.RskcProcessedContentDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRskcProcessedContentFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRskcProcessedContentFail", err))
		return
	}
	e.OK(req.GetId())
}

// Export 导出报文处理列表
// @Summary 导出报文处理列表
// @Description 导出报文处理列表
// @Tags 报文处理
// @Param id query int64 false "主键"
// @Param contentId query int64 false "外键(rskc_origin_content.id)"
// @Param content query string false "数据(json字符串格式)"
// @Param statusCode query int false "状态码"
// @Param created_at query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Router /api/v1/rskc-processed-content/export [get]
// @Security Bearer
func (e RskcProcessedContent) Export(c *gin.Context) {
	req := dto.RskcProcessedContentGetPageReq{}
	s := service.RskcProcessedContent{}
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
