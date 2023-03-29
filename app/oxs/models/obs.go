/**
* @Author: Akiraka
* @Date: 2022/10/17 11:03
 */

package models

// OBSErr 请求华为云临时秘钥错误返回
type OBSErr struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Title   string `json:"title"`
	} `json:"error"`
}

type ResponseOBS struct {
	// 开关
	Enable bool `json:"enable"`
	// 临时身份验证
	ProvisionalAuth bool `json:"provisionalAuth"`
	// 对象存储类型
	OxsType string `json:"oxsType"`
	// 端点 Endpoint 华为
	Endpoint string `json:"endpoint"`
	// 访问域名
	AccessDomain string `json:"accessDomain"`
	// 阿里 华为 腾讯 七牛
	Bucket string `json:"bucket"`
	// 凭证 Credential 阿里 华为 腾讯
	Credential interface{} `json:"credential"`
	//  Status
	Status bool `json:"status"`
}

type ResponseOBSAccessSecret struct {
	// 开关
	Enable bool `json:"enable"`
	// 临时身份验证
	ProvisionalAuth bool `json:"provisionalAuth"`
	// 对象存储类型
	OxsType string `json:"oxsType"`
	// 端点 Endpoint 华为
	Endpoint string `json:"endpoint"`
	// 访问域名
	AccessDomain string `json:"accessDomain"`
	AccessKey    string `json:"accessKey"`
	SecretKey    string `json:"secretKey"`
	// 阿里 华为 腾讯 华为
	Bucket string `json:"bucket"`
}
