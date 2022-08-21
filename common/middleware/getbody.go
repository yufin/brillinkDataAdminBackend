package middleware

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"io"
	"io/ioutil"
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
		bf := bytes.NewBuffer(nil)
		wt := bufio.NewWriter(bf)
		_, err := io.Copy(wt, c.Request.Body)
		if err != nil {
			log.Warnf("copy body error, %s", err.Error())
		}
		rb, err := ioutil.ReadAll(bf)
		if err != nil {
			log.Warnf("ReadAll error, %s", err.Error())
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
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
