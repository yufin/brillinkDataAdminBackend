package apis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"

	"go-admin/app/oxs/models"
	"go-admin/app/oxs/service"
	"go-admin/common/apis"
	_ "go-admin/common/response/antd"
)

type OXS struct {
	apis.Api
}

func (e OXS) GetOBS(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}
	res := s.GetOBS()

	e.OK(models.ResponseOXS{
		Enable:       true,
		OxsType:      "obs",
		Endpoint:     "端点",
		AccessDomain: "访问域名",
		Bucket:       "桶名称",
		Credential:   res,
	})
}

func (e OXS) GetOSS(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}
	res := s.GetOSS()
	e.OK(models.ResponseOXS{
		Enable:  true,
		OxsType: "oss",
		Region:  "区域",
		Bucket:  "go-quark",
		Credential: sts.Credentials{
			AccessKeyId:     res.Credentials.AccessKeyId,
			Expiration:      res.Credentials.Expiration,
			AccessKeySecret: res.Credentials.AccessKeySecret,
			SecurityToken:   res.Credentials.SecurityToken,
		},
	})
}

func (e OXS) GetCOS(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}
	res, _ := s.GetCOS()
	e.OK(models.ResponseOXS{
		Enable:      true,
		OxsType:     "cos",
		Region:      "ap-nanjing",
		Bucket:      "桶名称",
		Credential:  res.Response.Credentials,
		ExpiredTime: res.Response.ExpiredTime,
	})
}

func (e OXS) GetKodo(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}
	res, _ := s.GetKodo()
	e.OK(models.ResponseOXS{
		Enable:       true,
		OxsType:      "kodo",
		Region:       "区域",
		AccessDomain: "访问域名",
		Bucket:       "桶名称",
		Token:        res,
	})
}
