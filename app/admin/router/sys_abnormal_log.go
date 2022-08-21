package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/admin/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysAbnormalLogRouter)
}

// registerSysAbnormalLogRouter
func registerSysAbnormalLogRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.SysAbnormalLog{}
	r := v1.Group("/sys-abnormal-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
	}
}
