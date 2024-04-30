/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.reportEncryptedSpam_handler.go
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

//  messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;
//
func (s *MessagesServiceImpl) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReportEncryptedSpam %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesReportEncryptedSpam logic

	return mtproto.ToMTBool(true), nil
}
