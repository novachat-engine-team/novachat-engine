package rpc

import (
	"context"
	"errors"
	"fmt"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/pkg/log"
	"strings"
)

func (m *SfsServiceImpl) ReqGetFile(ctx context.Context, req *sfsService.GetFile) (*sfsService.FileInfo, error) {
	if req.VolumeId == 0 {
		log.Errorf("ReqGetFile req.VolumeId == 0 && req.Id == 0")
		return nil, errors.New("empty files req.Id == 0 && req.VolumeId == 0")
		//return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_FILE_PARTS_INVALID)
	}

	var fileType int32
	var date int32
	var filePath string
	switch req.FileType {
	case sfsService.GetFile_Photo:
		photoInfo, err := m.core.GetPhoto(req.VolumeId, req.LocalId)
		if err != nil {
			log.Errorf("ReqGetFile GetPhoto error:%s", err.Error())
			return nil, err
		}
		filePath = photoInfo.FilePath
		date = photoInfo.Date
		fileType = photoInfo.FileType
		if len(req.ThumbSize) > 0 {
			for _, v := range photoInfo.Detail {
				if strings.Compare(v.PhotoSize.Type, partfs.KPhotoSizePathSizeType) == 0 {
					filePath = v.FilePath
					break
				}
			}
		}

	case sfsService.GetFile_Document:
		fileInfo, err := m.core.GetDocument(req.VolumeId, req.LocalId)
		if err != nil {
			log.Errorf("ReqGetFile GetDocument error:%s", err.Error())
			return nil, err
		}
		fileType = fileInfo.FileType
		filePath = fileInfo.FilePath
		date = fileInfo.Date

	default:
		return nil, fmt.Errorf("ReqGetFile req fileType error:%v", req.GetFileType())
	}

	bytes, err := partfs.GetFileData(filePath, req.Offset, req.Limit)
	if err != nil {
		log.Warnf("ReqGetFile GetFileData req:%+v, error:%s", req, err.Error())
		bytes = []byte{}
	}

	log.Debugf("ReqGetFile GetFileData req:%+v, len()=%d", req, len(bytes))
	return &sfsService.FileInfo{
		FileType: fileType,
		Bytes:    bytes,
		Mtime:    date,
	}, nil
}
