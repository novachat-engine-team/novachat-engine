/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getScheduledHistory_handler.go
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

//  messages.getScheduledHistory#e2c2685b peer:InputPeer hash:int = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetScheduledHistory(ctx context.Context, request *mtproto.TLMessagesGetScheduledHistory) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetScheduledHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetScheduledHistory logic
	return mtproto.NewTLMessagesMessagesNotModified(nil).To_Messages_Messages(), nil
}
