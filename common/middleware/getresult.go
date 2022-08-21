package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
)

// GetResult 获取API中Result数据
func GetResult(c *gin.Context) string {
	rt, bl := c.Get("result")
	var result = ""
	if bl {
		rb, err := json.Marshal(rt)
		if err != nil {
			log.Warnf("json Marshal result error, %s", err.Error())
		} else {
			result = string(rb)
		}
	}
	if result == "" {
		result = "{}"
	}
	return result
}
