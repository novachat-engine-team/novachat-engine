/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/input"
)

func (c *Core) ReplyMessage(userId int64, replyToMsgId int32, inputPeer *input.InputPeer) (*mtproto.TLMessageReplyHeader, error) {

	if replyToMsgId == 0 {
		return nil, nil
	}

	replyTo := mtproto.NewTLMessageReplyHeader(&mtproto.MessageReplyHeader{
		ReplyToMsgId: replyToMsgId,
		//ReplyToPeerId:        input.MakePeer(inputPeer.ToInputPeer(), 0),
		ReplyToTopId: 0,
	})

	//if inputPeer.GetPeerType() == constants.PeerTypeSelf || inputPeer.GetPeerType() == constants.PeerTypeUser {
	//	//if inputPeer.GetPeerId() == userId {
	//	//
	//	//} else {
	//	//	peerMsgId, err := dbProxyClient.GetDBProxyClient().MessageMessageIdToPeerMessageId(context.Background(), &dbproxy_pb.MessageMessageIdToPeerMessageId{
	//	//		UserId:    md.UserId,
	//	//		MessageId: replyToMsgId,
	//	//		PeerId:    inputPeer.GetPeerId(),
	//	//		PeerType:  inputPeer.GetPeerType().ToInt32(),
	//	//	})
	//	//	if err != nil {
	//	//		log.Errorf("MessageMessageIdToPeerMessageId userId:%d replyToMsgId:%d error:%s", md.UserId, replyToMsgId, err.Error())
	//	//		//return nil, 0, err
	//	//	}
	//	//	if peerMsgId != nil && peerMsgId.Value != 0 {
	//	//		replyToMessageId = peerMsgId.Value
	//	//		message.SetReplyToMsgId(replyToMsgId)
	//	//	}
	//		//messageList, err := dbProxyClient.GetDBProxyClient().MessageGetMessagesByMessageIds(context.Background(), &dbproxy_pb.MessageGetMessagesByMessageIds{
	//		//	Value: []*dbproxy_pb.MessageMessageData{{
	//		//		UserId:        md.UserId,
	//		//		MessageIdList: []int32{replyToMsgId},
	//		//		InputPeer:     mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: md.UserId}).To_InputPeer(),
	//		//	}},
	//		//	UserId: md.UserId,
	//		//})
	//		//if err != nil {
	//		//	log.Errorf("MessageGetMessagesByMessageIds userId:%d replyToMsgId:%d error:%s", md.UserId, replyToMsgId, err.Error())
	//		//	return nil, nil, err
	//		//}
	//		//if messageList != nil && len(messageList.MessageList) > 0 {
	//		//	replyToMessage = messageList.MessageList[0]
	//		//}
	//	//}
	//} else {
	//	message.SetReplyToMsgId(replyToMsgId)
	//}
	// //
	//	replyTo.SetReplyToPeerId(mtproto.NewTLPeerUser(&mtproto.Peer{UserId: inputPeer.GetPeerId()}).To_Peer())
	//} else if inputPeer.GetPeerType() == constants.PeerTypeChat {
	//	replyTo.SetReplyToPeerId(mtproto.NewTLPeerChat(&mtproto.Peer{ChatId: inputPeer.GetPeerId()}).To_Peer())
	//} else if inputPeer.GetPeerType() == constants.PeerTypeChannel {
	//	replyTo.SetReplyToPeerId(mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: inputPeer.GetPeerId()}).To_Peer())
	//}

	return replyTo, nil
}
