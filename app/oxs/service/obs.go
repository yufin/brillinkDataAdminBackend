package service

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	"github.com/pkg/errors"
	"go-admin/app/oxs/utils"
)

func (e OXS) GetOBS() *model.Credential {
	token, err := e.KeystoneCreateUserTokenByPassword()
	if err != nil {
		return nil
	}
	credential, err := e.CreateTemporaryAccessKeyByAgency(token)
	if err != nil {
		return nil
	}
	return credential.Credential
}

// KeystoneCreateUserTokenByPassword 获取IAM用户Token(使用密码)
func (e OXS) KeystoneCreateUserTokenByPassword() (string, error) {
	nameDomain, ok := sdk.Runtime.GetConfig("oxs_obs_username").(string)
	if !ok {
		err := errors.New("获取 COS 的 oxs_obs_username 失败")
		return "", err
	}
	oxsObsPassword, ok := sdk.Runtime.GetConfig("oxs_obs_password").(string)
	if !ok {
		err := errors.New("获取 OBS 的 oxs_obs_password 失败")
		return "", err
	}
	domainScope := &model.AuthScopeDomain{
		Name: &nameDomain,
	}
	scopeAuth := &model.AuthScope{
		Domain: domainScope,
	}
	domainUser := &model.PwdPasswordUserDomain{
		Name: nameDomain, //"填写 子账号 或者主账号",
	}
	userPassword := &model.PwdPasswordUser{
		Domain:   domainUser,
		Name:     nameDomain, //"填写 子账号 或者主账号",
		Password: oxsObsPassword,
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

	return XSubjectToken, nil
}

// CreateTemporaryAccessKeyByAgency 通过委托获取临时访问密钥
func (e OXS) CreateTemporaryAccessKeyByAgency(XSubjectToken string) (model.CreateTemporaryAccessKeyByTokenResponse, error) {
	expires, ok := sdk.Runtime.GetConfig("oxs_duration_seconds").(uint64)
	if !ok {
		err := errors.New("获取 COS 的 oxs_duration_seconds 失败")
		return model.CreateTemporaryAccessKeyByTokenResponse{}, err
	}
	// 通过 Token 获取临时访问秘钥
	durationSecondsToken := int32(expires)
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
	authBody := &model.TokenAuth{
		Identity: identityAuth,
	}
	request := &model.CreateTemporaryAccessKeyByTokenRequestBody{
		Auth: authBody,
	}
	res, _ := utils.PostRequest(request, "https://iam.cn-east-2.myhuaweicloud.com/v3.0/OS-CREDENTIAL/securitytokens")
	// 序列化结果
	var credential model.CreateTemporaryAccessKeyByTokenResponse
	err := json.Unmarshal(res, &credential)
	if err != nil {
		return model.CreateTemporaryAccessKeyByTokenResponse{}, err
	}

	return credential, nil
}
