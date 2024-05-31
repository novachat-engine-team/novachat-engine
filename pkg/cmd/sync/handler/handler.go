/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/24 16:09
 * @File : session_handler.go
 */

package handler

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	sessionClient "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/cmd/sync/conf"
	"novachat_engine/pkg/cmd/sync/messageProducer"
	"novachat_engine/pkg/cmd/sync/npns"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/mq"
	"novachat_engine/pkg/snowflake"
	"novachat_engine/pkg/util"
	accountUpdate "novachat_engine/service/account/update"
	"novachat_engine/service/core/update"
)

type Handler struct {
	conf                 *conf.Conf
	consumerList         []mq.Consumer
	npnsHandler          *npns.Npns
	accountUpdateCore    *accountUpdate.Core
	updateKeyNode        *snowflake.Node
	updateChannelKeyNode *snowflake.Node
}

func NewHandler(conf *conf.Conf) *Handler {
	node, err := snowflake.NewNode(int64(conf.Server.ServerId))
	if err != nil {
		panic(err.Error())
	}
	channelNode, err := snowflake.NewNode(int64(conf.Server.ServerId))
	if err != nil {
		panic(err.Error())
	}

	s := &Handler{
		conf:                 conf,
		consumerList:         make([]mq.Consumer, 0, 3),
		npnsHandler:          npns.NewNpns(conf),
		accountUpdateCore:    accountUpdate.NewCore(update.NewCore(conf.MongoDB), nil),
		updateKeyNode:        node,
		updateChannelKeyNode: channelNode,
	}

	if index := util.Index(conf.Consumer, func(idx int) bool {
		return conf.Consumer[idx].Name == messageProducer.UsersProducer
	}); index > 0 {
		consumer := mq.NewKafkaConsumer(&conf.Consumer[index].MessageConsumer, mq.WithConsumerCB(s))
		s.consumerList = append(s.consumerList, consumer)
		go consumer.Run()
	}
	if index := util.Index(conf.Consumer, func(idx int) bool {
		return conf.Consumer[idx].Name == messageProducer.ChatsProducer
	}); index > 0 {
		consumer := mq.NewKafkaConsumer(&conf.Consumer[index].MessageConsumer, mq.WithConsumerCB(s))
		s.consumerList = append(s.consumerList, consumer)
		go consumer.Run()
	}
	if index := util.Index(conf.Consumer, func(idx int) bool {
		return conf.Consumer[idx].Name == messageProducer.ChannelsProducer
	}); index > 0 {
		consumer := mq.NewKafkaConsumer(&conf.Consumer[index].MessageConsumer, mq.WithConsumerCB(s))
		s.consumerList = append(s.consumerList, consumer)
		go consumer.Run()
	}

	return s
}

func (m *Handler) Close() {
	for _, v := range m.consumerList {
		v.Close()
	}

	m.npnsHandler.Close()
}

func (m *Handler) update(updateData *syncService.UpdateData, push bool) {
	userSessionInfo, err := authClient.GetAuthClientByAuthKey(updateData.UserId).
		ReqUserSession(context.Background(), &authClient.UserSession{
			UserIdList: []int64{updateData.UserId},
		})
	if err != nil {
		log.Errorf("sync handler update error:%s", err.Error())
		return
	}

	log.Debugf("update userId:%d list:%+v", updateData.UserId, userSessionInfo.SessionInfoList)
	var ok mtproto.Bool
	var resp *types.Any
	for _, sessionInfo := range userSessionInfo.SessionInfoList {
		if len(sessionInfo.SessionId) == 0 {
			continue
		}

		any, _ := types.MarshalAny(updateData.Updates)
		cli := sessionClient.GetSessionClientByKey(fmt.Sprintf("%d", sessionInfo.AuthKeyId))
		if cli == nil {
			log.Warnf("sync handler update GetSessionClientByKey authKeyId:%d error:%s", sessionInfo.AuthKeyId, err.Error())
			if push && sessionInfo.Device != nil {
				m.push(updateData, sessionInfo.Device)
			}
			continue
		}
		resp, _ = cli.PushUpdates(context.Background(), &sessionClient.PushUpdatesEvent{
			AuthKeyId: sessionInfo.AuthKeyId,
			SessionId: sessionInfo.SessionId,
			Updates:   any,
		})
		types.UnmarshalAny(resp, &ok)
		if mtproto.ToBool(&ok) == false && push && sessionInfo.Device != nil {
			m.push(updateData, sessionInfo.Device)
		}
	}
}

func (m *Handler) push(updateData *syncService.UpdateData, device *authClient.Device) {
	//TODO:(Coderxw npns)
	//m.npnsHandler.OnMessageData()
}
