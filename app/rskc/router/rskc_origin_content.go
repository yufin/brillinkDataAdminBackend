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
	r := v1.Group("/rskc/originContent")
	{
		r.GET("/temp", api.ParseJsonFile)
		r.GET("", api.GetPage)
		r.PUT("/:ContentId", api.Update)
		r.GET("/info", api.GetPageWithoutContent)
		r.GET("/task/sync", api.TaskSyncOriginContent)
		r.GET("/noContent", api.GetPageWithoutContent)
	}
}
