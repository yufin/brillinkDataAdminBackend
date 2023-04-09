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

// cmd note
// -a true # 同步接口信息
