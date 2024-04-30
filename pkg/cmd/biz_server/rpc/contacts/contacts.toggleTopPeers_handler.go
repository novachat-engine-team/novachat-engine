/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.toggleTopPeers_handler.go
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

//  contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
//
func (s *ContactsServiceImpl) ContactsToggleTopPeers(ctx context.Context, request *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsToggleTopPeers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsToggleTopPeers logic

	return mtproto.ToMTBool(true), nil
}
