/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getWebPagePreview_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/webpage"
)

//  messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
//  messages.getWebPagePreview#25223e24 message:string = MessageMedia;
//
func (s *MessagesServiceImpl) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetWebPagePreview %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	webpage := webpage.GetWebPagePreview(request.Message)
	media := mtproto.NewTLMessageMediaWebPage(&mtproto.MessageMedia{
		Webpage: webpage,
	})

	return media.To_MessageMedia(), nil
}
