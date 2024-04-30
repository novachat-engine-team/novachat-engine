/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getMessagesViews_handler.go
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

//  messages.getMessagesViews#5784d3e1 peer:InputPeer id:Vector<int> increment:Bool = messages.MessageViews;
//  messages.getMessagesViews#c4c8a55d peer:InputPeer id:Vector<int> increment:Bool = Vector<int>;
//
func (s *MessagesServiceImpl) MessagesGetMessagesViews(ctx context.Context, request *mtproto.TLMessagesGetMessagesViews) (*mtproto.Response_MessagesGetMessagesViews, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("MessagesGetMessagesViews %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    response := mtproto.Response_MessagesGetMessagesViews{
        MessagesGetMessagesViews5784D3E1: mtproto.NewTLMessagesMessageViews(nil).To_Messages_MessageViews(),
        MessagesGetMessagesViewsc4C8A55D: &mtproto.VectorInt{
            Datas: nil,
        },
    }

    if md.Layer >= 119 {
        response.Cmd = mtproto.Cmd_MessagesGetMessagesViews.ToUInt32()
    } else {
        response.Cmd = mtproto.Cmd_MessagesGetMessagesViews5784d3e1.ToUInt32()
    }
    return &response, nil
}
