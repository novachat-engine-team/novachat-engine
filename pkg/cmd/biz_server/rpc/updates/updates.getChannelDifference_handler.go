/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : updates.getChannelDifference_handler.go
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

//  updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
//
func (s *UpdatesServiceImpl) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UpdatesGetChannelDifference %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("UpdatesGetChannelDifference")
	//
	//chatLogic, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
	//if err != nil {
	//    log.Errorf("UpdatesGetChannelDifference %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatLogic.IsForbidden() == true {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHANNEL_PRIVATE)
	//    log.Errorf("UpdatesGetChannelDifference %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//updatesGetChannelDifference, err := chatLogic.ChatGetChannelDifference(md.UserId, request.Filter, request.Pts, request.Limit, request.Force, md.Layer, 0)
	//if err != nil {
	//    log.Errorf("ChannelsGetChannels %v, request: %v GetChats error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//return updatesGetChannelDifference, nil
}
