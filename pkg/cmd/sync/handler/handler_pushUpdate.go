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
)

func (m *Handler) pushUpdate(updateData *syncService.UpdateData) error {
	err := m.processUpdates(updateData.UserId, updateData.Updates)

	if err == nil {
		go func() {
			m.update(updateData, true)
		}()
	}
	return err
}
