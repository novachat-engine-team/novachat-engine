/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.toggleSignatures_handler.go
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

//  channels.toggleSignatures#1f69b606 channel:InputChannel enabled:Bool = Updates;
//
func (s *ChannelsServiceImpl) ChannelsToggleSignatures(ctx context.Context, request *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsToggleSignatures %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("ChannelsToggleSignatures")

	//
	//chatLogic, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
	//if err != nil {
	//    log.Errorf("ChannelsToggleSignatures %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatLogic.IsForbidden() == true {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
	//    log.Errorf("ChannelsToggleSignatures %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//updates, err := chatLogic.ToggleSignatures(md.UserId, mtproto.ToBool(request.Enabled))
	//if err != nil {
	//    log.Errorf("ChannelsToggleSignatures %v, request: %v ToggleSignatures error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//return updates, nil
}
