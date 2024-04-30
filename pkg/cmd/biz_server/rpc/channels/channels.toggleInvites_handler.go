/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.toggleInvites_handler.go
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

//  channels.toggleInvites#49609307 channel:InputChannel enabled:Bool = Updates;
//
func (s *ChannelsServiceImpl) ChannelsToggleInvites(ctx context.Context, request *mtproto.TLChannelsToggleInvites) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsToggleInvites %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("ChannelsToggleInvites")

	//chatLogic, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
	//if err != nil {
	//    log.Errorf("ChannelsToggleInvites %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatLogic.IsForbidden() == true {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
	//    log.Errorf("ChannelsToggleInvites %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//updates, err := chatLogic.ToggleInvites(md, md.UserId, mtproto.ToBool(request.Enabled))
	//if err != nil {
	//    log.Errorf("ChannelsToggleInvites %v, request: %v ToggleInvites error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//return updates, nil
}
