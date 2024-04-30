/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.uploadMedia_handler.go
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

//  messages.uploadMedia#519bc2b1 peer:InputPeer media:InputMedia = MessageMedia;
//
func (s *MessagesServiceImpl) MessagesUploadMedia(ctx context.Context, request *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesUploadMedia %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	media, err := s.accountMessageCore.SendUploadMedia(md.AuthKeyId, md.Layer, request.Media)
	if err != nil {
		log.Errorf("MessagesUploadMedia - request: %v error:%s", request, err.Error())
		return nil, err
	}

	log.Infof("MessagesUploadMedia - request: %v reply ok!!!!!!!!!!!!!!!", request)
	return media, nil
}
