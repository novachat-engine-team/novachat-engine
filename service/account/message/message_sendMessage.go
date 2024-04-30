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
	"context"
	"github.com/go-redsync/redsync"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	chatClient "novachat_engine/pkg/cmd/chat/rpc_client"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/account/contact"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/common/hash"
	"novachat_engine/service/constants"
	data_users "novachat_engine/service/data/users"
	"novachat_engine/service/input"
	"novachat_engine/service/rpc_context"
	"time"
)

func (c *Core) SendMessage(userId int64, authKeyId int64, request *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {

	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeEmpty {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("SendMessage constants.PeerTypeEmpty userId:%d request.Peer:%+v error:%s", userId, request.Peer, err.Error())
		return nil, err
	}

	var userData *data_users.Users
	if inputPeer.GetPeerType() == constants.PeerTypeSelf {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(userId).ToInt32(),
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())

		blocked, err := c.accountContactCore.CheckBlock(userId, inputPeer.GetPeerId(), int32(time.Now().Unix()))
		if err != nil {
			log.Errorf("SendMessage constants.PeerTypeEmpty userId:%d request.Peer:%+v error:%s", userId, request.Peer, err.Error())
			return nil, err
		}

		if blocked > contact.BlockedNone {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_IS_BLOCKED)
			return nil, err
		}
		userData, err = c.accountUsersCore.UserData(inputPeer.GetPeerId())
		if err != nil {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
			log.Errorf("UserData error:%s", err.Error())
			return nil, err
		}
	} else if inputPeer.GetPeerType() == constants.PeerTypeChat {
		log.Debugf("SendMessage userId:%d chatId", userId, request.Peer.ChatId)
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
			ChannelId:  request.Peer.ChatId,
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())
	} else if inputPeer.GetPeerType() == constants.PeerTypeChannel {
		log.Debugf("SendMessage userId:%d ChannelId", userId, request.Peer.ChannelId)
	}

	mutex := cache.GetRedSyncClient().NewMutex(
		makeMutexMessageName(userId, request.RandomId),
		redsync.SetExpiry(3*time.Second),
		redsync.SetTries(3))
	err := mutex.Lock()
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("SendMessage NewMutex error:%s", err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	var updates *mtproto.Updates
	updates, err = c.makeDuplicateMessage(userId, request.RandomId)
	if err != nil {
		return nil, err
	}
	if updates != nil {
		return updates, nil
	}

	message, err := c.makeMessage(userId, inputPeer, request, userData)
	if err != nil {
		log.Errorf("SendMessage ReplyMessage error:%s", err.Error())
		return nil, err
	}

	replyToMessage, err := c.ReplyMessage(userId, request.ReplyToMsgId, inputPeer)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("SendMessage ReplyMessage error:%s", err.Error())
		return nil, err
	}
	if replyToMessage != nil {
		message.ReplyToMsgId = replyToMessage.GetReplyToMsgId()
		message.ReplyTo = replyToMessage.To_MessageReplyHeader()
	}

	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
	)

	//TODO:Coder
	// sendMessage
	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser:
		updates, err = msgClient.GetMsgClient().ReqSendMessages(ctx, &msgClient.SendMessages{
			AuthKeyId:  authKeyId,
			FromUserId: userId,
			PeerId:     inputPeer.GetPeerId(),
			PeerType:   inputPeer.GetPeerType().ToInt32(),
			DataList: []*msgClient.SendMessageData{{
				RandomId:         request.RandomId,
				Message:          message,
				ReplyToMessageId: request.ReplyToMsgId,
			}},
			ClearDraft: request.ClearDraft,
		})
	case constants.PeerTypeChat:
		updates, err = chatClient.GetChatClientByKeyId(constants.PeerTypeFromChatIDType32(request.Peer.ChatId).ToInt()).
			ReqSendOutboxesMessages(ctx, &chatClient.SendOutboxesMessages{
				Message: &msgClient.SendMessages{
					AuthKeyId:  authKeyId,
					FromUserId: userId,
					PeerId:     inputPeer.GetPeerId(),
					PeerType:   inputPeer.GetPeerType().ToInt32(),
					DataList: []*msgClient.SendMessageData{{
						RandomId:         request.RandomId,
						Message:          message,
						ReplyToMessageId: request.ReplyToMsgId,
					}},
					ClearDraft: request.ClearDraft,
				}})
	case constants.PeerTypeChannel:
		panic("SendMessage Channel")
	}
	cancel()
	if err != nil {
		log.Errorf("SendMessage SendMessages error:%s", err.Error())
		return nil, err
	}
	ptsCount := int32(0)
	pts := int32(0)
	messageId := int32(0)
	for _, up := range updates.Updates {
		if up.ClassName == mtproto.ClassUpdateMessageID {
			messageId = up.To_UpdateMessageID().GetId4E90BFD671()
		} else if up.ClassName == mtproto.ClassUpdateNewMessage {
			updateNewMessage := up.To_UpdateNewMessage()
			pts = updateNewMessage.GetPts()
			ptsCount = updateNewMessage.GetPtsCount()
		}
	}

	c.putDuplicateMessage(userId, request.RandomId, messageId, pts, ptsCount)
	return updates, nil
}

