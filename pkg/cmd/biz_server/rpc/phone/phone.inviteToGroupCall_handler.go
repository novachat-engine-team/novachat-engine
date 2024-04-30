/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.inviteToGroupCall_handler.go
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

//  phone.inviteToGroupCall#7b393160 call:InputGroupCall users:Vector<InputUser> = Updates;
//
func (s *PhoneServiceImpl) PhoneInviteToGroupCall(ctx context.Context, request *mtproto.TLPhoneInviteToGroupCall) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneInviteToGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneInviteToGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneInviteToGroupCall")
}
