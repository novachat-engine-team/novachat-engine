/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.sendSignalingData_handler.go
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

//  phone.sendSignalingData#ff7a9383 peer:InputPhoneCall data:bytes = Bool;
//
func (s *PhoneServiceImpl) PhoneSendSignalingData(ctx context.Context, request *mtproto.TLPhoneSendSignalingData) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("PhoneSendSignalingData %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("PhoneSendSignalingData")
	//ok, err := s.phoneCallCore.SendSignalingData(md, md.UserId, request.Peer, request.Data)
	//if err != nil {
	//	log.Errorf("PhoneSendSignalingData - request: %v error:%s !!!!!!!!!!!!!!", request, err.Error())
	//	return nil, err
	//}
	//
	//log.Infof("PhoneSendSignalingData %v, request: %v reply ok:%v !!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request, ok)
	//return mtproto.ToMTBool(ok), nil
}
