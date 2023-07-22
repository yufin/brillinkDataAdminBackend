package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/rc/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRcTradesDetailRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerRcTradesDetailRouterNoCheck)
}

// registerRcTradesDetailRouter
func registerRcTradesDetailRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RcTradesDetail{}
	r := v1.Group("/rc/trades-detail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		//r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		//r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerRcTradesDetailRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.RcTradesDetail{}
	r := v1.Group("/rc/trades-detail")
	{
		//r.GET("/task/sync", api.TaskSyncTradesDetail)
		r.GET("", api.GetPage)
		r.PUT("/:id", api.Update)
		r.GET("/join-wait", api.GetJoin)
	}
}
