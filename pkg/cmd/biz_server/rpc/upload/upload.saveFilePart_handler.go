/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.saveFilePart_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
)

//  upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
//
func (s *UploadServiceImpl) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UploadSaveFilePart %v, fileId:%d filePart:%d", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart)

	var err error
	if len(request.Bytes) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD_FILE_PART_EMPTY)
		log.Errorf("UploadSaveFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	if request.FilePart < 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD_FILE_PART_INVALID)
		log.Errorf("UploadSaveFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	ok, err := s.uploadCore.SavePartFileData(&sfsService.UploadSavePartFileData{
		AuthKeyId: md.AuthKeyId,
		FileId:    request.FileId,
		FilePart:  request.FilePart,
		Bytes:     request.Bytes,
	})

	if err != nil {
		log.Errorf("UploadSaveFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	log.Debugf("UploadSaveFilePart %v, fileId:%d filePart:%d value:%v reply ok!!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, mtproto.ToBool(ok))
	return ok, nil
}
