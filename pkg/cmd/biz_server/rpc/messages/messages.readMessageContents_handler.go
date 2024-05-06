/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.readMessageContents_handler.go
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

//  messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
//
func (s *MessagesServiceImpl) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesReadMessageContents %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Id) == 0 {
		return mtproto.NewTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      -1,
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	resp, err := s.accountMessageCore.ReadMessageContents(
		md.UserId,
		request.GetId())
	if err != nil {
		log.Errorf("MessagesReadMessageContents request: %v error:%s", request, err.Error())
		return nil, err
	}
	return resp, nil
}
