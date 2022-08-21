package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/global"
	"go-admin/common/jwtauth/user"
)

// OperateLogger 操作日志
func OperateLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		rt, bl := c.Get("operateLog")
		if bl {
			SetLog(c, rt.(map[string]interface{}), global.OperateLog)
		}
	}
}

//SetContextOperateLog 设置上下文中操作日志信息
func SetContextOperateLog(c *gin.Context, typeX, desc, before, after string) {
	mp := make(map[string]interface{}, 0)
	mp["type"] = typeX
	mp["description"] = user.GetUserName(c) + desc
	mp["userName"] = user.GetUserName(c)
	mp["userId"] = int64(user.GetUserId(c))
	mp["updateBefore"] = before
	mp["updateAfter"] = after
	c.Set("operateLog", mp)
}
