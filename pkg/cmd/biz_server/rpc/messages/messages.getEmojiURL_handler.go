/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getEmojiURL_handler.go
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

//  messages.getEmojiURL#d5b10c26 lang_code:string = EmojiURL;
//
func (s *MessagesServiceImpl) MessagesGetEmojiURL(ctx context.Context, request *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetEmojiURL %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetEmojiURL logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetEmojiURL")
}
