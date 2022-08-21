package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysConfigRouter)
}

// 需认证的路由代码
func registerSysConfigRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.SysConfig{}
	r := v1.Group("/sys-config").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}

	r1 := v1.Group("/configKey").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/:configKey", api.GetSysConfigByKEYForService)
	}

	r2 := v1.Group("/app-config") //.Use(authMiddleware.MiddlewareFunc())
	{
		r2.GET("", api.Get2SysApp)
	}

	r3 := v1.Group("/set-config").Use(authMiddleware.MiddlewareFunc())
	{
		r3.PUT("", api.Update2Set)
		r3.GET("", api.Get2Set)
	}

}
