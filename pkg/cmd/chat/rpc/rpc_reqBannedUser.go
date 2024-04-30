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
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	"time"
)

func (impl *Impl) ReqBannedUser(ctx context.Context, request *chatService.BannedUser) (*types.Any, error) {
	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		log.Errorf("ReqBannedUser request:%+v", request, err.Error())
		return nil, err
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	chatInfo := chat.GetChatInfo()
	if chatInfo.ChatData.Creator != request.UserId && !request.IsLeft {
		participant := chat.GetChatInfo().GetParticipants(request.UserId)
		if participant == nil {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
		}
		if !participant.Admin {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	var rights int32
	var reason data_chat.ParticipantState
	if request.IsLeft {
		reason = data_chat.ParticipantState_ParticipantStateLeft
	} else if request.IsDeleted {
		reason = data_chat.ParticipantState_ParticipantStateDelete
	} else if request.IsBanned {
		reason = data_chat.ParticipantState_ParticipantStateBan
		rights = request.Rights
	}
	_, err = chat.ModifyParticipant(request.UserId, []int64{request.PeerId}, reason, rights)
	if err != nil {
		log.Errorf("ReqBannedUser ModifyParticipant error:%s", err.Error())
		return nil, err
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromUserIDType(request.ChatId).ToInt32(),
			}).To_Update(),
		},
	}).To_Updates()
	if request.IsDeleted {
		dataList := make([]*msgService.SendMessageData, 0, 1)
		dataList = append(dataList, &msgService.SendMessageData{
			RandomId: time.Now().UnixMicro(),
			Message: mtproto.NewTLMessageService(&mtproto.Message{
				Out:               true,
				FromId90DDDC1171:  constants.PeerTypeFromUserIDType(request.UserId).ToInt32(),
				FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(request.UserId).ToInt32()}).To_Peer(),
				PeerId:            mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.ChatId).ToInt32()}).To_Peer(),
				ToId:              mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.ChatId).ToInt32()}).To_Peer(),
				Action:            mtproto.NewTLMessageActionChatDeleteUser(&mtproto.MessageAction{Users: []int32{constants.PeerTypeFromUserIDType(request.PeerId).ToInt32()}}).To_MessageAction(),
				Date:              request.Date,
			}).To_Message(),
			ReplyToMessageId: 0,
		})

		updatesValue, err1 := msgService.GetMsgClient().ReqSendMessages(ctx, &msgService.SendMessages{
			AuthKeyId:  request.AuthKeyId,
			FromUserId: request.UserId,
			PeerId:     request.ChatId,
			PeerType:   constants.PeerTypeChannel.ToInt32(),
			DataList:   dataList,
			ClearDraft: false,
		})
		if err1 != nil {
			log.Errorf("ReqBannedUser ModifyParticipant SendMessage error:%s", err1.Error())
			return nil, err1
		}
		updates.Updates = append(updates.Updates, updatesValue.Updates...)
	}

	syncUpdates := &syncService.SyncUpdates{
		UpdateDataList: make([]*syncService.UpdateData, 0, chatInfo.Count),
		Push:           false,
		PeerType:       constants.PeerTypeUser.ToInt32(),
	}
	chatInfo.Iteration(func(participant *data_chat.ChatParticipant) {
		if participant.UserId == request.UserId {
			syncUpdates.UpdateDataList = append(syncUpdates.UpdateDataList, &syncService.UpdateData{
				UserId:          participant.UserId,
				IgnoreAuthKeyId: request.AuthKeyId,
				Updates:         updates,
			})
		} else {
			syncUpdates.UpdateDataList = append(syncUpdates.UpdateDataList, &syncService.UpdateData{
				UserId:  participant.UserId,
				Updates: updates,
			})
		}
	})

	_, err = syncService.GetSyncClientById(request.UserId).ReqSyncUpdates(ctx, syncUpdates)
	if err != nil {
		log.Warnf("ReqBannedRights ReqSyncUpdates request:%+v", request, err.Error())
	}
	return types.MarshalAny(updates)
}
