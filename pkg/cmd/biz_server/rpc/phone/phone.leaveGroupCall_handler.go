/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.leaveGroupCall_handler.go
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

//  phone.leaveGroupCall#500377f9 call:InputGroupCall source:int = Updates;
//
func (s *PhoneServiceImpl) PhoneLeaveGroupCall(ctx context.Context, request *mtproto.TLPhoneLeaveGroupCall) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneLeaveGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneLeaveGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneLeaveGroupCall")
}
