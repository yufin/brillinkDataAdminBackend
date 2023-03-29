package models

import (
	"go-admin/common/models"
	"time"
)

type TodoList struct {
	Id             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Owner          string    `json:"owner" `
	Title          string    `json:"title"`
	Avatar         string    `json:"avatar"`
	Cover          string    `json:"cover"`
	Status         string    `json:"status"`
	Percent        int       `json:"percent"`
	Logo           string    `json:"logo"`
	Href           string    `json:"href"`
	EndAt          time.Time `json:"endAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	SubDescription string    `json:"subDescription"`
	Description    string    `json:"description"`
	ActiveUser     int       `json:"activeUser"`
	NewUser        int       `json:"newUser"`
	Star           int       `json:"star"`
	Like           int       `json:"like"`
	Message        int       `json:"message"`
	Priority       string    `json:"priority"`
	Content        string    `json:"content"`
	Members        []Member  `json:"members" gorm:"many2many:tb_member_rule"`
	models.ControlBy
	models.ModelTime
}

type Member struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Id     int64  `json:"id" gorm:"primaryKey;autoIncrement"`
}

func (Member) TableName() string {
	return "tb_member"
}

func (TodoList) TableName() string {
	return "tb_todo"
}
