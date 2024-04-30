/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.search_handler.go
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

//400	CHANNEL_INVALID	The provided channel is invalid
//400	CHANNEL_PRIVATE	You haven't joined this channel/supergroup
//400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
//400	INPUT_CONSTRUCTOR_INVALID	The provided constructor is invalid
//400	INPUT_USER_DEACTIVATED	The specified user was deleted
//400	MSG_ID_INVALID	Invalid message ID provided
//400	PEER_ID_INVALID	The provided peer id is invalid
//400	PEER_ID_NOT_SUPPORTED	The provided peer ID is not supported
//400	SEARCH_QUERY_EMPTY	The search query is empty
//400	USER_ID_INVALID	The provided user ID is invalid

//  messages.search#c352eec flags:# peer:InputPeer q:string from_id:flags.0?InputPeer top_msg_id:flags.1?int filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:int = messages.Messages;
//  messages.search#39e9ea0 flags:# peer:InputPeer q:string from_id:flags.0?InputUser filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesSearch %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//var err error
	messages := mtproto.NewTLMessagesMessages(&mtproto.Messages_Messages{
		Messages:       []*mtproto.Message{},
		Chats:          []*mtproto.Chat{},
		Users:          []*mtproto.User{},
		Inexact:        false,
		Count:          0,
		NextRate:       0,
		OffsetIdOffset: 0,
		Pts:            0,
	})
	//
	//if len(request.Q) == 0 {
	//    //err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_SEARCH_QUERY_EMPTY, fmt.Sprintf("peer:%s", request.Peer.ClassName))
	//    //log.Errorf("MessagesSearch %v, request: %v error:%s", md, request, err.Error())
	//    log.Warnf("MessagesSearch %v, request: %v empty", md, request)
	//    return messages.To_Messages_Messages(), nil
	//}
	//
	////TODO:(Coder) Impl MessagesSearch logic
	//inputPeer := input.MakeInputPeer(request.Peer)
	//if inputPeer == nil {
	//    err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_INPUT_CONSTRUCTOR_INVALID, fmt.Sprintf("peer:%s", request.Peer.ClassName))
	//    log.Errorf("MessagesSearch %v, request: %v error:%s", md, request, err.Error())
	//    return nil, err
	//}
	//if inputPeer.GetPeerType() == constants.PeerTypeSelf {
	//    inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
	//        UserId:                 md.UserId,
	//        AccessHash:             request.Peer.AccessHash,
	//    }).To_InputPeer())
	//} else if inputPeer.GetPeerType() == constants.PeerTypeChat {
	//    inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
	//        ChannelId:                 request.Peer.ChatId,
	//        AccessHash:             request.Peer.AccessHash,
	//    }).To_InputPeer())
	//}
	//
	//switch inputPeer.GetPeerType() {
	//case constants.PeerTypeEmpty:
	//case constants.PeerTypeChat:
	//case constants.PeerTypeUser, constants.PeerTypeSelf:
	//    var accessHashStr string
	//    accessHashStr, err = s.AuthAccessHash.GetAccessHash(constants.AccessHashTypeUser, fmt.Sprintf("%d",inputPeer.GetPeerId()))
	//    if err != nil {
	//        err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//        log.Errorf("MessagesSearch %v, request: %v error:%s", md, request, err.Error())
	//        return nil, err
	//    }
	//    if strings.Compare(accessHashStr, fmt.Sprintf("%d", request.Peer.AccessHash)) != 0 {
	//        err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_NO_PROMISSION)
	//        log.Errorf("MessagesSearch %v, request: %v error:%s", md, request, err.Error())
	//        return nil, err
	//    }
	//
	//case constants.PeerTypeChannel:
	//case constants.PeerTypeUserFromMessage, constants.PeerTypeChannelFromMessage:
	//    err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_NO_PROMISSION, fmt.Sprintf("peer:%s", request.Peer.ClassName))
	//    log.Errorf("MessagesSearch %v, request: %v error:%s", md, request, err.Error())
	//    return nil, err
	//}
	//
	//messagesFilter := input.MakeMessagesFilterType(request.Filter)
	//switch messagesFilter {
	//case constants.InputMessagesFilterEmpty:
	//case constants.InputMessagesFilterPhotos:
	//case constants.InputMessagesFilterVideo:
	//case constants.InputMessagesFilterPhotoVideo:
	//case constants.InputMessagesFilterDocument:
	//case constants.InputMessagesFilterUrl:
	//case constants.InputMessagesFilterGif:
	//case constants.InputMessagesFilterVoice:
	//case constants.InputMessagesFilterMusic:
	//case constants.InputMessagesFilterChatPhotos:
	//case constants.InputMessagesFilterPhoneCalls:
	//case constants.InputMessagesFilterRoundVoice:
	//case constants.InputMessagesFilterRoundVideo:
	//case constants.InputMessagesFilterMyMentions:
	//case constants.InputMessagesFilterGeo:
	//case constants.InputMessagesFilterContacts:
	//case constants.InputMessagesFilterPinned:
	//case constants.InputMessagesFilterPhotoVideoDocuments:
	//case constants.InputMessagesFilterUnknown:
	//    err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_INPUT_CONSTRUCTOR_INVALID, fmt.Sprintf("peer:%s", request.Peer.ClassName))
	//    log.Errorf("MessagesSearch %v, request: %v error:%s", md, request, err.Error())
	//    return nil, err
	//}
	//
	log.Debugf("MessagesSearch %v, request: %v reply ok!!!!!!!!!!!", md, request)
	return messages.To_Messages_Messages(), nil

}
