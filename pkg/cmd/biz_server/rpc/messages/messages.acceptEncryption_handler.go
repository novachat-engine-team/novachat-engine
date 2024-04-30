/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.acceptEncryption_handler.go
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

//  messages.acceptEncryption#3dbc0415 peer:InputEncryptedChat g_b:bytes key_fingerprint:long = EncryptedChat;
//
func (s *MessagesServiceImpl) MessagesAcceptEncryption(ctx context.Context, request *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesAcceptEncryption %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl MessagesAcceptEncryption logic

    return nil, fmt.Errorf("%s", "Not impl MessagesAcceptEncryption")
}
