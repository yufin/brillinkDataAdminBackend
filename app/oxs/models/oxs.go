/**
* @Author: Akiraka
* @Date: 2022/8/17 11:52
 */

package models

type ResponseOXS struct {
	// 开关
	Enable bool `json:"enable"`
	// 对象存储类型
	OxsType string `json:"oxsType"`
	// 端点 Endpoint 华为
	Endpoint string `json:"endpoint"`
	// 地区 Region 阿里 腾讯 七牛
	Region string `json:"region"`
	// 访问域名
	AccessDomain string `json:"accessDomain"`
	// 阿里 华为 腾讯
	Bucket string `json:"bucket"`
	// 凭证 Credential 阿里 华为 腾讯
	Credential interface{} `json:"credential"`
	// 腾讯 Response
	//Response interface{} `json:"response"`
	// 腾讯 ExpiredTime 临时证书有效的时间，返回 Unix 时间戳，精确到秒
	ExpiredTime *uint64 `json:"expiredTime"`
	// 七牛 Token
	Token string `json:"token"`
	// 七牛 UseCdnDomain 表示是否使用 cdn 加速域名，为布尔值，true 表示使用，默认为 false。
	UseCdnDomain bool `json:"useCdnDomain"`
	// 七牛 UseHttpsDomain 是否使用https域名
	UseHttpsDomain bool `json:"useHttpsDomain"`
}

// PolicyCOS 腾讯云 授予该临时证书权限的CAM策略
type PolicyCOS struct {
	Version   string    `json:"version"`
	Statement Statement `json:"statement"`
}

type Statement struct {
	Effect   string   `json:"effect"`
	Action   []string `json:"action"`
	Resource []string `json:"resource"`
}
