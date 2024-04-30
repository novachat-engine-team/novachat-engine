/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.setTyping_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.setTyping#58943ee2 flags:# peer:InputPeer top_msg_id:flags.0?int action:SendMessageAction = Bool;
//  messages.setTyping#a3825e50 peer:InputPeer action:SendMessageAction = Bool;
//
func (s *MessagesServiceImpl) MessagesSetTyping(ctx context.Context, request *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesSetTyping %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	_ = request.TopMsgId

	peer := input.MakeInputPeer(request.Peer)
	if peer.GetPeerType() == constants.PeerTypeUser {
		peer.UpdatePeerId(peer.GetPeerId())
	}

	if peer.GetPeerType() == constants.PeerTypeUser || peer.GetPeerType() == constants.PeerTypeSelf {
		if md.UserId == peer.GetPeerId() {
			return mtproto.ToMTBool(true), nil
		}

		typing := mtproto.NewTLUpdateUserTyping(nil)
		typing.SetUserId(constants.PeerTypeFromUserIDType(md.UserId).ToInt32())
		typing.SetAction(request.GetAction())

		updates := mtproto.NewTLUpdates(nil)
		updates.SetUpdates([]*mtproto.Update{typing.To_Update()})
		updates.SetUsers([]*mtproto.User{})
		updates.SetChats([]*mtproto.Chat{})
		updates.SetDate(int32(time.Now().Unix()))
		updates.SetSeq(0)

		_, err := syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
			UserId:   peer.GetPeerId(),
			Updates:  updates.To_Updates(),
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Errorf("MessagesSetTyping %v, request: %v PushUpdates error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		}
	} else {
		//return mtproto.ToMTBool(true), nil
		//chatLogic, err := group.NewChatCoreById(peer.GetPeerId(), md.UserId)
		//if err != nil {
		//    log.Errorf("MessagesSetTyping NewChatCoreById userId:%d chatId:%d error:%s", md.UserId, peer.GetPeerId(), err.Error())
		//    return nil, err
		//}
		//if peer.GetPeerType() == constants.PeerTypeChat {
		//    peer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
		//        ChannelId:  request.Peer.ChatId,
		//        AccessHash: request.Peer.AccessHash,
		//    }).To_InputPeer())
		//}
		//
		//if chatLogic.ChatInfo.Creator != md.UserId {
		//    if chatLogic.ChatParticipant.Admin == false {
		//        rights := banned_right.RightsToChatBannedRight(chatLogic.ChatInfo.Right, 0)
		//        if rights.GetSendMessages() == true {
		//            err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_RESTRICTED)
		//            log.Errorf("MessagesSetTyping GetSendMessages() == false chatId:%d userId:%d error:%s", peer.GetPeerId(), md.UserId, err.Error())
		//            return nil, err
		//        }
		//
		//        rights = banned_right.RightsToChatBannedRight(chatLogic.ChatParticipant.Right, 0)
		//        if rights.GetSendMessages() == true {
		//            err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_RESTRICTED)
		//            log.Errorf("MessagesSetTyping ChatParticipant GetSendMessages() == false chatId:%d userId:%d error:%s", peer.GetPeerId(), md.UserId, err.Error())
		//            return nil, err
		//        }
		//    }
		//}
		//
		//typing := mtproto.NewTLUpdateChannelUserTyping(nil)
		//typing.SetUserId(md.UserId)
		//typing.SetAction(request.GetAction())
		//typing.SetChannelId(peer.GetPeerId())
		//
		//updates := mtproto.NewTLUpdates(nil)
		//updates.SetUpdates([]*mtproto.Update{typing.To_Update()})
		//updates.SetUsers([]*mtproto.User{})
		//updates.SetChats([]*mtproto.Chat{})
		//updates.SetDate(int32(time.Now().Unix()))
		//updates.SetSeq(0)
		//
		//ctx1, cancel := rpc_context.Context(nil, constants.RpcDefaultTimeout)
		//defer cancel()
		//participant, err := dbProxyClient.GetDBProxyClient().ChatGetParticipants(ctx1, &dbproxy_pb.ChatGetParticipants{
		//    ChatId: peer.GetPeerId(),
		//    UserId: md.UserId,
		//    Filter: constants.ChannelParticipantsTypeContacts.ToInt32(),
		//})
		//if err != nil {
		//    log.Errorf("MessagesSetTyping userId:%d peerId:%v ChatGetParticipants error:%s", md.UserId, peer, err.Error())
		//    return nil, err
		//}
		//
		//const chatSliceMax = int32(512)
		//idx := int32(0)
		//syncUpdatesList := make([]*sync_pb.SyncUpdates, chatSliceMax, chatSliceMax)
		//
		//for _, v := range participant.ChatParticipant {
		//    if v.Left == true {
		//        continue
		//    }
		//
		//    if v.UserId != 0 {
		//        if v.UserId == md.UserId {
		//            continue
		//        } else {
		//            syncUpdatesList[idx] = &sync_pb.SyncUpdates{
		//                UserId:               v.UserId,
		//                Updates:              updates.To_Updates(),
		//            }
		//            idx++
		//        }
		//    }
		//
		//    if idx >= chatSliceMax {
		//        _, err = syncClient.GetSyncClient().SyncChannelUpdatesList(context.Background(), &sync_pb.SyncChannelUpdatesList{
		//            SyncUpdatesList:      syncUpdatesList[:chatSliceMax],
		//        })
		//        if err != nil {
		//            log.Errorf("MessagesSetTyping SyncUpdates userId:%d peerId:%v", v.UserId, peer)
		//        }
		//        idx = 0
		//    }
		//}
		//
		//if idx > 0 {
		//    _, err = syncClient.GetSyncClient().SyncChannelUpdatesList(context.Background(), &sync_pb.SyncChannelUpdatesList{
		//        SyncUpdatesList: syncUpdatesList[:idx],
		//    })
		//    if err != nil {
		//        log.Errorf("MessagesSetTyping SyncUpdates userId:%d peerId:%v ", md.UserId, peer)
		//    }
		//}
	}

	log.Debugf("MessagesSetTyping %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return mtproto.ToMTBool(true), nil
}
