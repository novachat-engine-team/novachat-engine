/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.joinChannel_handler.go
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
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"time"
)

//  channels.joinChannel#24b524c5 channel:InputChannel = Updates;
//
func (s *ChannelsServiceImpl) ChannelsJoinChannel(ctx context.Context, request *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsJoinChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	addUserResp, err := chatService.GetChatClientByKeyId(chatId).ReqAddUser(
		ctx,
		&chatService.AddUser{
			UserId:    md.UserId,
			ChatId:    chatId,
			PeerId:    md.UserId,
			AuthKeyId: md.AuthKeyId,
			Date:      int32(time.Now().Unix()),
		})
	if err != nil {
		log.Errorf("ChannelsJoinChannel %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	types.UnmarshalAny(addUserResp, &updates)
	log.Infof("ChannelsJoinChannel %v, request: %v ok!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
