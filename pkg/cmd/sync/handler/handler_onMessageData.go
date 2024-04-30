package handler

import (
	"github.com/gogo/protobuf/proto"
	"novachat_engine/pkg/cmd/sync/messageProducer"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *Handler) OnMessageData(key string, value []byte) error {
	var err error
	updateData := &syncService.UpdateData{}
	switch key {
	case messageProducer.SyncUpdatesKey:
		err = proto.Unmarshal(value, updateData)
		if err != nil {
			log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
		} else {
			err = m.syncUpdate(updateData)
		}
	case messageProducer.PushUpdatesKey:
		err = proto.Unmarshal(value, updateData)
		if err != nil {
			log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
		} else {
			if err = m.pushUpdate(updateData); err != nil {
				log.Errorf("sync OnMessageData PushUpdatesKey value:%s error:%s", string(value), err.Error())
			}
		}
	default:
		log.Fatalf("sync OnMessageData key:%s unknown value:`%s`", key, string(value))
	}
	return err
}
