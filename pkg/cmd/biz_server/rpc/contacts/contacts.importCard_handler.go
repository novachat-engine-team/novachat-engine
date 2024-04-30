/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.importCard_handler.go
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

//  contacts.importCard#4fe196fe export_card:Vector<int> = User;
//
func (s *ContactsServiceImpl) ContactsImportCard(ctx context.Context, request *mtproto.TLContactsImportCard) (*mtproto.User, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsImportCard %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsImportCard logic

	return nil, fmt.Errorf("%s", "Not impl ContactsImportCard")
}
