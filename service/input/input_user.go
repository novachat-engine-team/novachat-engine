/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/25 18:19
 * @File : input_user.go
 */

package input

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

type InputUser struct {
	inputUserType constants.InputUserType
	id            int64
	accessHash    int64
}

func (m *InputUser) GetInputUserType() constants.InputUserType {
	return m.inputUserType
}

func (m *InputUser) GetId() int64 {
	return m.id
}

func (m *InputUser) GetAccessHash() int64 {
	return m.accessHash
}

func (m *InputUser) ToInputUser() *mtproto.InputUser {
	switch m.inputUserType {
	case constants.InputUserTypeEmpty:
		return mtproto.NewTLInputUserEmpty(nil).To_InputUser()
	case constants.InputUserTypeSelf:
		return mtproto.NewTLInputUserSelf(nil).To_InputUser()
	case constants.InputUserTypeUser:
		return mtproto.NewTLInputUser(nil).To_InputUser()
	case constants.InputUserTypeUserFromMessage:
		return mtproto.NewTLInputUserFromMessage(nil).To_InputUser()
	default:
		log.Errorf("InputUser ToInputUser error:%v", m.inputUserType)
		return nil
	}
}

func MakeInputUser(inputUser *mtproto.InputUser) *InputUser {
	p := &InputUser{}

	switch inputUser.ClassName {
	case mtproto.ClassInputUserEmpty:
		p.inputUserType = constants.InputUserTypeEmpty
	case mtproto.ClassInputUserSelf:
		p.inputUserType = constants.InputUserTypeSelf
		p.id = constants.PeerTypeFromUserIDType32(inputUser.UserId).ToInt()
		p.accessHash = inputUser.AccessHash
	case mtproto.ClassInputUser:
		p.inputUserType = constants.InputUserTypeUser
		p.id = constants.PeerTypeFromUserIDType32(inputUser.UserId).ToInt()
		p.accessHash = inputUser.AccessHash
	case mtproto.ClassInputUserFromMessage:
		p.inputUserType = constants.InputUserTypeUserFromMessage
		p.id = constants.PeerTypeFromUserIDType32(inputUser.MsgId).ToInt()
	default:
		log.Errorf("MakeInputUser error:%v", inputUser)
		return nil
	}
	return p
}
