/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.updateDialogFilter_handler.go
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

//  messages.updateDialogFilter#1ad4a04a flags:# id:int filter:flags.0?DialogFilter = Bool;
//
func (s *MessagesServiceImpl) MessagesUpdateDialogFilter(ctx context.Context, request *mtproto.TLMessagesUpdateDialogFilter) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesUpdateDialogFilter %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesUpdateDialogFilter logic
	//req := &dbproxy_pb.ConversationEditFolder{
	//    UserId:          md.UserId,
	//    FolderId:        request.Id,
	//    PinnedPeer:      nil,
	//    IncludePeer:     nil,
	//    ExcludePeer:     nil,
	//    Deleted:         request.Filter == nil,
	//}
	//if request.Filter != nil {
	//    req.Title = request.Filter.Title
	//    req.Emoticon = request.Filter.Emoticon
	//    req.Contacts = request.Filter.Contacts
	//    req.NonContacts = request.Filter.NonContacts
	//    req.Broadcast = request.Filter.Broadcasts
	//    req.Bots = request.Filter.Bots
	//    req.ExcludeMuted = request.Filter.ExcludeMuted
	//    req.ExcludeRead = request.Filter.ExcludeRead
	//    req.ExcludeArchived = request.Filter.ExcludeArchived
	//    req.Groups = request.Filter.Groups
	//
	//    req.PinnedPeer = make([]*dbproxy_pb.ConversationPeer, 0, len(request.Filter.PinnedPeers))
	//    for _, v := range request.Filter.PinnedPeers {
	//        inputPeer := input.MakeInputPeer(v)
	//        req.PinnedPeer = append(req.PinnedPeer, &dbproxy_pb.ConversationPeer{
	//            UserId:   md.UserId,
	//            PeerId:   inputPeer.GetPeerId(),
	//            PeerType: inputPeer.GetPeerType().ToInt32(),
	//        })
	//    }
	//    req.ExcludePeer = make([]*dbproxy_pb.ConversationPeer, 0, len(request.Filter.ExcludePeers))
	//    for _, v := range request.Filter.ExcludePeers {
	//        inputPeer := input.MakeInputPeer(v)
	//        req.ExcludePeer = append(req.PinnedPeer, &dbproxy_pb.ConversationPeer{
	//            UserId:   md.UserId,
	//            PeerId:   inputPeer.GetPeerId(),
	//            PeerType: inputPeer.GetPeerType().ToInt32(),
	//        })
	//    }
	//    req.IncludePeer = make([]*dbproxy_pb.ConversationPeer, 0, len(request.Filter.IncludePeers))
	//    for _, v := range request.Filter.IncludePeers {
	//        inputPeer := input.MakeInputPeer(v)
	//        req.IncludePeer = append(req.IncludePeer, &dbproxy_pb.ConversationPeer{
	//            UserId:   md.UserId,
	//            PeerId:   inputPeer.GetPeerId(),
	//            PeerType: inputPeer.GetPeerType().ToInt32(),
	//        })
	//    }
	//}
	//
	//ok, err := dbproxyClient.GetDBProxyClient().ConversationEditFolder(context.Background(), req)
	//if err != nil {
	//    log.Errorf("MessagesUpdateDialogFilter %v, request:%+v, error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	////mtproto.NewTLUpdateDialogFilters()
	//go func() {
	//    _, err = syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
	//        UserId:               md.UserId,
	//        ExcludeAuthKeyIdList: []int64{md.AuthKeyId},
	//        Updates:              mtproto.NewTLUpdates(&mtproto.Updates{
	//            Updates: []*mtproto.Update{
	//                mtproto.NewTLUpdateDialogFilter(&mtproto.Update{
	//                    Id4E90BFD671: request.Id,
	//                    Filter:   request.Filter,
	//                }).To_Update(),
	//            },
	//            Users:  []*mtproto.User{},
	//            Chats:  []*mtproto.Chat{},
	//            Date:   int32(time.Now().Unix()),
	//            Seq:    0,
	//        }).To_Updates(),
	//    })
	//    if err != nil {
	//        log.Errorf("MessagesUpdateDialogFilter syncClient error:%s", err.Error())
	//    }
	//}()
	//return ok, nil

	panic(fmt.Sprintf("%+v", request))
}
