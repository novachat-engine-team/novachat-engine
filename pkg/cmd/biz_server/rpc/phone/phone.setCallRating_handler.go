/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.setCallRating_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
)

//  phone.setCallRating#59ead627 flags:# user_initiative:flags.0?true peer:InputPhoneCall rating:int comment:string = Updates;
//  phone.setCallRating#1c536a34 peer:InputPhoneCall rating:int comment:string = Updates;
//
func (s *PhoneServiceImpl) PhoneSetCallRating(ctx context.Context, request *mtproto.TLPhoneSetCallRating) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("PhoneSetCallRating %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		UserId:  constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
		Updates: nil,
		Users:   nil,
		Chats:   nil,
	})

	log.Debugf("PhoneSetCallRating %v, request: %v reply ok!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return updates.To_Updates(), nil
}
