/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.startBot_handler.go
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

//  messages.startBot#e6df7378 bot:InputUser peer:InputPeer random_id:long start_param:string = Updates;
//
func (s *MessagesServiceImpl) MessagesStartBot(ctx context.Context, request *mtproto.TLMessagesStartBot) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesStartBot %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesStartBot logic

	return nil, fmt.Errorf("%s", "Not impl MessagesStartBot")
}
