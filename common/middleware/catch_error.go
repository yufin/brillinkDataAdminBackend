package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/common/exception"
	"go-admin/common/response/antd"
	"log"
	"net/http"
	"strings"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method
				log.Printf("| url [%s] | method | [%s] | error [%s] |", url, method, err)
				codeX := "-"
				funcX := "-"
				infoX := "-"
				e, ok := err.(exception.Exception)
				if !ok {
					if e, bl := err.(error); bl {
						infoX, codeX, funcX = caller(e)
						SetContextAbnormalLog(c, infoX, codeX, funcX, fmt.Sprintf("%+v", e))
					}
					antd.Failed(c, http.StatusOK, 500, "未知错误，请联系管理员！", 2)
					return
				}
				infoX, codeX, funcX = caller(e.Err)
				SetContextAbnormalLog(c, infoX, codeX, funcX, fmt.Sprintf("%+v", e.Err))
				errorMessage := "系统异常"
				if e.Msg != "" {
					errorMessage = e.Msg
				} else {
					errorMessage, _ = exception.StatusText(e.ErrCode)
				}
				antd.Failed(c, e.Code, e.ErrCode, errorMessage, 2)
			}
		}()
		c.Next()
	}
}

func caller(e error) (info, codeName, funcName string) {
	s := fmt.Sprintf("%+v", e)
	first := strings.Split(s, "$$$")
	if len(first) > 0 {
		list := strings.Split(first[0], "\n")
		if len(list) >= 3 {
			project := strings.Split(list[1], "/")[1]
			i := strings.Index(list[2], project)
			name := ""
			if i != 0 {
				name = list[2][i:]
			}
			return list[0], list[1], name
		}
	}
	return "-", "-", "-"
}
