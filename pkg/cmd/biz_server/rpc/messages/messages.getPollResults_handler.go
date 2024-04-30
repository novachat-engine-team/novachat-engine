/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getPollResults_handler.go
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

//  messages.getPollResults#73bb643b peer:InputPeer msg_id:int = Updates;
//
func (s *MessagesServiceImpl) MessagesGetPollResults(ctx context.Context, request *mtproto.TLMessagesGetPollResults) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetPollResults %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetPollResults logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetPollResults")
}
