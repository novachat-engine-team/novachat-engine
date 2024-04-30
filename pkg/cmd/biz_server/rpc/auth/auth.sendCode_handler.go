/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.sendCode_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/random2"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/app/help"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/phone"
	"novachat_engine/service/sms"
)

//Parameters
//Name	Type	Description
//phone_number	string	Phone number in international format
//api_id	int	Application identifier (see App configuration)
//api_hash	string	Application secret hash (see App configuration)
//settings	CodeSettings	Settings for the code type to send
//Result
//The method returns an auth.SentCode object with information on the message sent.
//
//Possible errors
//Code	Type	Description
//400	API_ID_INVALID	API ID invalid
//400	API_ID_PUBLISHED_FLOOD	This API id was published somewhere, you can't use it now
//401	AUTH_KEY_PERM_EMPTY	The temporary auth key must be binded to the permanent auth key to use these methods.
//400	INPUT_REQUEST_TOO_LONG	The request is too big
//303	NETWORK_MIGRATE_X	Repeat the query to data-center X
//303	PHONE_MIGRATE_X	Repeat the query to data-center X
//400	PHONE_NUMBER_APP_SIGNUP_FORBIDDEN	You can't sign up using this app
//400	PHONE_NUMBER_BANNED	The provided phone number is banned from telegram
//400	PHONE_NUMBER_FLOOD	You asked for the code too many times.
//400	PHONE_NUMBER_INVALID	Invalid phone number
//406	PHONE_PASSWORD_FLOOD	You have tried logging in too many times
//400	PHONE_PASSWORD_PROTECTED	This phone is password protected
//400	SMS_CODE_CREATE_FAILED	An error occurred while creating the SMS code

//  auth.sendCode#a677244f phone_number:string api_id:int api_hash:string settings:CodeSettings = auth.SentCode;
//  auth.sendCode#86aef0ec flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool api_id:int api_hash:string = auth.SentCode;
//
func (s *AuthServiceImpl) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthSendCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	conf := help.GetConfig()

	thisDC := dc.GetDCByApiID(request.ApiId, request.ApiHash)
	if conf.ThisDc != thisDC {
		log.Warnf("AuthSendCode - request: %v thisDC:%d ApiID dc:%d", request, conf.ThisDc, thisDC)
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_SEE_OTHER_NETWORK_MIGRATE.ToInt32(), fmt.Sprintf("NETWORK_MIGRATE_%d", thisDC))
	}

	phoneNumber, err := phone.Parse(request.PhoneNumber)
	if err != nil {
		log.Errorf("AuthSendCode - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
	}

	if s.banned.InvalidNumber(phoneNumber.NormalizeDigitsOnly()) {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_BANNED)
	}

	phoneCodeHash := random2.RandomAlphaOrNumeric(10, true, true)
	phoneCode := sms.DefaultPhoneCode

	if s.conf.EnvMode == constants.EnvModelTest {
		phoneCode = sms.DefaultPhoneCode
	} else {
		//if config.SendPhoneCodeIgnore(s.conf.Ignore.Phones, phoneNumber.NormalizeDigitsOnly()) {
		//    phoneCode = sms.DefaultPhoneCode
		//} else {
		//    if phone.IsVirtualPhone(request.PhoneNumber) {
		//        return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_BANNED)
		//    }
		phoneCode = random2.RandomAlphaOrNumeric(sms.CodeLength, false, true)
		//}
	}
	//if !isVirtualPhone && s.conf.BizLogic.EnvMode != constants.EnvModelTest || !config.SendPhoneCodeIgnore(s.conf.Ignore.Phones, phoneNumber.NormalizeDigitsOnly()){
	//    phoneCode = random2.RandomAlphaOrNumeric(sms.CodeLength, false,true)
	//}

	request.AllowFlashcall = false
	if request.Settings != nil {
		request.Settings.AllowFlashcall = false
	}
	reply := mtproto.NewTLAuthSentCode(nil)
	reply.SetPhoneCodeHash(phoneCodeHash)
	reply.SetTimeout(sms.TTL)

	if request.AllowFlashcall || request.Settings != nil && request.Settings.AllowFlashcall {
		reply.SetNextType(constants.MakeAuthCode(constants.CodeTypeFlashCall))
		reply.SetType(constants.MakeSendCodeType(constants.CodeTypeFlashCall, sms.CodeLength, "*"))
	} else {
		reply.SetNextType(constants.MakeAuthCode(constants.CodeTypeApp))
		reply.SetType(constants.MakeSendCodeType(constants.CodeTypeSms, sms.CodeLength, "*"))
	}

	log.Infof("AuthSendCode - authKeyId:%d phoneNumber:%s phoneCode:%s, phoneCodeHash:%s", md.AuthKeyId, request.PhoneNumber, phoneCode, phoneCodeHash)

	if !phone.IsVirtualPhone(request.PhoneNumber) {
		sms.SendSmsVerifyCode(sms.CodeSignIn, phoneNumber.NormalizeDigitsOnly(), phoneCode, phoneCodeHash)
	}

	if false {
		////TODO:(Coder)  check status
		//userId := constants.OfficialSupportUserID
		//peer := mtproto.NewTLPeerUser(nil)
		//peer.SetUserId(userId)
		//text := fmt.Sprintf(lang.GetCodeText(phoneNumber.GetRegionCode()), s.conf.SmsConfig.Title, phoneCode)
		//message := mtproto.NewTLMessage(&mtproto.Message{
		//	Out:              true,
		//	Date:             int32(time.Now().Unix()),
		//	Message:          text,
		//	FromId90DDDC1171: userId,
		//	ToId:             peer.To_Peer(),
		//})
		//_, err = msgClient.GetMsgClient().SendUserMessage(context.Background(), &msg_pb.SendUserMessage{
		//	Data: &msg_pb.SendMessage{
		//		AuthKeyId:  md.AuthKeyId,
		//		FromUserId: constants.OfficialSupportUserID,
		//		Peer: mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
		//			UserId: userId,
		//		}).To_InputPeer(),
		//		Data: &msg_pb.SendMessageData{
		//			RandomId: rand.Int63(),
		//			Message:  message.To_Message(),
		//		},
		//	},
		//})
		//if err != nil {
		//	log.Warnf("AuthSendCode SendUserMessage error:%s", err.Error())
		//}
	}
	log.Debugf("AuthSendCode - request: %v ok!!!", request)
	return reply.To_Auth_SentCode(), nil
}
