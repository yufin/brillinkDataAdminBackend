package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"

	"go-admin/common/apis"
)

type System struct {
	apis.Api
}

// GenerateCaptchaHandler 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登陆
// @Success 200 {object} antd.Response{data=string,id=string,msg=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/captcha [get]
func (e System) GenerateCaptchaHandler(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Error(500, "服务初始化失败！", "")
		return
	}
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		e.Logger.Errorf("DriverDigitFunc error, %s", err.Error())
		e.Error(500, "验证码获取失败", "")
		return
	}
	e.Custom(gin.H{
		"success": true,
		"data":    b64s,
		"id":      id,
		"msg":     "success",
	})
}
