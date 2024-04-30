/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.report_handler.go
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

const saveMessageLen = 10

//  messages.report#bd82b658 peer:InputPeer id:Vector<int> reason:ReportReason = Bool;
//
func (s *MessagesServiceImpl) MessagesReport(ctx context.Context, request *mtproto.TLMessagesReport) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReport %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic(fmt.Sprintf("%+v", request))
}
