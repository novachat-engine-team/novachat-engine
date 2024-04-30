/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : messages.sendBroadcast_handler.go
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

//  messages.sendBroadcast#bf73f4da contacts:Vector<InputUser> random_id:Vector<long> message:string media:InputMedia = Updates;
//
func (s *MessagesServiceImpl) MessagesSendBroadcast(ctx context.Context, request *mtproto.TLMessagesSendBroadcast) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSendBroadcast %v, request: %v", md, request)

	// Impl MessagesSendBroadcast logic

	return nil, fmt.Errorf("%s", "Not impl MessagesSendBroadcast")
}
