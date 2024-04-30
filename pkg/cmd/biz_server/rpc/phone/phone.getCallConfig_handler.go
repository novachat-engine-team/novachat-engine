/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.getCallConfig_handler.go
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

//  phone.getCallConfig#55451fa9 = DataJSON;
//
func (s *PhoneServiceImpl) PhoneGetCallConfig(ctx context.Context, request *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("PhoneGetCallConfig %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	panic("PhoneGetCallConfig")
	//reply := mtproto.NewTLDataJSON(&mtproto.DataJSON{
	//	Data: s.phoneCallConfig.GetPhoneCallConfigString(),
	//})
	//
	//log.Debugf("PhoneGetCallConfig - request: %v reply ok!!!!!!!!!!!!!!!!", request)
	//return reply.To_DataJSON(), nil
}
