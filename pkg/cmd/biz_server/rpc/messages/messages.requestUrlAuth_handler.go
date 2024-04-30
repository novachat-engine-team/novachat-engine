/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.requestUrlAuth_handler.go
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

//  messages.requestUrlAuth#e33f5613 peer:InputPeer msg_id:int button_id:int = UrlAuthResult;
//
func (s *MessagesServiceImpl) MessagesRequestUrlAuth(ctx context.Context, request *mtproto.TLMessagesRequestUrlAuth) (*mtproto.UrlAuthResult, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesRequestUrlAuth %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesRequestUrlAuth logic

	return nil, fmt.Errorf("%s", "Not impl MessagesRequestUrlAuth")
}
