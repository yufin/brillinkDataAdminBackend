/**
* @Author: Akiraka
* @Date: 2022/8/17 10:09
 */

package service

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	"go-admin/app/oxs/utils"
	"strconv"
)

func (e OXS) GetOBS() *model.Credential {
	token := e.KeystoneCreateUserTokenByPassword()
	credential := e.CreateTemporaryAccessKeyByAgency(token)

	return credential.Credential
}

// KeystoneCreateUserTokenByPassword 获取IAM用户Token(使用密码)
func (e OXS) KeystoneCreateUserTokenByPassword() string {
	nameDomain := sdk.Runtime.GetConfig("oxs_obs_main_username").(string)
	domainScope := &model.AuthScopeDomain{
		Name: &nameDomain,
	}
	scopeAuth := &model.AuthScope{
		Domain: domainScope,
	}
	domainUser := &model.PwdPasswordUserDomain{
		Name: sdk.Runtime.GetConfig("oxs_obs_main_username").(string),
	}
	userPassword := &model.PwdPasswordUser{
		Domain:   domainUser,
		Name:     sdk.Runtime.GetConfig("oxs_obs_iam_username").(string),
		Password: sdk.Runtime.GetConfig("oxs_obs_iam_password").(string),
	}
	passwordIdentity := &model.PwdPassword{
		User: userPassword,
	}
	var listMethodsIdentity = []model.PwdIdentityMethods{
		model.GetPwdIdentityMethodsEnum().PASSWORD,
	}
	identityAuth := &model.PwdIdentity{
		Methods:  listMethodsIdentity,
		Password: passwordIdentity,
	}
	authbody := &model.PwdAuth{
		Identity: identityAuth,
		Scope:    scopeAuth,
	}
	request := &model.KeystoneCreateUserTokenByPasswordRequestBody{
		Auth: authbody,
	}

	_, XSubjectToken := utils.PostRequest(request, "https://iam.cn-east-2.myhuaweicloud.com/v3/auth/tokens")

	return XSubjectToken
}

// CreateTemporaryAccessKeyByAgency 通过委托获取临时访问密钥
func (e OXS) CreateTemporaryAccessKeyByAgency(XSubjectToken string) model.CreateTemporaryAccessKeyByTokenResponse {

	// 字符串转 int类型
	durationSeconds, _ := strconv.Atoi(sdk.Runtime.GetConfig("oxs_duration_seconds").(string))

	// 通过 Token 获取临时访问秘钥
	durationSecondsToken := int32(durationSeconds)
	tokenIdentity := &model.IdentityToken{
		Id:              &XSubjectToken,
		DurationSeconds: &durationSecondsToken,
	}
	var listMethodsIdentity = []model.TokenAuthIdentityMethods{
		model.GetTokenAuthIdentityMethodsEnum().TOKEN,
	}
	identityAuth := &model.TokenAuthIdentity{
		Methods: listMethodsIdentity,
		Token:   tokenIdentity,
	}
	authbody := &model.TokenAuth{
		Identity: identityAuth,
	}
	request := &model.CreateTemporaryAccessKeyByTokenRequestBody{
		Auth: authbody,
	}
	res, _ := utils.PostRequest(request, "https://iam.cn-east-2.myhuaweicloud.com/v3.0/OS-CREDENTIAL/securitytokens")
	// 序列化结果
	var credential model.CreateTemporaryAccessKeyByTokenResponse
	json.Unmarshal(res, &credential)

	return credential
}
