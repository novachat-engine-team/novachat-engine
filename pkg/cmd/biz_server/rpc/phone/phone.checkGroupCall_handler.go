/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.checkGroupCall_handler.go
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

//  phone.checkGroupCall#b74a7bea call:InputGroupCall source:int = Bool;
//
func (s *PhoneServiceImpl) PhoneCheckGroupCall(ctx context.Context, request *mtproto.TLPhoneCheckGroupCall) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneCheckGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneCheckGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneCheckGroupCall")
}
