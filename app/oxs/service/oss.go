/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
)

func (e OXS) GetOSS(c *gin.Context) (status bool, result *sts.AssumeRoleResponse) {
	//构建一个阿里云客户端, 用于发起请求。
	//设置调用者（RAM用户或RAM角色）的AccessKey ID和AccessKey Secret。
	//第一个参数就是bucket所在位置，可查看oss对象储存控制台的概况获取
	//第二个参数就是步骤一获取的AccessKey ID
	//第三个参数就是步骤一获取的AccessKey Secret
	client, err := sts.NewClientWithAccessKey(
		sdk.Runtime.GetConfig(c.Request.Host, "oxs_region").(string),
		sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_key").(string),
		sdk.Runtime.GetConfig(c.Request.Host, "oxs_secret_key").(string))
	//if err != nil {
	//}

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	request.RoleArn = sdk.Runtime.GetConfig(c.Request.Host, "obx_oss_role_arn").(string)                               //步骤三获取的角色ARN
	request.RoleSessionName = sdk.Runtime.GetConfig(c.Request.Host, "obx_oss_role_session_name").(string)              //步骤三中的RAM角色名称
	request.DurationSeconds = requests.Integer(sdk.Runtime.GetConfig(c.Request.Host, "oxs_duration_seconds").(string)) // Token 有效期 默认3600，最小值900秒

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		return false, nil
		fmt.Print(err.Error())
	} else {
		return true, response
	}
	return false, nil
}
