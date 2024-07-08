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
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	data_fs "novachat_engine/service/data/fs"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/upload/photo"
)

///////////////////////////////////////////////////////////////////////////////
// MessageMedia <--
//  + TL_messageMediaEmpty
//  + TL_messageMediaPhoto
//  + TL_messageMediaGeo
//  + TL_messageMediaContact
//  + TL_messageMediaUnsupported
//  + TL_messageMediaDocument
//  + TL_messageMediaWebPage
//  + TL_messageMediaVenue
//  + TL_messageMediaGame
//  + TL_messageMediaInvoice
//  + TL_messageMediaGeoLive
//  + TL_messageMediaPoll
//  + TL_messageMediaDice

func toMessageMediaType(media *mtproto.MessageMedia) constants.MessageMedia {
	switch media.ClassName {
	case mtproto.ClassMessageMediaEmpty:
		return constants.MessageMediaEmpty
	case mtproto.ClassMessageMediaPhoto:
		return constants.MessageMediaPhoto
	case mtproto.ClassMessageMediaGeo:
		return constants.MessageMediaGeo
	case mtproto.ClassMessageMediaContact:
		return constants.MessageMediaContact
	case mtproto.ClassMessageMediaUnsupported:
		return constants.MessageMediaUnsupported
	case mtproto.ClassMessageMediaDocument:
		return constants.MessageMediaDocument
	case mtproto.ClassMessageMediaWebPage:
		return constants.MessageMediaWebPage
	case mtproto.ClassMessageMediaVenue:
		return constants.MessageMediaVenue
	case mtproto.ClassMessageMediaGame:
		return constants.MessageMediaGame
	case mtproto.ClassMessageMediaInvoice:
		return constants.MessageMediaInvoice
	case mtproto.ClassMessageMediaGeoLive:
		return constants.MessageMediaGeoLive
	case mtproto.ClassMessageMediaPoll:
		return constants.MessageMediaPoll
	case mtproto.ClassMessageMediaDice:
		return constants.MessageMediaDice
	default:
		panic(fmt.Sprintf("toMessageMediaType: %+v", media.ClassName))
	}
}

func toMessageMedia(media *data_message.Media, layer int32) *mtproto.MessageMedia {
	if media == nil {
		return mtproto.NewTLMessageMediaEmpty(nil).To_MessageMedia()
	}
	var m *mtproto.MessageMedia
	//TODO:Coder
	// media to mtproto.MessageMedia
	switch constants.MessageMedia(media.Type) {
	case constants.MessageMediaEmpty:
		return mtproto.NewTLMessageMediaEmpty(m).To_MessageMedia()
	case constants.MessageMediaPhoto:
		return mtproto.NewTLMessageMediaPhoto(&mtproto.MessageMedia{
			Title:           media.Caption,
			TtlSeconds:      media.TtlSeconds,
			PhotoB5223B0F71: photo.PhotoData2Photo(media.Photo),
		}).To_MessageMedia()
	case constants.MessageMediaGeo:
		return mtproto.NewTLMessageMediaGeo(&mtproto.MessageMedia{
			Geo: mtproto.NewTLGeoPoint(&mtproto.GeoPoint{
				Long:           media.GeoPoint.Long,
				Lat:            media.GeoPoint.Lat,
				AccessHash:     media.GeoPoint.AccessHash,
				AccuracyRadius: media.GeoPoint.AccuracyRadius,
			}).To_GeoPoint(),
		}).To_MessageMedia()
	case constants.MessageMediaContact:
		return mtproto.NewTLMessageMediaContact(&mtproto.MessageMedia{
			UserId:      constants.PeerTypeFromUserIDType(media.UserId).ToInt32(),
			LastName:    media.LastName,
			FirstName:   media.FirstName,
			PhoneNumber: media.PhoneNumber,
			Vcard:       media.Vcard,
		}).To_MessageMedia()
	case constants.MessageMediaUnsupported:
		return mtproto.NewTLMessageMediaUnsupported(m).To_MessageMedia()
	case constants.MessageMediaDocument:
		return mtproto.NewTLMessageMediaDocument(&mtproto.MessageMedia{
			Caption:    media.Caption,
			TtlSeconds: media.TtlSeconds,
			Document:   ToMessageDocument(media.Document, layer),
		}).To_MessageMedia()
	case constants.MessageMediaWebPage:
		return mtproto.NewTLMessageMediaWebPage(m).To_MessageMedia()
	case constants.MessageMediaVenue:
		return mtproto.NewTLMessageMediaVenue(m).To_MessageMedia()
	case constants.MessageMediaGame:
		return mtproto.NewTLMessageMediaGame(m).To_MessageMedia()
	case constants.MessageMediaInvoice:
		return mtproto.NewTLMessageMediaInvoice(m).To_MessageMedia()
	case constants.MessageMediaGeoLive:
		return mtproto.NewTLMessageMediaGeoLive(&mtproto.MessageMedia{
			Geo: mtproto.NewTLGeoPoint(&mtproto.GeoPoint{
				Long:           media.GeoPoint.Long,
				Lat:            media.GeoPoint.Lat,
				AccessHash:     media.GeoPoint.AccessHash,
				AccuracyRadius: media.GeoPoint.AccuracyRadius,
			}).To_GeoPoint(),
			Period:                      media.Period,
			Heading:                     media.Heading,
			ProximityNotificationRadius: media.NotificationRadius,
		}).To_MessageMedia()
	case constants.MessageMediaPoll:
		return mtproto.NewTLMessageMediaPoll(m).To_MessageMedia()
	case constants.MessageMediaDice:
		return mtproto.NewTLMessageMediaDice(m).To_MessageMedia()
	default:
		panic(fmt.Sprintf("toMessageMedia %+v", media.Type))
	}
}

