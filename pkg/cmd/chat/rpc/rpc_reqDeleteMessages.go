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
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/banned_right"
	"novachat_engine/service/common/errors"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	messageCore "novachat_engine/service/core/message/message"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqDeleteMessages(ctx context.Context, request *chatService.DeleteMessages) (*types.Any, error) {

	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		log.Errorf("ReqDeleteMessages request:%+v", request, err.Error())
		return nil, err
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	messageList, err := impl.accountMessageCore.GetChannelMessageList(request.UserId, &messageCore.ChannelMessageId{
		ChannelId: request.ChatId,
		IdList:    request.Ids,
	}, request.Layer)
	if err != nil {
		log.Errorf("ReqDeleteMessages GetChannelMessageList error:%s", err.Error())
		return nil, err
	}
	if len(messageList) == 0 {
		return types.MarshalAny(&mtproto.TLInt32{Value: 0})
	}

	log.Debugf("ReqDeleteMessages fromUserId:%d request.UserId:%d", messageList[0].FromId90DDDC1171, request.UserId)
	chatData := chat.GetChatInfo().ChatData
	if !(chatData.Creator == request.UserId ||
		constants.PeerTypeFromUserIDType32(messageList[0].FromId90DDDC1171).ToInt() == request.UserId) {
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

	pts, err := impl.accountMessageCore.DeleteChannelMessages(request.ChatId, request.UserId, request.Ids)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("ReqDeleteMessages DeleteMessages error:%s need retry", err.Error())
		return nil, err
	}

	log.Debugf("ReqDeleteMessages fromUserId:%d request.UserId:%d pts:%d", messageList[0].FromId90DDDC1171, request.UserId, pts)
	if pts > 0 {
		data, _ := proto.Marshal(&chatService.DeleteMessagesUpdates{
			PeerId: request.ChatId,
			UserId: request.UserId,
			Ids:    request.Ids,
			Date:   request.Date,
			Pts:    pts,
		})
		_, _, err = impl.chatInProducer.SendMessage(constants.DeleteChannelMessages, data)
		if err != nil {
			log.Errorf("ReqDeleteMessages SendMessage error:%s need retry", err.Error())
			return nil, err
		}
	}

	return types.MarshalAny(&mtproto.TLInt32{Value: pts})
}

func (impl *Impl) ReqDeleteMessagesUpdates(ctx context.Context, request *chatService.DeleteMessagesUpdates) (*types.Any, error) {
	log.Debugf("ReqDeleteMessagesUpdates request:%+v", request)
	chat, err := impl.chatManager.GetChat(request.PeerId)
	if err != nil {
		log.Errorf("ReqDeleteMessages request:%+v", request, err.Error())
		return nil, err
	}

	if chat.Invalid() || chat.Deleted() {
		return types.MarshalAny(mtproto.ToMTBool(true))
	}

	log.Debugf("ReqDeleteMessagesUpdates request:%+v count:%d", request, chat.GetChatInfo().Count)
	updateDataList := make([]*syncService.UpdateData, 0, chat.GetChatInfo().Count)
	chat.GetChatInfo().IterationCondition(func(participant *data_chat.ChatParticipant) bool {
		if participant.State == data_chat.ParticipantState_ParticipantStateNormal {
			var update *mtproto.Update
			_, err = impl.messageCore.DeleteMessagesIds(participant.UserId, request.PeerId, constants.PeerTypeChannel, request.Ids, request.Date, request.Pts)
			if err != nil {
				return false
			}
			update = mtproto.NewTLUpdateDeleteChannelMessages(&mtproto.Update{
				ChannelId: constants.PeerTypeFromChannelIDType(request.PeerId).ToInt32(),
				PeerId:    mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(request.PeerId).ToInt32()}).To_Peer(),
				Messages:  request.Ids,
				Pts:       request.Pts,
				PtsCount:  1,
			}).To_Update()

			updateDataList = append(updateDataList, &syncService.UpdateData{
				UserId: participant.UserId,
				Updates: mtproto.NewTLUpdates(&mtproto.Updates{
					Updates: []*mtproto.Update{update},
				}).To_Updates(),
			})
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	_, err = syncService.GetSyncClientById(request.UserId).ReqSyncUpdates(ctx, &syncService.SyncUpdates{
		UpdateDataList: updateDataList,
		Push:           false,
		PeerType:       constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("ReqDeleteMessages ReqSyncUpdates:%+v", request, err.Error())
	}

	return types.MarshalAny(mtproto.ToMTBool(true))
}
