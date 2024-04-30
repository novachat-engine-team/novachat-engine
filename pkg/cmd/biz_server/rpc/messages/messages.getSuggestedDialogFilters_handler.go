/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getSuggestedDialogFilters_handler.go
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

//  messages.getSuggestedDialogFilters#a29cd42c = Vector<DialogFilterSuggested>;
//
func (s *MessagesServiceImpl) MessagesGetSuggestedDialogFilters(ctx context.Context, request *mtproto.TLMessagesGetSuggestedDialogFilters) (*mtproto.Vector_DialogFilterSuggested, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetSuggestedDialogFilters %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetSuggestedDialogFilters logic
	return &mtproto.Vector_DialogFilterSuggested{
		Datas: []*mtproto.DialogFilterSuggested{},
	}, nil
}
