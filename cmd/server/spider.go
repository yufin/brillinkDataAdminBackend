package server

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/spider/router"
)

func init() {
	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	sdk.Runtime.SetAppRouters(router.InitRouter)
}
