package photo

import (
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/compat"
	data_sfs "novachat_engine/service/data/fs"
)

//  + TL_photoSizeEmpty
//  + TL_photoSize
//  + TL_photoCachedSize
//  + TL_photoStrippedSize
//  + TL_photoSizeProgressive
//  + TL_photoPathSize

func PhotoInfo2PhotoSize(photo *sfsService.PhotoInfo, layer int32) *mtproto.PhotoSize {
	if photo == nil {
		return mtproto.NewTLPhotoSizeEmpty(&mtproto.PhotoSize{
			Type: partfs.KPhotoSizeMediumType,
		}).To_PhotoSize()
	}

	if len(photo.Bytes) == 0 {
		return mtproto.NewTLPhotoSize(&mtproto.PhotoSize{
			Type:     photo.Type,
			W:        photo.Weight,
			H:        photo.Height,
			Size_:    photo.Size_,
			Location: compat.NewFileLocationByLayer(photo.VolumeId, photo.LocalId, layer, dc.DefaultDc),
		}).To_PhotoSize()
	} else {
		return mtproto.NewTLPhotoPathSize(&mtproto.PhotoSize{
			Type:  photo.Type,
			Bytes: photo.Bytes,
		}).To_PhotoSize()
	}
}

func PhotoInfo2PhotoStrippedSize(photo *sfsService.PhotoInfo) *mtproto.PhotoSize {
	if photo == nil {
		return nil
	}

	if len(photo.Bytes) == 0 {
		return nil
	}

	return mtproto.NewTLPhotoStrippedSize(&mtproto.PhotoSize{
		Type:  photo.Type,
		Bytes: photo.Bytes,
	}).To_PhotoSize()
}

func PhotoSize2PhotoData(photoSize *mtproto.PhotoSize) *data_sfs.Photo {
	if photoSize.ClassName == mtproto.ClassPhotoSizeEmpty {
		return nil
	}

	photo := &data_sfs.Photo{
		Detail: []*data_sfs.PhotoDetail{{
			PhotoSize: &data_sfs.PhotoSize{
				Width:  photoSize.W,
				Height: photoSize.H,
				Size_:  photoSize.Size_,
				Type:   photoSize.Type,
			},
			Bytes: photoSize.Bytes,
		}},
		Size_: photoSize.Size_,
	}
	if photoSize.Location != nil {
		photo.LocalId = photoSize.Location.LocalId
		photo.VolumeId = photoSize.Location.VolumeId
		photo.Detail[0].VolumeId = photoSize.Location.VolumeId
		photo.Detail[0].LocalId = photoSize.Location.LocalId
	}

	return photo
}

func PhotoData2PhotoSize(photo *data_sfs.Photo, layer int32) *mtproto.PhotoSize {
	if photo == nil || len(photo.Detail) == 0 {
		return mtproto.NewTLPhotoSizeEmpty(&mtproto.PhotoSize{
			Type: partfs.KPhotoSizeMediumType,
		}).To_PhotoSize()
	}

	if len(photo.Detail[0].Bytes) == 0 {
		return mtproto.NewTLPhotoSize(&mtproto.PhotoSize{
			Type:     photo.Detail[0].PhotoSize.Type,
			W:        photo.Detail[0].PhotoSize.Width,
			H:        photo.Detail[0].PhotoSize.Height,
			Size_:    photo.Size_,
			Location: compat.NewFileLocationByLayer(photo.VolumeId, photo.LocalId, layer, dc.DefaultDc),
		}).To_PhotoSize()
	} else {
		return mtproto.NewTLPhotoPathSize(&mtproto.PhotoSize{
			Type:  photo.Detail[0].PhotoSize.Type,
			Bytes: photo.Detail[0].Bytes,
		}).To_PhotoSize()
	}
}

func PhotoSize2PhotoSize(photo *data_sfs.PhotoDetail) *mtproto.PhotoSize {
	if photo == nil {
		return mtproto.NewTLPhotoSizeEmpty(&mtproto.PhotoSize{
			Type: partfs.KPhotoSizeMediumType,
		}).To_PhotoSize()
	}

	if len(photo.Bytes) == 0 {
		return mtproto.NewTLPhotoSize(&mtproto.PhotoSize{
			Type:     photo.PhotoSize.Type,
			W:        photo.PhotoSize.Width,
			H:        photo.PhotoSize.Height,
			Size_:    photo.PhotoSize.Size_,
			Location: compat.NewFileLocationByLayer(photo.VolumeId, photo.LocalId, 0, dc.DefaultDc),
		}).To_PhotoSize()
	} else {
		return mtproto.NewTLPhotoPathSize(&mtproto.PhotoSize{
			Type:  photo.PhotoSize.Type,
			Bytes: photo.Bytes,
		}).To_PhotoSize()
	}
}

func PhotoData2PhotoStrippedSize(photo *data_sfs.PhotoDetail) *mtproto.PhotoSize {
	if photo == nil {
		return nil
	}

	if len(photo.Bytes) == 0 {
		return nil
	}

	return mtproto.NewTLPhotoStrippedSize(&mtproto.PhotoSize{
		Type:  photo.PhotoSize.Type,
		Bytes: photo.Bytes,
	}).To_PhotoSize()
}

func PhotoProfileUserProfilePhoto(dataPhoto *data_sfs.PhotoProfile, layer int32) *mtproto.UserProfilePhoto {
	if dataPhoto != nil {
		userProfilePhoto := mtproto.NewTLUserProfilePhoto(&mtproto.UserProfilePhoto{
			PhotoId: dataPhoto.Photo.VolumeId,
			DcId:    dc.DefaultDc,
		}).To_UserProfilePhoto()
		userProfilePhoto.PhotoSmall = compat.NewFileLocationByLayer(
			dataPhoto.Photo.VolumeId,
			dataPhoto.Photo.LocalId,
			layer,
			dc.DefaultDc)
		userProfilePhoto.HasVideo = dataPhoto.Video != nil
		if dataPhoto.Video != nil {
			userProfilePhoto.PhotoBig = compat.NewFileLocationByLayer(
				dataPhoto.Video.VolumeId,
				dataPhoto.Video.LocalId,
				layer,
				dc.DefaultDc)
		} else {
			userProfilePhoto.PhotoBig = compat.NewFileLocationByLayer(
				dataPhoto.Photo.VolumeId,
				dataPhoto.Photo.LocalId,
				layer,
				dc.DefaultDc)
		}
		return userProfilePhoto
	}
	return nil
}
