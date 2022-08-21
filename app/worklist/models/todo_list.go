package models

import (
	"go-admin/common/models"
	"time"
)

type TbTodo struct {
	Id             int       `json:"id"`
	Owner          string    `json:"owner"`
	Title          string    `json:"title"`
	Avatar         string    `json:"avatar"`
	Cover          string    `json:"cover"`
	Status         string    `json:"status"`
	Percent        int       `json:"percent"`
	Logo           string    `json:"logo"`
	Href           string    `json:"href"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	EndAt          time.Time `json:"endAt"`
	SubDescription string    `json:"subDescription"`
	Description    string    `json:"description"`
	ActiveUser     int       `json:"activeUser"`
	NewUser        int       `json:"newUser"`
	Star           int       `json:"star"`
	Like           int       `json:"like"`
	Message        int       `json:"message"`
	Priority       string    `json:"priority"`
	Content        string    `json:"content"`
	Members        []Member  `json:"members"  gorm:"many2many:tb_member_rule"`
	models.ControlBy
	models.ModelTime
}

type Member struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Id     string `json:"id"`
}

func (Member) TableName() string {
	return "tb_member"
}

func (TbTodo) TableName() string {
	return "tb_todo"
}

func (e *TbTodo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbTodo) GetId() interface{} {
	return e.Id
}
