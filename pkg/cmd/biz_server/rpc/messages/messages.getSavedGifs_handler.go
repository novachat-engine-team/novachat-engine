/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getSavedGifs_handler.go
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

//  messages.getSavedGifs#83bf3d52 hash:int = messages.SavedGifs;
//
func (s *MessagesServiceImpl) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetSavedGifs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetSavedGifs logic
	return mtproto.NewTLMessagesSavedGifsNotModified(&mtproto.Messages_SavedGifs{}).To_Messages_SavedGifs(), nil
}
