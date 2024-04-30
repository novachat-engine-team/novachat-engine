/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getWebPage_handler.go
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

//  messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
//
func (s *MessagesServiceImpl) MessagesGetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetWebPage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	webpage := webpage.GetWebPagePreview(request.Url)
	return webpage, nil
}
