/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.deleteMessages_handler.go
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

//  messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
//
func (s *MessagesServiceImpl) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesDeleteMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	affectedMessages, err := s.accountMessageCore.DeleteMessages(md.UserId, md.AuthKeyId, request.Revoke, request.Id)
	if err != nil {
		log.Errorf("MessagesDeleteMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	log.Debugf("MessagesDeleteMessages %v, request: %v count:%d reply ok!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request, affectedMessages.GetPtsCount())
	return affectedMessages, nil
}
