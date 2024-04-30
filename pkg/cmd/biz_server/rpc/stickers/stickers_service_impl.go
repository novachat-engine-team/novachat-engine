/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stickers_service_impl.go
 * @Desc :
 *
 */

package rpc

import "novachat_engine/pkg/cmd/biz_server/conf"

type StickersServiceImpl struct {
	conf *conf.Config
}

func NewStickersServiceImpl(conf *conf.Config) *StickersServiceImpl {
	impl := &StickersServiceImpl{
		conf: conf,
	}

	return impl
}
