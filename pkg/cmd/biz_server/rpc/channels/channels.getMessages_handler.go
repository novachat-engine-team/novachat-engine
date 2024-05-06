/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getMessages_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message/message"
)

//  channels.getMessages#ad8c9a23 channel:InputChannel id:Vector<InputMessage> = messages.Messages;
//  channels.getMessages#93d7b347 channel:InputChannel id:Vector<int> = messages.Messages;
//
func (s *ChannelsServiceImpl) ChannelsGetMessages(ctx context.Context, request *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.IdAD8C9A2385) == 0 && len(request.Id93D7B34771) == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MSG_ID_INVALID)
		log.Errorf("ChannelsGetMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var idList []int32
	if len(request.IdAD8C9A2385) > 0 {
		for _, v := range request.IdAD8C9A2385 {
			idList = append(idList, v.Id)
		}
	} else {
		idList = request.Id93D7B34771
	}
	messageList, err := s.accountMessageCore.GetChannelMessageList(&message.ChannelMessageId{
		ChannelId: constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt(),
		IdList:    idList,
	})
	if err != nil {
		log.Errorf("ChannelsGetMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	messagesMessages := mtproto.NewTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: messageList,
		Chats:    nil,
		Users:    nil,
	})
	return messagesMessages.To_Messages_Messages(), nil
}