func (c *Core) SendMultiMessage(userId int64, authKeyId int64, request *mtproto.TLMessagesSendMultiMedia, layer int32) (*mtproto.Updates, error) {
	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeSelf {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(userId).ToInt32(),
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())

		blocked, err := c.accountContactCore.CheckBlock(userId, inputPeer.GetPeerId(), int32(time.Now().Unix()))
		if err != nil {
			log.Errorf("SendMultiMessage constants.PeerTypeEmpty userId:%d request.Peer:%+v error:%s", userId, request.Peer, err.Error())
			return nil, err
		}
		if blocked > contact.BlockedNone {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_IS_BLOCKED)
			return nil, err
		}
	} else if inputPeer.GetPeerType() == constants.PeerTypeChat {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
			ChannelId:  request.Peer.ChatId,
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())
	}

	messageList, err := c.makeMultiMedia(userId, authKeyId, layer, inputPeer, request)
	if err != nil {
		log.Errorf("SendMultiMessage makeMultiMedia makeMedia error:%s", err.Error())
		return nil, err
	}

	if len(messageList) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_EMPTY)
	}

	randomIdList := make([]int64, 0, len(request.MultiMedia))
	for _, v := range request.MultiMedia {
		randomIdList = append(randomIdList, v.RandomId)
	}

	mediaListRandomId := hash.HashId64ListNew(randomIdList)
	mutex := cache.GetRedSyncClient().NewMutex(
		makeMutexMessageName(userId, mediaListRandomId),
		redsync.SetExpiry(3*time.Second),
		redsync.SetTries(3))
	err = mutex.Lock()
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("SendMessageMedia NewMutex error:%s", err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	var updates *mtproto.Updates
	updates, err = c.makeDuplicateMessage(userId, mediaListRandomId)
	if err != nil {
		return nil, err
	}
	if updates != nil {
		return updates, nil
	}

	replyToMessage, err := c.ReplyMessage(userId, request.ReplyToMsgId, inputPeer)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("SendMessageMedia ReplyMessage error:%s", err.Error())
		return nil, err
	}
	if replyToMessage != nil {
		messageList[0].ReplyToMsgId = replyToMessage.GetReplyToMsgId()
		messageList[0].ReplyTo = replyToMessage.To_MessageReplyHeader()
	}

	_ = request.ClearDraft

	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
	)

	sendMessageList := make([]*msgClient.SendMessageData, 0, len(messageList))
	for idx := range messageList {
		sendMessageList = append(sendMessageList, &msgClient.SendMessageData{
			RandomId:         randomIdList[idx],
			Message:          messageList[idx],
			ReplyToMessageId: messageList[idx].ReplyToMsgId,
		})
	}

	//TODO:Coder
	// sendMessage
	updates, err = msgClient.GetMsgClient().ReqSendMessages(ctx, &msgClient.SendMessages{
		AuthKeyId:  authKeyId,
		FromUserId: userId,
		PeerId:     inputPeer.GetPeerId(),
		PeerType:   inputPeer.GetPeerType().ToInt32(),
		DataList:   sendMessageList,
		ClearDraft: request.ClearDraft,
	})
	cancel()
	if err != nil {
		log.Errorf("SendMessage SendMessages error:%s", err.Error())
		return nil, err
	}

	ptsCount := int32(0)
	pts := int32(0)
	messageId := int32(0)
	for _, up := range updates.Updates {
		if up.ClassName == mtproto.ClassUpdateMessageID {
			messageId = up.To_UpdateMessageID().GetId4E90BFD671()
		} else if up.ClassName == mtproto.ClassUpdateNewMessage {
			updateNewMessage := up.To_UpdateNewMessage()
			pts = updateNewMessage.GetPts()
			ptsCount = updateNewMessage.GetPtsCount()
		}
	}
	c.putDuplicateMessage(userId, mediaListRandomId, messageId, pts, ptsCount)
	return updates, nil
}

