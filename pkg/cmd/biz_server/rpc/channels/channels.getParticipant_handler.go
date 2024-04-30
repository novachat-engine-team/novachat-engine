/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getParticipant_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  channels.getParticipant#546dd7a6 channel:InputChannel user_id:InputUser = channels.ChannelParticipant;
//
func (s *ChannelsServiceImpl) ChannelsGetParticipant(ctx context.Context, request *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetParticipant %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("ChannelsGetParticipant")

	//chatLogic, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
	//if err != nil {
	//    log.Errorf("ChannelsGetParticipant %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatLogic.IsForbidden() == true {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHANNEL_PRIVATE)
	//    log.Errorf("ChannelsGetParticipant %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//participant, err := chatLogic.GetParticipant(md.UserId, request.UserId)
	//if err != nil {
	//    log.Errorf("ChannelsGetParticipant %v, request: %v GetParticipant error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//for idx := range participant.Users {
	//    participant.Users[idx] = compat.CompatUser(participant.Users[idx], md.Layer)
	//}
	//
	//return participant, nil
}
