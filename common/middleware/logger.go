package middleware

import (
	"encoding/json"
	gaConfig "go-admin/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"go-admin/common/global"
	"go-admin/common/jwtauth/user"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		// 开始时间
		startTime := time.Now()
		c.Next()
		if c.Request.Method == http.MethodOptions {
			return
		}
		url := c.Request.RequestURI
		if strings.Index(url, "logout") > -1 ||
			strings.Index(url, "login") > -1 {
			return
		}
		// 结束时间
		endTime := time.Now()
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		logData := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}
		log.WithFields(logData).Info()

		if c.Request.Method != "OPTIONS" && reqUri != "/api/v1/sys/notice" && !strings.Contains(reqUri, "/api/v1/sys-request-log") && !strings.Contains(reqUri, "/api/v1/sys-abnormal-log") && config.LoggerConfig.EnabledDB && statusCode != 404 {
			SetDBOperLog(c, latencyTime)
		}
	}
}

// SetDBOperLog 写入操作日志表 fixme 该方法后续即将弃用
func SetDBOperLog(c *gin.Context, latencyTime time.Duration) {
	log := api.GetRequestLogger(c)
	l := make(map[string]interface{})
	l["_fullPath"] = c.FullPath()
	l["operUrl"] = c.Request.RequestURI
	l["operIp"] = c.ClientIP()
	l["operLocation"] = pkg.GetLocation(c.ClientIP(), gaConfig.ExtConfig.AMap.Key)
	l["operName"] = user.GetUserName(c)
	l["requestMethod"] = c.Request.Method
	l["operHeaders"] = InterfaceToString(c.Request.Header)
	l["operParam"] = GetBody(c)
	l["operTime"] = time.Now()
	l["jsonResult"] = GetResult(c)
	l["latencyTime"] = strconv.FormatInt(latencyTime.Milliseconds(), 10)
	l["userAgent"] = c.Request.UserAgent()
	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := sdk.Runtime.GetStreamMessage("", string(global.RequestLog), l)
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

func InterfaceToString(m interface{}) string {
	headers, _ := json.Marshal(m)
	return string(headers)
}
