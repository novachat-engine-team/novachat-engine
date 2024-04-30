package rpc

import (
	"context"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
)

func (m *SfsServiceImpl) ReqGetDocument(ctx context.Context, req *sfsService.GetDocument) (*sfsService.DocumentInfo, error) {
	documentInfo, err := m.core.GetDocument(req.FileId, req.LocalId)
	if err != nil {
		log.Warnf("ReqGetDocument error:%s", err.Error())
	}

	return utils.ToDocument(documentInfo), nil
}
