/**
* @Author: Akiraka
* @Date: 2022/10/17 11:41
 */

package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/oxs/models"
	"go-admin/app/oxs/service"
	"go-admin/app/oxs/utils"
)

// GetOBS 华为云对象存储
func (e OXS) GetOBS(c *gin.Context) {
	e.MakeContext(c)
	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)) == false {
		e.OK(models.ResponseOBSAccessSecret{
			Enable:          utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_enable").(string)),
			ProvisionalAuth: utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
			Endpoint:        sdk.Runtime.GetConfig(c.Request.Host, "oxs_obs_endpoint").(string),
			AccessDomain:    sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_domain").(string),
			Bucket:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
			AccessKey:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_key").(string),
			SecretKey:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_secret_key").(string),
		})
	} else {
		s := service.OXS{}
		status, res := s.GetOBS(c)
		if status == true {
			e.OK(models.ResponseOBS{
				Enable:          utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_enable").(string)),
				ProvisionalAuth: utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)),
				OxsType:         sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
				Endpoint:        sdk.Runtime.GetConfig(c.Request.Host, "oxs_obs_endpoint").(string),
				AccessDomain:    sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_domain").(string),
				Bucket:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
				Credential:      res,
				Status:          status,
			})
		} else {
			e.OK(models.ResponseErr{
				Status:  status,
				Message: "请检查华为云OBS相关账号信息",
			})
		}
	}
}
