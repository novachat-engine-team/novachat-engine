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
	"novachat_engine/pkg/cmd/session"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "session"
	command.RootCmd.Short = "session *.yaml"
	command.RootCmd.Long = "session *.yaml"
}

func main() {
	app := session.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
