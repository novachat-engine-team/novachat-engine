/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getAutoDownloadSettings_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  account.getAutoDownloadSettings#56da0b3f = account.AutoDownloadSettings;
//
func (s *AccountServiceImpl) AccountGetAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountGetAutoDownloadSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//TODO:(Coder)
	log.Infof("AccountGetAutoDownloadSettings - request: %v", request)
	settings := mtproto.NewTLAccountAutoDownloadSettings(nil)

	settings.SetHigh(mtproto.NewTLAutoDownloadSettings(nil).To_AutoDownloadSettings())
	settings.SetLow(mtproto.NewTLAutoDownloadSettings(nil).To_AutoDownloadSettings())
	settings.SetMedium(mtproto.NewTLAutoDownloadSettings(nil).To_AutoDownloadSettings())

	return settings.To_Account_AutoDownloadSettings(), nil
}
