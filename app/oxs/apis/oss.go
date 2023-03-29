/**
* @Author: Akiraka
* @Date: 2022/10/17 11:41
 */

package apis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/oxs/models"
	"go-admin/app/oxs/service"
	"go-admin/app/oxs/utils"
)

// GetOSS 阿里云对象存储
func (e OXS) GetOSS(c *gin.Context) {
	e.MakeContext(c)
	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)) == false {
		e.OK(models.ResponseOSSAccessSecret{
			Enable:          utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_enable").(string)),
			ProvisionalAuth: utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
			Region:          "oss-" + sdk.Runtime.GetConfig(c.Request.Host, "oxs_region").(string),
			Bucket:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
			AccessKey:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_key").(string),
			SecretKey:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_secret_key").(string),
		})
	} else {
		s := service.OXS{}
		status, res := s.GetOSS(c)
		if status == true {
			e.OK(models.ResponseOSS{
				Enable:          utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_enable").(string)),
				ProvisionalAuth: utils.StringToBool(sdk.Runtime.GetConfig(c.Request.Host, "oxs_provisional_auth").(string)),
				OxsType:         sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
				Region:          "oss-" + sdk.Runtime.GetConfig(c.Request.Host, "oxs_region").(string),
				Bucket:          sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
				Credential: sts.Credentials{
					AccessKeyId:     res.Credentials.AccessKeyId,
					Expiration:      res.Credentials.Expiration,
					AccessKeySecret: res.Credentials.AccessKeySecret,
					SecurityToken:   res.Credentials.SecurityToken,
				},
				Status: status,
			})
		} else {
			e.OK(models.ResponseErr{
				Status:  status,
				Message: "请检查阿里云OSS相关账号信息",
			})
		}
	}
}
