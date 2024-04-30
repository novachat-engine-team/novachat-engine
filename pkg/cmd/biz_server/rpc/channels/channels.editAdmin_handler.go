/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editAdmin_handler.go
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

//  channels.editAdmin#d33c8902 channel:InputChannel user_id:InputUser admin_rights:ChatAdminRights rank:string = Updates;
//  channels.editAdmin#20b88214 channel:InputChannel user_id:InputUser admin_rights:ChannelAdminRights = Updates;
//
func (s *ChannelsServiceImpl) ChannelsEditAdmin(ctx context.Context, request *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsEditAdmin %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqEditAdmin(ctx, &chatService.EditAdmin{
		ChatId:    chatId,
		PeerId:    constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt(),
		IsAdmin:   banned_right.AdminBannedRightRemove(request.AdminRights70F893BA93),
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Date:      int32(time.Now().Unix()),
		Right:     banned_right.ChatAdminBannedRightToRights(request.AdminRights70F893BA93),
	})
	if err != nil {
		log.Warnf("ChannelsEditAdmin %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	log.Infof("ChannelsEditAdmin %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	err = types.UnmarshalAny(resp, &updates)
	return &updates, nil
}
