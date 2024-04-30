/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.readMessageContents_handler.go
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

//  channels.ChannelsReadMessageContents#eab5dc38 channel:InputChannel id:Vector<int> = Bool;
//
func (s *ChannelsServiceImpl) ChannelsReadMessageContents(ctx context.Context, request *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsReadMessageContents %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Id) == 0 {
		return mtproto.ToMTBool(true), nil
	}

	//chatId := request.Channel.ChannelId
	//
	////  updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
	//updates := mtproto.NewTLUpdates(&mtproto.Updates{
	//    Updates : []*mtproto.Update{mtproto.NewTLUpdateChannelReadMessagesContents(&mtproto.Update{
	//        ChannelId:          chatId,
	//        Messages:           request.Id,
	//    }).To_Update(),
	//    },
	//    Users:   []*mtproto.User{},
	//    Chats:   []*mtproto.Chat{},
	//    Date:    int32(time.Now().Unix()),
	//    Seq:     0,
	//})
	//
	//_, err := syncClient.GetSyncClient().PushChannelUpdatesList(context.Background(), &sync_pb.PushChannelUpdatesList{
	//    PushUpdatesList: []*sync_pb.PushUpdates{
	//        {
	//            UserId:               md.UserId,
	//            ExcludeAuthKeyIdList: []int64{md.AuthKeyId},
	//            Updates:              updates.To_Updates(),
	//        },
	//    },
	//    ChannelMultiUpdateDbOne: false,
	//})
	//if err != nil {
	//    log.Errorf("ChannelsReadMessageContents userId:%s PushUpdates error:%s", md.UserId, err.Error())
	//}
	//
	return mtproto.ToMTBool(true), nil
}
