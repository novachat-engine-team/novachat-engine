package attribute

import (
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	data_fs "novachat_engine/service/data/fs"
	"novachat_engine/service/upload/stickerset"
)

type DocumentAttributesType int32

func (d DocumentAttributesType) ToInt32() int32 {
	return int32(d)
}

const (
	DocumentAttributeFilename    DocumentAttributesType = 0
	DocumentAttributeImageSize   DocumentAttributesType = 1
	DocumentAttributeAnimated    DocumentAttributesType = 2
	DocumentAttributeSticker     DocumentAttributesType = 3
	DocumentAttributeVideo       DocumentAttributesType = 4
	DocumentAttributeAudio       DocumentAttributesType = 5
	DocumentAttributeHasStickers DocumentAttributesType = 6
)

func MtDocumentAttribute2Attributes(attribute *mtproto.DocumentAttribute) *sfsService.Attributes {
	attr := &sfsService.Attributes{}
	switch attribute.ClassName {
	case mtproto.ClassDocumentAttributeSticker:
		attr.AttributesType = DocumentAttributeSticker.ToInt32()
		std := stickerset.MtInputStickerSet2DataStickerSet(attribute.Stickerset)
		attr.StickerData = &sfsService.StickerSet{
			StickerSetType: std.StickerSetType,
			StickerSetId:   std.StickerSetId,
			AccessHash:     std.AccessHash,
			ShortName:      std.ShortName,
			Emoticon:       std.Emoticon,
		}

		attr.Mask = attribute.GetMask()
		attr.Alt = attribute.GetAlt()
		attr.MaskCoords = &sfsService.MaskCoords{
			N:    attribute.GetMaskCoords().GetN(),
			X:    attribute.GetMaskCoords().GetX(),
			Y:    attribute.GetMaskCoords().GetY(),
			Zoom: attribute.GetMaskCoords().GetZoom(),
		}
		break
	case mtproto.ClassDocumentAttributeAnimated:
		attr.AttributesType = DocumentAttributeAnimated.ToInt32()
		break
	case mtproto.ClassDocumentAttributeImageSize:
		attr.AttributesType = DocumentAttributeImageSize.ToInt32()
		attr.ImageW = attribute.GetW()
		attr.ImageH = attribute.GetH()
		break
	case mtproto.ClassDocumentAttributeHasStickers:
		attr.AttributesType = DocumentAttributeHasStickers.ToInt32()
		attr.Stickers = true
		break
	case mtproto.ClassDocumentAttributeFilename:
		attr.AttributesType = DocumentAttributeFilename.ToInt32()
		attr.Filename = attribute.GetFileName()
		break
	case mtproto.ClassDocumentAttributeAudio:
		attr.AttributesType = DocumentAttributeAudio.ToInt32()
		attr.Voice = attribute.GetVoice()
		attr.AudioDuration = attribute.GetDuration()
		attr.Title = attribute.GetTitle()
		attr.Performer = attribute.GetPerformer()
		attr.Waveform = attribute.GetWaveform()
		break
	case mtproto.ClassDocumentAttributeVideo:
		attr.AttributesType = DocumentAttributeVideo.ToInt32()
		attr.RoundMessage = attribute.GetRoundMessage()
		attr.SupportsStreaming = false
		attr.VideoDuration = attribute.GetDuration()
		attr.VideoW = attribute.GetW()
		attr.VideoH = attribute.GetH()
		break
	default:
		return nil
	}
	return attr
}

