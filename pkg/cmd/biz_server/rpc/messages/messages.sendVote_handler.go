/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.sendVote_handler.go
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

//  messages.sendVote#10ea6184 peer:InputPeer msg_id:int options:Vector<bytes> = Updates;
//
func (s *MessagesServiceImpl) MessagesSendVote(ctx context.Context, request *mtproto.TLMessagesSendVote) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSendVote %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesSendVote logic

	return nil, fmt.Errorf("%s", "Not impl MessagesSendVote")
}
