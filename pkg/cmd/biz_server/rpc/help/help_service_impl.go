/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"github.com/howeyc/fsnotify"
	"novachat_engine/pkg/cmd/biz_server/conf"
	"novachat_engine/service/app/appconfig"
	"novachat_engine/service/app/country"
	"novachat_engine/service/core/account/account"
)

type HelpServiceImpl struct {
	conf         *conf.Config
	watcher      *fsnotify.Watcher
	accountCache *account.Core
}

func NewHelpServiceImpl(conf *conf.Config) *HelpServiceImpl {
	impl := &HelpServiceImpl{
		conf: conf,
	}

	country.InstallCountryList(conf.CountryFile)
	appconfig.InstallAppConfig(conf.AppConfig)
	return impl
}
