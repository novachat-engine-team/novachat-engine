/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.readFeaturedStickers_handler.go
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

//  messages.readFeaturedStickers#5b118126 id:Vector<long> = Bool;
//
func (s *MessagesServiceImpl) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesReadFeaturedStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl MessagesReadFeaturedStickers logic

    return mtproto.ToMTBool(true), nil
}
