package middleware

import (
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/common/exception"

	"go-admin/app/admin/models"
	"go-admin/common/apis"
	"go-admin/common/jwtauth"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := apis.GetRequestLogger(c)
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		db := sdk.Runtime.GetDbByKey(c.Request.Host)

		e := sdk.Runtime.GetCasbinKey(c.Request.Host)
		var res, casbinExclude bool
		var err error
		//检查权限
		if v["roleKey"] == "admin" {
			res = true
			c.Next()
			return
		}
		for _, i := range CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			log.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		role := models.SysRole{}
		err = db.Model(&role).Preload("SysMenu").Where("role_key = ?", v["roleKey"]).First(&role).Error
		if err != nil {
			err = errors.WithStack(err)
			log.Errorf("Read role data error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			panic(err)
			return
		}
		menuList := *role.SysMenu
		for _, menu := range menuList {
			if menu.Permission != "" {
				res, err = e.Enforce(menu.Permission, c.Request.URL.Path, c.Request.Method)
				if res {
					goto success
				}
			}
		}
		res, err = e.Enforce(v["roleKey"], c.Request.URL.Path, c.Request.Method)
		if err != nil {
			err = errors.WithStack(err)
			log.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			panic(err)
			return
		}

	success:
		if res {
			log.Infof("isTrue: %v role: %s method: %s path: %s", res, v["roleKey"], c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, v["roleKey"], c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			panic(exception.New(exception.ApiNoPermissionFail, errors.New(c.Request.URL.Path+"接口访问权限")))
			return
		}

	}
}
