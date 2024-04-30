/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getGameHighScores_handler.go
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

//  messages.getGameHighScores#e822649d peer:InputPeer id:int user_id:InputUser = messages.HighScores;
//
func (s *MessagesServiceImpl) MessagesGetGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetGameHighScores %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetGameHighScores logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetGameHighScores")
}
