/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.acceptUrlAuth_handler.go
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

//  messages.acceptUrlAuth#f729ea98 flags:# write_allowed:flags.0?true peer:InputPeer msg_id:int button_id:int = UrlAuthResult;
//
func (s *MessagesServiceImpl) MessagesAcceptUrlAuth(ctx context.Context, request *mtproto.TLMessagesAcceptUrlAuth) (*mtproto.UrlAuthResult, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesAcceptUrlAuth %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesAcceptUrlAuth logic

	return nil, fmt.Errorf("%s", "Not impl MessagesAcceptUrlAuth")
}
