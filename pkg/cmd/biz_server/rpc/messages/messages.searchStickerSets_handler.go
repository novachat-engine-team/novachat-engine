/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.searchStickerSets_handler.go
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

//  messages.searchStickerSets#c2b7d08b flags:# exclude_featured:flags.0?true q:string hash:int = messages.FoundStickerSets;
//
func (s *MessagesServiceImpl) MessagesSearchStickerSets(ctx context.Context, request *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSearchStickerSets %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return mtproto.NewTLMessagesFoundStickerSetsNotModified(nil).To_Messages_FoundStickerSets(), nil
}
