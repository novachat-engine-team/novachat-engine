/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : upload.getFile_handler.go
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

//  upload.getFile#b15a9afc flags:# precise:flags.0?true cdn_supported:flags.1?true location:InputFileLocation offset:int limit:int = upload.File;
//  upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
//
func (s *UploadServiceImpl) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UploadGetFile %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	var ret *mtproto.Upload_File
	switch request.Location.ClassName {
	case mtproto.ClassInputFileLocation, mtproto.ClassInputTakeoutFileLocation:
		//  inputFileLocation#dfdaabe1 volume_id:long local_id:int secret:long file_reference:bytes = InputFileLocation;
		//  inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
		inputFileLocation := request.Location.To_InputFileLocation()

		//  inputTakeoutFileLocation#29be5899 = InputFileLocation;
		inputTakeoutFileLocation := request.Location.To_InputTakeoutFileLocation()
		_ = inputTakeoutFileLocation

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType: sfsService.GetFile_Document,
			VolumeId: inputFileLocation.GetVolumeId(),
			LocalId:  inputFileLocation.GetLocalId(),
			Offset:   request.Offset,
			Limit:    request.Limit,
		})

	case mtproto.ClassInputEncryptedFileLocation:
		//  inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
		fileLocation := request.Location.To_InputEncryptedFileLocation()
		_ = fileLocation

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType: sfsService.GetFile_Document,
			VolumeId: fileLocation.GetId(),
			Offset:   request.Offset,
			Limit:    request.Limit,
		})
	case mtproto.ClassInputDocumentFileLocation:
		//  inputDocumentFileLocation#bad07584 id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;
		//  inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
		fileLocation := request.Location.To_InputDocumentFileLocation()
		_ = fileLocation

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType:  sfsService.GetFile_Document,
			VolumeId:  fileLocation.GetId(),
			Offset:    request.Offset,
			Limit:     request.Limit,
			ThumbSize: fileLocation.GetThumbSize(),
		})
	case mtproto.ClassInputSecureFileLocation:
		//  inputSecureFileLocation#cbc7ee28 id:long access_hash:long = InputFileLocation;
		fileLocation := request.Location.To_InputSecureFileLocation()
		_ = fileLocation

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType: sfsService.GetFile_Document,
			VolumeId: fileLocation.GetId(),
			Offset:   request.Offset,
			Limit:    request.Limit,
		})
	case mtproto.ClassInputPhotoFileLocation:
		//  inputPhotoFileLocation#40181ffe id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;
		photoFile := request.Location.To_InputPhotoFileLocation()
		_ = photoFile

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType:  sfsService.GetFile_Photo,
			VolumeId:  photoFile.GetId(),
			Offset:    request.Offset,
			Limit:     request.Limit,
			ThumbSize: photoFile.GetThumbSize(),
		})
	case mtproto.ClassInputPhotoLegacyFileLocation:
		//  inputPhotoLegacyFileLocation#d83466f3 id:long access_hash:long file_reference:bytes volume_id:long local_id:int secret:long = InputFileLocation;
		fileLocation := request.Location.To_InputPhotoLegacyFileLocation()
		_ = fileLocation

		id := fileLocation.GetId()
		if id == 0 {
			id = fileLocation.GetVolumeId()
		}
		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType: sfsService.GetFile_Photo,
			VolumeId: id,
			LocalId:  fileLocation.GetLocalId(),
			Offset:   request.Offset,
			Limit:    request.Limit,
		})
	case mtproto.ClassInputPeerPhotoFileLocation:
		//  inputPeerPhotoFileLocation#27d69997 flags:# big:flags.0?true peer:InputPeer volume_id:long local_id:int = InputFileLocation;
		photoFile := request.Location.To_InputPeerPhotoFileLocation()
		_ = photoFile

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType: sfsService.GetFile_Photo,
			VolumeId: photoFile.GetVolumeId(),
			LocalId:  photoFile.GetLocalId(),
			Offset:   request.Offset,
			Limit:    request.Limit,
		})

	case mtproto.ClassInputStickerSetThumb:
		//  inputStickerSetThumb#dbaeae9 stickerset:InputStickerSet volume_id:long local_id:int = InputFileLocation;
		inputStickerSetThumb := request.Location.To_InputStickerSetThumb()
		_ = inputStickerSetThumb

		ret, err = s.uploadCore.GetFile(&sfsService.GetFile{
			FileType: sfsService.GetFile_Photo,
			VolumeId: inputStickerSetThumb.GetVolumeId(),
			LocalId:  inputStickerSetThumb.GetLocalId(),
			Offset:   request.Offset,
			Limit:    request.Limit,
		})

	default:
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD_FILE_LOCATION_INVALID)
		log.Errorf("UploadGetFileHashes %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if err != nil {
		log.Warnf("UploadGetFileHashes %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	return ret, nil
}
