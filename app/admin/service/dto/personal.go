package dto

import (
	"go-admin/app/admin/models"
	common "go-admin/common/models"
	"time"
)

type Personal struct {
	Name        string     `json:"name"`
	Avatar      string     `json:"avatar"`
	Userid      string     `json:"userid"`
	Email       string     `json:"email"`
	Signature   string     `json:"signature"`
	Title       string     `json:"title"`
	Group       string     `json:"group"`
	Tags        []Tag      `json:"tags"`
	NotifyCount int        `json:"notifyCount"`
	UnreadCount int        `json:"unreadCount"`
	Country     string     `json:"country"`
	Access      string     `json:"access"`
	AccessList  []string   `json:"accessList"`
	Geographic  Geographic `json:"geographic"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
	Mobile      string     `json:"mobile"`
}

type Tag struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}
type Geographic struct {
	Province Province `json:"province"`
	City     City     `json:"city"`
}

type Province struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

type City struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

type UpdatePersonalReq struct {
	UserId   int    `json:"userId" comment:"用户ID"` // 用户ID
	NickName string `json:"nickName" comment:"昵称" vd:"len($)>0"`
	Phone    string `json:"phone" comment:"手机号"`
	Email    string `json:"email" comment:"邮箱"`
	Avatar   string `json:"avatar" comment:"头像"`
	Sex      string `json:"sex" comment:"性别"`
	//Country    string `json:"country" comment:"国家"`
	Address string `json:"address" comment:"地址"`
	common.ControlBy
}

func (s *UpdatePersonalReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.Avatar = s.Avatar
	model.Sex = s.Sex
	model.Email = s.Email
	//model.Country = s.Country
	model.Address = s.Address
	model.UpdatedAt = time.Now()
}

func (s *UpdatePersonalReq) GetId() interface{} {
	return s.UserId
}

type UpdatePersonalAvatarReq struct {
	UserId int    `json:"userId" comment:"用户ID" vd:"len($)>0"` // 用户ID
	Avatar string `json:"avatar" comment:"头像" vd:"len($)>0"`
	common.ControlBy
}

func (s *UpdatePersonalAvatarReq) GetId() interface{} {
	return s.UserId
}

func (s *UpdatePersonalAvatarReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Avatar = s.Avatar
}
