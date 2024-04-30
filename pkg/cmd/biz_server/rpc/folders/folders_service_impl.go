/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : folders_service_impl.go
 * @Desc :
 *
 */

package rpc

import "novachat_engine/pkg/cmd/biz_server/conf"

type FoldersServiceImpl struct {
	conf *conf.Config
}

func NewFoldersServiceImpl(conf *conf.Config) *FoldersServiceImpl {
	impl := &FoldersServiceImpl{
		conf: conf,
	}

	return impl
}
