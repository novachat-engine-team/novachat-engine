/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.updateDialogFiltersOrder_handler.go
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

//  messages.updateDialogFiltersOrder#c563c1e4 order:Vector<int> = Bool;
//
func (s *MessagesServiceImpl) MessagesUpdateDialogFiltersOrder(ctx context.Context, request *mtproto.TLMessagesUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesUpdateDialogFiltersOrder %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//_, err := syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
	//    UserId:               md.UserId,
	//    ExcludeAuthKeyIdList: []int64{md.AuthKeyId},
	//    Updates:              mtproto.NewTLUpdates(&mtproto.Updates{
	//        Updates: []*mtproto.Update{
	//            mtproto.NewTLUpdateDialogFilterOrder(&mtproto.Update{
	//                OrderA5D72105122: request.Order,
	//            }).To_Update(),
	//        },
	//        Users:  []*mtproto.User{},
	//        Chats:  []*mtproto.Chat{},
	//        Date:   int32(time.Now().Unix()),
	//        Seq:    0,
	//    }).To_Updates(),
	//})
	//if err != nil {
	//    log.Errorf("MessagesUpdateDialogFiltersOrder syncClient error:%s", err.Error())
	//}

	return mtproto.ToMTBool(true), nil
}
