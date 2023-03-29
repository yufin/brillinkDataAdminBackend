package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/oxs/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerAppGrayscaleStrategyRouter)
}

// registerAppGrayscaleStrategyRouter
func registerAppGrayscaleStrategyRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.OXS{}
	r := v1.Group("/oxs").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.OXS)
	}
}
