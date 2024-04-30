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
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/banned_right"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/common/hash"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	data_fs "novachat_engine/service/data/fs"
	"novachat_engine/service/upload/photo"
)

func (impl *Impl) ReqEditTitle(ctx context.Context, request *chatService.EditTitle) (*types.Any, error) {
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
		if !participant.Admin && banned_right.RightsToChatBannedRight(participant.Rights, 0).ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
		}

		bannedRight := banned_right.RightsToChatBannedRight(chatData.BannedRights, chatData.RightsUtilDate)
		if bannedRight.ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	_, err = chat.EditTitle(request.Title)
	if err != nil {
		log.Errorf("ReqEditTitle request:%+v", request, err.Error())
		return nil, err
	}
	chatPeer := mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChatIDType(chatData.ChatId).ToInt32()}).To_Peer()
	message := mtproto.NewTLMessageService(&mtproto.Message{
		Date:              request.Date,
		FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(request.UserId).ToInt32()}).To_Peer(),
		FromId90DDDC1171:  constants.PeerTypeFromUserIDType(request.UserId).ToInt32(),
		ToId:              chatPeer,
		PeerId:            chatPeer,
		Action: mtproto.NewTLMessageActionChatEditTitle(&mtproto.MessageAction{
			Title: request.Title,
		}).To_MessageAction(),
	})

	updates, err := impl.ReqSendOutboxesMessages(ctx, &chatService.SendOutboxesMessages{
		Message: &msgClient.SendMessages{
			AuthKeyId:  request.AuthKeyId,
			FromUserId: request.UserId,
			PeerId:     chatData.ChatId,
			PeerType:   constants.PeerTypeChannel.ToInt32(),
			DataList: []*msgService.SendMessageData{{
				RandomId: hash.HashId64ListNew([]int64{request.UserId, int64(request.Date)}),
				Message:  message.To_Message(),
			}},
		},
	})
	if err != nil {
		log.Warnf("ReqEditTitle ReqSendOutboxesMessages request:%+v", request, err.Error())
		return nil, err
	} else {
		return types.MarshalAny(updates)
	}
}

func (impl *Impl) ReqEditPhoto(ctx context.Context, request *chatService.EditPhoto) (*types.Any, error) {
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
		if !participant.Admin && banned_right.RightsToChatBannedRight(participant.Rights, 0).ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
		}

		bannedRight := banned_right.RightsToChatBannedRight(chatData.BannedRights, chatData.RightsUtilDate)
		if bannedRight.ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	_, err = chat.EditPhoto(request.Photo)
	if err != nil {
		log.Errorf("ReqEditPhoto request:%+v", request, err.Error())
		return nil, err
	}
	chatPeer := mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChatIDType(chatData.ChatId).ToInt32()}).To_Peer()
	message := mtproto.NewTLMessageService(&mtproto.Message{
		Date:              request.Date,
		FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(request.UserId).ToInt32()}).To_Peer(),
		FromId90DDDC1171:  constants.PeerTypeFromUserIDType(request.UserId).ToInt32(),
		ToId:              chatPeer,
		PeerId:            chatPeer,
		Action: mtproto.NewTLMessageActionChatEditPhoto(&mtproto.MessageAction{
			Photo: photo.PhotoData2Photo(&data_fs.Photo{
				VolumeId: request.Photo.Photo.VolumeId,
				LocalId:  request.Photo.Photo.LocalId,
				Detail:   []*data_fs.PhotoDetail{request.Photo.Photo},
			}),
		}).To_MessageAction(),
	})

	updates, err := impl.ReqSendOutboxesMessages(ctx, &chatService.SendOutboxesMessages{
		Message: &msgClient.SendMessages{
			AuthKeyId:  request.AuthKeyId,
			FromUserId: request.UserId,
			PeerId:     chatData.ChatId,
			PeerType:   constants.PeerTypeChannel.ToInt32(),
			DataList: []*msgService.SendMessageData{{
				RandomId: hash.HashId64ListNew([]int64{request.UserId, int64(request.Date)}),
				Message:  message.To_Message(),
			}},
		},
	})
	if err != nil {
		log.Warnf("ReqEditPhoto ReqSendOutboxesMessages request:%+v", request, err.Error())
		return nil, err
	} else {
		updates.Updates = append(updates.Updates, mtproto.NewTLUpdateChannel(&mtproto.Update{ChannelId: constants.PeerTypeFromChatIDType(request.ChatId).ToInt32()}).To_Update())
		return types.MarshalAny(updates)
	}
}

func (impl *Impl) ReqEditAbout(ctx context.Context, request *chatService.EditAbout) (*types.Any, error) {
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
		if !participant.Admin && banned_right.RightsToChatBannedRight(participant.Rights, 0).ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
		}

		bannedRight := banned_right.RightsToChatBannedRight(chatData.BannedRights, chatData.RightsUtilDate)
		if bannedRight.ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	_, err = chat.EditAbout(request.About)
	if err != nil {
		log.Errorf("ReqEditAbout request:%+v", request, err.Error())
		return nil, err
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromChatIDType(request.ChatId).ToInt32(),
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
		log.Warnf("ReqEditAbout ReqSyncUpdates:%+v", request, err.Error())
	}

	return types.MarshalAny(updates)
}

func (impl *Impl) ReqEditAdmin(ctx context.Context, request *chatService.EditAdmin) (*types.Any, error) {
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

		if banned_right.RightsToChatAdminBannedRight(participant.AdminRights).AddAdmins {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
		}
	}

	_, err = chat.EditAdmin(request.UserId, request.PeerId, request.IsAdmin)
	if err != nil {
		log.Errorf("ReqEditAdmin request:%+v", request, err.Error())
		return nil, err
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{ChannelId: constants.PeerTypeFromChatIDType(request.ChatId).ToInt32()}).To_Update(),
			mtproto.NewTLUpdateChatParticipantAdmin(&mtproto.Update{
				ChatId:  constants.PeerTypeFromChatIDType(request.ChatId).ToInt32(),
				UserId:  constants.PeerTypeFromUserIDType(request.PeerId).ToInt32(),
				IsAdmin: mtproto.ToMTBool(request.IsAdmin),
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
		log.Warnf("ReqEditAdmin ReqSyncUpdates:%+v", request, err.Error())
	}

	return types.MarshalAny(updates)
}

func (impl *Impl) ReqEditGeoPoint(ctx context.Context, request *chatService.EditGeoPoint) (*types.Any, error) {

	chat, err := impl.chatManager.GetChat(request.PeerId)
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
		if !participant.Admin && banned_right.RightsToChatBannedRight(participant.Rights, 0).ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
		}

		bannedRight := banned_right.RightsToChatBannedRight(chatData.BannedRights, chatData.RightsUtilDate)
		if bannedRight.ChangeInfo {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	var geopoint *data_fs.GeoPoint
	var address string
	if request.GeoPoint != nil {
		geopoint = &data_fs.GeoPoint{
			Long:           request.GeoPoint.Long,
			Lat:            request.GeoPoint.Lat,
			AccessHash:     0,
			AccuracyRadius: 0,
		}
		address = request.GeoPoint.Address
	}

	_, err = chat.EditGeoPoint(geopoint, address)
	if err != nil {
		log.Errorf("ReqEditAbout request:%+v", request, err.Error())
		return nil, err
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromChatIDType(request.PeerId).ToInt32(),
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
		log.Warnf("ReqEditAbout ReqSyncUpdates:%+v", request, err.Error())
	}

	return types.MarshalAny(mtproto.ToMTBool(true))
}
