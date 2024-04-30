/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/7 19:00
 * @File : mimetype_util.go
 */

package partfs

import "strings"

//audio/  mp3 mov
//video/  mp4 mkv
//image/   png jpeg gif webp
//application/x-tgwallpattern
//application/x-tgsticker		 *.tgs
//application/x-tgtheme-android  *.attheme

var _suffixOther = map[string]SubPathsType{
	".pdf": SubPathsTypePdf,
	".mkv": SubPathsTypeMp4,
}

func MimeTypeToSubPathType(mimeType string, name string) SubPathsType {
	if strings.HasPrefix(mimeType, "audio/") {
		if strings.HasPrefix(mimeType, "audio/mp3") {
			return SubPathsTypeMp3
		} else if strings.HasPrefix(mimeType, "audio/mov") {
			return SubPathsTypeMov
		}
		return SubPathsTypeMp3
	} else if strings.HasPrefix(mimeType, "video/") {
		//if strings.HasPrefix(mimeType, "video/mp4") {
		//} else if strings.HasPrefix(mimeType, "video/mkv") {
		//}
		return SubPathsTypeMp4
	} else if strings.HasPrefix(mimeType, "image/") {
		if strings.HasPrefix(mimeType, "image/png") {
		} else if strings.HasPrefix(mimeType, "image/jpeg") {
		} else if strings.HasPrefix(mimeType, "image/gif") {
			return SubPathsTypeFGif
		} else if strings.HasPrefix(mimeType, "image/webp") {
			return SubPathsTypeWebp
		}
		return SubPathsTypePng
	} else if strings.Compare(strings.ToLower(mimeType), "application/x-tgsticker") == 0 {
		return SubPathsTypeStickerSet
	} else {
		for k, v := range _suffixOther {
			if strings.HasSuffix(name, k) {
				return v
			}
		}
		return SubPathsTypeDat
	}
}
