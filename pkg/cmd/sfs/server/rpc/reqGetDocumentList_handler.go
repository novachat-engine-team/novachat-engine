package rpc

import (
	"context"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
)

func (m *SfsServiceImpl) ReqGetDocumentList(ctx context.Context, req *sfsService.GetDocumentList) (*sfsService.DocumentInfoList, error) {
	documentInfoList, err := m.core.GetDocumentList(req.FileId)
	if err != nil {
		log.Warnf("ReqGetDocumentList error:%s", err.Error())
	}

	ret := &sfsService.DocumentInfoList{
		Values: make([]*sfsService.DocumentInfo, 0, len(documentInfoList)),
	}

	for _, v := range documentInfoList {
		ret.Values = append(ret.Values, utils.ToDocument(v))
	}
	return ret, nil
}
