package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CopyBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		setBody(c)
		c.Next()
	}
}

func setBody(c *gin.Context) {
	body := ""
	switch c.Request.Method {
	case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
		rb, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
		body = string(rb)
	}
	c.Set("body", body)
}

// GetBody 获取API中Body数据
func GetBody(c *gin.Context) string {
	body, _ := c.Get("body")
	if body.(string) == "" {
		body = "{}"
	}
	return body.(string)
}
