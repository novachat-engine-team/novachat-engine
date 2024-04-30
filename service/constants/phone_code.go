/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/25 18:10
 * @File : phone_code.go
 */

package constants

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
)

//# udp_relay 2.4.4/2.5.0
//# webrtc 2.7.7 3.0.0
//# rtc 2.8.8
const (
	PhoneVersionUdpRelay int32 = 0
	PhoneVersionWebrtc   int32 = 1
	PhoneVersionRtc      int32 = 2

	PhoneVersion2_4_4 string = "2.4.4"
	PhoneVersion2_5_0 string = "2.5.0"
	PhoneVersion2_7_7 string = "2.7.7"
	PhoneVersion2_8_8 string = "2.8.8"
	PhoneVersion3_0_0 string = "3.0.0"
)

// auth_SentCodeType <--
//  + TL_auth_sentCodeTypeApp
//  + TL_auth_sentCodeTypeSms
//  + TL_auth_sentCodeTypeCall
//  + TL_auth_sentCodeTypeFlashCall

type CodeType int32

const (
	CodeTypeNone      CodeType = 0
	CodeTypeApp       CodeType = 1
	CodeTypeSms       CodeType = 2
	CodeTypeCall      CodeType = 3
	CodeTypeFlashCall CodeType = 4
)

func (m CodeType) ToInt32() int32 {
	return int32(m)
}

func CodeTypeFromInt32(v int32) CodeType {
	return CodeType(v)
}

func MakeAuthCode(t CodeType) *mtproto.Auth_CodeType {
	switch t {
	case CodeTypeSms, CodeTypeApp:
		return mtproto.NewTLAuthCodeTypeSms(nil).To_Auth_CodeType()
	case CodeTypeCall:
		return mtproto.NewTLAuthCodeTypeCall(nil).To_Auth_CodeType()
	case CodeTypeFlashCall:
		return mtproto.NewTLAuthCodeTypeFlashCall(nil).To_Auth_CodeType()
	default:
		log.Errorf("MakeAuthCode CodeType value:%d error", t.ToInt32())
		return nil
	}
}

func MakeSendCodeType(t CodeType, codeLength int32, pattern string) *mtproto.Auth_SentCodeType {
	v := &mtproto.Auth_SentCodeType{}
	v.Length = codeLength
	v.Pattern = pattern

	switch t {
	case CodeTypeApp:
		return v.To_AuthSentCodeTypeApp().To_Auth_SentCodeType()
	case CodeTypeSms:
		return v.To_AuthSentCodeTypeSms().To_Auth_SentCodeType()
	case CodeTypeCall:
		return v.To_AuthSentCodeTypeCall().To_Auth_SentCodeType()
	case CodeTypeFlashCall:
		return v.To_AuthSentCodeTypeFlashCall().To_Auth_SentCodeType()
	default:
		log.Errorf("MakeSendCodeType CodeType:%d", t.ToInt32())
		return nil
	}
}
