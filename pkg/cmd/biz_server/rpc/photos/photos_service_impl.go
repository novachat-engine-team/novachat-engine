/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : photos_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cmd/biz_server/conf"
	"novachat_engine/service/core/account/users/sql_core"
)

type PhotosServiceImpl struct {
	conf     *conf.Config
	userCore *sql_core.Core
}

func NewPhotosServiceImpl(conf *conf.Config) *PhotosServiceImpl {
	impl := &PhotosServiceImpl{
		conf:     conf,
		userCore: sql_core.NewUsersCore(&conf.MySQL),
	}

	return impl
}
