package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/rc/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerRcOriginContentRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerRcOriginContentRouterNoAuth)
}

// registerRcOriginContentRouter
func registerRcOriginContentRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RcOriginContent{}
	r := v1.Group("/rc/origin-content").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerRcOriginContentRouterNoAuth(v1 *gin.RouterGroup) {
	api := apis.RcOriginContent{}
	r := v1.Group("/rc/origin-content")
	{
		r.PUT("/re-sync", api.RePubNewContentId)
	}
}
