package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/jwtauth"

	"go-admin/app/admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDashboardRouter)
}

// registerSysApiRouter
func registerDashboardRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	api := apis.Dashboard{}
	r := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/fake_analysis_chart_data", api.DashboardA)
		r.GET("/fake_workplace_chart_data", api.DashboardA)
		r.GET("/activities", api.Activities)
		r.GET("/notice", api.Notice)
		r.GET("/notices", api.Notices)

		r.GET("/currentUserDetail", api.CurrentUserDetail)
		r.GET("/fake_list_Detail", api.FakeListDetail)
		r.GET("/accountSettingCurrentUser", api.CurrentUserDetail)
	}
}
