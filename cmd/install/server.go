package install

import (
	"github.com/spf13/cobra"
	"go-admin/common/startup"

	_ "go-admin/cmd/migrate/migration/version"
	_ "go-admin/cmd/migrate/migration/version-local"
)

var (
	//configYml  string
	StartCmd = &cobra.Command{
		Use:     "install",
		Short:   "项目初始化安装；针对配置文件和数据结构初始化；",
		Example: "go run main.go install",
		Run: func(cmd *cobra.Command, args []string) {
			setup()
		},
	}
)

func init() {
	// TODO: 是否需要待定...
	//StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	// 检查是否是首次登陆使用
	startup.Startup()
}
