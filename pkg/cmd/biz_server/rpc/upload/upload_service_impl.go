/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cmd/biz_server/conf"
	"novachat_engine/service/upload"
)

type UploadServiceImpl struct {
	conf       *conf.Config
	uploadCore *upload.Core
}

func NewUploadServiceImpl(conf *conf.Config) *UploadServiceImpl {
	impl := &UploadServiceImpl{
		conf: conf,
	}
	impl.uploadCore = upload.NewUploadCore(&conf.SfsClient)

	return impl
}
