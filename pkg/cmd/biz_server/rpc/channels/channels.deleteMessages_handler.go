/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.deleteMessages_handler.go
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

//  channels.deleteMessages#84c1fd4e channel:InputChannel id:Vector<int> = messages.AffectedMessages;
//
func (s *ChannelsServiceImpl) ChannelsDeleteMessages(ctx context.Context, request *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsDeleteMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqDeleteMessages(ctx, &chatService.DeleteMessages{
		ChatId:    chatId,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Ids:       request.Id,
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Warnf("ChannelsDeleteMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var pts mtproto.TLInt32
	err = types.UnmarshalAny(resp, &pts)
	ptsCount := int32(0)
	if pts.Value > 0 {
		ptsCount = 1
	}
	affectedMessage := mtproto.NewTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts.Value,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages()
	return affectedMessage, nil
}
