/**
* @Author: Akiraka
* @Date: 2022/8/17 10:55
 */

package service

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/common/service"
)

type OXS struct {
	service.Service
}

func (e OXS) OXS() string {
	sdk.Runtime.GetConfig("        \"oxs_provisional_auth\": \"false\",\n")

	if sdk.Runtime.GetConfig("oxs_provisional_auth") == "true" {
		switch sdk.Runtime.GetConfig("oxs_type") {
		case "obs":
		case "oss":
		case "cos":
		case "kodo":
		}
	} else {
		switch sdk.Runtime.GetConfig("oxs_type") {
		case "obs":
			e.GetOBS()
		case "oss":
			e.GetOSS()
		case "cos":
			e.GetCOS()
		case "kodo":
			e.GetKodo()
		}
	}

	return "sss"
}
