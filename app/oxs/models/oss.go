/**
* @Author: Akiraka
* @Date: 2022/10/17 10:56
 */

package models

type ResponseOSS struct {
	// 开关
	Enable bool `json:"enable"`
	// 临时身份验证
	ProvisionalAuth bool `json:"provisionalAuth"`
	// 对象存储类型
	OxsType string `json:"oxsType"`
	// 地区 Region 阿里 腾讯 七牛
	Region string `json:"region"`
	// 阿里 华为 腾讯 七牛
	Bucket string `json:"bucket"`
	// 凭证 Credential 阿里 华为 腾讯
	Credential interface{} `json:"credential"`
	// 状态
	Status bool `json:"status"`
}

type ResponseOSSAccessSecret struct {
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
