/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getNotifyExceptions_handler.go
 * @Desc :
 *
 */

package rpc

import (
    "context"
    "novachat_engine/mtproto"
    "novachat_engine/pkg/log"
    "novachat_engine/pkg/rpc/metadata"
    "time"
)

//  account.getNotifyExceptions#53577479 flags:# compare_sound:flags.1?true peer:flags.0?InputNotifyPeer = Updates;
//
func (s *AccountServiceImpl) AccountGetNotifyExceptions(ctx context.Context, request *mtproto.TLAccountGetNotifyExceptions) (*mtproto.Updates, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountGetNotifyExceptions %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountGetNotifyExceptions logic

    //  updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
    return mtproto.NewTLUpdates(&mtproto.Updates{
        Updates: []*mtproto.Update{},
        Users:   []*mtproto.User{},
        Chats:   []*mtproto.Chat{},
        Date:    int32(time.Now().Unix()),
        Seq:     0,
    }).To_Updates(), nil
}
