/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	v20180813 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"go-admin/app/oxs/models"
	"strconv"
)

func (e OXS) GetCOS() (status bool, message string, result *v20180813.GetFederationTokenResponse) {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		sdk.Runtime.GetConfig("oxs_access_key").(string),
		sdk.Runtime.GetConfig("oxs_secret_key").(string),
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sts.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := v20180813.NewClient(credential, sdk.Runtime.GetConfig("oxs_region").(string), cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := v20180813.NewGetFederationTokenRequest()
	// 您可以自定义调用方英文名称，由字母组成。
	request.Name = common.StringPtr("quark")
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

	// 字符串转 uint类型
	durationSeconds, _ := strconv.ParseUint(sdk.Runtime.GetConfig("oxs_duration_seconds").(string), 10, 64)

	request.DurationSeconds = common.Uint64Ptr(durationSeconds)
	// 返回的resp是一个GetFederationTokenResponse的实例，与请求对象对应
	response, err := client.GetFederationToken(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		//return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	//fmt.Printf("%s", response.ToJsonString())

	return true, "", response
}
