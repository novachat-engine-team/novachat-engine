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
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/banned_right"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	"time"
)

func (impl *Impl) ReqAddUser(ctx context.Context, request *chatService.AddUser) (*types.Any, error) {
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

	participant := chat.GetChatInfo().GetParticipants(request.PeerId)
	if participant != nil {
		if participant.State == data_chat.ParticipantState_ParticipantStateBan {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_NOT_MUTUAL_CONTACT)
		} else {
			if chat.GetChatInfo().Count >= chat.GetChatMax() {
				return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERS_TOO_MUCH)
			}
			addUserList, modifyUserList, err := chat.AddUsers(request.UserId, []int64{request.PeerId})
			if err != nil {
				log.Errorf("ReqAddUser AddUsers error:%s", err.Error())
				return nil, err
			}
			if len(addUserList) == 0 && len(modifyUserList) == 0 {
				return types.MarshalAny(mtproto.NewTLUpdates(nil).To_Updates())
			}
		}
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromUserIDType(request.ChatId).ToInt32(),
			}).To_Update(),
		},
	}).To_Updates()

	updatesValue, err := msgService.GetMsgClient().ReqSendMessages(ctx, &msgService.SendMessages{
		AuthKeyId:  request.AuthKeyId,
		FromUserId: request.PeerId,
		PeerId:     request.ChatId,
		PeerType:   constants.PeerTypeChannel.ToInt32(),
		DataList: []*msgService.SendMessageData{{
			RandomId: time.Now().UnixMicro(),
			Message: mtproto.NewTLMessageService(&mtproto.Message{
				Out:               true,
				FromId90DDDC1171:  constants.PeerTypeFromUserIDType(request.PeerId).ToInt32(),
				FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(request.PeerId).ToInt32()}).To_Peer(),
				PeerId:            mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.ChatId).ToInt32()}).To_Peer(),
				ToId:              mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.ChatId).ToInt32()}).To_Peer(),
				Action:            mtproto.NewTLMessageActionChatAddUser(&mtproto.MessageAction{Users: []int32{constants.PeerTypeFromUserIDType(request.PeerId).ToInt32()}}).To_MessageAction(),
				Date:              request.Date,
			}).To_Message(),
			ReplyToMessageId: 0,
		}},
		ClearDraft: false,
	})
	if err != nil {
		log.Errorf("ReqAddUser AddUsers SendMessage error:%s", err.Error())
		return nil, err
	}
	updates.Updates = append(updates.Updates, updatesValue.Updates...)
	return types.MarshalAny(updates)
}

func (impl *Impl) ReqDeleteUser(ctx context.Context, request *chatService.DeleteUser) (*types.Any, error) {
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

	chatData := chat.GetChatInfo().ChatData

	if chatData.Creator != request.UserId {
		participant := chat.GetChatInfo().GetParticipants(request.UserId)
		if participant == nil {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
		}
		if !participant.Admin {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}

		if !banned_right.RightsToChatAdminBannedRight(participant.AdminRights).AddAdmins {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	modifyUserList, err := chat.ModifyParticipant(request.UserId, []int64{request.PeerId}, data_chat.ParticipantState_ParticipantStateDelete, 0)
	if err != nil {
		log.Errorf("ReqDeleteUser DeleteUsers error:%s", err.Error())
		return nil, err
	}

	if len(modifyUserList) == 0 {
		return types.MarshalAny(mtproto.NewTLUpdates(nil).To_Updates())
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{ChannelId: constants.PeerTypeFromChatIDType(request.ChatId).ToInt32()}).To_Update(),
			mtproto.NewTLUpdateChatParticipantDelete(&mtproto.Update{
				ChatId:  constants.PeerTypeFromChatIDType(request.ChatId).ToInt32(),
				UserId:  constants.PeerTypeFromUserIDType(request.PeerId).ToInt32(),
				Version: 0,
			}).To_Update(),
		},
	}).To_Updates()
	updateDataList := make([]*syncService.UpdateData, 0, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		if request.UserId == participant.UserId {
			updateDataList = append(updateDataList, &syncService.UpdateData{
				UserId:          participant.UserId,
				IgnoreAuthKeyId: request.AuthKeyId,
				Updates:         updates,
			})
		} else {
			updateDataList = append(updateDataList, &syncService.UpdateData{
				UserId:  participant.UserId,
				Updates: updates,
			})
		}
	})

	_, err = syncService.GetSyncClientById(request.UserId).ReqSyncUpdates(ctx, &syncService.SyncUpdates{
		UpdateDataList: updateDataList,
		Push:           false,
		PeerType:       constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("ReqDeleteUser ReqSyncUpdates:%+v", request, err.Error())
	}

	return types.MarshalAny(updates)
}
