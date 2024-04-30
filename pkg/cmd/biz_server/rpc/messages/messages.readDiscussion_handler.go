/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.readDiscussion_handler.go
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

//  messages.readDiscussion#f731a9f4 peer:InputPeer msg_id:int read_max_id:int = Bool;
//
func (s *MessagesServiceImpl) MessagesReadDiscussion(ctx context.Context, request *mtproto.TLMessagesReadDiscussion) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReadDiscussion %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesReadDiscussion logic

	return nil, fmt.Errorf("%s", "Not impl MessagesReadDiscussion")
}
