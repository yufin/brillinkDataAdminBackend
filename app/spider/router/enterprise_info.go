package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerEnterpriseInfoRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerEnterpriseInfoRouterNoCheck)
}

// registerEnterpriseInfoRouter
func registerEnterpriseInfoRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.EnterpriseInfo{}
	r := v1.Group("/spider/enterprise-info").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		//r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerEnterpriseInfoRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.EnterpriseInfo{}
	r := v1.Group("/spider/enterprise-info")
	{
		r.POST("", api.Insert)
	}
}
