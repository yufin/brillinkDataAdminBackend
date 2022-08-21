package router

import (
	"go-admin/app/jobs/apis"
	"go-admin/common/jwtauth"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysJobRouter)
}

// 需认证的路由代码
func registerSysJobRouter(v1 *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {

	api := apis.SysJob{}
	r := v1.Group("/sys-job").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("/:id", api.Delete)
	}
	//r := v1.Group("/sys-jobs").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	//{
	//	sysJob := &models.SysJob{}
	//	r.GET("", actions.PermissionAction(), actions.IndexAction(sysJob, new(dto.SysJobSearch), func() interface{} {
	//		list := make([]models.SysJob, 0)
	//		return &list
	//	}))
	//	r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.SysJobById), func() interface{} {
	//		return &dto.SysJobItem{}
	//	}))
	//	r.POST("", actions.CreateAction(new(dto.SysJobControl)))
	//	r.PUT("", actions.PermissionAction(), actions.UpdateAction(new(dto.SysJobControl)))
	//	r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.SysJobById)))
	//}
	sysJob := apis.SysJob{}

	v1.GET("/job/remove/:id", sysJob.RemoveJobForService)
	v1.GET("/job/start/:id", sysJob.StartJobForService)
}
