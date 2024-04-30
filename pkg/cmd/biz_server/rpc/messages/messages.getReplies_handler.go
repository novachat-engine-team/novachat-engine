/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getReplies_handler.go
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

//  messages.getReplies#24b581ba peer:InputPeer msg_id:int offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:int = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetReplies(ctx context.Context, request *mtproto.TLMessagesGetReplies) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetReplies %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetReplies logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetReplies")
}
