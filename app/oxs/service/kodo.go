/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"strconv"
)

func (e OXS) GetKodo() string {
	accessKey := sdk.Runtime.GetConfig("oxs_access_key").(string)
	secretKey := sdk.Runtime.GetConfig("oxs_secret_key").(string)
	bucket := sdk.Runtime.GetConfig("oxs_bucket").(string)

	// 字符串转 uint类型
	durationSeconds, _ := strconv.ParseUint(sdk.Runtime.GetConfig("oxs_duration_seconds").(string), 10, 64)

	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 示例2小时有效期
	putPolicy.Expires = durationSeconds
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return upToken
}
