package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerEnterpriseCertificationRouter)
}

// registerEnterpriseCertificationRouter
func registerEnterpriseCertificationRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.EnterpriseCertification{}
	r := v1.Group("/enterprise-certification").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
