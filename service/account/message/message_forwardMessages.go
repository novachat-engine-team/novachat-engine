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
	"novachat_engine/mtproto"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"novachat_engine/service/rpc_context"
	"time"
)

func (c *Core) makeForwardMessages(userId int64, messageList []*mtproto.Message, toPeer *input.InputPeer, silent bool, layer int32) ([]*mtproto.Message, error) {

	userIdList := make([]int64, 0, len(messageList))
	for _, v := range messageList {
		if v.FromId90DDDC1171 != 0 {
			userIdList = append(userIdList, constants.PeerTypeFromUserIDType32(v.FromId90DDDC1171).ToInt())
		} else {
			if v.FromId286FA604119 != nil && v.FromId286FA604119.UserId != 0 {
				userIdList = append(userIdList, constants.PeerTypeFromUserIDType32(v.FromId286FA604119.UserId).ToInt())
			} else {
				log.Warnf("makeForwardMessages v.FromId58AE39C9122 is not PeerUser value:%+v  message:%+v", v.FromId286FA604119, v)
			}
		}
	}

	userList, err := c.accountUsersCore.GetUserList(userId, userIdList, layer)
	if err != nil {
		log.Warnf("makeForwardMessages GetUserList user:%d userIdList:%+v error:%s", userIdList, userIdList, err.Error())
	}
	userNameMap := make(map[int32]*mtproto.User, len(userIdList))
	for _, v := range userList {
		userNameMap[v.Id] = v
	}

	channelId := int32(0)
	fromName := ""
	forwardMessageList := make([]*mtproto.Message, 0, len(messageList))
	for _, v := range messageList {
		sendUserId := v.FromId90DDDC1171
		if sendUserId == 0 {
			if v.FromId286FA604119 != nil && v.FromId286FA604119.UserId != 0 {
				sendUserId = v.FromId286FA604119.UserId
			} else {
				log.Warnf("makeForwardMessages v.FromId286FA604119 is not PeerUser value:%+v  message:%+v", v.FromId286FA604119, v)
				continue
			}
		}

		if v.PeerId != nil && v.PeerId.ChannelId != 0 {
			channelId = v.PeerId.ChannelId
		} else {
			if v.ToId != nil && v.ToId.ChannelId != 0 {
				channelId = v.ToId.ChannelId
			}
		}

		u, exists := userNameMap[sendUserId]
		if exists == false {
			continue
		}
		fromName = u.FirstName
		if len(u.LastName) > 0 {
			fromName += " " + u.LastName
		}

		message := mtproto.NewTLMessage(&mtproto.Message{
			Out:               true,
			Silent:            silent,
			FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer(),
			FromId90DDDC1171:  constants.PeerTypeFromUserIDType(userId).ToInt32(),
			Date:              int32(time.Now().Unix()),
			PeerId:            toPeer.ToPeer(),
			Message:           v.Message,
			Media:             v.Media,
			Entities:          v.Entities,
			ToId:              toPeer.ToPeer(),
			ReplyToMsgId:      0,
			FwdFrom: mtproto.NewTLMessageFwdHeader(&mtproto.MessageFwdHeader{
				FromId5F777DCE119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: sendUserId}).To_Peer(),
				FromName:          fromName,
				Date:              v.Date,
				ChannelPost:       0,
				PostAuthor:        "",
				SavedFromPeer:     nil,
				SavedFromMsgId:    0,
				PsaType:           "",
				FromIdFADFF4AC71:  sendUserId,
				ChannelId:         channelId,
			}).To_MessageFwdHeader(),
		})

		forwardMessageList = append(forwardMessageList, message.To_Message())
	}

	return forwardMessageList, nil
}

func (c *Core) ForwardMessages(userId int64,
	authKeyId int64,
	fromPeer *input.InputPeer,
	toPeer *input.InputPeer,
	idList []int32, randomIdList []int64,
	silent bool,
	layer int32) (*mtproto.Updates, error) {

	var messageList []*mtproto.Message
	var err error
	messageList, err = c.GetMessages(userId, idList, false)
	if err != nil {
		log.Errorf("ForwardMessages GetMessages userId:%d fromPeer:%v idList:%+v error:%s", userId, fromPeer, idList, err.Error())
		return nil, err
	}

	forwardMessageList, err := c.makeForwardMessages(userId, messageList, toPeer, silent, layer)
	if err != nil {
		log.Errorf("ForwardMessages makeForwardMessages userId:%d fromPeer:%v idList:%+v error:%s", userId, fromPeer, idList, err.Error())
		return nil, err
	}

	reqSendMessages := &msgClient.SendMessages{
		AuthKeyId:  authKeyId,
		FromUserId: userId,
		PeerId:     toPeer.GetPeerId(),
		PeerType:   toPeer.GetPeerType().ToInt32(),
		DataList:   make([]*msgClient.SendMessageData, 0, len(forwardMessageList)),
	}
	for idx, v := range forwardMessageList {
		reqSendMessages.DataList = append(reqSendMessages.DataList, &msgClient.SendMessageData{
			RandomId:         randomIdList[idx],
			Message:          v,
			ReplyToMessageId: 0,
		})
	}
	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
	)
	defer cancel()
	updates, err := msgClient.GetMsgClient().ReqSendMessages(ctx, reqSendMessages)
	if err != nil {
		log.Errorf("ForwardMessages ReqSendMessages userId:%d fromPeer:%v idList:%+v error:%s", userId, fromPeer, idList, err.Error())
		return nil, err
	}

	log.Infof("ForwardMessages ReqSendMessages userId:%d fromPeer:%v idList:%+v updates:%+v", userId, fromPeer, idList, updates)
	return updates, nil
}
