package config

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/logger"

	"github.com/go-admin-team/go-admin-core/config/source/file"
	"github.com/spf13/cobra"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "go-admin config -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func run() {
	config.Setup(file.NewSource(file.WithPath(configYml)))

	application, err := json.MarshalIndent(config.ApplicationConfig, "", "   ") //转换成JSON返回的是byte[]
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("application:", string(application))

	jwt, err := json.MarshalIndent(config.JwtConfig, "", "   ") //转换成JSON返回的是byte[]
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("jwt:", string(jwt))

	// todo 需要兼容
	database, err := json.MarshalIndent(config.DatabasesConfig, "", "   ") //转换成JSON返回的是byte[]
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("database:", string(database))

	gen, errs := json.MarshalIndent(config.GenConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		logger.Error(err)
	}
	fmt.Println("gen:", string(gen))

	loggerConfig, errs := json.MarshalIndent(config.LoggerConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		logger.Error(err)
	}
	fmt.Println("logger:", string(loggerConfig))

}
