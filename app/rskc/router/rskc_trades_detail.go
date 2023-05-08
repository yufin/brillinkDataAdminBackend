package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/rskc/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRskcTradesDetailRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerRskcTradesDetailRouterNoCheck)
}

// registerRskcTradesDetailRouter
func registerRskcTradesDetailRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RskcTradesDetail{}
	r := v1.Group("/rskc/trades-detail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		//r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		//r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerRskcTradesDetailRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.RskcTradesDetail{}
	r := v1.Group("/rskc/trades-detail")
	{
		r.GET("/task/sync", api.TaskSyncTradesDetail)
		r.GET("/task/sync-wait", api.TaskSyncWaitList)
		r.GET("", api.GetPage)
		r.PUT("/:id", api.Update)
	}
}
