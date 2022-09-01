package install

import (
	"github.com/spf13/cobra"
	"go-admin/common/startup"

	_ "go-admin/cmd/migrate/migration/version"
	_ "go-admin/cmd/migrate/migration/version-local"
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
		Use:     "install",
		Short:   "项目初始化安装；针对配置文件和数据结构初始化；",
		Example: "go run main.go install",
		Run: func(cmd *cobra.Command, args []string) {
			setup()
		},
	}
)

// fixme 在您看不见代码的时候运行迁移，我觉得是不安全的，所以编译后最好不要去执行迁移
func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	// 检查是否是首次登陆使用
	startup.Startup(configYml)
}
