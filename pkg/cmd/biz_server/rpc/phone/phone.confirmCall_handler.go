/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.confirmCall_handler.go
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

//  phone.confirmCall#2efe1722 peer:InputPhoneCall g_a:bytes key_fingerprint:long protocol:PhoneCallProtocol = phone.PhoneCall;
//
func (s *PhoneServiceImpl) PhoneConfirmCall(ctx context.Context, request *mtproto.TLPhoneConfirmCall) (*mtproto.Phone_PhoneCall, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("PhoneConfirmCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("PhoneConfirmCall")
	//replyPhoneCall, err := s.phoneCallCore.ConfirmPhoneCall(md.UserId, md.AuthKeyId, request.Peer.Id, request.GA, request.KeyFingerprint, request.Protocol)
	//if err != nil {
	//	log.Errorf("PhoneConfirmCall - request: %v error:%s", request, err.Error())
	//	return nil, err
	//}
	//
	//log.Debugf("PhoneConfirmCall - request: %v reply ok!!!!!!!!!!!!!!!!!", request)
	//return replyPhoneCall.To_Phone_PhoneCall(), nil
}
