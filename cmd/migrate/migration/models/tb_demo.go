package models

type TbDemo struct {
	Model
	Name string `json:"name" gorm:"size:128;comment:名称"`
	ModelTime
	ControlBy
}

func (TbDemo) TableName() string {
	return "tb_demo"
}
