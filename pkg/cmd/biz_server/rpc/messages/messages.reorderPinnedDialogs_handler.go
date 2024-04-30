/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.reorderPinnedDialogs_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
	data_message "novachat_engine/service/data/messages/message"
	"sort"
	"time"
)

//  messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
//  messages.reorderPinnedDialogs#959ff644 flags:# force:flags.0?true order:Vector<InputPeer> = Bool;
//
func (s *MessagesServiceImpl) MessagesReorderPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReorderPinnedDialogs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var peerList []*data_message.Conversation
	if len(request.Order959FF64471) > 0 {
		peerList = make([]*data_message.Conversation, 0, len(request.Order959FF64471))
	} else if len(request.Order5B51D63F85) > 0 {
		peerList = make([]*data_message.Conversation, 0, len(request.Order5B51D63F85))
	}

	//orderPeer := make([]*mtproto.Peer, 0, len(request.Order959FF64471))
	//orderDialogPeer := make([]*mtproto.DialogPeer, 0, len(request.Order5B51D63F85))

	//peerList := make([]*dbproxy_pb.ConversationPeer, 0, 1)
	for _, v := range request.Order959FF64471 {
		if v.ClassName == mtproto.ClassInputPeerUser {
			peerList = append(peerList, &data_message.Conversation{
				UserId:   md.UserId,
				PeerId:   constants.PeerTypeFromUserIDType32(v.UserId).ToInt(),
				PeerType: constants.PeerTypeUser.ToInt32(),
			})
		} else if v.ClassName == mtproto.ClassInputPeerChat {
			peerList = append(peerList, &data_message.Conversation{
				UserId:   md.UserId,
				PeerId:   constants.PeerTypeFromChatIDType32(v.ChatId).ToInt(),
				PeerType: constants.PeerTypeChannel.ToInt32(),
			})
		} else if v.ClassName == mtproto.ClassInputPeerChannel {
			peerList = append(peerList, &data_message.Conversation{
				UserId:   md.UserId,
				PeerId:   constants.PeerTypeFromChannelIDType32(v.ChannelId).ToInt(),
				PeerType: constants.PeerTypeChannel.ToInt32(),
			})
		}
	}
	for _, v := range request.Order5B51D63F85 {
		if v.Peer.ClassName == mtproto.ClassInputPeerUser {
			peerList = append(peerList, &data_message.Conversation{
				UserId:   md.UserId,
				PeerId:   constants.PeerTypeFromUserIDType32(v.Peer.UserId).ToInt(),
				PeerType: constants.PeerTypeUser.ToInt32(),
			})
		} else if v.Peer.ClassName == mtproto.ClassInputPeerChat {
			peerList = append(peerList, &data_message.Conversation{
				UserId:   md.UserId,
				PeerId:   constants.PeerTypeFromChatIDType32(v.Peer.ChatId).ToInt(),
				PeerType: constants.PeerTypeChannel.ToInt32(),
			})
		} else if v.Peer.ClassName == mtproto.ClassInputPeerChannel {
			peerList = append(peerList, &data_message.Conversation{
				UserId:   md.UserId,
				PeerId:   constants.PeerTypeFromChannelIDType32(v.Peer.ChannelId).ToInt(),
				PeerType: constants.PeerTypeChannel.ToInt32(),
			})
		}
	}

	var err error
	var conversationList []*data_message.Conversation
	if len(peerList) > 0 {
		conversationList, err = s.messageCore.GetConversationByPeerList(md.UserId, peerList)
	} else {
		conversationList, err = s.messageCore.GetConversationList(md.UserId, false, constants.PeerTypeEmpty, 0, 0, request.FolderId)
	}
	if err != nil {
		log.Errorf("GetConversationList %v, request: %v ConversationFolders error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	sort.Slice(conversationList, func(i, j int) bool {
		return conversationList[i].Date > conversationList[j].Date
	})

	orderDialogPeer := make([]*mtproto.DialogPeer, 0, len(conversationList))
	orderPeer := make([]*mtproto.Peer, 0, len(conversationList))
	for _, v := range conversationList {
		var peer *mtproto.Peer
		switch constants.PeerTypeFromInt32(v.PeerType) {
		case constants.PeerTypeUser:
			peer = mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(v.PeerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChat:
			peer = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChatIDType(v.PeerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChannel:
			peer = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(v.PeerId).ToInt32()}).To_Peer()
		default:
			continue
		}
		orderDialogPeer = append(orderDialogPeer, mtproto.NewTLDialogPeer(&mtproto.DialogPeer{
			Peer:     peer,
			FolderId: request.FolderId,
		}).To_DialogPeer())
		orderPeer = append(orderPeer, peer)
	}

	go func() {
		_, err = syncClient.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncClient.SyncUpdate{
			UserId: md.UserId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{mtproto.NewTLUpdatePinnedDialogs(&mtproto.Update{
					FolderId:        request.FolderId,
					OrderEA4CB65B85: orderDialogPeer,
					OrderD8CAF68D71: orderPeer,
				}).To_Update()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
				Date:  int32(time.Now().Unix()),
				Seq:   0,
			}).To_Updates(),
			Push:     false,
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("MessagesReorderPinnedDialogs error:%s", err.Error())
		}
	}()

	return mtproto.ToMTBool(true), nil
}
