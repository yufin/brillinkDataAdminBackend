package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/rc/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRcDecisionParamRouter)
}

// registerRcDecisionParamRouter
func registerRcDecisionParamRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RcDecisionParam{}
	r := v1.Group("/rc-decision-param").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/export", api.Export)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
