/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.editGroupCallMember_handler.go
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

//  phone.editGroupCallMember#63146ae4 flags:# muted:flags.0?true call:InputGroupCall user_id:InputUser = Updates;
//
func (s *PhoneServiceImpl) PhoneEditGroupCallMember(ctx context.Context, request *mtproto.TLPhoneEditGroupCallMember) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneEditGroupCallMember %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneEditGroupCallMember logic

	return nil, fmt.Errorf("%s", "Not impl PhoneEditGroupCallMember")
}
