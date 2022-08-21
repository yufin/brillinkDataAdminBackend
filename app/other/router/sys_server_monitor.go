package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/jwtauth"

	"go-admin/app/other/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysServerMonitorRouter)
}

// 需认证的路由代码
func registerSysServerMonitorRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.ServerMonitor{}
	r := v1.Group("/server-monitor").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.ServerInfo)
	}
}
