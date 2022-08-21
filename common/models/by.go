package models

import (
	"gorm.io/gorm"
	"time"
)

type ControlBy struct {
	CreateBy int64 `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int64 `json:"updateBy" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int64) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int64) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

type ModelDriver interface {
	TableName() string
	Generate() ActiveRecord
	GetId() interface{}
}
