/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : channels.toggleComments_handler.go
 * @Desc :
 *
 */

package rpc

import (
    "context"
    "fmt"
    "novachat_engine/mtproto"
    "novachat_engine/pkg/log"
    "novachat_engine/pkg/rpc/metadata"
)

//  channels.toggleComments#aaa29e88 channel:InputChannel enabled:Bool = Updates;
//
func (s *ChannelsServiceImpl) ChannelsToggleComments(ctx context.Context, request *mtproto.TLChannelsToggleComments) (*mtproto.Updates, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("ChannelsToggleComments %v, request: %v", md, request)

    // Impl ChannelsToggleComments logic

    return nil, fmt.Errorf("%s", "Not impl ChannelsToggleComments")
}
