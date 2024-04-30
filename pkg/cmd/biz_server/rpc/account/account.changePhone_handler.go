/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.changePhone_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/phone"
	"novachat_engine/service/sms"
	"strings"
	"time"
)

//  account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
//
func (s *AccountServiceImpl) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountChangePhone %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	phoneNumber := strings.TrimSpace(request.PhoneNumber)
	var err error
	if len(phoneNumber) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EMPTY)
		log.Errorf("AccountChangePhone - request:%v error:%s", request, err.Error())
		return nil, err
	}

	phoneNumberParser, err := phone.Parse(phoneNumber)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
		log.Errorf("AccountChangePhone - request:%v error:%s", request, err.Error())
		return nil, err
	}

	phoneNumber = phoneNumberParser.NormalizeDigitsOnly()
	userInfo, err := s.accountUsersCore.PhoneUser(phoneNumber)
	if err != nil {
		log.Errorf("AccountChangePhone UserByPhoneNumber request:%v  error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}

	if userInfo.Id == md.UserId {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
		log.Errorf("AccountChangePhone - request:%s error:%s", request, err.Error())
		return nil, err
	}

	if userInfo.Id != 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_OCCUPIED)
		log.Errorf("AccountChangePhone - request:%s SendSmsCode error:%s", request, err.Error())
		return nil, err
	}

	err = sms.VerifySmsCode(sms.CodeChange, phoneNumber, request.PhoneCode, request.PhoneCodeHash)
	if err != nil {
		log.Errorf("AccountChangePhone - request: %s VerifySmsCode error:%s", request, err.Error())
		return nil, err
	}
	defer func() {
		sms.ClearSmsCode(sms.CodeChange, phoneNumber, request.PhoneCode, request.PhoneCodeHash)
	}()

	_, err = s.usersCore.UpdateBindPhone(md.UserId, phoneNumber)
	if err != nil {
		log.Errorf("AccountChangePhone %s, request: %s UpdateBindPhone phone:%s phoneCode:%s error:%s", metadata.RpcMetaDataDebug(md), request, phoneNumber, request.PhoneCode, err.Error())
		return nil, err
	}

	user := usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(userInfo))
	user.Phone = phoneNumber

	//TODO:Coder
	go func() {
		_, err = syncClient.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncClient.SyncUpdate{
			UserId:          md.UserId,
			IgnoreAuthKeyId: md.AuthKeyId,
			Push:            true,
			PeerType:        constants.PeerTypeUser.ToInt32(),
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{
					mtproto.NewTLUpdateUserPhone(&mtproto.Update{
						UserId: constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
						Phone:  phoneNumber,
					}).To_Update(),
				},
				Chats: []*mtproto.Chat{},
				Users: []*mtproto.User{},
				Date:  int32(time.Now().Unix()),
				Seq:   0,
			}).To_Updates(),
		})
		if err != nil {
			log.Warnf("AccountChangePhone - sync error:%s", err.Error())
		}
	}()
	return user, nil
}
