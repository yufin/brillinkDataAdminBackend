package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerEnterpriseCertificationRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerEnterpriseCertificationRouterNoCheck)
}

// registerEnterpriseCertificationRouter
func registerEnterpriseCertificationRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.EnterpriseCertification{}
	r := v1.Group("/spider/enterprise-certification").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		//r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerEnterpriseCertificationRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.EnterpriseCertification{}
	r := v1.Group("/spider/enterprise-certification")
	{
		r.POST("", api.Insert)
	}
}
