/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getPollVotes_handler.go
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

//  messages.getPollVotes#b86e380e flags:# peer:InputPeer id:int option:flags.0?bytes offset:flags.1?string limit:int = messages.VotesList;
//
func (s *MessagesServiceImpl) MessagesGetPollVotes(ctx context.Context, request *mtproto.TLMessagesGetPollVotes) (*mtproto.Messages_VotesList, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetPollVotes %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetPollVotes logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetPollVotes")
}
