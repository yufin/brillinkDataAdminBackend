package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"go-admin/app/notice/models"
	"go-admin/app/notice/service"
	"go-admin/app/notice/service/dto"
	"go-admin/common/apis"
	_ "go-admin/common/response/antd"
)

type TbNotice struct {
	apis.Api
}

func (e TbNotice) GetList(c *gin.Context) {
	s := service.TbNotice{}
	req := dto.TbNoticeGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	//数据权限检查
	//p := actions.GetPermission(c)
	list := make([]models.TbNotice, 0)
	var count int64
	err = s.GetList(&req, &list, &count)
	if err != nil {
		panic(err)
		return
	}

	e.OK(list)
}
