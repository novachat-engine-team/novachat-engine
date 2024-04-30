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
	"novachat_engine/pkg/cmd/msg"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "msg"
	command.RootCmd.Short = "msg *.yaml"
	command.RootCmd.Long = "msg *.yaml"
}

func main() {
	app := msg.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
