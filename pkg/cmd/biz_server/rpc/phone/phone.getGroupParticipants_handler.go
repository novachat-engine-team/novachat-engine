/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.getGroupParticipants_handler.go
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

//  phone.getGroupParticipants#c9f1d285 call:InputGroupCall ids:Vector<int> sources:Vector<int> offset:string limit:int = phone.GroupParticipants;
//
func (s *PhoneServiceImpl) PhoneGetGroupParticipants(ctx context.Context, request *mtproto.TLPhoneGetGroupParticipants) (*mtproto.Phone_GroupParticipants, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneGetGroupParticipants %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneGetGroupParticipants logic

	return nil, fmt.Errorf("%s", "Not impl PhoneGetGroupParticipants")
}
