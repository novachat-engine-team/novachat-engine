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
	"math/rand"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/app/help"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/phone"
	"novachat_engine/service/data/users"
	"novachat_engine/service/input"
	"novachat_engine/service/upload/attribute"
	"time"
)

func InputGeoPointToGeoPoint(geoPoint *mtproto.InputGeoPoint) *mtproto.GeoPoint {
	switch geoPoint.ClassName {
	case mtproto.ClassInputGeoPointEmpty:
		return mtproto.NewTLGeoPointEmpty(nil).To_GeoPoint()
	case mtproto.ClassInputGeoPoint:
		return mtproto.NewTLGeoPoint(&mtproto.GeoPoint{
			Long: geoPoint.Long,
			Lat:  geoPoint.Lat,
		}).To_GeoPoint()
	default:
		log.Errorf("InputGeoPointToGeoPoint invalid  %v", geoPoint.ClassName)
	}
	return mtproto.NewTLGeoPointEmpty(nil).To_GeoPoint()
}

func (c *Core) makeInputMedia(authKeyId int64, layer int32, medias []*mtproto.InputMedia) ([]*mtproto.MessageMedia, error) {

	if len(medias) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MEDIA_INVALID)
	}

	ret := make([]*mtproto.MessageMedia, 0, len(medias))
	for _, media := range medias {
		switch media.ClassName {
		case mtproto.ClassInputMediaUploadedDocument:
			//  inputMediaUploadedDocument#5b38c6c1 flags:# nosound_video:flags.3?true force_file:flags.4?true file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
			//  inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;

			uploadedDocument := media.To_InputMediaUploadedDocument()
			stickers := make([]int64, 0, len(media.Stickers))
			for _, sticker := range media.Stickers {
				stickers = append(stickers, sticker.Id)
			}
			var thumb *sfsService.UploadedDocument_Thumb
			if uploadedDocument.GetThumb() != nil {
				thumb = &sfsService.UploadedDocument_Thumb{
					FileId:   uploadedDocument.GetThumb().GetId(),
					FileName: uploadedDocument.GetThumb().GetName(),
					Parts:    uploadedDocument.GetThumb().GetParts(),
				}
			}

			inputFile := uploadedDocument.GetFile()
			attributes := make([]*sfsService.Attributes, 0, len(uploadedDocument.GetAttributes()))
			for _, v := range uploadedDocument.GetAttributes() {
				attributes = append(attributes, attribute.MtDocumentAttribute2Attributes(v))
			}

			document, err := c.uploadCore.UploadDocument(&sfsService.UploadedDocument{
				AuthKeyId:  authKeyId,
				FileId:     inputFile.GetId(),
				FileName:   inputFile.GetName(),
				Parts:      inputFile.GetParts(),
				MimeType:   uploadedDocument.GetMimeType(),
				Attributes: attributes,
				Thumb:      thumb,
				Video:      uploadedDocument.GetNosoundVideo(),
				TtlSecond:  uploadedDocument.GetTtlSeconds(),
			}, layer)
			if err != nil {
				log.Errorf("ClassInputMediaUploadedDocument inputFile:%v error:%s", media, err.Error())
				return nil, err
			}

			ret = append(ret, mtproto.NewTLMessageMediaDocument(&mtproto.MessageMedia{
				TtlSeconds: media.TtlSeconds,
				Caption:    media.Caption,
				Document:   document,
			}).To_MessageMedia())
			break

		case mtproto.ClassInputMediaDocument:
			//  inputMediaDocument#33473058 flags:# id:InputDocument ttl_seconds:flags.0?int query:flags.1?string = InputMedia;
			//  inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;

			// messageMediaDocument#9cb070d7 flags:# document:flags.0?Document ttl_seconds:flags.2?int = MessageMedia;
			// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
			id := media.To_InputMediaDocument().GetId5ACB668E71()
			document, err := c.uploadCore.GetDocument(&sfsService.GetDocument{
				FileId:  id.GetId(),
				LocalId: 0,
			}, layer)
			if err != nil {
				log.Errorf("ClassInputMediaDocument inputFile:%v error:%s", id, err.Error())
				return nil, err
			}

			ret = append(ret, mtproto.NewTLMessageMediaDocument(&mtproto.MessageMedia{
				TtlSeconds: media.TtlSeconds,
				Caption:    media.Caption,
				Document:   document,
			}).To_MessageMedia())
			break

		case mtproto.ClassInputMediaUploadedPhoto:
			//  inputMediaUploadedPhoto#1e287d04 flags:# file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
			//  inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;

			// messageMediaPhoto#695150d7 flags:# photo:flags.0?Photo ttl_seconds:flags.2?int = MessageMedia;
			// messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
			uploadedPhoto := media.To_InputMediaUploadedPhoto()
			getFile := uploadedPhoto.GetFile()
			stickers := make([]int64, 0, len(media.Stickers))
			for _, sticker := range media.Stickers {
				stickers = append(stickers, sticker.Id)
			}

			photo, err := c.uploadCore.UploadPhoto(&sfsService.UploadPhoto{
				AuthKeyId: authKeyId,
				FileId:    getFile.GetId(),
				FileName:  getFile.GetName(),
				Parts:     getFile.GetParts(),
				Stickers:  stickers,
			}, layer)
			if err != nil {
				log.Errorf("ClassInputMediaUploadedPhoto inputFile:%v error:%s", media, err.Error())
				return nil, err
			}
			mediaPhoto := mtproto.NewTLMessageMediaPhoto(&mtproto.MessageMedia{
				PhotoB5223B0F71: photo,
				TtlSeconds:      uploadedPhoto.GetTtlSeconds(),
				Caption:         uploadedPhoto.GetCaption(),
			})

			ret = append(ret, mediaPhoto.To_MessageMedia())
			break
		case mtproto.ClassInputMediaPhoto:
			//  inputMediaPhoto#b3ba0635 flags:# id:InputPhoto ttl_seconds:flags.0?int = InputMedia;
			//  inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;

			// messageMediaPhoto#695150d7 flags:# photo:flags.0?Photo ttl_seconds:flags.2?int = MessageMedia;
			// messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
			inputFile := media.To_InputMediaPhoto()
			inputPhoto := inputFile.GetId81FA373A71()

			photo, err := c.uploadCore.GetPhoto(&sfsService.GetPhoto{
				PhotoId: inputPhoto.GetId(),
				LocalId: 0,
			}, layer)
			if err != nil {
				log.Errorf("ClassInputMediaPhoto inputFile:%v error:%s", media, err.Error())
				return nil, err
			}
			ret = append(ret, mtproto.NewTLMessageMediaPhoto(&mtproto.MessageMedia{
				PhotoB5223B0F71: photo,
				Caption:         media.Caption,
				TtlSeconds:      inputFile.GetTtlSeconds(),
			}).To_MessageMedia())
			break

		case mtproto.ClassInputMediaGeoPoint:
			//  messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
			messageMedia := mtproto.NewTLMessageMediaGeo(&mtproto.MessageMedia{
				Geo: InputGeoPointToGeoPoint(media.To_InputMediaGeoPoint().GetGeoPoint()),
			})
			ret = append(ret, messageMedia.To_MessageMedia())
			break
		case mtproto.ClassInputMediaDice:
			// messageMediaDice#3f7ee58b value:int emoticon:string = MessageMedia;
			ret = append(ret, mtproto.NewTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31n(6) + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia())
			break

		case mtproto.ClassInputMediaContact:
			// messageMediaContact#cbf24940 phone_number:string first_name:string last_name:string vcard:string user_id:int = MessageMedia;
			contact := media.To_InputMediaContact()
			messageMedia := mtproto.NewTLMessageMediaContact(&mtproto.MessageMedia{
				PhoneNumber: contact.GetPhoneNumber(),
				FirstName:   contact.GetFirstName(),
				LastName:    contact.GetLastName(),
				Vcard:       contact.GetVcard(),
			})

			phoneNumber, err := phone.Parse(contact.GetPhoneNumber())
			if err == nil {
				var userInfo *data_users.Users
				userCache, err := c.accountCore.PhoneUser(phoneNumber.NormalizeDigitsOnly())
				if err != nil || len(userCache) == 0 {
					userInfo, err = c.userCore.UserByPhoneNumber(phoneNumber.NormalizeDigitsOnly())
					if err != nil {
						log.Errorf("ClassInputMediaContact contact:%v error:%s", contact, err.Error())
						return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CONTACT_ID_INVALID)
					}
				} else {
					userInfo, _ = account.CacheInfo2User(userCache)
				}
				if userInfo != nil {
					messageMedia.SetUserId(constants.PeerTypeFromUserIDType(userInfo.Id).ToInt32())
				} else {
					return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CONTACT_ID_INVALID)
				}
			}

			ret = append(ret, messageMedia.To_MessageMedia())
			break
		case mtproto.ClassInputMediaVenue:
			//  inputMediaVenue#c13d1c11 geo_point:InputGeoPoint title:string address:string provider:string venue_id:string venue_type:string = InputMedia;
			//  inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;

			// messageMediaVenue#2ec0533f geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string = MessageMedia;
			// messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
			venue := media.To_InputMediaVenue()
			messageMedia := mtproto.NewTLMessageMediaVenue(&mtproto.MessageMedia{
				Geo:       InputGeoPointToGeoPoint(venue.GetGeoPoint()),
				Title:     venue.GetTitle(),
				Address:   venue.GetAddress(),
				Provider:  venue.GetProvider(),
				VenueId:   venue.GetVenueId(),
				VenueType: venue.GetVenueType(),
			})

			ret = append(ret, messageMedia.To_MessageMedia())
			break
		case mtproto.ClassInputMediaGifExternal, mtproto.ClassInputMediaDocumentExternal, mtproto.ClassInputMediaPhotoExternal:
			// inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
			//mediaGif := mtproto.NewTLMessageMediaWebPage(&mtproto.MessageMedia{
			//	Webpage:
			//}).To_MessageMedia()
			ret = append(ret, mtproto.NewTLMessageMediaUnsupported(nil).To_MessageMedia())
			break
		case mtproto.ClassInputMediaGame:
			// inputMediaGame#d33f43f3 id:InputGame = InputMedia;
			// game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
			// messageMediaGame#fdb19008 game:Game = MessageMedia;
			inputMediaGame := media.To_InputMediaGame()
			inputGameId := inputMediaGame.GetIdD33F43F371()

			//TODO: GetGameBy ID
			ret = append(ret, mtproto.NewTLMessageMediaGame(&mtproto.MessageMedia{
				Game: mtproto.NewTLGame(&mtproto.Game{
					Id:          inputGameId.GetId(),
					AccessHash:  inputGameId.GetAccessHash(),
					ShortName:   inputGameId.GetShortName(),
					Title:       "",
					Description: "",
					Photo:       mtproto.NewTLPhotoEmpty(nil).To_Photo(),
					Document:    mtproto.NewTLDocumentEmpty(nil).To_Document(),
				}).To_Game(),
			}).To_MessageMedia())

		case mtproto.ClassInputMediaInvoice:
			//  inputMediaInvoice#f4e096c3 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string provider_data:DataJSON start_param:string = InputMedia;
			//  inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;

			// messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;

			var webDocument *mtproto.WebDocument
			if media.Photo != nil {
				webDocument = mtproto.NewTLWebDocument(&mtproto.WebDocument{
					Url:        media.Photo.Url,
					AccessHash: 0, //TODO:(Coder)
					Size_:      media.Photo.Size_,
					MimeType:   media.Photo.MimeType,
					Attributes: media.Photo.Attributes,
					DcId:       help.GetConfig().ThisDc,
				}).To_WebDocument()
			}

			// TODO: Invoice
			ret = append(ret, mtproto.NewTLMessageMediaInvoice(&mtproto.MessageMedia{
				ShippingAddressRequested: false,
				Test:                     false,
				Title:                    media.Title,
				Description:              media.Description,
				Photo8455134771:          webDocument,
				TotalAmount:              0,
				StartParam:               media.StartParam,
			}).To_MessageMedia())
			break

		case mtproto.ClassInputMediaGeoLive:
			//  inputMediaGeoLive#971fa843 flags:# stopped:flags.0?true geo_point:InputGeoPoint heading:flags.2?int period:flags.1?int proximity_notification_radius:flags.3?int = InputMedia;
			//  messageMediaGeoLive#b940c666 flags:# geo:GeoPoint heading:flags.0?int period:int proximity_notification_radius:flags.1?int = MessageMedia;
			messageMedia := mtproto.NewTLMessageMediaGeoLive(&mtproto.MessageMedia{
				Geo:                         InputGeoPointToGeoPoint(media.To_InputMediaGeoLive().GetGeoPoint()),
				Period:                      media.To_InputMediaGeoLive().GetPeriod(),
				Heading:                     media.Heading,
				ProximityNotificationRadius: media.ProximityNotificationRadius,
			})
			ret = append(ret, messageMedia.To_MessageMedia())
			break

		case mtproto.ClassInputMediaPoll:
			//TODO:(Coder) 投票

			//  pollResults#badcc1a3 flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int recent_voters:flags.3?Vector<int> solution:flags.4?string solution_entities:flags.4?Vector<MessageEntity> = PollResults;
			mediaPoll := mtproto.NewTLMessageMediaPoll(&mtproto.MessageMedia{
				Poll:    media.Poll,
				Results: mtproto.NewTLPollResults(nil).To_PollResults(),
			}).To_MessageMedia()
			mediaPoll.Poll.Closed = false

			ret = append(ret, mediaPoll)
			break
		case mtproto.ClassInputMediaEmpty:
			ret = append(ret, mtproto.NewTLMessageMediaEmpty(nil).To_MessageMedia())
			break
		}
	}

	return ret, nil
}

