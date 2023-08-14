package apis

import (
	"encoding/binary"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/rc/models"
	"go-admin/app/rc/service"
	"go-admin/app/rc/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
	"go-admin/pkg/natsclient"
)

type RcOriginContent struct {
	apis.Api
}

// RePubNewContentId 刷新依赖表
func (e RcOriginContent) RePubNewContentId(c *gin.Context) {
	//contentIdsStr := c.QueryArray("contentId")
	contentIds := dto.IdsReq{}
	//for _, contentIdStr := range contentIdsStr {
	//	contentId, err := strconv.ParseInt(contentIdStr, 10, 64)
	//	if err != nil {
	//		e.Logger.Error(err)
	//		panic(exception.WithMsg(50000, "RePubNewContentIdRcOriginContentFail", err))
	//		return
	//	}
	//	contentIds = append(contentIds, contentId)
	//}

	s := service.RcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&contentIds).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcOriginContentFail", err))
		return
	}

	modelRoc := models.RcOriginContent{}
	db := sdk.Runtime.GetDbByKey(modelRoc.TableName())

	for _, contentId := range contentIds.Ids {
		msg := make([]byte, 8)
		binary.BigEndian.PutUint64(msg, uint64(contentId))
		_, err := natsclient.TaskJs.Publish(natsclient.TopicContentNew, msg)
		if err != nil {
			e.Logger.Error(err)
			panic(exception.WithMsg(50000, "RePubNewContentIdRcOriginContentFail", err))
			return
		}
		err = db.Model(&models.RcOriginContent{}).Where("id = ?", contentId).Update("status_code", 0).Error
		if err != nil {
			e.Logger.Error(err)
			panic(exception.WithMsg(50000, "RePubNewContentIdRcOriginContentFail", err))
			return
		}
	}

	e.OK(contentIds)
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
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.RcOriginContent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-origin-content [get]
// @Security Bearer
func (e RcOriginContent) GetPage(c *gin.Context) {
	req := dto.RcOriginContentGetPageReq{}
	s := service.RcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcOriginContentFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.RcOriginContent, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageRcOriginContentFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取微众json存储
// @Summary 获取微众json存储
// @Description 获取微众json存储
// @Tags 微众json存储
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.RcOriginContent} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-origin-content/{id} [get]
// @Security Bearer
func (e RcOriginContent) Get(c *gin.Context) {
	req := dto.RcOriginContentGetReq{}
	resp := dto.RcOriginContentGetResp{}
	s := service.RcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetRcOriginContentFail", err))
		return
	}
	var object models.RcOriginContent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetRcOriginContentFail", err))
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
// @Param data body dto.RcOriginContentInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-origin-content [post]
// @Security Bearer
func (e RcOriginContent) Insert(c *gin.Context) {
	req := dto.RcOriginContentInsertReq{}
	s := service.RcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertRcOriginContentFail", err))
		return
	}
	// 设置创建人
	req.SetCreateBy(int64(user.GetUserId(c)))

	err = s.Insert(&req)
	if err != nil {
		panic(exception.WithMsg(50000, "InsertRcOriginContentFail", err))
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
// @Param data body dto.RcOriginContentUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-origin-content/{id} [put]
// @Security Bearer
func (e RcOriginContent) Update(c *gin.Context) {
	req := dto.RcOriginContentUpdateReq{}
	s := service.RcOriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateRcOriginContentFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateRcOriginContentFail", err))
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
// @Router /api/v1/rc-origin-content [delete]
// @Security Bearer
func (e RcOriginContent) Delete(c *gin.Context) {
	s := service.RcOriginContent{}
	req := dto.RcOriginContentDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteRcOriginContentFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteRcOriginContentFail", err))
		return
	}
	e.OK(req.GetId())
}
