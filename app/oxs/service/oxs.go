/**
* @Author: Akiraka
* @Date: 2022/8/17 10:55
 */

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/common/service"
)

type OXS struct {
	service.Service
}

func (e OXS) OXS(c *gin.Context) string {
	sdk.Runtime.GetConfig(c.Request.Host, "        \"oxs_provisional_auth\": \"false\",\n")

	if sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth") == "true" {
		switch sdk.Runtime.GetConfig(c.Request.Host, "oxs_type") {
		case "obs":
		case "oss":
		case "cos":
		case "kodo":
		}
	} else {
		switch sdk.Runtime.GetConfig(c.Request.Host, "oxs_type") {
		case "obs":
			e.GetOBS(c)
		case "oss":
			e.GetOSS(c)
		case "cos":
			e.GetCOS(c)
		case "kodo":
			e.GetKodo(c)
		}
	}

	return "sss"
}
