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
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1644509367537Test)
}

func _1644509367537Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		err := tx.Migrator().AutoMigrate(
			new(models.TbNotice),
		)
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
