package server

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/notice/models"
	"go-admin/app/notice/router"
	"go-admin/common/global"
)

func init() {
	sdk.Runtime.SetAppRouters(router.InitRouter)
	//注册监听函数
	Queue.Register(string(global.Notice), models.SaveTbNotice)
}
