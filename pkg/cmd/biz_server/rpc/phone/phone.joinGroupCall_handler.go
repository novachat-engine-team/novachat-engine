/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.joinGroupCall_handler.go
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

//  phone.joinGroupCall#5f9c8e62 flags:# muted:flags.0?true call:InputGroupCall params:DataJSON = Updates;
//
func (s *PhoneServiceImpl) PhoneJoinGroupCall(ctx context.Context, request *mtproto.TLPhoneJoinGroupCall) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneJoinGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneJoinGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneJoinGroupCall")
}
