/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.resetTopPeerRating_handler.go
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

//  contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
//
func (s *ContactsServiceImpl) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsResetTopPeerRating %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsResetTopPeerRating logic

	return mtproto.ToMTBool(true), nil
}
