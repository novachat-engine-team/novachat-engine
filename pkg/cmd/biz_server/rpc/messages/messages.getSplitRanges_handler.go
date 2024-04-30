/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getSplitRanges_handler.go
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

//  messages.getSplitRanges#1cff7e08 = Vector<MessageRange>;
//
func (s *MessagesServiceImpl) MessagesGetSplitRanges(ctx context.Context, request *mtproto.TLMessagesGetSplitRanges) (*mtproto.Vector_MessageRange, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetSplitRanges %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetSplitRanges logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetSplitRanges")
}
