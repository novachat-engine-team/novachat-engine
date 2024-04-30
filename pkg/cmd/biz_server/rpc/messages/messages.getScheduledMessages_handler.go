/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getScheduledMessages_handler.go
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

//  messages.getScheduledMessages#bdbb0464 peer:InputPeer id:Vector<int> = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetScheduledMessages(ctx context.Context, request *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetScheduledMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetScheduledMessages logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetScheduledMessages")
}
