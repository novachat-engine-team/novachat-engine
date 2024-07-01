/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package handler

import (
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *Handler) pushUpdates(updateDataList []*syncService.UpdateData) error {
	log.Debugf("pushUpdates updateDataList:%+v", updateDataList)
	var err error
	for idx, updateData := range updateDataList {
		log.Debugf("pushUpdates userId:%d", updateData.UserId)
		err = m.processUpdates(updateData.UserId, updateData.Updates)
		if err == nil {
			go func(v *syncService.UpdateData) {
				m.update(v, true)
			}(updateDataList[idx])
		} else {
			log.Errorf("processUpdates error:%s", err.Error())
			break
		}
	}
	return err
}

func (m *Handler) pushUpdate(updateData *syncService.UpdateData) error {
	err := m.processUpdates(updateData.UserId, updateData.Updates)

	if err == nil {
		go func() {
			m.update(updateData, true)
		}()
	}
	return err
}
