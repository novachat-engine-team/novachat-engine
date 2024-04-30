/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.setContentSettings_handler.go
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

//  account.setContentSettings#b574b16b flags:# sensitive_enabled:flags.0?true = Bool;
//
func (s *AccountServiceImpl) AccountSetContentSettings(ctx context.Context, request *mtproto.TLAccountSetContentSettings) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSetContentSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSetContentSettings logic

	return nil, fmt.Errorf("%s", "Not impl AccountSetContentSettings")
}
