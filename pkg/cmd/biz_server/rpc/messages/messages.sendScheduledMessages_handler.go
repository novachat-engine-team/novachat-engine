/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.sendScheduledMessages_handler.go
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
	"novachat_engine/service/common/errors"
)

//  messages.sendScheduledMessages#bd38850a peer:InputPeer id:Vector<int> = Updates;
//
func (s *MessagesServiceImpl) MessagesSendScheduledMessages(ctx context.Context, request *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSendScheduledMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesSendScheduledMessages logic
	return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_DIRTY_WORD)

	return nil, fmt.Errorf("%s", "Not impl MessagesSendScheduledMessages")
}
