/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.toggleChatAdmins_handler.go
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

//  messages.toggleChatAdmins#ec8bd9e1 chat_id:int enabled:Bool = Updates;
//
func (s *MessagesServiceImpl) MessagesToggleChatAdmins(ctx context.Context, request *mtproto.TLMessagesToggleChatAdmins) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesToggleChatAdmins %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesToggleChatAdmins logic

	return nil, fmt.Errorf("%s", "Not impl MessagesToggleChatAdmins")
}
