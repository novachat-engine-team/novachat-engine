/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getOnlines_handler.go
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

//  messages.getOnlines#6e2be050 peer:InputPeer = ChatOnlines;
//
func (s *MessagesServiceImpl) MessagesGetOnlines(ctx context.Context, request *mtproto.TLMessagesGetOnlines) (*mtproto.ChatOnlines, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesGetOnlines %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    return mtproto.NewTLChatOnlines(&mtproto.ChatOnlines{
        Onlines: 0,
    }).To_ChatOnlines(), nil
}
