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
	"novachat_engine/service/data/messages/message"
)

///////////////////////////////////////////////////////////////////////////////
// MessageAction <--
//  + TL_messageActionEmpty
//  + TL_messageActionChatCreate
//  + TL_messageActionChatEditTitle
//  + TL_messageActionChatEditPhoto
//  + TL_messageActionChatDeletePhoto
//  + TL_messageActionChatAddUser
//  + TL_messageActionChatDeleteUser
//  + TL_messageActionChatJoinedByLink
//  + TL_messageActionChannelCreate
//  + TL_messageActionChatMigrateTo
//  + TL_messageActionChannelMigrateFrom
//  + TL_messageActionPinMessage
//  + TL_messageActionHistoryClear
//  + TL_messageActionGameScore
//  + TL_messageActionPaymentSentMe
//  + TL_messageActionPaymentSent
//  + TL_messageActionScreenshotTaken
//  + TL_messageActionCustomAction
//  + TL_messageActionBotAllowed
//  + TL_messageActionSecureValuesSentMe
//  + TL_messageActionSecureValuesSent
//  + TL_messageActionContactSignUp
//  + TL_messageActionPhoneCall
//  + TL_messageActionGeoProximityReached
//  + TL_messageActionGroupCall
//  + TL_messageActionInviteToGroupCall
//

func toMessageActionType(a *mtproto.MessageAction) constants.MessageAction {
	switch a.ClassName {
	case mtproto.ClassMessageActionEmpty:
		return constants.MessageActionEmpty
	case mtproto.ClassMessageActionChatCreate:
		return constants.MessageActionChatCreate
	case mtproto.ClassMessageActionChatEditTitle:
		return constants.MessageActionChatEditTitle
	case mtproto.ClassMessageActionChatEditPhoto:
		return constants.MessageActionChatEditPhoto
	case mtproto.ClassMessageActionChatDeletePhoto:
		return constants.MessageActionChatDeletePhoto
	case mtproto.ClassMessageActionChatAddUser:
		return constants.MessageActionChatAddUser
	case mtproto.ClassMessageActionChatDeleteUser:
		return constants.MessageActionChatDeleteUser
	case mtproto.ClassMessageActionChatJoinedByLink:
		return constants.MessageActionChatJoinedByLink
	case mtproto.ClassMessageActionChannelCreate:
		return constants.MessageActionChannelCreate
	case mtproto.ClassMessageActionChatMigrateTo:
		return constants.MessageActionChatMigrateTo
	case mtproto.ClassMessageActionChannelMigrateFrom:
		return constants.MessageActionChannelMigrateFrom
	case mtproto.ClassMessageActionPinMessage:
		return constants.MessageActionPinMessage
	case mtproto.ClassMessageActionHistoryClear:
		return constants.MessageActionHistoryClear
	case mtproto.ClassMessageActionGameScore:
		return constants.MessageActionGameScore
	case mtproto.ClassMessageActionPaymentSentMe:
		return constants.MessageActionPaymentSentMe
	case mtproto.ClassMessageActionPaymentSent:
		return constants.MessageActionPaymentSent
	case mtproto.ClassMessageActionScreenshotTaken:
		return constants.MessageActionScreenshotTaken
	case mtproto.ClassMessageActionCustomAction:
		return constants.MessageActionCustomAction
	case mtproto.ClassMessageActionBotAllowed:
		return constants.MessageActionBotAllowed
	case mtproto.ClassMessageActionSecureValuesSentMe:
		return constants.MessageActionSecureValuesSentMe
	case mtproto.ClassMessageActionSecureValuesSent:
		return constants.MessageActionSecureValuesSent
	case mtproto.ClassMessageActionContactSignUp:
		return constants.MessageActionContactSignUp
	case mtproto.ClassMessageActionPhoneCall:
		return constants.MessageActionPhoneCall
	case mtproto.ClassMessageActionGeoProximityReached:
		return constants.MessageActionGeoProximityReached
	case mtproto.ClassMessageActionGroupCall:
		return constants.MessageActionGroupCall
	case mtproto.ClassMessageActionInviteToGroupCall:
		return constants.MessageActionInviteToGroupCall
	default:
		panic(a.ClassName)
	}
}

