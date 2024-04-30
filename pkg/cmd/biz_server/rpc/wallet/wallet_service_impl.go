/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : wallet_service_impl.go
 * @Desc :
 *
 */

package rpc

import "novachat_engine/pkg/cmd/biz_server/conf"

type WalletServiceImpl struct {
	conf *conf.Config
}

func NewWalletServiceImpl(conf *conf.Config) *WalletServiceImpl {
	impl := &WalletServiceImpl{
		conf: conf,
	}

	return impl
}
