/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func (e OXS) GetOSS() *sts.AssumeRoleResponse {
	//构建一个阿里云客户端, 用于发起请求。
	//设置调用者（RAM用户或RAM角色）的AccessKey ID和AccessKey Secret。
	//第一个参数就是bucket所在位置，可查看oss对象储存控制台的概况获取
	//第二个参数就是步骤一获取的AccessKey ID
	//第三个参数就是步骤一获取的AccessKey Secret
	client, err := sts.NewClientWithAccessKey("获取区域", "填写 ak", "填写 sk")

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见《API参考》。
	request.RoleArn = "acs:ram::12021278814:role/oss-quark  这个需要到阿里云去拿" //步骤三获取的角色ARN
	request.RoleSessionName = "OSS-QUARK  填写角色名称"                       //步骤三中的RAM角色名称
	request.DurationSeconds = requests.Integer("获取秒")                   // Token 有效期 默认3600，最小值900秒

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	return response
}
