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

type Operate string

const (
	// OperateAdd 新增
	OperateAdd Operate = "新增"
	// OperateUpdate 修改
	OperateUpdate Operate = "修改"
	// OperateDelete 删除
	OperateDelete Operate = "删除"
)

//SetContextOperateLog 设置上下文中操作日志信息
func SetContextOperateLog(c *gin.Context, typeX Operate, desc, before, after string) {
	mp := make(map[string]interface{}, 0)
	mp["type"] = typeX
	mp["description"] = user.GetUserName(c) + desc
	mp["userName"] = user.GetUserName(c)
	mp["userId"] = user.GetUserId(c)
	mp["updateBefore"] = before
	mp["updateAfter"] = after
	c.Set("operateLog", mp)
}
