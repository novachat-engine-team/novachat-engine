/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.saveBigFilePart_handler.go
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

//  upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
//
func (s *UploadServiceImpl) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UploadSaveBigFilePart %v,", metadata.RpcMetaDataDebug(md))

	var err error
	if request.FileTotalParts == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD_FILE_PARTS_INVALID)
		log.Errorf("UploadSaveBigFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	if len(request.Bytes) <= 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD_FILE_PART_EMPTY)
		log.Errorf("UploadSaveBigFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	if request.FilePart < 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD_FILE_PART_INVALID)
		log.Errorf("UploadSaveBigFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	ok, err := s.uploadCore.UploadSaveMultiPartFileData(&sfsService.UploadSaveMultiPartFileData{
		AuthKeyId:      md.AuthKeyId,
		FileId:         request.FileId,
		FilePart:       request.FilePart,
		FileTotalParts: request.FileTotalParts,
		Bytes:          request.Bytes,
	})
	if err != nil {
		log.Errorf("UploadSaveBigFilePart %v, fileId:%d filePart:%d error:%s", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, err.Error())
		return nil, err
	}

	log.Debugf("UploadSaveBigFilePart %v, fileId:%d filePart:%d value:%v reply ok!!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request.FileId, request.FilePart, mtproto.ToBool(ok))
	return ok, nil
}
