/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.getWebFile_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
//
func (s *UploadServiceImpl) UploadGetWebFile(ctx context.Context, request *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UploadGetWebFile %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLUploadWebFile(&mtproto.Upload_WebFile{
		Size_:    0,
		FileType: mtproto.NewTLStorageFileUnknown(nil).To_Storage_FileType(),
	}).To_Upload_WebFile(), nil
}