func Attributes2MtDocumentAttribute(attribute *sfsService.Attributes) *mtproto.DocumentAttribute {
	var attr *mtproto.DocumentAttribute
	switch attribute.AttributesType {
	case DocumentAttributeFilename.ToInt32():
		attr = mtproto.NewTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
			FileName: attribute.Filename,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeImageSize.ToInt32():
		attr = mtproto.NewTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
			W: attribute.ImageW,
			H: attribute.ImageH,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeAnimated.ToInt32():
		attr = mtproto.NewTLDocumentAttributeAnimated(nil).To_DocumentAttribute()
		break
	case DocumentAttributeSticker.ToInt32():
		attr = mtproto.NewTLDocumentAttributeSticker(&mtproto.DocumentAttribute{
			Mask: attribute.Mask,
			Alt:  attribute.Alt,
			//TODO: StickerSet
			Stickerset: stickerset.DataStickerSet2MtInputStickerSet(&data_fs.StickerSet{
				StickerSetType: attribute.StickerData.StickerSetType,
				StickerSetId:   attribute.StickerData.StickerSetId,
				AccessHash:     attribute.StickerData.AccessHash,
				ShortName:      attribute.StickerData.ShortName,
				Emoticon:       attribute.StickerData.Emoticon,
			}),
		}).To_DocumentAttribute()
		if attribute.MaskCoords != nil {
			attr.MaskCoords = mtproto.NewTLMaskCoords(&mtproto.MaskCoords{
				N:    attribute.MaskCoords.N,
				X:    attribute.MaskCoords.X,
				Y:    attribute.MaskCoords.Y,
				Zoom: attribute.MaskCoords.Zoom,
			}).To_MaskCoords()
		}
		break
	case DocumentAttributeVideo.ToInt32():
		attr = mtproto.NewTLDocumentAttributeVideo(&mtproto.DocumentAttribute{
			RoundMessage: attribute.RoundMessage,
			Duration:     attribute.VideoDuration,
			W:            attribute.VideoW,
			H:            attribute.VideoH,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeAudio.ToInt32():
		attr = mtproto.NewTLDocumentAttributeAudio(&mtproto.DocumentAttribute{
			Voice:     attribute.Voice,
			Duration:  attribute.AudioDuration,
			Title:     attribute.Title,
			Performer: attribute.Performer,
			Waveform:  attribute.Waveform,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeHasStickers.ToInt32():
		break
	default:
		return nil
	}
	return attr
}

func MtDocumentAttribute2DataAttributes(attribute *mtproto.DocumentAttribute) *data_fs.Attributes {
	attr := &data_fs.Attributes{}
	switch attribute.ClassName {
	case mtproto.ClassDocumentAttributeSticker:
		attr.AttributesType = DocumentAttributeSticker.ToInt32()
		if attribute.Stickerset != nil {
			attr.StickerData = stickerset.MtInputStickerSet2DataStickerSet(attribute.Stickerset)
		}
		attribute.GetStickerset()
		attr.Mask = attribute.GetMask()
		attr.Alt = attribute.GetAlt()
		attr.MaskCoords = &data_fs.MaskCoords{
			N:    attribute.GetMaskCoords().GetN(),
			X:    attribute.GetMaskCoords().GetX(),
			Y:    attribute.GetMaskCoords().GetY(),
			Zoom: attribute.GetMaskCoords().GetZoom(),
		}
		break
	case mtproto.ClassDocumentAttributeAnimated:
		attr.AttributesType = DocumentAttributeAnimated.ToInt32()
		break
	case mtproto.ClassDocumentAttributeImageSize:
		attr.AttributesType = DocumentAttributeImageSize.ToInt32()
		attr.ImageW = attribute.GetW()
		attr.ImageH = attribute.GetH()
		break
	case mtproto.ClassDocumentAttributeHasStickers:
		attr.AttributesType = DocumentAttributeHasStickers.ToInt32()
		attr.Stickers = true
		break
	case mtproto.ClassDocumentAttributeFilename:
		attr.AttributesType = DocumentAttributeFilename.ToInt32()
		attr.Filename = attribute.GetFileName()
		break
	case mtproto.ClassDocumentAttributeAudio:
		attr.AttributesType = DocumentAttributeAudio.ToInt32()
		attr.Voice = attribute.GetVoice()
		attr.AudioDuration = attribute.GetDuration()
		attr.Title = attribute.GetTitle()
		attr.Performer = attribute.GetPerformer()
		attr.Waveform = attribute.GetWaveform()
		break
	case mtproto.ClassDocumentAttributeVideo:
		attr.AttributesType = DocumentAttributeVideo.ToInt32()
		attr.RoundMessage = attribute.GetRoundMessage()
		attr.SupportsStreaming = false
		attr.VideoDuration = attribute.GetDuration()
		attr.VideoW = attribute.GetW()
		attr.VideoH = attribute.GetH()
		break
	default:
		return nil
	}
	return attr
}

func DataAttributes2MtDocumentAttribute(attribute *data_fs.Attributes) *mtproto.DocumentAttribute {
	var attr *mtproto.DocumentAttribute
	switch attribute.AttributesType {
	case DocumentAttributeFilename.ToInt32():
		attr = mtproto.NewTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
			FileName: attribute.Filename,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeImageSize.ToInt32():
		attr = mtproto.NewTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
			W: attribute.ImageW,
			H: attribute.ImageH,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeAnimated.ToInt32():
		attr = mtproto.NewTLDocumentAttributeAnimated(nil).To_DocumentAttribute()
		break
	case DocumentAttributeSticker.ToInt32():
		attr = mtproto.NewTLDocumentAttributeSticker(&mtproto.DocumentAttribute{
			Mask:       attribute.Mask,
			Alt:        attribute.Alt,
			Stickerset: nil,
			MaskCoords: nil,
		}).To_DocumentAttribute()
		if attribute.StickerData != nil {
			attr.Stickerset = stickerset.DataStickerSet2MtInputStickerSet(attribute.StickerData)
		}
		if attribute.MaskCoords != nil {
			attr.MaskCoords = mtproto.NewTLMaskCoords(&mtproto.MaskCoords{
				N:    attribute.MaskCoords.N,
				X:    attribute.MaskCoords.X,
				Y:    attribute.MaskCoords.Y,
				Zoom: attribute.MaskCoords.Zoom,
			}).To_MaskCoords()
		}
		break
	case DocumentAttributeVideo.ToInt32():
		attr = mtproto.NewTLDocumentAttributeVideo(&mtproto.DocumentAttribute{
			RoundMessage: attribute.RoundMessage,
			Duration:     attribute.VideoDuration,
			W:            attribute.VideoW,
			H:            attribute.VideoH,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeAudio.ToInt32():
		attr = mtproto.NewTLDocumentAttributeAudio(&mtproto.DocumentAttribute{
			Voice:     attribute.Voice,
			Duration:  attribute.AudioDuration,
			Title:     attribute.Title,
			Performer: attribute.Performer,
			Waveform:  attribute.Waveform,
		}).To_DocumentAttribute()
		break
	case DocumentAttributeHasStickers.ToInt32():
		break
	default:
		return nil
	}
	return attr
}
