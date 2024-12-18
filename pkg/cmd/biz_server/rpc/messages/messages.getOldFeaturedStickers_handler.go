/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getOldFeaturedStickers_handler.go
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

//  messages.getOldFeaturedStickers#5fe7025b offset:int limit:int hash:int = messages.FeaturedStickers;
//
func (s *MessagesServiceImpl) MessagesGetOldFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetOldFeaturedStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLMessagesFeaturedStickersNotModified(nil).To_Messages_FeaturedStickers(), nil
}
