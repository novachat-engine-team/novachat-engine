/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.deleteScheduledMessages_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"time"
)

//  messages.deleteScheduledMessages#59ae2b16 peer:InputPeer id:Vector<int> = Updates;
//
func (s *MessagesServiceImpl) MessagesDeleteScheduledMessages(ctx context.Context, request *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesDeleteScheduledMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesDeleteScheduledMessages logic

	return mtproto.NewTLUpdates(&mtproto.Updates{
		Date: int32(time.Now().Unix()),
		Seq:  0,
	}).To_Updates(), nil
}
