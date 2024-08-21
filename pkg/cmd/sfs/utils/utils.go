/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package utils

import (
	"github.com/json-iterator/go"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/service/data/fs"
)

func ToDocument(fileInfo *data_fs.Document) *sfsService.DocumentInfo {
	if fileInfo == nil {
		return &sfsService.DocumentInfo{}
	}

	var thumbs []*sfsService.PhotoInfo
	for _, v := range fileInfo.Thumbs {
		thumbs = append(thumbs, ToPhoto(v))
	}
	attributes := make([]*sfsService.Attributes, 0, len(fileInfo.Attributes))
	for _, v := range fileInfo.Attributes {
		attributes = append(attributes, ToAttributes(v))
	}
	documentInfo := sfsService.DocumentInfo{
		VolumeId:     fileInfo.VolumeId,
		LocalId:      fileInfo.LocalId,
		MimeType:     fileInfo.MimeType,
		Size_:        fileInfo.Size_,
		Thumbs:       thumbs,
		Attributes:   attributes,
		Date:         fileInfo.Date,
		AccessHash:   fileInfo.AccessHash,
		VideoStartTs: fileInfo.VideoStartTs,
	}

	return &documentInfo
}

func DocumentToVideo(fileInfo *sfsService.DocumentInfo) *data_fs.Video {
	v := &data_fs.Video{
		VolumeId: fileInfo.VolumeId,
		LocalId:  fileInfo.LocalId,
		Detail:   make([]*data_fs.VideoDetail, 0, len(fileInfo.Thumbs)),
		FilePath: "",
	}

	for _, vv := range fileInfo.Thumbs {
		for _, vvv := range vv.PhotoSize {
			v.Detail = append(v.Detail, &data_fs.VideoDetail{
				VolumeId: vvv.VolumeId,
				LocalId:  vvv.LocalId,
				VideoSize: &data_fs.VideoSize{
					Type:   vvv.Type,
					Width:  vvv.Weight,
					Height: vvv.Height,
					Size_:  vvv.Size_,
				},
			})
		}
	}

	return v
}

func ToDataPhoto(info *sfsService.PhotoInfo) *data_fs.Photo {
	photo := &data_fs.Photo{
		VolumeId:   info.VolumeId,
		LocalId:    info.LocalId,
		Filename:   info.Filename,
		Md5Sum:     info.Md5Sum,
		Date:       info.Date,
		Detail:     make([]*data_fs.PhotoDetail, 0, len(info.PhotoSize)),
		FileType:   info.FileType,
		Size_:      info.Size_,
		AccessHash: info.AccessHash,
	}

	for _, v := range info.PhotoSize {
		photo.Detail = append(photo.Detail, &data_fs.PhotoDetail{
			VolumeId: v.VolumeId,
			LocalId:  v.LocalId,
			PhotoSize: &data_fs.PhotoSize{
				Width:  v.Weight,
				Height: v.Height,
				Size_:  v.Size_,
				Type:   v.Type,
			},
			Bytes: v.Bytes,
		})
	}

	return photo
}

func ToPhoto(photo *data_fs.Photo) *sfsService.PhotoInfo {
	if photo == nil || len(photo.Detail) == 0 {
		return &sfsService.PhotoInfo{}
	}

	photoInfo := &sfsService.PhotoInfo{
		VolumeId:   photo.VolumeId,
		LocalId:    photo.LocalId,
		Filename:   photo.Filename,
		Md5Sum:     photo.Md5Sum,
		Date:       photo.Date,
		Size_:      photo.Size_,
		Height:     photo.Detail[0].PhotoSize.Height,
		Weight:     photo.Detail[0].PhotoSize.Width,
		FileType:   photo.FileType,
		PhotoSize:  make([]*sfsService.PhotoInfo, 0, len(photo.Detail)),
		AccessHash: photo.AccessHash,
	}

	for _, v := range photo.Detail {
		photoInfo.PhotoSize = append(photoInfo.PhotoSize, &sfsService.PhotoInfo{
			VolumeId: v.VolumeId,
			LocalId:  v.LocalId,
			PhotoSize: []*sfsService.PhotoInfo{{
				VolumeId: v.VolumeId,
				LocalId:  v.LocalId,
				Date:     photo.Date,
				Size_:    v.PhotoSize.Size_,
				Height:   v.PhotoSize.Height,
				Weight:   v.PhotoSize.Width,
				FileType: photo.FileType,
				Type:     v.PhotoSize.Type,
				Bytes:    v.Bytes,
			}},
		})
	}

	return photoInfo
}

func FromAttributes(attributes *sfsService.Attributes) *data_fs.Attributes {
	var dataAttributes data_fs.Attributes
	if attributes != nil {
		s, _ := jsoniter.Marshal(attributes)
		jsoniter.Unmarshal(s, &dataAttributes)
	}
	return &dataAttributes
}

func ToAttributes(attributes *data_fs.Attributes) *sfsService.Attributes {
	var sfsAttributes sfsService.Attributes
	if attributes != nil {
		s, _ := jsoniter.Marshal(attributes)
		jsoniter.Unmarshal(s, &sfsAttributes)
	}
	return &sfsAttributes
}
