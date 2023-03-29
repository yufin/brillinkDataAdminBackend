package handler

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
	Type     string `form:"-" json:"type"`
}

func (u *Login) GetUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("username = ?  and status = 2", u.Username).First(&user).Error
	if err != nil {
		err = errors.WithStack(err)
		log.Errorf("get user error, %s", err.Error())
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		log.Errorf("user login error, %s", err.Error())
		err = errors.New("账号或密码错误")
		return
	}
	err = tx.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		err = errors.WithStack(err)
		log.Errorf("get role error, %s", err.Error())
		return
	}
	return
}
