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

// GetCOS 腾讯云对象存储
func (e OXS) GetCOS(c *gin.Context) {
	e.MakeContext(c)
	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)) == false {
		e.OK(models.ResponseCOSAccessSecret{
			Enable:          utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_enable").(string)),
			ProvisionalAuth: utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
			Region:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_region").(string),
			Bucket:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
			AccessKey:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_key").(string),
			SecretKey:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_secret_key").(string),
		})
	} else {
		s := service.OXS{}
		status, res := s.GetCOS(c)
		if status == true {
			e.OK(models.ResponseCOS{
				Enable:          utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_enable").(string)),
				ProvisionalAuth: utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)),
				OxsType:         sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
				Region:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_region").(string),
				Bucket:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
				Credential:      res.Response.Credentials,
				ExpiredTime:     res.Response.ExpiredTime,
				Status:          status,
			})
		} else {
			e.OK(models.ResponseErr{
				Status:  status,
				Message: "请检查腾讯云COS相关账号信息",
			})
		}
	}
}
