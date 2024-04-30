/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getRecentLocations_handler.go
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

//  messages.getRecentLocations#bbc45b09 peer:InputPeer limit:int hash:int = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetRecentLocations(ctx context.Context, request *mtproto.TLMessagesGetRecentLocations) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetRecentLocations %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetRecentLocations logic
	return mtproto.NewTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Users:    []*mtproto.User{},
		Chats:    []*mtproto.Chat{},
	}).To_Messages_Messages(), nil
}
