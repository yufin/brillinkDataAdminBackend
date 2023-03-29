package app

import (
	"bytes"
	log "github.com/go-admin-team/go-admin-core/logger"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"text/template"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:     "app",
		Short:   "Create a new app",
		Long:    "Use when you need to create a new app",
		Example: "go-admin app -n admin",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "", "Start server with provided configuration file")
}

func run() {
	log.Info(`开始执行...`)
	//1. 读取配置
	log.Info(`生成文件中...`)
	_ = genFile()
	log.Info(`执行完毕！`)
}

func genFile() error {
	if appName == "" {
		return errors.New("arg `name` invalid ：name is empty")
	}
	path := "app/"
	appPath := path + appName
	err := utils.IsNotExistMkDir(appPath)
	if err != nil {
		return err
	}
	apiPath := appPath + "/apis/"
	err = utils.IsNotExistMkDir(apiPath)
	if err != nil {
		return err
	}
	modelsPath := appPath + "/models/"
	err = utils.IsNotExistMkDir(modelsPath)
	if err != nil {
		return err
	}
	routerPath := appPath + "/router/"
	err = utils.IsNotExistMkDir(routerPath)
	if err != nil {
		return err
	}
	servicePath := appPath + "/service/"
	err = utils.IsNotExistMkDir(servicePath)
	if err != nil {
		return err
	}
	dtoPath := appPath + "/service/dto/"
	err = utils.IsNotExistMkDir(dtoPath)
	if err != nil {
		log.Error(err)
		return err
	}

	t1, err := template.ParseFiles("template/cmd_api.tpl")
	if err != nil {
		log.Error(err)
		return err
	}
	m := map[string]string{}
	m["appName"] = appName
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	pkg.FileCreate(b1, "./cmd/server/"+appName+".go")
	t2, err := template.ParseFiles("template/router.tpl")
	var b2 bytes.Buffer
	err = t2.Execute(&b2, nil)
	pkg.FileCreate(b2, appPath+"/router/router.go")
	return nil
}
