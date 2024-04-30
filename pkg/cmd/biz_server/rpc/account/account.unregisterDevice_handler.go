/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.unregisterDevice_handler.go
 * @Desc :
 *
 */

package rpc

import (
    "context"
    "novachat_engine/mtproto"
    "novachat_engine/pkg/log"
    "novachat_engine/pkg/rpc/metadata"
    "novachat_engine/service/constants"
)

//  account.unregisterDevice#3076c4bf token_type:int token:string other_uids:Vector<int> = Bool;
//  account.unregisterDevice#65c55b40 token_type:int token:string = Bool;
//
func (s *AccountServiceImpl) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Debugf("AccountUnregisterDevice %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    //var err error
    var ok bool
    tokenType := constants.RegisterDeviceTypeFromInt32(request.TokenType)
    if constants.RegisterDeviceTypeDefault == tokenType {
        ok = true
        log.Infof("AccountUnregisterDevice %v, request: %v RegisterDeviceType:%v reply:%v", metadata.RpcMetaDataDebug(md), request, tokenType, true)
    } else {
        panic("AccountUnregisterDevice")
        //ok, err = s.authCore.UnregisterDevice(md.AuthKeyId, request.Token, request.OtherUids)
        //if err != nil {
        //	log.Warnf("AccountUnregisterDevice - request: %v RegisterDeviceType:%v reply:%v err:%s", request, tokenType, ok, err.Error())
        //} else {
        //	log.Infof("AccountUnregisterDevice - request: %v RegisterDeviceType:%v reply:%v", request, tokenType, ok)
        //}
    }

    return mtproto.ToMTBool(ok), nil
}
