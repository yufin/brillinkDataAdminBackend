package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerEnterpriseWaitListRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerEnterpriseWaitListRouterNoCheck)
}

// registerEnterpriseWaitListRouter
func registerEnterpriseWaitListRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.EnterpriseWaitList{}
	r := v1.Group("/enterprise-wait-list").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerEnterpriseWaitListRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.EnterpriseWaitList{}
	r := v1.Group("/spider/enterprise/wait-list")
	{
		r.GET("/to-collect", api.GetPageWaitingForCollect)
		r.GET("/to-ident", api.GetPageWaitingForIdent)
		r.PUT("/to-ident/:id", api.UpdateMatchedIdent)
		r.PUT("/as-illegal/:id", api.UpdateAsIllegal)
		r.GET("/snowflake", api.GetSnowFlakeId)
		//r.POST("", api.Insert)
		//r.GET("", api.GetPage)
		//r.GET("/:id", api.Get)
		//r.PUT("/:id", api.Update)
		//r.DELETE("", api.Delete)
	}
}
