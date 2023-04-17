package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerEnterpriseProductRouter)
}

// registerEnterpriseProductRouter
func registerEnterpriseProductRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.EnterpriseProduct{}
	r := v1.Group("/enterprise-product").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