func (c *Core) SendMessageMedia(userId int64, authKeyId int64, request *mtproto.TLMessagesSendMedia, layer int32) (*mtproto.Updates, error) {

	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeSelf {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(userId).ToInt32(),
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())

		blocked, err := c.accountContactCore.CheckBlock(userId, inputPeer.GetPeerId(), int32(time.Now().Unix()))
		if err != nil {
			log.Errorf("SendMessageMedia constants.PeerTypeEmpty userId:%d request.Peer:%+v error:%s", userId, request.Peer, err.Error())
			return nil, err
		}

		if blocked > contact.BlockedNone {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_IS_BLOCKED)
			return nil, err
		}

	} else if inputPeer.GetPeerType() == constants.PeerTypeChat || inputPeer.GetPeerType() == constants.PeerTypeChannel {
		if inputPeer.GetPeerType() == constants.PeerTypeChat {
			inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
				ChannelId:  request.Peer.ChatId,
				AccessHash: request.Peer.AccessHash,
			}).To_InputPeer())
		}
	}

	message, err := c.makeMedia(userId, authKeyId, layer, inputPeer, request)
	if err != nil {
		log.Errorf("SendMessageMedia makeMedia error:%s", err.Error())
		return nil, err
	}

	mutex := cache.GetRedSyncClient().NewMutex(
		makeMutexMessageName(userId, request.RandomId),
		redsync.SetExpiry(3*time.Second),
		redsync.SetTries(3))
	err = mutex.Lock()
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("SendMessageMedia NewMutex error:%s", err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	var updates *mtproto.Updates
	updates, err = c.makeDuplicateMessage(userId, request.RandomId)
	if err != nil {
		return nil, err
	}
	if updates != nil {
		return updates, nil
	}

	replyToMessage, err := c.ReplyMessage(userId, request.ReplyToMsgId, inputPeer)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("SendMessageMedia ReplyMessage error:%s", err.Error())
		return nil, err
	}
	if replyToMessage != nil {
		message.ReplyToMsgId = replyToMessage.GetReplyToMsgId()
		message.ReplyTo = replyToMessage.To_MessageReplyHeader()
	}

	_ = request.ClearDraft

	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
	)

	//TODO:Coder
	// sendMessage
	updates, err = msgClient.GetMsgClient().ReqSendMessages(ctx, &msgClient.SendMessages{
		AuthKeyId:  authKeyId,
		FromUserId: userId,
		PeerId:     inputPeer.GetPeerId(),
		PeerType:   inputPeer.GetPeerType().ToInt32(),
		DataList: []*msgClient.SendMessageData{{
			RandomId:         request.RandomId,
			Message:          message,
			ReplyToMessageId: request.ReplyToMsgId,
		}},
		ClearDraft: request.ClearDraft,
	})
	cancel()
	if err != nil {
		log.Errorf("SendMessageMedia SendMessages error:%s", err.Error())
		return nil, err
	}

	ptsCount := int32(0)
	pts := int32(0)
	messageId := int32(0)
	for _, up := range updates.Updates {
		if up.ClassName == mtproto.ClassUpdateMessageID {
			messageId = up.To_UpdateMessageID().GetId4E90BFD671()
		} else if up.ClassName == mtproto.ClassUpdateNewMessage {
			updateNewMessage := up.To_UpdateNewMessage()
			pts = updateNewMessage.GetPts()
			ptsCount = updateNewMessage.GetPtsCount()
		}
	}
	c.putDuplicateMessage(userId, request.RandomId, messageId, pts, ptsCount)
	return updates, nil
}

func (c *Core) SendUploadMedia(authKeyId int64, layer int32, media *mtproto.InputMedia) (*mtproto.MessageMedia, error) {
	mediaList, err := c.makeInputMedia(authKeyId, layer, []*mtproto.InputMedia{media})
	if err != nil {
		log.Errorf("SendUploadMedia makeInputMedia error:%s", err.Error())
		return nil, err
	}

	if len(mediaList) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MEDIA_INVALID)
	}
	return mediaList[0], nil
}
