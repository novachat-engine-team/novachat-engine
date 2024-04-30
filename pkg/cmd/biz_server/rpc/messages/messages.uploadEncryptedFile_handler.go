/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.uploadEncryptedFile_handler.go
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

//  messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
//
func (s *MessagesServiceImpl) MessagesUploadEncryptedFile(ctx context.Context, request *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesUploadEncryptedFile %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesUploadEncryptedFile logic

	return nil, fmt.Errorf("%s", "Not impl MessagesUploadEncryptedFile")
}
