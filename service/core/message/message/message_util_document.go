/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/compat"
	data_fs "novachat_engine/service/data/fs"
	"novachat_engine/service/upload/attribute"
	"novachat_engine/service/upload/photo"
)

const (
	DocumentTypeEmpty    int32 = 0
	DocumentTypeDocument int32 = 1
)

func ToMessageDocument(d *data_fs.Document, layer int32) *mtproto.Document {
	if d == nil || d.VolumeId == 0 {
		return mtproto.NewTLDocumentEmpty(nil).To_Document()
	}

	document := &mtproto.Document{}
	document.Id = d.VolumeId
	document.AccessHash = d.AccessHash
	document.Date = d.Date
	document.MimeType = d.MimeType
	document.Size_ = d.Size_
	document.DcId = dc.DefaultDc
	document.Version = d.Version
	document.FileReference = d.FileReference
	if d.VideoStartTs > 0.0 {
		//TODO:(Coderxw VideoThumbs)
		document.VideoThumbs = make([]*mtproto.VideoSize, 0, len(d.Thumbs))
		for _, v := range d.Thumbs {
			videoSize := mtproto.NewTLVideoSize(nil).To_VideoSize()
			videoSize.Type = v.Detail[0].PhotoSize.Type
			videoSize.H = v.Detail[0].PhotoSize.Height
			videoSize.W = v.Detail[0].PhotoSize.Width
			videoSize.Size_ = v.Detail[0].PhotoSize.Size_
			videoSize.Location = compat.NewFileLocationByLayer(
				v.VolumeId,
				v.LocalId,
				layer,
				dc.DefaultDc)
			videoSize.VideoStartTs = d.VideoStartTs
			document.VideoThumbs = append(document.VideoThumbs, videoSize)
		}
	} else {
		if len(d.Thumbs) > 0 {
			document.Thumbs = make([]*mtproto.PhotoSize, 0, len(d.Thumbs))
			for _, v := range d.Thumbs {
				document.Thumbs = append(document.Thumbs, photo.PhotoData2PhotoSize(v, layer))
			}
			document.Thumb = document.Thumbs[0]
		}
	}
	if len(d.Attributes) > 0 {
		document.Attributes = make([]*mtproto.DocumentAttribute, 0, len(d.Attributes))
		for _, v := range d.Attributes {
			document.Attributes = append(document.Attributes, attribute.DataAttributes2MtDocumentAttribute(v))
		}
	}
	return mtproto.NewTLDocument(document).To_Document()
}

func toMessageDocumentType(document *mtproto.Document) int32 {
	switch document.ClassName {
	case mtproto.ClassDocumentEmpty:
		return DocumentTypeEmpty
	case mtproto.ClassDocument:
		return DocumentTypeDocument
	default:
		panic(document.ClassName)
	}
}

func ToMessageDocumentData(document *mtproto.Document) *data_fs.Document {
	if document == nil {
		return nil
	}
	documentType := toMessageDocumentType(document)
	m := &data_fs.Document{}

	switch documentType {
	case DocumentTypeEmpty:
		return nil
	case DocumentTypeDocument:
		m.VolumeId = document.Id
		m.AccessHash = document.AccessHash
		m.Date = document.Date
		m.MimeType = document.MimeType
		m.Size_ = document.Size_
		m.FileReference = document.FileReference
		m.Version = document.Version
		if document.Thumb != nil {
			m.Thumbs = []*data_fs.Photo{photo.PhotoSize2PhotoData(document.Thumb)}
		} else if len(document.Thumbs) > 0 {
			m.Thumbs = make([]*data_fs.Photo, 0, len(document.Thumbs))
			for _, v := range document.Thumbs {
				m.Thumbs = append(m.Thumbs, photo.PhotoSize2PhotoData(v))
			}
		}
		if len(document.Attributes) > 0 {
			m.Attributes = make([]*data_fs.Attributes, 0, len(document.Attributes))
			for _, v := range document.Attributes {
				m.Attributes = append(m.Attributes, attribute.MtDocumentAttribute2DataAttributes(v))
			}
		}
		if len(document.VideoThumbs) > 0 {
			//TODO:(Coderxw VideoThumbs)
			//m.VideoThumbs = make([]*data_fs.Video, 0, len(document.VideoThumbs))
			//for _, v := range document.VideoThumbs {
			//	m.VideoThumbs = append(m.VideoThumbs, videoSizeToMessageVideoData(v))
			//}
		}
	}
	return m
}
