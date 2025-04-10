package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/rc/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRcProcessedContentRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerRcProcessedContentRouterNoAuth)
}

// registerRcProcessedContentRouter
func registerRcProcessedContentRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RcProcessedContent{}
	r := v1.Group("/rc/processed-content").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/export", api.Export)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

func registerRcProcessedContentRouterNoAuth(v1 *gin.RouterGroup) {
	api := apis.RcProcessedContent{}
	r := v1.Group("/rc/processed-content")
	{
		r.GET("/report-builder-test", api.ReportBuilderTest)
	}
}
