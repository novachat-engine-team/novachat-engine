/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package photo

import (
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/data/fs"
)

func PhotoInfo2Photo(photoInfo *sfsService.PhotoInfo, layer int32) *mtproto.Photo {

	var sizes []*mtproto.PhotoSize
	if len(photoInfo.PhotoSize) > 0 {
		sizes = make([]*mtproto.PhotoSize, 0, len(photoInfo.PhotoSize))
		for _, v := range photoInfo.PhotoSize {
			for _, vv := range v.PhotoSize {
				if len(vv.PhotoSize) == 0 {
					tmpPhotoSize := PhotoInfo2PhotoSize(vv, layer)
					if tmpPhotoSize.Location != nil {
						tmpPhotoSize.Location.VolumeId = v.VolumeId
					}
					sizes = append(sizes, tmpPhotoSize)
					if strippedSize := PhotoInfo2PhotoStrippedSize(vv); strippedSize != nil {
						sizes = append(sizes, strippedSize)
					}
				} else {
					for _, vvv := range vv.PhotoSize {
						tmpPhotoSize := PhotoInfo2PhotoSize(vvv, layer)
						if tmpPhotoSize.Location != nil {
							tmpPhotoSize.Location.VolumeId = v.VolumeId
						}
						sizes = append(sizes, tmpPhotoSize)
						if strippedSize := PhotoInfo2PhotoStrippedSize(vvv); strippedSize != nil {
							sizes = append(sizes, strippedSize)
						}
					}
				}
			}
		}
	} else {
		sizes = make([]*mtproto.PhotoSize, 1)
		sizes[0] = mtproto.NewTLPhotoSizeEmpty(&mtproto.PhotoSize{
			Type: partfs.KPhotoSizeMediumType,
		}).To_PhotoSize()
	}

	return mtproto.NewTLPhoto(&mtproto.Photo{
		Id:          photoInfo.VolumeId,
		Sizes:       sizes,
		DcId:        dc.DefaultDc,
		AccessHash:  photoInfo.AccessHash,
		HasStickers: len(photoInfo.PhotoSize) > 0,
		Date:        photoInfo.Date,
	}).To_Photo()
}

func PhotoData2Photo(photoInfo *data_fs.Photo) *mtproto.Photo {

	var sizes []*mtproto.PhotoSize
	if len(photoInfo.Detail) > 0 {
		sizes = make([]*mtproto.PhotoSize, 0, len(photoInfo.Detail))
		for _, v := range photoInfo.Detail {
			sizes = append(sizes, PhotoSize2PhotoSize(v))
			if strippedSize := PhotoData2PhotoStrippedSize(v); strippedSize != nil {
				sizes = append(sizes, strippedSize)
			}
		}
	} else {
		sizes = make([]*mtproto.PhotoSize, 1)
		sizes[0] = mtproto.NewTLPhotoSizeEmpty(&mtproto.PhotoSize{
			Type: partfs.KPhotoSizeMediumType,
		}).To_PhotoSize()
	}

	return mtproto.NewTLPhoto(&mtproto.Photo{
		Id:          photoInfo.VolumeId,
		Sizes:       sizes,
		DcId:        dc.DefaultDc,
		AccessHash:  photoInfo.AccessHash,
		HasStickers: len(photoInfo.Detail) > 1,
		Date:        photoInfo.Date,
	}).To_Photo()
}

func Photo2PhotoData(photoSize *mtproto.Photo) *data_fs.Photo {
	if photoSize.ClassName == mtproto.ClassPhotoEmpty {
		return nil
	}

	photo := &data_fs.Photo{
		VolumeId:   photoSize.Id,
		LocalId:    0,
		Date:       photoSize.Date,
		FileType:   0,
		AccessHash: photoSize.AccessHash,
	}

	for idx, v := range photoSize.Sizes {

		detail := data_fs.PhotoDetail{}
		if v.Location != nil {
			detail.VolumeId = v.Location.VolumeId
			detail.LocalId = v.Location.LocalId
			photo.VolumeId = v.Location.VolumeId
			photo.LocalId = v.Location.LocalId
		}
		if idx == 0 {
			photo.Size_ = v.Size_
		}

		detail.PhotoSize = &data_fs.PhotoSize{
			Width:  v.W,
			Height: v.H,
			Size_:  v.Size_,
			Type:   v.Type,
		}
		detail.Bytes = v.Bytes
		photo.Detail = append(photo.Detail, &detail)
	}

	return photo
}
