/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getLocated_handler.go
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

//  contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
//
func (s *ContactsServiceImpl) ContactsGetLocated(ctx context.Context, request *mtproto.TLContactsGetLocated) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsGetLocated %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsGetLocated logic

	return nil, fmt.Errorf("%s", "Not impl ContactsGetLocated")
}
