/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.finishTakeoutSession_handler.go
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

//  account.finishTakeoutSession#1d2652ee flags:# success:flags.0?true = Bool;
//
func (s *AccountServiceImpl) AccountFinishTakeoutSession(ctx context.Context, request *mtproto.TLAccountFinishTakeoutSession) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountFinishTakeoutSession %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountFinishTakeoutSession logic

	return nil, fmt.Errorf("%s", "Not impl AccountFinishTakeoutSession")
}
