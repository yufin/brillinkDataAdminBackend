package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/admin/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysOperateLogRouter)
}

// registerSysOperateLogRouter
func registerSysOperateLogRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.SysOperateLog{}
	r := v1.Group("/sys-operate-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
	}
}
