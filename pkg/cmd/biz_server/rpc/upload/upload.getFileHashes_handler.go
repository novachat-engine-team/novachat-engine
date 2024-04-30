/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.getFileHashes_handler.go
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

//  upload.getFileHashes#c7025931 location:InputFileLocation offset:int = Vector<FileHash>;
//
func (s *UploadServiceImpl) UploadGetFileHashes(ctx context.Context, request *mtproto.TLUploadGetFileHashes) (*mtproto.Vector_FileHash, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UploadGetFileHashes %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// unused
	///////////////////////////////////////////////////////////////////////////////
	// InputFileLocation <--
	//  + TL_inputFileLocation
	//  + TL_inputEncryptedFileLocation
	//  + TL_inputDocumentFileLocation
	//  + TL_inputSecureFileLocation
	//  + TL_inputTakeoutFileLocation
	//  + TL_inputPhotoFileLocation
	//  + TL_inputPhotoLegacyFileLocation
	//  + TL_inputPeerPhotoFileLocation
	//  + TL_inputStickerSetThumb
	panic("UploadGetFileHashes")
}
