/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package main

import (
	"fmt"
	"novachat_engine/pkg/cmd/biz_server"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "biz_server"
	command.RootCmd.Short = "biz_server *.yaml"
	command.RootCmd.Long = "biz_server *.yaml"
}

func main() {
	app := biz_server.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
