package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/cms/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTbCmsFriendlinkRouter)
}

// registerTbCmsFriendlinkRouter
func registerTbCmsFriendlinkRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.TbCmsFriendlink{}
	r := v1.Group("/tb-cms-friendlink").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
