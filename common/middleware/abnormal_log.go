package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"go-admin/common/global"
	"go-admin/common/jwtauth/user"
	"go-admin/utils"
)

// AbnormalLogger 异常日志
func AbnormalLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		rt, bl := c.Get("abnormalLog")
		if bl {
			mp := rt.(map[string]interface{})
			SetLog(c, mp, global.AbnormalLog)
		}
	}

}

// SetContextAbnormalLog 设置上下文中操作日志信息
func SetContextAbnormalLog(c *gin.Context, abInfo, abSource, abFunc, stackTrace string) {
	mp := make(map[string]interface{}, 0)
	mp["method"] = c.Request.Method                     // 请求方式
	mp["url"] = c.Request.RequestURI                    // 请求地址
	mp["ip"] = c.ClientIP()                             // ip
	mp["abInfo"] = abInfo                               // 异常信息
	mp["abSource"] = abSource                           // 异常来源
	mp["abFunc"] = abFunc                               // 异常方法
	mp["userName"] = user.GetUserName(c)                // 操作人
	mp["userId"] = int64(user.GetUserId(c))             // 用户id
	mp["headers"] = InterfaceToString(c.Request.Header) // 请求头
	mp["body"], _ = c.Get("body")                       // 请求数据
	mp["stackTrace"] = stackTrace                       // 堆栈追踪
	c.Set("abnormalLog", mp)
	utils.SendWechatAlertWithErr(c.GetHeader(pkg.TrafficKey), c.Request.Method, c.Request.RequestURI, c.ClientIP(), abInfo, abSource, abFunc)
}

// SetLog 写入操作日志表
func SetLog(c *gin.Context, result map[string]interface{}, topicType global.TopicType) {
	log := api.GetRequestLogger(c)
	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	resp, _ := c.Get("result")
	result["resp"] = InterfaceToString(resp)
	message, err := sdk.Runtime.GetStreamMessage("", string(topicType), result)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}
