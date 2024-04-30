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
	"errors"
	"fmt"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
)

func (impl *Impl) ReqSendOutboxesMessages(ctx context.Context, request *chatService.SendOutboxesMessages) (*mtproto.Updates, error) {
	peerType := constants.PeerTypeFromInt32(request.Message.PeerType)
	switch peerType {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("ReqSendMessages peerType:%d error", peerType)
		return nil, fmt.Errorf("ReqSendMessages peerType:%d error", peerType)
	}

	chat, err := impl.chatManager.GetChat(request.Message.PeerId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}
	if !chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
	}

	participantOfMe := chat.GetChatInfo().GetParticipants(request.Message.FromUserId)
	if participantOfMe == nil {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
	}

	//TODO:Coderxw chat rights
	updates, err := msgService.GetMsgClient().ReqSendMessages(context.Background(), request.Message)
	if err != nil {
		log.Errorf("ReqSendMessages userId:%d peerId:%d error:%s", request.Message.FromUserId, request.Message.PeerId, err.Error())
		return nil, err
	}

	return updates, nil
}
