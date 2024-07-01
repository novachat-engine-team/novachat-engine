/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/10 13:38
 * @File :
 */

package handler

import (
	"github.com/gogo/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/pkg/cmd/sync/messageProducer"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *Handler) OnMessageData(key string, value []byte) error {
	var err error

	switch key {
	case messageProducer.SyncUpdatesKey:
		var updateData syncService.UpdateData
		err = proto.Unmarshal(value, &updateData)
		if err != nil {
			log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
		} else {
			err = m.syncUpdate(&updateData)
		}
	case messageProducer.SyncUpdatesListKey:
		var updateData []*syncService.UpdateData
		err = jsoniter.Unmarshal(value, &updateData)
		if err != nil {
			log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
		} else {
			err = m.syncUpdates(updateData)
		}
	case messageProducer.PushUpdatesKey:
		var updateData syncService.UpdateData
		err = proto.Unmarshal(value, &updateData)
		if err != nil {
			log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
		} else {
			if err = m.pushUpdate(&updateData); err != nil {
				log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
			}
		}
	case messageProducer.PushUpdatesListKey:
		var updateData []*syncService.UpdateData
		err = jsoniter.Unmarshal(value, &updateData)
		if err != nil {
			log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
		} else {
			if err = m.pushUpdates(updateData); err != nil {
				log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
			}
		}
	default:
		log.Fatalf("sync OnMessageData key:%s unknown value:`%s`", key, string(value))
	}
	return err
}
