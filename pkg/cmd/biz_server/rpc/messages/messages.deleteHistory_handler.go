/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.deleteHistory_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/input"
	"time"
)

//  messages.deleteHistory#1c015b09 flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int = messages.AffectedHistory;
//
func (s *MessagesServiceImpl) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesDeleteHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	inputPeer := input.MakeInputPeer(request.Peer)
	affectHistory, err := s.accountMessageCore.DeleteHistory(md.UserId, md.AuthKeyId, request.JustClear, request.Revoke, inputPeer, request.MaxId, int32(time.Now().Unix()))
	if err != nil {
		log.Errorf("MessagesDeleteHistory %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	return affectHistory, nil
}
