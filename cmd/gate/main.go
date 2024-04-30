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
	"novachat_engine/pkg/cmd/gate"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "gate"
	command.RootCmd.Short = "gate *.yaml"
	command.RootCmd.Long = "gate *.yaml"
}

func main() {
	app := gate.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
