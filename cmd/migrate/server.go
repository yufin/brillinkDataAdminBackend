package migrate

import (
	"bytes"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"strconv"
	"text/template"
	"time"

	"github.com/go-admin-team/go-admin-core/config/source/file"
	"github.com/spf13/cobra"

	"github.com/go-admin-team/go-admin-core/sdk/config"
	"go-admin/cmd/migrate/migration"
	_ "go-admin/cmd/migrate/migration/version"
	_ "go-admin/cmd/migrate/migration/version-local"
	"go-admin/common/database"
	"go-admin/common/models"
)

var (
	configYml  string
	generate   bool
	goAdmin    bool
	host       string
	SystemName string
	Username   string
	Password   string
	StartCmd   = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "go-admin migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
)

// fixme 在您看不见代码的时候运行迁移，我觉得是不安全的，所以编译后最好不要去执行迁移
func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
	StartCmd.PersistentFlags().BoolVarP(&goAdmin, "goAdmin", "a", false, "generate go-admin migration file")
	StartCmd.PersistentFlags().StringVarP(&host, "domain", "d", "*", "select tenant host")
	StartCmd.PersistentFlags().StringVarP(&SystemName, "systemName", "s", "go-admin管理系统", "管理系统名称")
	StartCmd.PersistentFlags().StringVarP(&Username, "username", "u", "admin", "系统超级管理员登录用户名")
	StartCmd.PersistentFlags().StringVarP(&Password, "password", "p", "123456", "系统超级管理员登录用户密码")
}

func Run() {

	if !generate {
		logger.Info(`start init`)
		//1. 读取配置
		config.Setup(
			file.NewSource(file.WithPath(configYml)),
			initDB,
		)
	} else {
		logger.Info(`generate migration file`)
		_ = genFile()
	}
}

func migrateModel() error {
	if host == "" {
		host = "*"
	}
	db := sdk.Runtime.GetDbByKey(host)
	if config.DatabasesConfig[host].Driver == "mysql" {
		//初始化数据库时候用
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	err := db.AutoMigrate(&models.Migration{})
	if err != nil {
		logger.Error("数据库迁移失败", err)
		return err
	}
	migration.Migrate.SetDb(db)
	migration.Migrate.Migrate()
	return err
}
func initDB() {
	//3. 初始化数据库链接
	database.Setup()
	//4. 数据库迁移
	logger.Info("数据库迁移开始")
	_ = migrateModel()
	logger.Info(`数据库基础数据初始化成功`)
}

func genFile() error {
	t1, err := template.ParseFiles("template/migrate.tpl")
	if err != nil {
		logger.Error("parse template error", err)
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version_local"
	if goAdmin {
		m["Package"] = "version"
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if goAdmin {
		pkg.FileCreate(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
	} else {
		pkg.FileCreate(b1, "./cmd/migrate/migration/version-local/"+m["GenerateTime"]+"_migrate.go")
	}
	return nil
}
