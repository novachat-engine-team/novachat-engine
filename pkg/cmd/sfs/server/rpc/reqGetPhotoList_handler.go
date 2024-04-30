package rpc

import (
	"context"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
)

func (m *SfsServiceImpl) ReqGetPhotoList(ctx context.Context, req *sfsService.GetPhotoList) (*sfsService.PhotoInfoList, error) {
	photoInfoList, err := m.core.GetPhotoList(req.PhotoIdList)
	if err != nil {
		log.Warnf("ReqGetPhotoList error:%s", err.Error())
	}

	ret := &sfsService.PhotoInfoList{
		Values: make([]*sfsService.PhotoInfo, 0, len(photoInfoList)),
	}

	for _, v := range photoInfoList {
		ret.Values = append(ret.Values, utils.ToPhoto(v))
	}
	return ret, nil

}
