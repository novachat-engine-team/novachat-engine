/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package outboxes

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	"novachat_engine/service/mgo"
)

func (c *Core) SendMessages(userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData, draft bool) ([]int32, []int64, error) {

	st := options.Transaction()
	st.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	st.SetReadConcern(readconcern.Majority())

	var ptsList []int32
	var globalMessageIdList []int64
	var sented bool
	var err error
	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction(st)
		if err != nil {
			log.Fatalf("SendMessages StartTransaction error:%s", err.Error())
			return err
		}

		ptsList, sented, err = c.messageCore.FillMessageBox(sessionContext, userId, peerId, list, message.OutboxType)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("SendMessages FillMessageBox false error:%s", err.Error())
			return err
		}
		if sented {
			sessionContext.AbortTransaction(sessionContext)
			return nil
		}

		globalMessageIdList = make([]int64, 0, len(ptsList))
		switch peerType {
		case constants.PeerTypeUser, constants.PeerTypeSelf:
			err = c.messageCore.SendMessages(sessionContext, userId, peerId, peerType, list, message.OutboxType, &globalMessageIdList)
			//case constants.PeerTypeChannel, constants.PeerTypeChat:
			//	err = c.messageCore.SendChannelMessages(context.Background(), userId, peerId, peerType, list, &globalMessageIdList)
		}
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("SendMessages SendMessages userId:%v, peerId:%v peerType:%v error:%+v", userId, peerId, peerType, err.Error())
			return err
		}
		pts := ptsList[0]
		for _, ptsV := range ptsList {
			if pts < ptsV {
				pts = ptsV
			}
		}

		err = c.updateOutboxConversation(sessionContext, userId, peerId, peerType, list, draft, pts)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("SendMessages updateOutboxConversation userId:%v, peerId:%v peerType:%v error:%+v", userId, peerId, peerType, err.Error())
			return err
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		return nil, nil, err
	}
	return ptsList, globalMessageIdList, nil
}

func (c *Core) SendChannelMessages(userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData, draft bool) ([]int32, error) {

	st := options.Transaction()
	st.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	st.SetReadConcern(readconcern.Majority())

	var ptsList []int32
	ctx := context.Background()

	var sented bool
	var err error
	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction(st)
		if err != nil {
			log.Fatalf("SendChannelMessages StartTransaction error:%s", err.Error())
			return err
		}

		ptsList, sented, err = c.messageCore.FillChannelMessageBox(ctx, userId, peerId, list)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("SendChannelMessages FillMessageBox false error:%s", err.Error())
			return err
		}
		if sented {
			sessionContext.AbortTransaction(sessionContext)
			return nil
		}

		switch peerType {
		case constants.PeerTypeChannel, constants.PeerTypeChat:
			err = c.messageCore.SendChannelMessages(context.Background(), userId, peerId, peerType, list)
		}
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("SendChannelMessages SendMessages userId:%v, peerId:%v peerType:%v error:%+v", userId, peerId, peerType, err.Error())
			return err
		}
		pts := ptsList[0]
		for _, ptsV := range ptsList {
			if pts < ptsV {
				pts = ptsV
			}
		}

		err = c.updateOutboxConversation(sessionContext, userId, peerId, peerType, list, draft, pts)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("SendChannelMessages updateOutboxConversation userId:%v, peerId:%v peerType:%v error:%+v", userId, peerId, peerType, err.Error())
			return err
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return ptsList, nil
}
