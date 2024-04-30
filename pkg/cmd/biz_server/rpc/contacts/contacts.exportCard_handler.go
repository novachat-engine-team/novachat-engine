/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.exportCard_handler.go
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

//  contacts.exportCard#84e53737 = Vector<int>;
//
func (s *ContactsServiceImpl) ContactsExportCard(ctx context.Context, request *mtproto.TLContactsExportCard) (*mtproto.VectorInt, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsExportCard %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsExportCard logic

	return nil, fmt.Errorf("%s", "Not impl ContactsExportCard")
}
