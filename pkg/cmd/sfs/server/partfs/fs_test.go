/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/1 19:24
 * @File : fs_test.go
 */

package partfs

import (
	"fmt"
	"os"
	"testing"
)

func TestInstallPartFS(t *testing.T) {
	//InstallPartFS("1")

	offset := int32(0)
	limit := int32(32768)
	filePathName := "E:\\test\\novachat_engine\\server\\sfs\\sfs_dir\\jpeg/A1/1437656537387503616.dat"
	fileSize, _ := FileSize(filePathName)

	data := make([]byte, 0, 10)
	for idx := 0; idx < 10; idx++ {
		d, err := GetFileData(filePathName, offset, limit)
		if err != nil {
			panic(err)
		}

		data = append(data, d...)
		fmt.Printf("offset:%d limit:%d len()= %d\n", offset, limit, len(d))
		offset+=limit
	}

	t.Logf("len()= %d %d", fileSize, len(data))

	f, err := os.Create(filePathName + "1")
	if err != nil {
		panic(err)
	}
	f.Write(data)
	defer f.Close()
}
