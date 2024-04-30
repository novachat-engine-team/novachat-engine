/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.unpinAllMessages_handler.go
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

//  messages.unpinAllMessages#f025bc8b peer:InputPeer = messages.AffectedHistory;
//
func (s *MessagesServiceImpl) MessagesUnpinAllMessages(ctx context.Context, request *mtproto.TLMessagesUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesUnpinAllMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesUnpinAllMessages logic

	return nil, fmt.Errorf("%s", "Not impl MessagesUnpinAllMessages")
}
