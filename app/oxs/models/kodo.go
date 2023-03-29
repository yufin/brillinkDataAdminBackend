/**
* @Author: Akiraka
* @Date: 2022/10/17 11:05
 */

package models

type ResponseKODO struct {
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
	// 访问域名
	AccessDomain string `json:"accessDomain"`
	// 七牛 Token
	Token string `json:"token"`
	// 七牛 UseCdnDomain 表示是否使用 cdn 加速域名，为布尔值，true 表示使用，默认为 false。
	UseCdnDomain bool `json:"useCdnDomain"`
	// 七牛 UseHttpsDomain 是否使用https域名
	UseHttpsDomain bool `json:"useHttpsDomain"`
	// 状态
	Status bool `json:"status"`
}
