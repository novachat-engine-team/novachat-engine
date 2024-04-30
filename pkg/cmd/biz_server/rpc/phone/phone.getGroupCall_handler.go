/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.getGroupCall_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  phone.getGroupCall#c7cb017 call:InputGroupCall = phone.GroupCall;
//
func (s *PhoneServiceImpl) PhoneGetGroupCall(ctx context.Context, request *mtproto.TLPhoneGetGroupCall) (*mtproto.Phone_GroupCall, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneGetGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneGetGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneGetGroupCall")
}
