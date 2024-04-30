/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getEmojiKeywordsLanguages_handler.go
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

//  messages.getEmojiKeywordsLanguages#4e9963b2 lang_codes:Vector<string> = Vector<EmojiLanguage>;
//
func (s *MessagesServiceImpl) MessagesGetEmojiKeywordsLanguages(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetEmojiKeywordsLanguages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetEmojiKeywordsLanguages logic

	return &mtproto.Vector_EmojiLanguage{
		Datas: []*mtproto.EmojiLanguage{},
	}, nil
}
