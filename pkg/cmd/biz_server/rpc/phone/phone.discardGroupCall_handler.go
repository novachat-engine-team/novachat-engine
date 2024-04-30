/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.discardGroupCall_handler.go
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

//  phone.discardGroupCall#7a777135 call:InputGroupCall = Updates;
//
func (s *PhoneServiceImpl) PhoneDiscardGroupCall(ctx context.Context, request *mtproto.TLPhoneDiscardGroupCall) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneDiscardGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneDiscardGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneDiscardGroupCall")
}
