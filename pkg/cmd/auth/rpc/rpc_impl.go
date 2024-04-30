/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"novachat_engine/pkg/cmd/auth/conf"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	accountAuth "novachat_engine/service/account/auth"
	"novachat_engine/service/core/account/auth"
)

type AuthImpl struct {
	conf            *conf.Config
	accountAuthCore *accountAuth.Core
}

func NewAuthImpl(conf *conf.Config) *AuthImpl {

	authClient.InstallAuthClient(conf.AuthClient)

	return &AuthImpl{
		conf:            conf,
		accountAuthCore: accountAuth.NewAuthCore(auth.NewAuthCore(&conf.Mongodb)),
	}
}
