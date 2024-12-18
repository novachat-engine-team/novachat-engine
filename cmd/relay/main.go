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
	"novachat_engine/pkg/cmd/relay"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "relay"
	command.RootCmd.Short = "relay *.yaml"
	command.RootCmd.Long = "relay *.yaml"
}

func main() {
	app := relay.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
