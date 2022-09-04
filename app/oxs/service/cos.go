/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tcErrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	v20180813 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"go-admin/app/oxs/models"
)

func (e OXS) GetCOS() (*v20180813.GetFederationTokenResponse, error) {
	accessKey, ok := sdk.Runtime.GetConfig("oxs_access_key").(string)
	if !ok {
		err := errors.New("获取 COS 的 oxs_access_key 失败")
		return nil, err
	}
	secretKey, ok := sdk.Runtime.GetConfig("oxs_secret_key").(string)
	if !ok {
		err := errors.New("获取 COS 的 oxs_secret_key 失败")
		return nil, err
	}
	bucket, ok := sdk.Runtime.GetConfig("oxs_bucket").(string)
	if !ok {
		err := errors.New("获取 COS 的 oxs_bucket 失败")
		return nil, err
	}
	endpoint, ok := sdk.Runtime.GetConfig("oxs_obs_endpoint").(string)
	if !ok {
		err := errors.New("获取 COS 的 oxs_obs_endpoint 失败")
		return nil, err
	}
	expires, ok := sdk.Runtime.GetConfig("oxs_duration_seconds").(uint64)
	if !ok {
		err := errors.New("获取 COS 的 oxs_duration_seconds 失败")
		return nil, err
	}
	region, ok := sdk.Runtime.GetConfig("oxs_region").(string)
	if !ok {
		err := errors.New("获取 COS 的 oxs_region 失败")
		return nil, err
	}
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		accessKey,
		secretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := v20180813.NewClient(credential, region, cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := v20180813.NewGetFederationTokenRequest()
	// 您可以自定义调用方英文名称，由字母组成。
	request.Name = common.StringPtr(bucket)
	policy := models.PolicyCOS{
		Version: "2.0",
		Statement: models.Statement{
			Effect: "allow",
			Action: []string{
				"cos:ListParts",
				"cos:PutObject",
				"cos:GetObject",
				"cos:HeadObject",
				"cos:OptionsObject",
				"cos:GetObjectTagging",
			},
			Resource: []string{
				"*",
			},
		},
	}
	// 转换成 JSON 格式
	policyJson, _ := json.Marshal(policy)
	request.Policy = common.StringPtr(string(policyJson))
	request.DurationSeconds = common.Uint64Ptr(expires)
	// 返回的resp是一个GetFederationTokenResponse的实例，与请求对象对应
	response, err := client.GetFederationToken(request)
	if _, ok := err.(*tcErrors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		//return
	}
	if err != nil {
		err = errors.WithStack(err)
		panic(err)
	}
	return response, err
}
