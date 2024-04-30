/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getSearchCounters_handler.go
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

//  messages.getSearchCounters#732eef00 peer:InputPeer filters:Vector<MessagesFilter> = Vector<messages.SearchCounter>;
//
func (s *MessagesServiceImpl) MessagesGetSearchCounters(ctx context.Context, request *mtproto.TLMessagesGetSearchCounters) (*mtproto.VectorMessages_SearchCounter, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetSearchCounters %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//
	//TODO:(Coder) Impl MessagesGetSearchCounters logic
	vectorCounter := &mtproto.VectorMessages_SearchCounter{
		Datas: make([]*mtproto.Messages_SearchCounter, 0, len(request.Filters)),
	}

	for _, r := range request.Filters {
		vectorCounter.Datas = append(vectorCounter.Datas, mtproto.NewTLMessagesSearchCounter(&mtproto.Messages_SearchCounter{
			Inexact: true,
			Filter:  r,
			Count:   0,
		}).To_Messages_SearchCounter())
	}

	log.Debugf("MessagesGetSearchCounters %v, request: %v reply ok!!!!!!!!!!!!!!!!", md, request)
	return vectorCounter, nil
}
