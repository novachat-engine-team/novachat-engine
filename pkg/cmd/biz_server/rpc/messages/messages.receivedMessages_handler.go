/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.receivedMessages_handler.go
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

//  messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;
//
func (s *MessagesServiceImpl) MessagesReceivedMessages(ctx context.Context, request *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector_ReceivedNotifyMessage, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReceivedMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesReceivedMessages logic

	return &mtproto.Vector_ReceivedNotifyMessage{
		Datas: []*mtproto.ReceivedNotifyMessage{},
	}, nil
}
