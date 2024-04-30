/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : users.getFullUser_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	accountPrivacy "novachat_engine/service/account/privacy"
	"novachat_engine/service/account/setting"
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/app/help"
	"novachat_engine/service/app/support"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/privacy"
	data_setting "novachat_engine/service/data/setting"
	"novachat_engine/service/input"
	"novachat_engine/service/notify_setting"
	"novachat_engine/service/upload/photo"
)

//  users.getFullUser#ca30a5b1 id:InputUser = UserFull;
//
func (s *UsersServiceImpl) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.UserFull, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UsersGetFullUser %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	inputUser := input.MakeInputUser(request.Id)
	switch inputUser.GetInputUserType() {
	case constants.InputUserTypeEmpty, constants.InputUserTypeUserFromMessage:
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("UsersGetFullUser - request: %v inputPeer:%v error:%s", request, err.Error())
		return nil, err
	default:
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("UsersGetFullUser - request: %v inputPeer:%v error:%s", request, err.Error())
		return nil, err

	case constants.InputUserTypeSelf, constants.InputUserTypeUser:
		if inputUser.GetInputUserType() == constants.InputUserTypeSelf {
			inputUser = input.MakeInputUser(mtproto.NewTLInputUser(&mtproto.InputUser{
				UserId:     constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
				AccessHash: request.Id.AccessHash,
			}).To_InputUser())
		}
	}

	user, err := s.accountUserCore.GetUser(md.UserId, inputUser.GetId())
	if err != nil {
		log.Errorf("UsersGetFullUser - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}
	userInfo, err := s.accountUserCore.UserData(inputUser.GetId())
	if err != nil {
		log.Errorf("UsersGetFullUser - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}
	contact, err := s.accountContactCore.GetContactById(md.UserId, inputUser.GetId())
	if err != nil {
		log.Errorf("UsersGetFullUser - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	userFull := mtproto.NewTLUserFull(nil)
	var myLink, foreignLink *mtproto.ContactLink
	if constants.PeerTypeFromUserIDType32(user.Id).ToInt() == md.UserId {
		user = usersUtil.UserCoreSelfUsers(user)
		myLink = mtproto.NewTLContactLinkContact(nil).To_ContactLink()
		foreignLink = mtproto.NewTLContactLinkContact(nil).To_ContactLink()
		userFull.SetPhoneCallsAvailable(false)
		userFull.SetPhoneCallsPrivate(false)

		var userPhoto *mtproto.Photo
		if user.Photo != nil {
			photoInfo, err1 := sfsService.GetSfsClient(fmt.Sprintf("%d", user.Photo.PhotoId)).
				ReqGetPhoto(ctx, &sfsService.GetPhoto{PhotoId: user.Photo.PhotoId})
			if err1 != nil {
				log.Warnf("UsersGetFullUser - request: %v GetPhotoFile error:%s", request, err1.Error())
			}

			if photoInfo != nil {
				userPhoto = photo.PhotoInfo2Photo(photoInfo, md.Layer)
			}
		}
		userFull.SetProfilePhoto(userPhoto)
	} else {
		privacyConf, err := s.accountPrivacyCore.GetPrivacyKeys(inputUser.GetId(), []int32{constants.KeyPrivacyAddByPhone.Int32(), constants.KeyPrivacyPhoneCall.Int32(), constants.KeyPrivacyProfilePhoto.Int32()})
		if err != nil {
			log.Errorf("UsersGetFullUser - request: %v GetPrivacy error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		}

		if accountPrivacy.CheckPrivacy(md.UserId, 0, privacy.TokenPrivacy(privacyConf, constants.KeyPrivacyAddByPhone.Int32()), user.Contact) {
			foreignLink = mtproto.NewTLContactLinkHasPhone(nil).To_ContactLink()
		} else {
			foreignLink = mtproto.NewTLContactLinkUnknown(nil).To_ContactLink()
		}
		myLink = mtproto.NewTLContactLinkHasPhone(nil).To_ContactLink()
		userFull.SetBlocked(contact.Block)

		canPhoneCall := accountPrivacy.CheckPrivacy(md.UserId, 0, privacy.TokenPrivacy(privacyConf, constants.KeyPrivacyPhoneCall.Int32()), user.Contact)
		userFull.SetPhoneCallsAvailable(canPhoneCall && s.phoneCallConfig.PhoneCallsAvailableConf() && help.GetConfig().PhonecallsEnabled)
		userFull.SetPhoneCallsPrivate(canPhoneCall && s.phoneCallConfig.PhoneCallsPrivateConf() && help.GetConfig().PhonecallsEnabled)
		userFull.SetVideoCallsAvailable(canPhoneCall && s.videoCall.VideoCallAvailableConf() && help.GetConfig().PhonecallsEnabled)

		if accountPrivacy.CheckPrivacy(md.UserId, 0, privacy.TokenPrivacy(privacyConf, constants.KeyPrivacyProfilePhoto.Int32()), user.Contact) {
			var userPhoto *mtproto.Photo
			if user.Photo != nil {
				photoInfo, err1 := sfsService.GetSfsClient(fmt.Sprintf("%d", user.Photo.PhotoId)).
					ReqGetPhoto(ctx, &sfsService.GetPhoto{PhotoId: user.Photo.PhotoId})
				if err1 != nil {
					log.Warnf("UsersGetFullUser - request: %v GetPhotoFile error:%s", request, err1.Error())
				}

				if photoInfo != nil {
					userPhoto = photo.PhotoInfo2Photo(photoInfo, md.Layer)
				}
			}
			userFull.SetProfilePhoto(userPhoto)
		} else {
			userFull.SetProfilePhoto(nil)
		}
	}

	userFull.SetLink(mtproto.NewTLContactsLink(&mtproto.Contacts_Link{
		User:        user,
		MyLink:      myLink,
		ForeignLink: foreignLink,
	}).To_Contacts_Link())
	userFull.SetAbout(userInfo.About)

	userFull.SetUser(compat.CompatUser(user, md.Layer))

	if support.IsSupportUser(inputUser.GetId()) {
		userFull.SetPhoneCallsAvailable(false)
		userFull.SetPhoneCallsPrivate(false)
	}

	var settingV *data_setting.NotifySetting
	if md.UserId == userInfo.Id {
		userFull.SetBotInfo(nil)
		userFull.SetSettings(s.accountSettingCore.PeerSettings(md.UserId, inputUser.GetId()))
		settingV, err = s.accountSettingCore.GetNotifySetting(
			&setting.NotifyPeerSettingType{
				UserId:         md.UserId,
				PeerId:         inputUser.GetId(),
				PeerType:       constants.PeerTypeUser,
				PeerNotifyType: constants.PeerGlobalSetting,
			})
		if err != nil {
			log.Errorf("UsersGetFullUser request:%v error:%s", request, err.Error())
			return nil, err
		}
	} else {
		settingV, err = s.accountSettingCore.GetNotifySetting(
			&setting.NotifyPeerSettingType{
				UserId:         md.UserId,
				PeerId:         inputUser.GetId(),
				PeerType:       constants.PeerTypeUser,
				PeerNotifyType: constants.PeerNotifyInputNotifyPeer,
			})
		if err != nil {
			log.Errorf("UsersGetFullUser request:%v error:%s", request, err.Error())
			return nil, err
		}

		peerSetting := s.accountSettingCore.PeerSettings(md.UserId, inputUser.GetId())
		if contact != nil {
			peerSetting.AddContact = contact.PeerId != inputUser.GetId()
			peerSetting.ShareContact = contact.PeerId == inputUser.GetId()
			peerSetting.BlockContact = contact.PeerId == inputUser.GetId()
		}
		userFull.SetSettings(peerSetting)
	}
	userFull.SetNotifySettings(notify_setting.ToPeerNotifySettings(settingV, md.Layer))

	//TODO:(Coder)
	// bot
	if userInfo.GetBot() {
		userFull.SetBotInfo(s.accountBotCore.BotInfo(inputUser.GetId()))
		userFull.SetCanPinMessage(false)
	} else {
		userFull.SetCanPinMessage(md.UserId != inputUser.GetId())
	}

	//TODO:Coder
	// conversation FolderId
	//userFull.SetFolderId(conversation.FolderId)
	//TODO:(Coder)
	// show scheduled.md
	userFull.SetHasScheduled(false)

	//TODO:Coder
	// Chats Common
	userFull.SetCommonChatsCount(0)

	log.Infof("UsersGetFullUser - request: %v  userFull:%+v reply ok!!!!!!!!!!!!", request, userFull)
	return userFull.To_UserFull(), nil
}
