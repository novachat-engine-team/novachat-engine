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
	"novachat_engine/pkg/cmd/sync"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "sync"
	command.RootCmd.Short = "sync *.yaml"
	command.RootCmd.Long = "sync *.yaml"
}

func main() {
	app := sync.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
