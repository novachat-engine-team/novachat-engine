/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.clearAllDrafts_handler.go
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

//  messages.clearAllDrafts#7e58ee9c = Bool;
//
func (s *MessagesServiceImpl) MessagesClearAllDrafts(ctx context.Context, request *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesClearAllDrafts %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	err := s.accountMessageCore.ClearAllDrafts(md.UserId)
	if err != nil {
		log.Errorf("MessagesClearAllDrafts %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	}
	return mtproto.ToMTBool(true), nil
}
