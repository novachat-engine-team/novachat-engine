/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getDialogUnreadMarks_handler.go
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

//  messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
//
func (s *MessagesServiceImpl) MessagesGetDialogUnreadMarks(ctx context.Context, request *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetDialogUnreadMarks %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetDialogUnreadMarks logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetDialogUnreadMarks")
}
