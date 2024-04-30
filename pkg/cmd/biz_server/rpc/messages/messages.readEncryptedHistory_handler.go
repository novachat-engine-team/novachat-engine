/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.readEncryptedHistory_handler.go
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

//  messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;
//
func (s *MessagesServiceImpl) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReadEncryptedHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesReadEncryptedHistory logic

	return nil, fmt.Errorf("%s", "Not impl MessagesReadEncryptedHistory")
}
