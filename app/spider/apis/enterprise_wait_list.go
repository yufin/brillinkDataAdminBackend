package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"go-admin/app/spider/models"
	"go-admin/app/spider/service"
	"go-admin/app/spider/service/dto"
	"go-admin/app/spider/task"
	"go-admin/common/actions"
	"go-admin/common/apis"
	dtoCommon "go-admin/common/dto"
	"go-admin/common/exception"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
	"go-admin/utils"
	"math"
	"net/url"
	"strconv"
)

type EnterpriseWaitList struct {
	apis.Api
}

func (e EnterpriseWaitList) GetSnowFlakeId(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetSnowFlakeIdFail", err))
		return
	}
	// amount to int
	amount, errConv := strconv.Atoi(c.DefaultQuery("amount", "1"))
	if errConv != nil {
		amount = 1
	}
	// max amount = 99999
	amount = int(math.Min(float64(amount), 99999))
	resp := make([]int64, 0)
	for i := 0; i < amount; i++ {
		resp = append(resp, utils.NewFlakeId())
	}
	e.OK(resp)
}

// GetPageWaitingForCollect 获取等待采集的等待列表(statusCode=2)
func (e EnterpriseWaitList) GetPageWaitingForCollect(c *gin.Context) {
	paginationReq := dtoCommon.Pagination{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().Bind(&paginationReq).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}
	req := dto.EnterpriseWaitListGetPageReq{
		Pagination: paginationReq,
		StatusCode: 2,
		EnterpriseWaitListPageOrder: dto.EnterpriseWaitListPageOrder{
			PriorityOrder: "desc",
		},
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseWaitList, 0)

	var count int64
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}

	e.PageOK(list, count, paginationReq.GetPageIndex(), paginationReq.GetPageSize())
}

// UpdateMatchedIdent 插入匹配的url,并检查数据完整性，修改为对应的状态码
// @Summary 通过id与名称插入匹配的url,
// @Description 修改待爬取列表
// @Tags 待爬取列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseWaitListUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-wait-list/qccUrl/{id} [put]
// @Security Bearer
func (e EnterpriseWaitList) UpdateMatchedIdent(c *gin.Context) {
	req := dto.EnterpriseWaitListUpdateReq{}
	sWait := service.EnterpriseWaitList{}
	if err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&sWait.Service).Errors; err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}

	// assert qccUrl is "" or nil
	if req.UscId == "" || req.UscId == "-" || len(req.UscId) != 18 {
		err := errors.Errorf("字段校验失败: 统一社会信用代码(uscId)不能为空/格式不正确")
		panic(exception.WithMsg(500, "UpdateEnterpriseWaitListFail", err))
		return
	}
	p := actions.GetPermissionFromContext(c)
	req.SetUpdateBy(int64(user.GetUserId(c)))
	err := sWait.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	if err := task.SyncTaskDetail(req.UscId); err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail-syncTaskDetailFail", err))
		return
	}
	e.OK(req.GetId())
}

// GetPageWaitingForIdent 获取企业等待匹配列表
// @Summary 获取企业等待匹配url的列表: 条件: qccUrl为空字符串, statusCode=0
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Param statusCode query int 状态码: 1.等待匹配url 2.等待爬取主体信息(enterprise_info), 3.等待爬取其他信息(tag,industry...)，4.完成爬取
// @Router /api/v1/enterprise-wait-lit/waiting [get]
func (e EnterpriseWaitList) GetPageWaitingForIdent(c *gin.Context) {
	paginationReq := dtoCommon.Pagination{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().Bind(&paginationReq).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}
	req := dto.EnterpriseWaitListGetPageReq{
		Pagination: paginationReq,
		StatusCode: 1,
		EnterpriseWaitListPageOrder: dto.EnterpriseWaitListPageOrder{
			PriorityOrder: "desc",
		},
	}
	req.EnterpriseWaitListPageOrder.PriorityOrder = "desc"

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseWaitList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}
	e.PageOK(list, count, paginationReq.GetPageIndex(), paginationReq.GetPageSize())
}

// UpdateAsIllegal 通过Id标记该行为非法公司主体(update statusCode=9)
func (e EnterpriseWaitList) UpdateAsIllegal(c *gin.Context) {
	reqTemp := dto.EnterpriseWaitListUpdateReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&reqTemp).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}

	req := dto.EnterpriseWaitListUpdateReq{}
	req.Id = reqTemp.Id
	req.StatusCode = 9

	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	e.OK(req.GetId())
}

