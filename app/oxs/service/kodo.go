/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func (e OXS) GetKodo() (string, error) {
	accessKey, ok := sdk.Runtime.GetConfig("oxs_access_key").(string)
	if !ok {
		err := errors.New("获取 Kodo 的 oxs_access_key 失败")
		return "", err
	}
	secretKey, ok := sdk.Runtime.GetConfig("oxs_secret_key").(string)
	if !ok {
		err := errors.New("获取 Kodo 的 oxs_secret_key 失败")
		return "", err
	}
	bucket, ok := sdk.Runtime.GetConfig("oxs_bucket").(string)
	if !ok {
		err := errors.New("获取 Kodo 的 oxs_bucket 失败")
		return "", err
	}

	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 示例2小时有效期
	expires, ok := sdk.Runtime.GetConfig("oxs_duration_seconds").(uint64)
	if !ok {
		err := errors.New("获取kodo的oxs_duration_seconds失败")
		return "", err
	}
	putPolicy.Expires = expires
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return upToken, nil
}
