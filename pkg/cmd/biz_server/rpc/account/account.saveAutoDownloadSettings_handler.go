/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.saveAutoDownloadSettings_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  account.saveAutoDownloadSettings#76f36233 flags:# low:flags.0?true high:flags.1?true settings:AutoDownloadSettings = Bool;
//
func (s *AccountServiceImpl) AccountSaveAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSaveAutoDownloadSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSaveAutoDownloadSettings logic

	return nil, fmt.Errorf("%s", "Not impl AccountSaveAutoDownloadSettings")
}
