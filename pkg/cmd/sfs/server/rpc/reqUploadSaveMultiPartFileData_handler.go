/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/3 15:57
 * @File : reqUploadSaveMultiPartFileData_handler.go
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *SfsServiceImpl) ReqUploadSaveMultiPartFileData(ctx context.Context, req *sfsService.UploadSaveMultiPartFileData) (*types.Any, error) {

	err := m.partCacheFS.WriteFilePartData(req.AuthKeyId, req.FileId, req.FilePart, req.FileTotalParts, req.Bytes)
	if err != nil {
		log.Warnf("ReqUploadSaveMultiPartFileData fileId:%d filePart:%d fileTotalParts:%d error:%s", req.FileId, req.FilePart, req.FileTotalParts, err.Error())
		return types.MarshalAny(mtproto.ToMTBool(false))
	}

	return types.MarshalAny(mtproto.ToMTBool(true))
}
