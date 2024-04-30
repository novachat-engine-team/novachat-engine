/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.searchGifs_handler.go
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

//  messages.searchGifs#bf9a776b q:string offset:int = messages.FoundGifs;
//
func (s *MessagesServiceImpl) MessagesSearchGifs(ctx context.Context, request *mtproto.TLMessagesSearchGifs) (*mtproto.Messages_FoundGifs, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSearchGifs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesSearchGifs logic

	return nil, fmt.Errorf("%s", "Not impl MessagesSearchGifs")
}
