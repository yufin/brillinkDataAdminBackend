package apis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/oxs/models"
	"go-admin/app/oxs/service"
	"go-admin/common/apis"
	_ "go-admin/common/response/antd"
)

type OXS struct {
	apis.Api
}

func (e OXS) OXS(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}

	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if sdk.Runtime.GetConfig("oxs_provisional_auth").(string) == "false" {
		if sdk.Runtime.GetConfig("oxs_type") == "obs" {
			e.OK(models.OXSAkSk{
				Enable:       e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
				OxsType:      sdk.Runtime.GetConfig("oxs_type").(string),
				Endpoint:     sdk.Runtime.GetConfig("oxs_obs_endpoint").(string),
				AccessDomain: sdk.Runtime.GetConfig("oxs_access_domain").(string),
				Bucket:       sdk.Runtime.GetConfig("oxs_bucket").(string),
				AccessKey:    sdk.Runtime.GetConfig("oxs_access_key").(string),
				SecretKey:    sdk.Runtime.GetConfig("oxs_secret_key").(string),
			})
		} else {
			e.OK(models.OXSAkSk{
				Enable:    e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
				OxsType:   sdk.Runtime.GetConfig("oxs_type").(string),
				Region:    "oss-" + sdk.Runtime.GetConfig("oxs_region").(string),
				Bucket:    sdk.Runtime.GetConfig("oxs_bucket").(string),
				AccessKey: sdk.Runtime.GetConfig("oxs_access_key").(string),
				SecretKey: sdk.Runtime.GetConfig("oxs_secret_key").(string),
			})
		}
	} else {
		switch sdk.Runtime.GetConfig("oxs_type") {
		case "obs":
			res := s.GetOBS()
			e.OK(models.ResponseOXS{
				Enable:       e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
				OxsType:      sdk.Runtime.GetConfig("oxs_type").(string),
				Endpoint:     sdk.Runtime.GetConfig("oxs_obs_endpoint").(string),
				AccessDomain: sdk.Runtime.GetConfig("oxs_access_domain").(string),
				Bucket:       sdk.Runtime.GetConfig("oxs_bucket").(string),
				Credential:   res,
			})
		case "oss":
			res := s.GetOSS()
			e.OK(models.ResponseOXS{
				Enable:  e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
				OxsType: sdk.Runtime.GetConfig("oxs_type").(string),
				Region:  "oss-" + sdk.Runtime.GetConfig("oxs_region").(string),
				Bucket:  sdk.Runtime.GetConfig("oxs_bucket").(string),
				Credential: sts.Credentials{
					AccessKeyId:     res.Credentials.AccessKeyId,
					Expiration:      res.Credentials.Expiration,
					AccessKeySecret: res.Credentials.AccessKeySecret,
					SecurityToken:   res.Credentials.SecurityToken,
				},
			})
		case "cos":
			res := s.GetCOS()
			e.OK(models.ResponseOXS{
				Enable:      e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
				OxsType:     sdk.Runtime.GetConfig("oxs_type").(string),
				Region:      sdk.Runtime.GetConfig("oxs_region").(string),
				Bucket:      sdk.Runtime.GetConfig("oxs_bucket").(string),
				Credential:  res.Response.Credentials,
				ExpiredTime: res.Response.ExpiredTime,
			})
		case "kodo":
			res := s.GetKodo()
			e.OK(models.ResponseOXS{
				Enable:  e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
				OxsType: sdk.Runtime.GetConfig("oxs_type").(string),
				Region:  sdk.Runtime.GetConfig("oxs_region").(string),
				Bucket:  sdk.Runtime.GetConfig("oxs_bucket").(string),
				Token:   res,
			})
		}
	}
}

// IfBool 转行布尔值
func (e OXS) IfBool(value string) bool {
	if value == "true" {
		return true
	} else {
		return false
	}
}
