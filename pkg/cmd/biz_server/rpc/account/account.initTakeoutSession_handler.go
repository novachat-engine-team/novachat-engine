/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.initTakeoutSession_handler.go
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

//  account.initTakeoutSession#f05b4804 flags:# contacts:flags.0?true message_users:flags.1?true message_chats:flags.2?true message_megagroups:flags.3?true message_channels:flags.4?true files:flags.5?true file_max_size:flags.5?int = account.Takeout;
//
func (s *AccountServiceImpl) AccountInitTakeoutSession(ctx context.Context, request *mtproto.TLAccountInitTakeoutSession) (*mtproto.Account_Takeout, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountInitTakeoutSession %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountInitTakeoutSession logic

    return nil, fmt.Errorf("%s", "Not impl AccountInitTakeoutSession")
}
