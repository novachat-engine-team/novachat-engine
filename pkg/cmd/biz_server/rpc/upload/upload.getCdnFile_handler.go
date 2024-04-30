/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.getCdnFile_handler.go
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

//  upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
//
func (s *UploadServiceImpl) UploadGetCdnFile(ctx context.Context, request *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UploadGetCdnFile %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl UploadGetCdnFile logic

	return nil, fmt.Errorf("%s", "Not impl UploadGetCdnFile")
}
