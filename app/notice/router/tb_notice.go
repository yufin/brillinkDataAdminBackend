package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/jwtauth"

	"go-admin/app/notice/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

// 需认证的路由代码
func registerSysDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.TbNotice{}

	r := v1.Group("/sys/notice") //.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetList)
	}
}
