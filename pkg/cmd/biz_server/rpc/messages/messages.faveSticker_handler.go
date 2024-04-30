/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.faveSticker_handler.go
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

//  messages.faveSticker#b9ffc55b id:InputDocument unfave:Bool = Bool;
//
func (s *MessagesServiceImpl) MessagesFaveSticker(ctx context.Context, request *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesFaveSticker %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	err := s.accountStickerSetCore.Faved(md.UserId, request.Id.Id, mtproto.ToBool(request.Unfave) == false)
	if err != nil {
		log.Errorf("MessagesFaveSticker %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	return mtproto.ToMTBool(true), nil
}
