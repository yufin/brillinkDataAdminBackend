package models

import "go-admin/common/models"

type OssMetadata struct {
	models.Model
	ObjName    string `json:"objName" gorm:"comment:对象名称"`
	BucketName string `json:"bucketName" gorm:"comment:存储桶名称"`
	Endpoint   string `json:"endpoint" gorm:"comment:存储桶地址"`
	App        int    `json:"app" gorm:"comment:app"`
	models.ModelTime
	models.ControlBy
}

func (*OssMetadata) TableName() string {
	return "oss_metadata"
}

func (e *OssMetadata) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OssMetadata) GetId() interface{} {
	return e.Id
}
