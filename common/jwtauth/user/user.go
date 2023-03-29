package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	jwt "go-admin/common/jwtauth"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get(jwt.JwtPayloadKey)
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func Get(c *gin.Context, key string) interface{} {
	data := ExtractClaims(c)
	if data[key] != nil {
		return data[key]
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " Get 缺少 " + key)
	return nil
}

func GetUserId(c *gin.Context) int64 {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return int64((data["identity"]).(float64))
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserId 缺少 identity")
	return 0
}

func GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return pkg.Int64ToString(int64((data["identity"]).(float64)))
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserIdStr 缺少 identity")
	return ""
}

func GetUserName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["username"] != nil {
		return (data["username"]).(string)
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserName 缺少 username")
	return ""
}

func GetNick(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["nick"] != nil {
		return (data["nick"]).(string)
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetNike 缺少 nick")
	return ""
}

func GetRoleKey(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["roleKey"] != nil {
		return (data["roleKey"]).(string)
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetRoleName 缺少 roleKey")
	return ""
}
func GetRoleName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["roleName"] != nil {
		return (data["roleName"]).(string)
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetRoleName 缺少 roleName")
	return ""
}

func GetRoleId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["roleId"] != nil {
		i := int((data["roleId"]).(float64))
		return i
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetRoleId 缺少 roleId")
	return 0
}

func GetDeptId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["deptid"] != nil {
		i := int((data["deptid"]).(float64))
		return i
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetDeptId 缺少 deptid")
	return 0
}

func GetDeptName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["deptkey"] != nil {
		return (data["deptkey"]).(string)
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetDeptName 缺少 deptkey")
	return ""
}
