/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getDiscussionMessage_handler.go
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

//  messages.getDiscussionMessage#446972fd peer:InputPeer msg_id:int = messages.DiscussionMessage;
//
func (s *MessagesServiceImpl) MessagesGetDiscussionMessage(ctx context.Context, request *mtproto.TLMessagesGetDiscussionMessage) (*mtproto.Messages_DiscussionMessage, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetDiscussionMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetDiscussionMessage logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetDiscussionMessage")
}
