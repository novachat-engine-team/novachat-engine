/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/banned_right"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
)

func (impl *Impl) ReqDeleteHistory(ctx context.Context, request *chatService.DeleteHistory) (*types.Any, error) {
	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		log.Errorf("ReqDeleteHistory request:%+v", request, err.Error())
		return nil, err
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	chatData := chat.GetChatInfo().ChatData
	if chatData.Creator != request.UserId {
		participant := chat.GetChatInfo().GetParticipants(request.UserId)
		if participant == nil {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHANNEL_PRIVATE)
		}
		if !participant.Admin {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}

		if banned_right.RightsToChatAdminBannedRight(participant.AdminRights).DeleteMessages {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_ID_INVALID)
		}
	}

	pts, idList, err := impl.accountMessageCore.DeleteChannelHistory(request.ChatId, request.UserId, request.MaxId, request.Date)
	if err != nil {
		log.Errorf("ReqDeleteHistory DeleteChannelHistory error:%s", err.Error())
		return types.MarshalAny(&mtproto.TLInt32{Value: 0})
	}

	if len(idList) == 0 {
		return types.MarshalAny(&mtproto.TLInt32{Value: 0})
	}

	data, _ := proto.Marshal(&chatService.DeleteMessagesUpdates{
		PeerId: request.ChatId,
		UserId: request.UserId,
		Ids:    idList,
		Date:   request.Date,
		Pts:    pts,
	})
	_, _, err = impl.chatInProducer.SendMessage(constants.DeleteChannelMessages, data)
	if err != nil {
		log.Errorf("ReqDeleteHistory SendMessage error:%s need retry", err.Error())
		return nil, err
	}

	return types.MarshalAny(&mtproto.TLInt32{Value: pts})
}
