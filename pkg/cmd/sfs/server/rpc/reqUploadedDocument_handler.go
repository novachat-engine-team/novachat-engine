package rpc

import (
	"context"
	"fmt"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/data/fs"
	"time"
)

func (m *SfsServiceImpl) ReqUploadedDocument(ctx context.Context, req *sfsService.UploadedDocument) (*sfsService.DocumentInfo, error) {
	log.Debugf("ReqUploadedDocument req:%v", req)
	var err error

	if req.Video && len(req.MimeType) == 0 {
		req.MimeType = "video/mpt4"
	}

	id := m.node.Generate().Int64()
	datPath := partfs.LocationPartDatDir(req.AuthKeyId, req.FileId)
	fileType := partfs.MimeTypeToSubPathType(req.MimeType, req.FileName)

	dstPathName := partfs.LocationDir(id, fileType)
	size, md5sum, err := m.partCacheFS.CalcDatPartsHash(datPath, req.Parts)
	if err != nil {
		log.Errorf("ReqUploadedDocument req:%v error:%s", req, err.Error())
		return nil, fmt.Errorf("ReqUploadedDocument CalcDatPartsHash error:%s", err.Error())
	}

	var fileInfo *data_fs.Document
	if len(md5sum) > 0 {
		md5sumInfo, err1 := m.core.GetMd5sum(md5sum)
		if err1 != nil || md5sumInfo.Md5Sum != md5sum {
			log.Warnf("ReqUploadedDocument GetMd5sum not found Md5sum md5sum:%s", md5sum)
		} else {
			if md5sumInfo.DocumentInfo != nil {
				fileInfo = md5sumInfo.DocumentInfo
			} else {
				log.Errorf("ReqUploadedDocument is document req:%v md5sum:`%s` error:%s", req, md5sum, err.Error())
				return nil, fmt.Errorf("ReqUploadedDocument is document")
			}
		}
	}

	if fileInfo == nil {
		var photo *data_fs.Photo
		if req.Thumb != nil {
			photo, err = m.uploadPhoto(m.node.Generate().Int64(), &sfsService.UploadPhoto{
				AuthKeyId: req.AuthKeyId,
				FileId:    req.Thumb.FileId,
				FileName:  req.Thumb.GetFileName(),
				Parts:     req.Thumb.GetParts(),
			})
			if err != nil {
				log.Debugf("UploadedDocument req:%v UploadedPhotoFile error:%s", req, err.Error())
				return nil, err
			}
		}
		attributes := make([]*data_fs.Attributes, 0, len(req.Attributes))
		for _, v := range req.Attributes {
			attributes = append(attributes, utils.FromAttributes(v))
		}
		fileInfo = &data_fs.Document{
			VolumeId:   id,
			LocalId:    0,
			MimeType:   req.MimeType,
			Size_:      int32(size),
			Thumbs:     []*data_fs.Photo{photo},
			Attributes: attributes,
			Date:       int32(time.Now().Unix()),
			FileType:   fileType.ToInt32(),
			FilePath:   dstPathName,
			AccessHash: crypto.GenerateAccessHash(),
		}

		if photo == nil && fileType == partfs.SubPathsTypePng {
			photo = &data_fs.Photo{
				VolumeId:   id,
				LocalId:    0,
				FilePath:   dstPathName,
				Filename:   req.GetFileName(),
				Md5Sum:     md5sum,
				Date:       fileInfo.Date,
				Detail:     nil,
				Size_:      fileInfo.Size_,
				FileType:   fileType.ToInt32(),
				AccessHash: fileInfo.AccessHash,
			}
			var photoList []*data_fs.Photo
			photoList, err = datPhotoHandler(datPath, dstPathName, req.Parts, photo, m)
			if err != nil {
				log.Debugf("UploadedDocument req:%v datPhotoHandler error:%s", req, err.Error())
				return nil, err
			}

			fileInfo.Thumbs = photoList
			m.core.SavePhoto(photoList)
			m.core.SaveMd5sum(&data_fs.Md5Sum{Md5Sum: md5sum, PhotoInfo: photo})
		} else {
			if req.Video {
				fileInfo.VideoStartTs = req.VideoStartTs
			}
			go datHandler(datPath, dstPathName, req.Parts, fileInfo, m)
			m.core.SaveDocument([]*data_fs.Document{fileInfo})
			m.core.SaveMd5sum(&data_fs.Md5Sum{
				Md5Sum:       md5sum,
				DocumentInfo: fileInfo,
			})
		}
	}

	log.Debugf("UploadedDocument req:%v reply ok!!!!!!!!!", req)
	return utils.ToDocument(fileInfo), nil
}

func datHandler(datPath, dstPathName string, parts int32, info *data_fs.Document, m *SfsServiceImpl) {
	err := m.partCacheFS.MergeDatFileParts(datPath, parts, dstPathName)
	if err != nil {
		m.partCacheFS.UnClearDir(datPath)
		log.Errorf("MergeDatFileParts datPath:%s dstPathName:%s info:%v error:%s", datPath, dstPathName, info, err.Error())
		return
	}

	log.Infof("datHandler datPath:%s dstPathName:%s info:%v completed!!!!!!!!!!!!!!!!!", datPath, dstPathName, info)
}
