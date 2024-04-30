/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.markDialogUnread_handler.go
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

//  messages.markDialogUnread#c286d98f flags:# unread:flags.0?true peer:InputDialogPeer = Bool;
//
func (s *MessagesServiceImpl) MessagesMarkDialogUnread(ctx context.Context, request *mtproto.TLMessagesMarkDialogUnread) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesMarkDialogUnread %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesMarkDialogUnread logic

	return mtproto.ToMTBool(true), nil
}
