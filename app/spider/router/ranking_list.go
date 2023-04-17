package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/spider/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRankingListRouter)
}

// registerRankingListRouter
func registerRankingListRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.RankingList{}
	r := v1.Group("/ranking-list").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
