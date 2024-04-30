/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/3 21:34
 * @File : syncfs.go
 */

package syncfs

type SyncFS interface {
	PartsComplete(pathName string, partsTotal int32) error
	SyncFile(filePathName string) error
}

func GetDefaultSyncFS() SyncFS {
	return &syncFS{}
}

type syncFS struct {}

func (m *syncFS) PartsComplete(pathName string, partsTotal int32) error {
	return nil
}

func (m *syncFS) SyncFile(filePathName string) error {
	return nil
}
