package template

import (
	"bytes"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/errors"
	"text/template"
)

var Yml = `settings:
  application:
    enabledp: true
    host: 0.0.0.0
    mode: dev
    name: testApp
    port: 8888
    readtimeout: 10000
    writertimeout: 20000
  database:
    driver: mysql
    source: {{.User}}:{{.Password}}@tcp({{.Host}}:{{.Port}})/{{.DatabaseName}}?charset=utf8mb4&parseTime=True&loc=Local&timeout=100000ms
  gen:
    dbname: {{.DatabaseName}}
    frontpath: ../go-admin-ui/src
  jwt:
    secret: go-admin
    timeout: 3600
  logger:
    # 日志存放路径
    path: temp/logs
    # 日志输出，file：文件，default：命令行，其他：命令行
    stdout: '' #控制台日志，启用后，不输出到文件
    # 日志等级, trace, debug, info, warn, error, fatal
    level: trace
    # 数据库日志开关
    enableddb: true
  queue:
    memory:
      poolSize: 100
  extend:
    amap:
      key: `

func GenCodeForString(t interface{}, temp string, path string, fileName string, codeName string) error {
	_ = pkg.PathCreate(path)
	filePath := path + fileName
	t4, err := template.New(codeName).Parse(temp)
	if err != nil {
		err = errors.WithStack(err)
		err = fmt.Errorf("读取%s模版失败！错误详情：%v", codeName, err)
		return err
	}
	var b4 bytes.Buffer
	err = t4.Execute(&b4, t)
	if err != nil {
		err = errors.WithStack(err)
		err = fmt.Errorf("解析%s模版失败！错误详情：%v", codeName, err)
		return err
	}
	pkg.FileCreate(b4, filePath)
	return nil
}