func (c *Core) makeMedia(userId int64, authKeyId int64, layer int32, peer *input.InputPeer, request *mtproto.TLMessagesSendMedia) (*mtproto.Message, error) {
	media, err := c.makeInputMedia(authKeyId, layer, []*mtproto.InputMedia{request.GetMedia()})
	if err != nil {
		log.Errorf("makeSendMedia peer:%v error:%s", peer, err.Error())
		return nil, err
	}

	message := mtproto.NewTLMessage(&mtproto.Message{
		Out:              true,
		Silent:           request.GetSilent(),
		FromId90DDDC1171: constants.PeerTypeFromUserIDType(userId).ToInt32(),
		FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{
			UserId: constants.PeerTypeFromUserIDType(userId).ToInt32(),
		}).To_Peer(),
		PeerId:      peer.ToPeer(),
		ToId:        peer.ToPeer(),
		Message:     request.GetMessage(),
		Media:       media[0],
		ReplyMarkup: request.GetReplyMarkup(),
		Date:        int32(time.Now().Unix()),
		Entities:    request.Entities,
	})
	return message.To_Message(), nil
}

func (c *Core) makeMultiMedia(userId int64, authKeyId int64, layer int32, peer *input.InputPeer, request *mtproto.TLMessagesSendMultiMedia) ([]*mtproto.Message, error) {

	multiMedia := request.MultiMedia
	messages := make([]*mtproto.Message, 0, len(multiMedia))

	inputMedias := make([]*mtproto.InputMedia, 0, len(multiMedia))
	for _, v := range multiMedia {
		inputMedias = append(inputMedias, v.Media)
	}
	medias, err := c.makeInputMedia(authKeyId, layer, inputMedias)
	if err != nil {
		return nil, err
	}

	for idx, media := range multiMedia {

		message := mtproto.NewTLMessage(&mtproto.Message{
			Out:               true,
			Silent:            request.GetSilent(),
			FromId90DDDC1171:  constants.PeerTypeFromUserIDType(userId).ToInt32(),
			FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer(),
			PeerId:            peer.ToPeer(),
			ToId:              peer.ToPeer(),
			Media:             medias[idx],
			Date:              int32(time.Now().Unix()),
			GroupedId:         multiMedia[0].RandomId,
			Entities:          media.Entities,
		})

		messages = append(messages, message.To_Message())
	}

	return messages, nil
}
