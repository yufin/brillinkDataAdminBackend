package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/rskc/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRskcOriginContentRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerRskcOriginContentRouterNoAuth)
}

// registerRskcOriginContentRouter
func registerRskcOriginContentRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RskcOriginContent{}
	r := v1.Group("/rskc/origin-content").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerRskcOriginContentRouterNoAuth(v1 *gin.RouterGroup) {
	api := apis.RskcOriginContent{}
	r := v1.Group("/rskc/origin-content")
	{
		r.GET("/task/sync", api.TaskSyncOriginContent)
	}
}
