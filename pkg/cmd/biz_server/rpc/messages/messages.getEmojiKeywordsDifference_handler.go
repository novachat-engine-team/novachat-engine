/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getEmojiKeywordsDifference_handler.go
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

//  messages.getEmojiKeywordsDifference#1508b6af lang_code:string from_version:int = EmojiKeywordsDifference;
//
func (s *MessagesServiceImpl) MessagesGetEmojiKeywordsDifference(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesGetEmojiKeywordsDifference %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl MessagesGetEmojiKeywordsDifference logic

    return mtproto.NewTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
        LangCode:    request.LangCode,
        FromVersion: request.FromVersion,
        Version:     request.FromVersion,
        Keywords:    []*mtproto.EmojiKeyword{},
    }).To_EmojiKeywordsDifference(), nil
}
