package main

import "go-admin/cmd"

//go:generate swag init --parseDependency --parseDepth=6

// @title GraphAdmin
// @version 1.0.0
// @description 朗链图数据权限管理系统的接口文档
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}

// __________________________________________________________________________

// cmd note
// -a true # 同步接口信息
// go run main.go app -n appname

// CROSS PLATFORM Compile
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o server-win-amd64.exe main.go
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o graph-server-linux-amd64 main.go

// git operation
// git pull upstream
// git merge upstream/main