func toMessageAction(dataAction *data_message.Action) *mtproto.MessageAction {
	if dataAction == nil {
		return nil
	}

	var action *mtproto.MessageAction
	//TODO:Coder
	// dataAction to action
	switch constants.MessageAction(dataAction.Type) {
	case constants.MessageActionEmpty:
		return mtproto.NewTLMessageActionEmpty(action).To_MessageAction()
	case constants.MessageActionChatCreate:
		return mtproto.NewTLMessageActionChatCreate(action).To_MessageAction()
	case constants.MessageActionChatEditTitle:
		return mtproto.NewTLMessageActionChatEditTitle(action).To_MessageAction()
	case constants.MessageActionChatEditPhoto:
		return mtproto.NewTLMessageActionChatEditPhoto(action).To_MessageAction()
	case constants.MessageActionChatDeletePhoto:
		return mtproto.NewTLMessageActionChatDeletePhoto(action).To_MessageAction()
	case constants.MessageActionChatAddUser:
		return mtproto.NewTLMessageActionChatAddUser(action).To_MessageAction()
	case constants.MessageActionChatDeleteUser:
		return mtproto.NewTLMessageActionChatDeleteUser(action).To_MessageAction()
	case constants.MessageActionChatJoinedByLink:
		return mtproto.NewTLMessageActionChatJoinedByLink(action).To_MessageAction()
	case constants.MessageActionChannelCreate:
		return mtproto.NewTLMessageActionChannelCreate(action).To_MessageAction()
	case constants.MessageActionChatMigrateTo:
		return mtproto.NewTLMessageActionChatMigrateTo(action).To_MessageAction()
	case constants.MessageActionChannelMigrateFrom:
		return mtproto.NewTLMessageActionChannelMigrateFrom(action).To_MessageAction()
	case constants.MessageActionPinMessage:
		return mtproto.NewTLMessageActionPinMessage(action).To_MessageAction()
	case constants.MessageActionHistoryClear:
		return mtproto.NewTLMessageActionHistoryClear(action).To_MessageAction()
	case constants.MessageActionGameScore:
		return mtproto.NewTLMessageActionGameScore(action).To_MessageAction()
	case constants.MessageActionPaymentSentMe:
		return mtproto.NewTLMessageActionPaymentSentMe(action).To_MessageAction()
	case constants.MessageActionPaymentSent:
		return mtproto.NewTLMessageActionPaymentSent(action).To_MessageAction()
	case constants.MessageActionScreenshotTaken:
		return mtproto.NewTLMessageActionScreenshotTaken(action).To_MessageAction()
	case constants.MessageActionCustomAction:
		return mtproto.NewTLMessageActionCustomAction(action).To_MessageAction()
	case constants.MessageActionBotAllowed:
		return mtproto.NewTLMessageActionBotAllowed(action).To_MessageAction()
	case constants.MessageActionSecureValuesSentMe:
		return mtproto.NewTLMessageActionSecureValuesSentMe(action).To_MessageAction()
	case constants.MessageActionSecureValuesSent:
		return mtproto.NewTLMessageActionSecureValuesSent(action).To_MessageAction()
	case constants.MessageActionContactSignUp:
		return mtproto.NewTLMessageActionContactSignUp(action).To_MessageAction()
	case constants.MessageActionPhoneCall:
		return mtproto.NewTLMessageActionPhoneCall(action).To_MessageAction()
	case constants.MessageActionGeoProximityReached:
		return mtproto.NewTLMessageActionGeoProximityReached(action).To_MessageAction()
	case constants.MessageActionGroupCall:
		return mtproto.NewTLMessageActionGroupCall(action).To_MessageAction()
	case constants.MessageActionInviteToGroupCall:
		return mtproto.NewTLMessageActionInviteToGroupCall(action).To_MessageAction()
	default:
		panic(fmt.Sprintf("toMessageAction type:%v", dataAction.Type))
	}
}

func messageActonUtil(action *mtproto.MessageAction) *data_message.Action {
	if action == nil {
		return nil
	}

	//TODO:Coder
	// Action
	return &data_message.Action{
		Type: toMessageActionType(action).ToInt32(),
	}
}
