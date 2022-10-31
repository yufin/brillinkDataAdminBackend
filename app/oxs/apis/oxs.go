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
	switch sdk.Runtime.GetConfig("oxs_type") {
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

// GetOBS 华为云对象存储
func (e OXS) GetOBS(c *gin.Context) {
	e.MakeContext(c)
	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)) == false {
		e.OK(models.OXSAkSk{
			Enable:          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
			ProvisionalAuth: e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig("oxs_type").(string),
			Endpoint:        sdk.Runtime.GetConfig("oxs_obs_endpoint").(string),
			AccessDomain:    sdk.Runtime.GetConfig("oxs_access_domain").(string),
			Bucket:          sdk.Runtime.GetConfig("oxs_bucket").(string),
			AccessKey:       sdk.Runtime.GetConfig("oxs_access_key").(string),
			SecretKey:       sdk.Runtime.GetConfig("oxs_secret_key").(string),
		})
	} else {
		s := service.OXS{}
		status, message, res := s.GetOBS()
		e.OK(models.ResponseOXS{
			Enable:          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
			ProvisionalAuth: e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig("oxs_type").(string),
			Endpoint:        sdk.Runtime.GetConfig("oxs_obs_endpoint").(string),
			AccessDomain:    sdk.Runtime.GetConfig("oxs_access_domain").(string),
			Bucket:          sdk.Runtime.GetConfig("oxs_bucket").(string),
			Credential:      res,
			Status:          status,
			Message:         message,
		})
	}

}

// GetOSS 阿里云对象存储
func (e OXS) GetOSS(c *gin.Context) {
	e.MakeContext(c)
	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)) == false {
		e.OK(models.OXSAkSk{
			Enable:          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
			ProvisionalAuth: e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig("oxs_type").(string),
			Region:          "oss-" + sdk.Runtime.GetConfig("oxs_region").(string),
			Bucket:          sdk.Runtime.GetConfig("oxs_bucket").(string),
			AccessKey:       sdk.Runtime.GetConfig("oxs_access_key").(string),
			SecretKey:       sdk.Runtime.GetConfig("oxs_secret_key").(string),
		})
	} else {
		s := service.OXS{}
		status, message, res := s.GetOSS()
		//e.Custom(models.ResponseOXS{
		//	Enable:          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
		//	ProvisionalAuth: e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
		//	OxsType:         sdk.Runtime.GetConfig("oxs_type").(string),
		//	Region:          "oss-" + sdk.Runtime.GetConfig("oxs_region").(string),
		//	Bucket:          sdk.Runtime.GetConfig("oxs_bucket").(string),
		//	Credential: sts.Credentials{
		//		AccessKeyId:     res.Credentials.AccessKeyId,
		//		Expiration:      res.Credentials.Expiration,
		//		AccessKeySecret: res.Credentials.AccessKeySecret,
		//		SecurityToken:   res.Credentials.SecurityToken,
		//	},
		//	Status:  status,
		//	Message: message,
		//})
		e.Custom(gin.H{
			"enable":          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
			"provisionalAuth": e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
			"oxsType":         sdk.Runtime.GetConfig("oxs_type").(string),
			"region":          "oss-" + sdk.Runtime.GetConfig("oxs_region").(string),
			"bucket":          sdk.Runtime.GetConfig("oxs_bucket").(string),
			"credential": sts.Credentials{
				AccessKeyId:     res.Credentials.AccessKeyId,
				Expiration:      res.Credentials.Expiration,
				AccessKeySecret: res.Credentials.AccessKeySecret,
				SecurityToken:   res.Credentials.SecurityToken,
			},
			"status":  status,
			"message": message,
		})
	}

}

// GetCOS 腾讯云对象存储
func (e OXS) GetCOS(c *gin.Context) {
	e.MakeContext(c)
	// 判断是否临时授权访问
	// false 通过ak/sk访问
	// true 通过临时授权访问
	if e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)) == false {
		e.OK(models.OXSAkSk{
			Enable:          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
			ProvisionalAuth: e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig("oxs_type").(string),
			Region:          sdk.Runtime.GetConfig("oxs_region").(string),
			Bucket:          sdk.Runtime.GetConfig("oxs_bucket").(string),
			AccessKey:       sdk.Runtime.GetConfig("oxs_access_key").(string),
			SecretKey:       sdk.Runtime.GetConfig("oxs_secret_key").(string),
		})
	} else {
		s := service.OXS{}
		status, _, res := s.GetCOS()
		e.OK(models.ResponseOXS{
			Enable:          e.IfBool(sdk.Runtime.GetConfig("oxs_enable").(string)),
			ProvisionalAuth: e.IfBool(sdk.Runtime.GetConfig("oxs_provisional_auth").(string)),
			OxsType:         sdk.Runtime.GetConfig("oxs_type").(string),
			Region:          sdk.Runtime.GetConfig("oxs_region").(string),
			Bucket:          sdk.Runtime.GetConfig("oxs_bucket").(string),
			Credential:      res.Response.Credentials,
			ExpiredTime:     res.Response.ExpiredTime,
			Status:          status,
			//Message: string(message),
		})
	}

}

// GetKodo 七牛云对象存储
func (e OXS) GetKodo(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}
	res := s.GetKodo()
	e.OK(models.ResponseOXS{
		Enable:       true,
		OxsType:      sdk.Runtime.GetConfig("oxs_type").(string),
		Region:       sdk.Runtime.GetConfig("oxs_region").(string),
		AccessDomain: sdk.Runtime.GetConfig("oxs_access_domain").(string),
		Bucket:       sdk.Runtime.GetConfig("oxs_bucket").(string),
		Token:        res,
	})
}

// IfBool 判断布尔值
func (e OXS) IfBool(value string) bool {
	if value == "true" {
		return true
	} else {
		return false
	}
}
