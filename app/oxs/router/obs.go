package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/oxs/apis"
	"go-admin/common/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerAppGrayscaleStrategyRouter)
}

// registerAppGrayscaleStrategyRouter
func registerAppGrayscaleStrategyRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.OXS{}
	r := v1.Group("/oxs") //.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetOBS)
	}
	r2 := v1.Group("/oss") //.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r2.GET("", api.GetOSS)
	}
	r3 := v1.Group("/cos") //.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r3.GET("", api.GetCOS)
	}
	r4 := v1.Group("/kodo") //.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r4.GET("", api.GetKodo)
	}
}