// GetPage 获取待爬取列表列表
// @Summary 获取待爬取列表列表
// @Description 获取待爬取列表列表
// @Tags 待爬取列表
// @Param id query int64 false "主键id"
// @Param enterpriseName query string false "企业名称"
// @Param uscId query string false "社会统一信用代码"
// @Param priority query int false "优先级"
// @Param qccUrl query string false "qcc主体网址"
// @Param statusCode query int false "数据爬取状态码"
// @Param source query string false "来源备注"
// @Param deleted_at query time.Time false "删除时间"
// @Param create_by query string false "创建人"
// @Param update_by query string false "更新人"
// @Param createdAt query time.Time false "创建时间"
// @Param updated_at query time.Time false "更新时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} antd.Response{data=antd.Pages{list=[]models.EnterpriseWaitList}} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-wait-list [get]
// @Security Bearer
func (e EnterpriseWaitList) GetPage(c *gin.Context) {
	req := dto.EnterpriseWaitListGetPageReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnterpriseWaitList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseWaitListFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get 获取待爬取列表
// @Summary 获取待爬取列表
// @Description 获取待爬取列表
// @Tags 待爬取列表
// @Param id path string false "id"
// @Success 200 {object} antd.Response{data=models.EnterpriseWaitList} "{"code": 200, "data": [...]}"
// @Router /api/v1/enterprise-wait-list/{id} [get]
// @Security Bearer
func (e EnterpriseWaitList) Get(c *gin.Context) {
	req := dto.EnterpriseWaitListGetReq{}
	resp := dto.EnterpriseWaitListGetResp{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetEnterpriseWaitListFail", err))
		return
	}
	var object models.EnterpriseWaitList

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.WithMsg(50000, "GetEnterpriseWaitListFail", err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert 创建待爬取列表
// @Summary 创建待爬取列表
// @Description 创建待爬取列表
// @Tags 待爬取列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseWaitListInsertReq true "data"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/enterprise-wait-list [post]
// @Security Bearer
func (e EnterpriseWaitList) Insert(c *gin.Context) {
	req := dto.EnterpriseWaitListInsertReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "InsertEnterpriseWaitListFail", err))
		return
	}

	if req.Priority == 0 {
		req.Priority = 1
	}

	u, err := url.Parse(req.QccUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		req.StatusCode = 1 // not valid url
	} else {
		req.StatusCode = 2
	}
	req.SetCreateBy(int64(user.GetUserId(c)))

	p := actions.GetPermissionFromContext(c)
	err = s.Insert(&req)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			// column not unique
			list := make([]models.EnterpriseWaitList, 0)
			var count int64
			err = s.GetPage(&dto.EnterpriseWaitListGetPageReq{
				EnterpriseName: req.EnterpriseName, UscId: req.UscId}, p, &list, &count)
			if err != nil {
				e.Logger.Error(err)
				panic(exception.WithMsg(50000, "InsertEnterpriseWaitListFail", err))
				return
			}
			if len(list) > 0 {
				e.OK(list[0])
				return
			}
		} else {
			e.Logger.Error(err)
			panic(exception.WithMsg(500, "InsertEnterpriseWaitListFail", err))
			return
		}
	}
	e.OK(req)
}

// Update 修改待爬取列表
// @Summary 修改待爬取列表
// @Description 修改待爬取列表
// @Tags 待爬取列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EnterpriseWaitListUpdateReq true "body"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/enterprise-wait-list/{id} [put]
// @Security Bearer
func (e EnterpriseWaitList) Update(c *gin.Context) {
	req := dto.EnterpriseWaitListUpdateReq{}
	s := service.EnterpriseWaitList{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "UpdateEnterpriseWaitListFail", err))
		return
	}
	e.OK(req.GetId())
}

// Delete 删除待爬取列表
// @Summary 删除待爬取列表
// @Description 删除待爬取列表
// @Tags 待爬取列表
// @Param ids body []int false "ids"
// @Success 200 {object} antd.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/enterprise-wait-list [delete]
// @Security Bearer
func (e EnterpriseWaitList) Delete(c *gin.Context) {
	s := service.EnterpriseWaitList{}
	req := dto.EnterpriseWaitListDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "DeleteEnterpriseWaitListFail", err))
		return
	}

	// req.SetUpdateBy(int64(user.GetUserId(c)))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		panic(exception.WithMsg(50000, "DeleteEnterpriseWaitListFail", err))
		return
	}
	e.OK(req.GetId())
}
