/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : channels.kickFromChannel_handler.go
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
	"novachat_engine/service/constants"
	"time"
)

//  channels.kickFromChannel#a672de14 channel:InputChannel user_id:InputUser kicked:Bool = Updates;
//
func (s *ChannelsServiceImpl) ChannelsKickFromChannel(ctx context.Context, request *mtproto.TLChannelsKickFromChannel) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsKickFromChannel %v, request: %v", md, request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqBannedUser(ctx, &chatService.BannedUser{
		ChatId:    chatId,
		PeerId:    constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt(),
		UserId:    md.UserId,
		IsDeleted: mtproto.ToBool(request.Kicked),
		Date:      int32(time.Now().Unix()),
		AuthKeyId: md.AuthKeyId,
	})
	if err != nil {
		log.Errorf("ChannelsKickFromChannel %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	err = types.UnmarshalAny(resp, &updates)
	log.Infof("ChannelsKickFromChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return &updates, nil
}
