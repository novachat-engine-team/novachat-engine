/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editInlineBotMessage_handler.go
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

//  messages.editInlineBotMessage#83557dba flags:# no_webpage:flags.1?true id:InputBotInlineMessageID message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Bool;
//  messages.editInlineBotMessage#130c2c85 flags:# no_webpage:flags.1?true id:InputBotInlineMessageID message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Bool;
//
func (s *MessagesServiceImpl) MessagesEditInlineBotMessage(ctx context.Context, request *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesEditInlineBotMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl MessagesEditInlineBotMessage logic

    return nil, fmt.Errorf("%s", "Not impl MessagesEditInlineBotMessage")
}
