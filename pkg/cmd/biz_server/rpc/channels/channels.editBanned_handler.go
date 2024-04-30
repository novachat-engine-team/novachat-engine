/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editBanned_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/banned_right"
	"novachat_engine/service/constants"
	"time"
)

//  channels.editBanned#72796912 channel:InputChannel user_id:InputUser banned_rights:ChatBannedRights = Updates;
//  channels.editBanned#bfd915cd channel:InputChannel user_id:InputUser banned_rights:ChannelBannedRights = Updates;
//
func (s *ChannelsServiceImpl) ChannelsEditBanned(ctx context.Context, request *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsEditBanned %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	isBanned := banned_right.BannedRightToBanned(request.BannedRights7279691293)
	deleted := banned_right.BannedRightToDeleted(request.BannedRights7279691293)
	right, date := banned_right.ChatBannedRightToRights(request.BannedRights7279691293)
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqBannedUser(ctx, &chatService.BannedUser{
		ChatId:    chatId,
		PeerId:    constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt(),
		UserId:    md.UserId,
		IsBanned:  isBanned,
		IsDeleted: deleted,
		Rights:    right,
		UtilDate:  date,
		Date:      int32(time.Now().Unix()),
		AuthKeyId: md.AuthKeyId,
	})
	if err != nil {
		log.Errorf("ChannelsEditBanned %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	err = types.UnmarshalAny(resp, &updates)
	log.Infof("ChannelsEditBanned %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return &updates, nil
}
