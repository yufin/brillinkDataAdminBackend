package version

import (
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
	"runtime"

	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1683285782265Test)
}

func _1683285782265Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Migrator().AddColumn(&SysOperateLog{}, "project")
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}

type SysOperateLog struct {
	LogId        int64  `json:"logId" gorm:"primaryKey;autoIncrement;comment:编码"`
	Type         string `json:"type" gorm:"type:varchar(128);comment:操作类型"`
	Description  string `json:"description" gorm:"type:varchar(128);comment:操作说明"`
	Project      string `json:"project" gorm:"type:varchar(128);comment:项目"`
	UserName     string `json:"userName" gorm:"type:varchar(128);comment:用户"`
	UserId       int64  `json:"userId" gorm:"type:int(11);comment:用户id"`
	UpdateBefore string `json:"updateBefore" gorm:"type:json;comment:更新前"`
	UpdateAfter  string `json:"updateAfter" gorm:"type:json;comment:更新后"`
	models.ModelTime
	models.ControlBy
}

func (SysOperateLog) TableName() string {
	return "sys_operate_log"
}
