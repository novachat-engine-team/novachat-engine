/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/3 15:57
 * @File : reqUploadSavePartFileData_handler.go
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *SfsServiceImpl) ReqUploadSavePartFileData(ctx context.Context, req *sfsService.UploadSavePartFileData) (*types.Any, error) {

	err := m.partCacheFS.WriteFileData(req.AuthKeyId, req.FileId, req.FilePart, req.Bytes)
	if err != nil {
		log.Warnf("ReqUploadSavePartFileData fileId:%d filePart:%d error:%s", req.FileId, req.FilePart, err.Error())
		return types.MarshalAny(mtproto.ToMTBool(false))
	}

	return types.MarshalAny(mtproto.ToMTBool(true))
}
