package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"
	"go-admin/app/rskc/task"
	"go-admin/app/rskc/utils"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/exception"
)

type RskcOriginContent struct {
	apis.Api
}

func (e RskcOriginContent) ParseJsonFile(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	var contentMap map[string]any
	utils.ParseJsonBytes(&contentMap)
	fmt.Println(contentMap)
	e.OK(nil)
}

func (e RskcOriginContent) GetPage(c *gin.Context) {
	req := dto.OriginContentGetPageReq{}
	s := service.OriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.OriginContent, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

func (e RskcOriginContent) GetPageWithoutContent(c *gin.Context) {
	req := dto.OriginContentGetPageReq{}
	s := service.OriginContent{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseFail", err))
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.OriginContentInfo, 0)
	var count int64

	err = s.GetPageWithoutContent(&req, p, &list, &count)
	//err = s.CountByInfo(&req, p, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "GetPageEnterpriseFail", err))
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

func (e RskcOriginContent) TaskSyncOriginContent(c *gin.Context) {
	s := service.OriginContent{}
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
