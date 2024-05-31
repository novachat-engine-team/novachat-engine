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
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
)

func (impl *Impl) ReqReadHistory(ctx context.Context, request *chatService.ChatReadHistory) (*types.Any, error) {
	peerType := constants.PeerTypeFromInt32(request.Value.PeerType)
	switch peerType {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("ReqReadHistory peerType:%d error", peerType)
		return nil, fmt.Errorf("ReqReadHistory peerType:%d error", peerType)
	}

	chat, err := impl.chatManager.GetChat(request.Value.PeerId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}
	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
	}

	participantOfMe := chat.GetChatInfo().GetParticipants(request.Value.UserId)
	if participantOfMe == nil {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
	}

	updates, err := msgService.GetMsgClient().ReqReadHistory(context.Background(), request.Value)
	if err != nil {
		log.Errorf("ReqReadHistory userId:%d peerId:%d error:%s", request.Value.UserId, request.Value.PeerId, err.Error())
		return nil, err
	}

	return updates, nil
}
