/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.receivedCall_handler.go
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

//  phone.receivedCall#17d54f61 peer:InputPhoneCall = Bool;
//
func (s *PhoneServiceImpl) PhoneReceivedCall(ctx context.Context, request *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Debugf("PhoneReceivedCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)
    panic("PhoneReceivedCall")
    //ok, err := s.phoneCallCore.ReceivedPhoneCall(md.UserId, request.Peer.Id)
    //if err != nil {
    //	log.Errorf("PhoneReceivedCall - request: %v error:%s", request, err.Error())
    //	return nil, err
    //}
    //
    //log.Infof("PhoneReceivedCall - request: %v reply:%v ok!!!!!!!!!!!!!!!!!", request, ok)
    //return mtproto.ToMTBool(ok), nil
}
