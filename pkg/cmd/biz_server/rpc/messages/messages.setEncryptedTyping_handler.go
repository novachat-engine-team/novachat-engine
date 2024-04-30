/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.setEncryptedTyping_handler.go
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

//  messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;
//
func (s *MessagesServiceImpl) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSetEncryptedTyping %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesSetEncryptedTyping logic

	return nil, fmt.Errorf("%s", "Not impl MessagesSetEncryptedTyping")
}
