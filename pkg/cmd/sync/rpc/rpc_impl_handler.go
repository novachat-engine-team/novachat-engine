/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/10 13:38
 * @File :
 */

package rpc

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cmd/sync/messageProducer"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

func (m *Impl) ReqSyncUpdate(ctx context.Context, request *syncService.SyncUpdate) (*types.Any, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("ReqSyncUpdate context done")
	default:
		log.Debugf("ReqSyncUpdate userId:%d authKeyId:%d ignoreAuthKeyId:%d updates:%+v",
			request.UserId,
			request.AuthKeyId,
			request.IgnoreAuthKeyId,
			request.Updates)
		break
	}

	data, _ := proto.Marshal(&syncService.UpdateData{
		UserId:          request.UserId,
		AuthKeyId:       request.AuthKeyId,
		IgnoreAuthKeyId: request.IgnoreAuthKeyId,
		Updates:         request.Updates,
	})

	var producerType messageProducer.ProducerType
	switch constants.PeerType(request.PeerType) {
	case constants.PeerTypeUser, constants.PeerTypeSelf:
		producerType = messageProducer.UsersProducer
	case constants.PeerTypeChat:
		producerType = messageProducer.ChatsProducer
	case constants.PeerTypeChannel:
		producerType = messageProducer.ChannelsProducer
	default:
		return nil, fmt.Errorf("ReqSyncUpdate PeerType:%v error", request.PeerType)
	}

	key := messageProducer.SyncUpdatesKey
	if request.Push {
		key = messageProducer.PushUpdatesKey
	}
	partition, offset, err := m.producer.SendMessage(producerType, key, data)
	if err != nil {
		return nil, fmt.Errorf("ReqSyncUpdate producer SendMessage error:%s", err.Error())
	}

	log.Infof("ReqSyncUpdate partition:%d, offset:%d", partition, offset)
	return types.MarshalAny(mtproto.ToMTBool(true))
}

func (m *Impl) ReqSyncUpdates(ctx context.Context, request *syncService.SyncUpdates) (*types.Any, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("ReqSyncUpdates context done")
	default:
		log.Debugf("ReqSyncUpdates updateDataList:%+v", request.UpdateDataList)
		break
	}

	if len(request.UpdateDataList) == 0 {
		log.Warnf("ReqSyncUpdates UpdateDataList is empty")
		return types.MarshalAny(mtproto.ToMTBool(false))
	}

	key := messageProducer.SyncUpdatesListKey
	if request.Push {
		key = messageProducer.PushUpdatesListKey
	}

	var producerType messageProducer.ProducerType
	switch constants.PeerType(request.PeerType) {
	case constants.PeerTypeUser, constants.PeerTypeSelf:
		producerType = messageProducer.UsersProducer
	case constants.PeerTypeChat:
		producerType = messageProducer.ChatsProducer
	case constants.PeerTypeChannel:
		producerType = messageProducer.ChannelsProducer
	default:
		return nil, fmt.Errorf("ReqSyncUpdates PeerType:%v error", request.PeerType)
	}

	dataList := make([]*syncService.UpdateData, 0, len(request.UpdateDataList))
	for _, v := range request.UpdateDataList {
		//data, _ := proto.Marshal(v)
		//dataList = append(dataList, data)
		dataList = append(dataList, v)
	}
	data, _ := jsoniter.Marshal(dataList)

	_, _, err := m.producer.SendMessage(producerType, key, data)
	//err := m.producer.SendMessages(producerType, key, dataList)
	if err != nil {
		return nil, fmt.Errorf("ReqSyncUpdates producer SendMessages error:%s", err.Error())
	}

	log.Infof("ReqSyncUpdates updateDataList:%+v", request.UpdateDataList)
	return types.MarshalAny(mtproto.ToMTBool(true))
}
