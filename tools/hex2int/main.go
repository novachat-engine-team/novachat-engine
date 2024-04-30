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
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("input hex number\n")
		return
	}

	for _, v := range os.Args[1:] {
		if strings.Index(v, "0x") == 0 {
			v = v[2:]
		}

		res, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			fmt.Printf("%s, error: %s\n", v, err.Error())
		} else {
			fmt.Printf("%s, result:%d\n", v, int32(res))
		}
	}
}
