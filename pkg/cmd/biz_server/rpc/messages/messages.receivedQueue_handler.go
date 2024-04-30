/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.receivedQueue_handler.go
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

//  messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;
//
func (s *MessagesServiceImpl) MessagesReceivedQueue(ctx context.Context, request *mtproto.TLMessagesReceivedQueue) (*mtproto.VectorLong, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReceivedQueue %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesReceivedQueue logic

	return nil, fmt.Errorf("%s", "Not impl MessagesReceivedQueue")
}
