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
	"novachat_engine/pkg/cmd/chat"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "chat"
	command.RootCmd.Short = "chat *.yaml"
	command.RootCmd.Long = "chat *.yaml"
}

func main() {
	app := chat.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
