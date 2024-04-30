/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getEmojiKeywords_handler.go
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

//  messages.getEmojiKeywords#35a0e062 lang_code:string = EmojiKeywordsDifference;
//
func (s *MessagesServiceImpl) MessagesGetEmojiKeywords(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetEmojiKeywords %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetEmojiKeywords logic

	return mtproto.NewTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
		LangCode:    "",
		FromVersion: 0,
		Version:     0,
		Keywords:    []*mtproto.EmojiKeyword{},
	}).To_EmojiKeywordsDifference(), nil
}
