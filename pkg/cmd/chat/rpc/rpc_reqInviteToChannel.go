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
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	"time"
)

func (impl *Impl) ReqInviteToChannel(ctx context.Context, request *chatService.InviteToChannel) (*types.Any, error) {
	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	inRemovedList := false
	chat.GetChatInfo().IterationCondition(func(participant *data_chat.ChatParticipant) bool {
		if idx := util.Index(request.PeerIdList, func(idx int) bool {
			return request.PeerIdList[idx] == participant.UserId
		}); idx >= 0 {
			if participant.State != data_chat.ParticipantState_ParticipantStateNormal {
				inRemovedList = true
				return false
			}
			copy(request.PeerIdList[idx:], request.PeerIdList[:idx+1])
		}
		return true
	})

	if inRemovedList {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_NOT_MUTUAL_CONTACT)
	}

	if len(request.PeerIdList) == 0 {
		return types.MarshalAny(mtproto.NewTLUpdates(nil).To_Updates())
	}

	if chat.GetChatInfo().Count >= chat.GetChatMax() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERS_TOO_MUCH)
	}

	addUserList, modifyUserList, err := chat.AddUsers(request.UserId, request.PeerIdList)
	if err != nil {
		log.Errorf("ReqAddUser AddUsers error:%s", err.Error())
		return nil, err
	}

	if len(addUserList) == 0 && len(modifyUserList) == 0 {
		return types.MarshalAny(mtproto.NewTLUpdates(nil).To_Updates())
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromUserIDType(request.ChatId).ToInt32(),
			}).To_Update(),
		},
	}).To_Updates()

	dataList := make([]*msgService.SendMessageData, 0, len(request.PeerIdList))
	for idx, peerId := range request.PeerIdList {
		dataList = append(dataList, &msgService.SendMessageData{
			RandomId: time.Now().UnixMicro() + int64(idx),
			Message: mtproto.NewTLMessageService(&mtproto.Message{
				Out:               true,
				FromId90DDDC1171:  constants.PeerTypeFromUserIDType(request.UserId).ToInt32(),
				FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(request.UserId).ToInt32()}).To_Peer(),
				PeerId:            mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.ChatId).ToInt32()}).To_Peer(),
				ToId:              mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.ChatId).ToInt32()}).To_Peer(),
				Action:            mtproto.NewTLMessageActionChatAddUser(&mtproto.MessageAction{Users: []int32{constants.PeerTypeFromUserIDType(peerId).ToInt32()}}).To_MessageAction(),
				Date:              request.Date,
			}).To_Message(),
			ReplyToMessageId: 0,
		})
	}
	updatesValue, err := msgService.GetMsgClient().ReqSendMessages(ctx, &msgService.SendMessages{
		AuthKeyId:  request.AuthKeyId,
		FromUserId: request.UserId,
		PeerId:     request.ChatId,
		PeerType:   constants.PeerTypeChannel.ToInt32(),
		DataList:   dataList,
		ClearDraft: false,
	})
	if err != nil {
		log.Errorf("ReqAddUser AddUsers SendMessage error:%s", err.Error())
		return nil, err
	}
	updates.Updates = append(updates.Updates, updatesValue.Updates...)
	return types.MarshalAny(updates)
}
