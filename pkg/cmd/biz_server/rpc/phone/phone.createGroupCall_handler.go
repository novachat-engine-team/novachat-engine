/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.createGroupCall_handler.go
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

//  phone.createGroupCall#bd3dabe0 peer:InputPeer random_id:int = Updates;
//
func (s *PhoneServiceImpl) PhoneCreateGroupCall(ctx context.Context, request *mtproto.TLPhoneCreateGroupCall) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneCreateGroupCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneCreateGroupCall logic

	return nil, fmt.Errorf("%s", "Not impl PhoneCreateGroupCall")
}
