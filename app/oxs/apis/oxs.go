package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/oxs/models"
	"go-admin/common/apis"
	_ "go-admin/common/response/antd"
)

type OXS struct {
	apis.Api
}

func (e OXS) OXS(c *gin.Context) {
	e.MakeContext(c)
	switch sdk.Runtime.GetConfig("oxs_type") {
	case "obs":
		e.GetOBS(c)
	case "oss":
		e.GetOSS(c)
	case "cos":
		e.GetCOS(c)
	case "kodo":
		e.GetKodo(c)
	default:
		e.OK(models.ResponseErr{
			Status:  false,
			Message: "未配置对象存储",
		})
	}
}
