/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.setDiscussionGroup_handler.go
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

//  channels.setDiscussionGroup#40582bb2 broadcast:InputChannel group:InputChannel = Bool;
//
func (s *ChannelsServiceImpl) ChannelsSetDiscussionGroup(ctx context.Context, request *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsSetDiscussionGroup %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.ToMTBool(false), nil
}
