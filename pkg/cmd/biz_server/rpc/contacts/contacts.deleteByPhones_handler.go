/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.deleteByPhones_handler.go
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

//  contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
//
func (s *ContactsServiceImpl) ContactsDeleteByPhones(ctx context.Context, request *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsDeleteByPhones %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsDeleteByPhones logic

	return nil, fmt.Errorf("%s", "Not impl ContactsDeleteByPhones")
}
