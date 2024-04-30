/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editCreator_handler.go
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

//  channels.editCreator#8f38cd1f channel:InputChannel user_id:InputUser password:InputCheckPasswordSRP = Updates;
//
func (s *ChannelsServiceImpl) ChannelsEditCreator(ctx context.Context, request *mtproto.TLChannelsEditCreator) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsEditCreator %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("ChannelsEditCreator")

	//var err error
	//if request.Channel == nil || request.Channel.ChannelId == 0 {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHANNEL_INVALID)
	//    log.Warnf("ChannelsEditCreator %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if request.UserId == nil {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_ID_INVALID)
	//    log.Warnf("ChannelsEditCreator %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if request.UserId.ClassName == mtproto.ClassInputUserSelf  || request.UserId.ClassName == mtproto.ClassInputUser && request.UserId.UserId == md.UserId{
	//    return mtproto.NewTLUpdates(&mtproto.Updates{
	//        Updates: []*mtproto.Update{},
	//        Users:   []*mtproto.User{},
	//        Chats:   []*mtproto.Chat{},
	//        Date:    int32(time.Now().Unix()),
	//        Seq:     0,
	//    }).To_Updates(), nil
	//}
	//
	//chatCore, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
	//if err != nil {
	//    log.Errorf("ChannelsEditCreator %v, request: %v NewChatCoreById error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatCore.IsForbidden() {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
	//    log.Warnf("ChannelsEditCreator %v, request: %v NewChatCoreById error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatCore.ChatInfo.Creator != md.UserId {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
	//    log.Warnf("ChannelsEditCreator %v, request: %v NewChatCoreById error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if true {
	//    chatCore1, err := group.NewChatCoreById(request.Channel.ChannelId, request.UserId.UserId)
	//    if err != nil {
	//        log.Errorf("ChannelsEditCreator %v, request: %v NewChatCoreById error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//        return nil, err
	//    }
	//    if chatCore1.IsForbidden() || chatCore1.ChatParticipant.Banned == true {
	//        err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_ID_INVALID)
	//        log.Warnf("ChannelsEditCreator %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//        return nil, err
	//    }
	//}
	//
	//updates, err := chatCore.EditCreator(md.UserId, request.UserId, request.Password)
	//if err != nil {
	//    log.Errorf("ChannelsEditCreator %v, request: %v NewChatCoreById EditCreator error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//log.Debugf("ChannelsEditCreator %v, request: %v OK!!!!!!!!!!!!!!! reply", metadata.RpcMetaDataDebug(md), request)
	//return updates, nil
}