func ToMessageMedia(media *data_message.Media, layer int32) *mtproto.MessageMedia {
	return toMessageMedia(media, layer)
}

func ToDataMessageMedia(media *mtproto.MessageMedia) *data_message.Media {
	return messageMediaUtil(media)
}

func messageMediaUtil(media *mtproto.MessageMedia) *data_message.Media {
	if media == nil {
		return nil
	}

	m := &data_message.Media{
		Type: toMessageMediaType(media).ToInt32(),
	}
	return toMessageMediaData(m, media)
}

func toMessageMediaData(m *data_message.Media, mMedia *mtproto.MessageMedia) *data_message.Media {
	media := m
	switch constants.MessageMediaFromInt32(m.Type) {
	case constants.MessageMediaEmpty:

	case constants.MessageMediaPhoto:
		media.Caption = mMedia.Title
		media.TtlSeconds = mMedia.TtlSeconds
		media.Photo = photo.Photo2PhotoData(mMedia.PhotoB5223B0F71)
	case constants.MessageMediaGeo:
		media.GeoPoint = &data_fs.GeoPoint{
			Long:           media.GeoPoint.Long,
			Lat:            media.GeoPoint.Lat,
			AccessHash:     media.GeoPoint.AccessHash,
			AccuracyRadius: media.GeoPoint.AccuracyRadius,
		}
	case constants.MessageMediaContact:
		media.UserId = constants.PeerTypeFromUserIDType32(mMedia.UserId).ToInt()
		media.LastName = mMedia.LastName
		media.FirstName = mMedia.FirstName
		media.PhoneNumber = mMedia.PhoneNumber
		media.Vcard = mMedia.Vcard
	case constants.MessageMediaUnsupported:
	case constants.MessageMediaDocument:
		media.Caption = mMedia.Caption
		media.TtlSeconds = mMedia.TtlSeconds
		media.Document = ToMessageDocumentData(mMedia.Document)
	case constants.MessageMediaWebPage:
	case constants.MessageMediaVenue:
	case constants.MessageMediaGame:
	case constants.MessageMediaInvoice:
	case constants.MessageMediaGeoLive:
		media.GeoPoint = &data_fs.GeoPoint{
			Long:           mMedia.Geo.Long,
			Lat:            mMedia.Geo.Lat,
			AccessHash:     mMedia.Geo.AccessHash,
			AccuracyRadius: mMedia.Geo.AccuracyRadius,
		}
		media.Period = mMedia.Period
		media.Heading = mMedia.Heading
		media.NotificationRadius = mMedia.ProximityNotificationRadius
	case constants.MessageMediaPoll:
	case constants.MessageMediaDice:
	default:
		panic(fmt.Sprintf("toMessageMediaData %+v", m.Type))
	}
	return media
}
