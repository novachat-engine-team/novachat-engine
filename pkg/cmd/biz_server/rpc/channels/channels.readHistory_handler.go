/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.readHistory_handler.go
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

//  channels.readHistory#cc104937 channel:InputChannel max_id:int = Bool;
//
func (s *ChannelsServiceImpl) ChannelsReadHistory(ctx context.Context, request *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("ChannelsReadHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    panic("ChannelsReadHistory")

    //tlUserId, err := authClient.GetAuthClient().AuthGetUserId(context.Background(), &auth_pb.AuthGetUserId{
    //    AuthKeyId: md.AuthKeyId,
    //})
    //if err != nil {
    //    log.Errorf("ChannelsReadHistory %v, request: %v AuthGetAuthKey error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
    //    return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, err.Error())
    //}
    //
    //if tlUserId.Value == 0 || tlUserId.Value != md.UserId{
    //    log.Errorf("ChannelsReadHistory %v, request: %v AuthGetUserId tlUserId.Value:%+v", md, request, tlUserId.Value)
    //    return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
    //}
    //
    //chatLogic, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
    //if err != nil {
    //    log.Errorf("ChannelsReadHistory %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
    //    return nil, err
    //}
    //
    //if chatLogic.IsForbidden() == true {
    //    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
    //    log.Errorf("ChannelsReadHistory %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
    //    return nil, err
    //}
    //
    //ok, err := chatLogic.ReadHistory(md, md.UserId, request.MaxId, chatLogic.ChatInfo.Count < s.conf.BizLogic.ReadMessageParticipantsCount)
    //if err != nil {
    //    log.Errorf("ChannelsReadHistory %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
    //    return nil, err
    //}
    //
    //return mtproto.ToMTBool(ok), nil
}
