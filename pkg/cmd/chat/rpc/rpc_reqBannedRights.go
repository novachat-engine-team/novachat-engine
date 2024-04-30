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
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqBannedRights(ctx context.Context, request *chatService.BannedRights) (*types.Any, error) {
	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		log.Errorf("ReqBannedRights request:%+v", request, err.Error())
		return nil, err
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	chatInfo := chat.GetChatInfo()
	if chatInfo.ChatData.Creator != request.UserId {
		participant := chat.GetChatInfo().GetParticipants(request.UserId)
		if participant == nil {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
		}
		if !participant.Admin {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	_, err = chat.EditBannedRights(request.Rights)
	if err != nil {
		log.Errorf("ReqBannedRights EditBannedRights request:%+v", request, err.Error())
		return nil, err
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromChatIDType(request.ChatId).ToInt32(),
			}).To_Update(),
		},
	})
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
				Updates:         updates.To_Updates(),
			})
		} else {
			syncUpdates.UpdateDataList = append(syncUpdates.UpdateDataList, &syncService.UpdateData{
				UserId:  participant.UserId,
				Updates: updates.To_Updates(),
			})
		}
	})

	_, err = syncService.GetSyncClientById(request.UserId).ReqSyncUpdates(ctx, syncUpdates)
	if err != nil {
		log.Warnf("ReqBannedRights ReqSyncUpdates request:%+v", request, err.Error())
	}

	return types.MarshalAny(updates.To_Updates())
}
