/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.discardCall_handler.go
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

//  phone.discardCall#b2cbc1c0 flags:# video:flags.0?true peer:InputPhoneCall duration:int reason:PhoneCallDiscardReason connection_id:long = Updates;
//  phone.discardCall#78d413a6 peer:InputPhoneCall duration:int reason:PhoneCallDiscardReason connection_id:long = Updates;
//
func (s *PhoneServiceImpl) PhoneDiscardCall(ctx context.Context, request *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Debugf("PhoneDiscardCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    panic("PhoneDiscardCall")
    //updates, err := s.phoneCallCore.DiscardCall(md, md.UserId, md.AuthKeyId, request.Peer.Id, request.Reason, request.Duration, request.Video, request.ConnectionId)
    //if err != nil {
    //	log.Errorf("PhoneDiscardCall - request: %v error:%s", request, err.Error())
    //	return nil, err
    //}
    //
    //log.Debugf("PhoneDiscardCall - request: %v reply ok!!!!!!!!!!!!!!!!!", request)
    //return updates.To_Updates(), nil
}
