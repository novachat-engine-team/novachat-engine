/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getUnreadMentions_handler.go
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

//  messages.getUnreadMentions#46578472 peer:InputPeer offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetUnreadMentions(ctx context.Context, request *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetUnreadMentions %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetUnreadMentions logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetUnreadMentions")
}
