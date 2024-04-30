/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getMessages_handler.go
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

//  channels.getMessages#ad8c9a23 channel:InputChannel id:Vector<InputMessage> = messages.Messages;
//  channels.getMessages#93d7b347 channel:InputChannel id:Vector<int> = messages.Messages;
//
func (s *ChannelsServiceImpl) ChannelsGetMessages(ctx context.Context, request *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	panic("ChannelsGetMessages")
	//
	//if len(request.IdAD8C9A23122) == 0 && len(request.Id93D7B34771) == 0 {
	//    err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MSG_ID_INVALID)
	//    log.Errorf("ChannelsGetMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//chatLogic, err := group.NewChatCoreById(request.Channel.ChannelId, md.UserId)
	//if err != nil {
	//    log.Errorf("ChannelsGetMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//if chatLogic.IsForbidden() == true {
	//    err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHANNEL_PRIVATE)
	//    log.Errorf("ChannelsGetMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//pinnedIdListData := &dbproxy_pb.MessagePinnedMessageData{
	//    UserId:               md.UserId,
	//    InputPeer:            mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: md.UserId}).To_InputPeer(),
	//}
	//
	//idListData := &dbproxy_pb.MessageMessageData{
	//    UserId:               md.UserId,
	//    MessageIdList:        []int32{},
	//    InputPeer:            mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: md.UserId}).To_InputPeer(),
	//}
	//id0 := int32(0)
	//isPinned := false
	//for idx, v := range request.IdAD8C9A23122 {
	//    if idx == 0 {
	//        id0 = v.Id
	//    }
	//    if v.ClassName == mtproto.ClassInputMessagePinned {
	//        isPinned = true
	//    } else {
	//        idListData.MessageIdList = append(idListData.MessageIdList, v.Id)
	//    }
	//}
	//for idx, v := range request.Id93D7B34771 {
	//    if idx == 0 {
	//        id0 = v
	//    }
	//    idListData.MessageIdList = append(idListData.MessageIdList, v)
	//}
	//
	//var messageGetMessageList *dbproxy_pb.MessageGetMessageList
	//if isPinned {
	//    messageGetMessageList, err = dbproxyClient.GetDBProxyClient().MessageGetPinnedMessages(context.Background(), &dbproxy_pb.MessageGetPinnedMessages{
	//       Value: []*dbproxy_pb.MessagePinnedMessageData{pinnedIdListData},
	//       UserId: md.UserId,
	//   })
	//} else {
	//    messageGetMessageList, err = dbproxyClient.GetDBProxyClient().MessageGetMessagesByMessageIds(context.Background(), &dbproxy_pb.MessageGetMessagesByMessageIds{
	//       Value:                []*dbproxy_pb.MessageMessageData{idListData},
	//       UserId:               md.UserId,
	//   })
	//}
	//if err != nil {
	//    log.Errorf("ChannelsGetMessages %v, request: %v isPinned:%v error:%s", metadata.RpcMetaDataDebug(md), request, isPinned, err.Error())
	//    return nil, err
	//}
	//
	//offsetIdOffset := int32(0)
	//if id0 != 0 {
	//    for idx, v := range messageGetMessageList.MessageList {
	//        if v.Id == id0 {
	//            offsetIdOffset = int32(idx)
	//            break
	//        }
	//    }
	//}
	//
	//// messages.channelMessages#64479808 flags:# inexact:flags.1?true pts:int count:int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
	//// messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
	//messagesMessages := mtproto.NewTLMessagesChannelMessages(&mtproto.Messages_Messages{
	//    Inexact: false,
	//    Pts:     messageGetMessageList.MaxPts,
	//    Count:  int32(len(messageGetMessageList.MessageList)),
	//    OffsetIdOffset: offsetIdOffset,
	//    Messages: messageGetMessageList.MessageList,
	//    Chats: []*mtproto.Chat{},
	//    Users: []*mtproto.User{},
	//})
	//
	//return messagesMessages.To_Messages_Messages(), nil
}
