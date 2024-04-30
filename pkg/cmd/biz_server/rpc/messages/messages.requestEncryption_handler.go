/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.requestEncryption_handler.go
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

//  messages.requestEncryption#f64daf43 user_id:InputUser random_id:int g_a:bytes = EncryptedChat;
//
func (s *MessagesServiceImpl) MessagesRequestEncryption(ctx context.Context, request *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesRequestEncryption %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl MessagesRequestEncryption logic

    return nil, fmt.Errorf("%s", "Not impl MessagesRequestEncryption")
}
