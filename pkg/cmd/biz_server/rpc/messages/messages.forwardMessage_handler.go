/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.forwardMessage_handler.go
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

//  messages.forwardMessage#33963bf9 peer:InputPeer id:int random_id:long = Updates;
//
func (s *MessagesServiceImpl) MessagesForwardMessage(ctx context.Context, request *mtproto.TLMessagesForwardMessage) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesForwardMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesForwardMessage logic

	return nil, fmt.Errorf("%s", "Not impl MessagesForwardMessage")
}
