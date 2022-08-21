package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysApiRouter)
}

// registerSysApiRouter
func registerSysApiRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.SysApi{}
	r := v1.Group("/sys-api").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/list", api.GetList)
		r.GET("/:id", api.Get)
		r.PUT("/:id", api.Update)
		r.PUT("", api.Delete)
	}
}
