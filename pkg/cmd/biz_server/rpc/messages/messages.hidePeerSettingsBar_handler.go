/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.hidePeerSettingsBar_handler.go
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

//  messages.hidePeerSettingsBar#4facb138 peer:InputPeer = Bool;
//
func (s *MessagesServiceImpl) MessagesHidePeerSettingsBar(ctx context.Context, request *mtproto.TLMessagesHidePeerSettingsBar) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesHidePeerSettingsBar %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesHidePeerSettingsBar logic

	return mtproto.ToMTBool(true), nil
	//return nil, fmt.Errorf("%s", "Not impl MessagesHidePeerSettingsBar")
}
