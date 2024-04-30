package handler

import (
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
)

func (m *Handler) pushUpdate(updateData *syncService.UpdateData) error {
	err := m.processUpdates(updateData.UserId, updateData.Updates)

	go func() {
		m.update(updateData, true)
	}()
	return err
}
