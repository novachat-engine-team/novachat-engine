/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.sendInlineBotResult_handler.go
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

//  messages.sendInlineBotResult#220815b0 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true hide_via:flags.11?true peer:InputPeer reply_to_msg_id:flags.0?int random_id:long query_id:long id:string schedule_date:flags.10?int = Updates;
//  messages.sendInlineBotResult#b16e06fe flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int random_id:long query_id:long id:string = Updates;
//
func (s *MessagesServiceImpl) MessagesSendInlineBotResult(ctx context.Context, request *mtproto.TLMessagesSendInlineBotResult) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSendInlineBotResult %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesSendInlineBotResult logic

	return nil, fmt.Errorf("%s", "Not impl MessagesSendInlineBotResult")
}
