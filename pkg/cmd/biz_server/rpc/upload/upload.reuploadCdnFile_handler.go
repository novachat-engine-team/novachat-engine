/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.reuploadCdnFile_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  upload.reuploadCdnFile#9b2754a8 file_token:bytes request_token:bytes = Vector<FileHash>;
//  upload.reuploadCdnFile#1af91c09 file_token:bytes request_token:bytes = Vector<CdnFileHash>;
//
func (s *UploadServiceImpl) UploadReuploadCdnFile(ctx context.Context, request *mtproto.TLUploadReuploadCdnFile) (*mtproto.Response_UploadReuploadCdnFile, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UploadReuploadCdnFile %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl UploadReuploadCdnFile logic

	return nil, fmt.Errorf("%s", "Not impl UploadReuploadCdnFile")
}
