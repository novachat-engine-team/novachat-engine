/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : channels.getImportantHistory_handler.go
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

//  channels.getImportantHistory#8f494bb2 channel:InputChannel offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
//
func (s *ChannelsServiceImpl) ChannelsGetImportantHistory(ctx context.Context, request *mtproto.TLChannelsGetImportantHistory) (*mtproto.Messages_Messages, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("ChannelsGetImportantHistory %v, request: %v", md, request)

    // Impl ChannelsGetImportantHistory logic

    return nil, fmt.Errorf("%s", "Not impl ChannelsGetImportantHistory")
}
