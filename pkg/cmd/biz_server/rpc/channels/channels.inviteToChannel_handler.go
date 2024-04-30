/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.inviteToChannel_handler.go
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

//  channels.inviteToChannel#199f3a6c channel:InputChannel users:Vector<InputUser> = Updates;
//
func (s *ChannelsServiceImpl) ChannelsInviteToChannel(ctx context.Context, request *mtproto.TLChannelsInviteToChannel) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsInviteToChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	peerIdList := make([]int64, 0, len(request.Users))
	for _, v := range request.Users {
		if v.UserId != 0 {
			peerIdList = append(peerIdList, constants.PeerTypeFromUserIDType32(v.UserId).ToInt())
		}
	}
	if len(peerIdList) == 0 {
		return mtproto.NewTLUpdates(nil).To_Updates(), nil
	}

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	addUserResp, err := chatService.GetChatClientByKeyId(chatId).ReqInviteToChannel(
		ctx,
		&chatService.InviteToChannel{
			UserId:     md.UserId,
			ChatId:     chatId,
			PeerIdList: peerIdList,
			AuthKeyId:  md.AuthKeyId,
			Date:       int32(time.Now().Unix()),
		})
	if err != nil {
		log.Errorf("ChannelsInviteToChannel %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	err = types.UnmarshalAny(addUserResp, &updates)
	log.Infof("ChannelsInviteToChannel %v, request: %v ok!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
