/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/2 10:23
 * @File : sfs_service_impl.go
 */

package rpc

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/disintegration/imaging"
	"novachat_engine/pkg/cmd/sfs/conf"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	partfs2 "novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/pkg/cmd/sfs/server/syncfs"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/core/sfs"
	"novachat_engine/service/data/fs"
	"time"
)

type SfsServiceImpl struct {
	partCacheFS *partfs2.PartFS
	//redisClient *cache.RedisClient
	node *snowflake.Node
	core *sfs.Core
}

func NewSfsServiceImpl(conf *conf.Config) *SfsServiceImpl {
	impl := new(SfsServiceImpl)
	impl.partCacheFS = partfs2.NewPartFS(partfs2.DefaultFRU(), syncfs.GetDefaultSyncFS())

	var err error
	impl.node, err = snowflake.NewNode(int64(conf.Server.ServerId))
	if err != nil {
		panic(err)
	}

	impl.core = sfs.NewSfsCore(&conf.MongoDB)
	return impl
}

func (m *SfsServiceImpl) uploadPhoto(id int64, req *sfsService.UploadPhoto) (*data_fs.Photo, error) {
	log.Debugf("uploadPhoto req:%v", req)
	datPath := partfs.LocationPartDatDir(req.AuthKeyId, req.FileId)

	//mimeType := "image/png"
	subPathType := partfs.SubPathsTypePng
	if len(req.FileName) > 0 {
		format, _, err1 := partfs.GetImageFormat(req.FileName)
		if format == -1 {
			log.Warnf("uploadPhoto GetImageFormat fileName:%s error:%s", req.FileName, err1.Error())
		} else {
			//mimeType = fmt.Sprintf("image/%s", ext[1:])
			if imaging.Format(format) == imaging.JPEG {
				subPathType = partfs.SubPathsTypeJpeg
			} else if imaging.Format(format) == imaging.PNG {
				subPathType = partfs.SubPathsTypePng
			} else if imaging.Format(format) == imaging.GIF {
				subPathType = partfs.SubPathsTypeFGif
			} else {
				subPathType = partfs.SubPathsTypePng
			}
		}
	}

	dstPathName := partfs.LocationDir(id, subPathType)
	size, md5sum, err := m.partCacheFS.CalcDatPartsHash(datPath, req.Parts)
	if err != nil {
		log.Errorf("uploadPhoto req:%v error:%s", req, err.Error())
		return nil, fmt.Errorf("uploadPhoto CalcDatPartsHash error:%s", err.Error())
	}

	var photo *data_fs.Photo
	if len(md5sum) > 0 {
		md5sumInfo, err1 := m.core.GetMd5sum(md5sum)
		if err1 != nil || md5sumInfo.Md5Sum != md5sum {
			log.Warnf("uploadPhoto GetMd5sum not found Md5sum md5sum:%s", md5sum)
		} else {
			if md5sumInfo.DocumentInfo != nil {
				log.Errorf("uploadPhoto is document req:%v md5sum:`%s` error:%s", req, md5sum, err.Error())
				return nil, fmt.Errorf("uploadPhoto is document")
			} else {
				photo = md5sumInfo.PhotoInfo
			}
		}
	}
	if photo == nil {
		photo = &data_fs.Photo{
			VolumeId:   id,
			LocalId:    0,
			FilePath:   dstPathName,
			Filename:   req.GetFileName(),
			Md5Sum:     md5sum,
			Date:       int32(time.Now().Unix()),
			Detail:     nil,
			Size_:      int32(size),
			FileType:   int32(subPathType),
			AccessHash: crypto.GenerateAccessHash(),
		}

		var photoList []*data_fs.Photo
		// 创建缩略图
		if photoList, err = datPhotoHandler(datPath, dstPathName, req.Parts, photo, m); err != nil {
			fmt.Errorf("uploadPhoto datPhotoHandler is error:%s", err.Error())
			return nil, fmt.Errorf("uploadPhoto create thumber error:%s", err.Error())
		}

		m.core.SavePhoto(photoList)
		m.core.SaveMd5sum(&data_fs.Md5Sum{Md5Sum: md5sum, PhotoInfo: photo})
	}

	return photo, nil
}
