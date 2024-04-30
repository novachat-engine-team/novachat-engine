/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.exportMessageLink_handler.go
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

//  channels.exportMessageLink#e63fadeb flags:# grouped:flags.0?true thread:flags.1?true channel:InputChannel id:int = ExportedMessageLink;
//  channels.exportMessageLink#c846d22d channel:InputChannel id:int = ExportedMessageLink;
//
func (s *ChannelsServiceImpl) ChannelsExportMessageLink(ctx context.Context, request *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsExportMessageLink %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	exportedMessageLink := mtproto.NewTLExportedMessageLink(nil)
	return exportedMessageLink.To_ExportedMessageLink(), nil
}
