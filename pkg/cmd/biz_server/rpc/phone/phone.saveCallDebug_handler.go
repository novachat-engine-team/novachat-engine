/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.saveCallDebug_handler.go
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

//  phone.saveCallDebug#277add7e peer:InputPhoneCall debug:DataJSON = Bool;
//
func (s *PhoneServiceImpl) PhoneSaveCallDebug(ctx context.Context, request *mtproto.TLPhoneSaveCallDebug) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("PhoneSaveCallDebug %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("PhoneSaveCallDebug")
	//s.phoneCallCore.SaveCallDebug(md.UserId, request.Peer, request.Debug)
	//
	//return mtproto.ToMTBool(true), nil
}
