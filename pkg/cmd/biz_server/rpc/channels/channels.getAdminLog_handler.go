/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getAdminLog_handler.go
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

//  channels.getAdminLog#33ddf480 flags:# channel:InputChannel q:string events_filter:flags.0?ChannelAdminLogEventsFilter admins:flags.1?Vector<InputUser> max_id:long min_id:long limit:int = channels.AdminLogResults;
//
func (s *ChannelsServiceImpl) ChannelsGetAdminLog(ctx context.Context, request *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsGetAdminLog %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	adminLogResults := mtproto.NewTLChannelsAdminLogResults(&mtproto.Channels_AdminLogResults{
		Events: []*mtproto.ChannelAdminLogEvent{},
		Chats:  []*mtproto.Chat{},
		Users:  []*mtproto.User{},
	})
	return adminLogResults.To_Channels_AdminLogResults(), nil
}
