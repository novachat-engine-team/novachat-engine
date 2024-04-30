/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.resetSaved_handler.go
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

//  contacts.resetSaved#879537f1 = Bool;
//
func (s *ContactsServiceImpl) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsResetSaved %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsResetSaved logic

	return mtproto.ToMTBool(true), nil
}
