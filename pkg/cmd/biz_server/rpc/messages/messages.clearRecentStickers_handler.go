/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.clearRecentStickers_handler.go
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

//  messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool;
//
func (s *MessagesServiceImpl) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesClearRecentStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesClearRecentStickers logic
	err := s.accountStickerSetCore.ClearRecent(md.UserId)
	if err != nil {
		log.Errorf("MessagesClearRecentStickers %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	log.Infof("MessagesClearRecentStickers %v, request: %v ok", metadata.RpcMetaDataDebug(md), request)
	return mtproto.ToMTBool(true), nil
}
