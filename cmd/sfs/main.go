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
	"novachat_engine/pkg/cmd/sfs"
	"novachat_engine/pkg/command"
	"os"
)

func init() {
	command.RootCmd.Use = "sfs"
	command.RootCmd.Short = "sfs *.yaml"
	command.RootCmd.Long = "sfs *.yaml"
}

func main() {
	app := sfs.Application{}
	if err := command.Run(os.Args[1:], &app); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
