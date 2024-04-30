/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getInlineGameHighScores_handler.go
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

//  messages.getInlineGameHighScores#f635e1b id:InputBotInlineMessageID user_id:InputUser = messages.HighScores;
//
func (s *MessagesServiceImpl) MessagesGetInlineGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetInlineGameHighScores) (*mtproto.Messages_HighScores, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetInlineGameHighScores %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetInlineGameHighScores logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetInlineGameHighScores")
}
