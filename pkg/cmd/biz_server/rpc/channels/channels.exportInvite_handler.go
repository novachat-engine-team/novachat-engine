/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.exportInvite_handler.go
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

//  channels.exportInvite#c7560885 channel:InputChannel = ExportedChatInvite;
//
func (s *ChannelsServiceImpl) ChannelsExportInvite(ctx context.Context, request *mtproto.TLChannelsExportInvite) (*mtproto.ExportedChatInvite, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsExportInvite %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("ChannelsExportInvite")

	//
	//chatLogic := group.NewChatCore()
	//link, err := chatLogic.ExportInvite(request.Channel.ChannelId)
	//if err != nil {
	//    log.Errorf("ChannelsExportInvite %v, request: %v ExportInvite error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//log.Infof("ChannelsExportInvite %v, request: %v link:%s, len(link)= %d", metadata.RpcMetaDataDebug(md), request, link, len(link))
	//
	//if len(link) == 0 {
	//    return mtproto.NewTLChatInviteEmpty(nil).To_ExportedChatInvite(), nil
	//}
	//
	//return mtproto.NewTLChatInviteExported(&mtproto.ExportedChatInvite{
	//    Link:                 link,
	//}).To_ExportedChatInvite(), nil
}
