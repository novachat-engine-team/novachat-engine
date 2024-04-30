/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/15 14:03
 * @File : photo_file_location.go
 */

package compat

import (
	"novachat_engine/mtproto"
)

func NewFileLocationByLayer(volumeId int64, localId int32, layer int32, dcId int32) *mtproto.FileLocation {
	if layer >= 98 {
		return mtproto.NewTLFileLocationToBeDeprecated(&mtproto.FileLocation{
			VolumeId:             volumeId,
			LocalId:              localId,
			DcId: 			      dcId,
		}).To_FileLocation()
	} else {
		return mtproto.NewTLFileLocation(&mtproto.FileLocation{
			VolumeId:             volumeId,
			LocalId:              localId,
			DcId:                 dcId,
		}).To_FileLocation()
	}
}
