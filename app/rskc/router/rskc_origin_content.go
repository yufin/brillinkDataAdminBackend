package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/rskc/apis"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerEnterpriseProductRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerRskcOriginContentRouterNoCheck)
}

func registerRskcOriginContentRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.RskcOriginContent{}
	r := v1.Group("/rskc")
	{
		r.GET("/temp", api.ParseJsonFile)
		r.GET("", api.GetPage)
		r.GET("/info", api.GetPageWithoutContent)
		r.GET("/task/syncOriginContent", api.TaskSyncOriginContent)
		r.GET("/noContent", api.GetPageWithoutContent)
	}
}
