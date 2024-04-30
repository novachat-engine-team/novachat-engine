/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getDocumentByHash_handler.go
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

//  messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
//
func (s *MessagesServiceImpl) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetDocumentByHash %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetDocumentByHash logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetDocumentByHash")
}
