package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerEnterpriseIndustryRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerEnterpriseIndustryRouterNoCheck)
}

// registerEnterpriseIndustryRouter
func registerEnterpriseIndustryRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.EnterpriseIndustry{}
	r := v1.Group("/enterprise-industry").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		//r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerEnterpriseIndustryRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.EnterpriseIndustry{}
	r := v1.Group("/enterprise-industry")
	{
		r.POST("", api.Insert)
	}
}
