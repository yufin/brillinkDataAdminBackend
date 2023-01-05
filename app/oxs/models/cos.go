/**
* @Author: Akiraka
* @Date: 2022/10/17 10:53
 */

package models

// PolicyCOS 腾讯云 授予该临时证书权限的CAM策略
type PolicyCOS struct {
	Version   string    `json:"version"`
	Statement Statement `json:"statement"`
}

// Statement 腾讯云
type Statement struct {
	Effect   string   `json:"effect"`
	Action   []string `json:"action"`
	Resource []string `json:"resource"`
}

type ResponseCOS struct {
	// 开关
	Enable bool `json:"enable"`
	// 临时身份验证
	ProvisionalAuth bool `json:"provisionalAuth"`
	// 对象存储类型
	OxsType string `json:"oxsType"`
	// 地区 Region 阿里 腾讯 七牛
	Region string `json:"region"`
	// 阿里 华为 腾讯
	Bucket string `json:"bucket"`
	// 凭证 Credential 阿里 华为 腾讯
	Credential interface{} `json:"credential"`
	// 腾讯 ExpiredTime 临时证书有效的时间，返回 Unix 时间戳，精确到秒
	ExpiredTime *uint64 `json:"expiredTime"`
	// 状态
	Status bool `json:"status"`
}

type ResponseCOSAccessSecret struct {
	// 开关
	Enable bool `json:"enable"`
	// 临时身份验证
	ProvisionalAuth bool `json:"provisionalAuth"`
	// 对象存储类型
	OxsType string `json:"oxsType"`
	// 地区 Region 阿里 腾讯 七牛
	Region    string `json:"region"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	// 阿里 华为 腾讯
	Bucket string `json:"bucket"`
}
