/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package inboxes

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

func (c *Core) SendMessages(userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData, globalMessageIdList []int64) ([]int32, error) {

	log.Debugf("inboxes SendMessages userId:%d peerId:%d peerType:%v len()=%d", userId, peerId, peerType, len(list))

	st := options.Transaction()
	st.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	st.SetReadConcern(readconcern.Majority())

	var ptsList []int32
	var sented bool
	var err error
	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction(st)
		if err != nil {
			log.Fatalf("SendMessages StartTransaction error:%s", err.Error())
			return err
		}

		ptsList, sented, err = c.messageCore.FillMessageBox(sessionContext, userId, peerId, list, message.InboxType)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)

			log.Errorf("SendMessages FillMessageBox false error:%s", err.Error())
			return err
		}

		if sented {
			sessionContext.AbortTransaction(sessionContext)
			return nil
		}

		err = c.messageCore.SendMessages(sessionContext, userId, peerId, peerType, list, message.InboxType, &globalMessageIdList)
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

		err = c.updateInboxConversation(sessionContext, userId, peerId, peerType, list, pts)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)

			log.Errorf("SendMessages updateInboxConversation userId:%v, peerId:%v peerType:%v error:%+v", userId, peerId, peerType, err.Error())
			return err
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Infof("inboxes SendMessages userId:%d peerId:%d peerType:%v len()=%d OK!!!", userId, peerId, peerType, len(list))
	return ptsList, nil
}

func (c *Core) SendChannelMessages(userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData, pts int32) error {
	log.Debugf("inboxes SendChannelMessages userId:%d peerId:%d peerType:%v len()=%d", userId, peerId, peerType, len(list))

	st := options.Transaction()
	st.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	st.SetReadConcern(readconcern.Majority())

	err := c.updateInboxConversation(context.Background(), userId, peerId, peerType, list, pts)
	if err != nil {
		log.Errorf("inboxes SendChannelMessages updateInboxConversation userId:%v, peerId:%v peerType:%v error:%+v", userId, peerId, peerType, err.Error())
		return err
	}
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Infof("inboxes SendChannelMessages userId:%d peerId:%d peerType:%v len()=%d OK!!!", userId, peerId, peerType, len(list))
	return nil
}
