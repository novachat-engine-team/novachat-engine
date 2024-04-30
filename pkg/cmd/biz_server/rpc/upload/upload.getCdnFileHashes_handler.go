/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.getCdnFileHashes_handler.go
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

//  upload.getCdnFileHashes#4da54231 file_token:bytes offset:int = Vector<FileHash>;
//  upload.getCdnFileHashes#f715c87b file_token:bytes offset:int = Vector<CdnFileHash>;
//
func (s *UploadServiceImpl) UploadGetCdnFileHashes(ctx context.Context, request *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Response_UploadGetCdnFileHashes, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UploadGetCdnFileHashes %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl UploadGetCdnFileHashes logic

	return nil, fmt.Errorf("%s", "Not impl UploadGetCdnFileHashes")
}
