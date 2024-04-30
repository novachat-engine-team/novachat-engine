package rpc

import (
	"context"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
)

func (m *SfsServiceImpl) ReqGetPhoto(ctx context.Context, req *sfsService.GetPhoto) (*sfsService.PhotoInfo, error) {
	photoInfo, err := m.core.GetPhoto(req.PhotoId, req.LocalId)
	if err != nil {
		log.Warnf("ReqGetPhoto error:%s", err.Error())
	}

	return utils.ToPhoto(photoInfo), nil
}
