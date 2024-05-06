/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/10 15:19
 * @File : const_util.go
 */

package constants

import (
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
)

func ArrayInt64To32(v []int64) []int32 {
	vv := make([]int32, 0, len(v))
	for _, v1 := range v {
		vv = append(vv, PeerTypeFromUserIDType(v1).ToInt32())
	}
	return vv
}

func ArrayInt32To64(v []int32) []int64 {
	vv := make([]int64, 0, len(v))
	for _, v1 := range v {
		vv = append(vv, PeerTypeFromUserIDType32(v1).ToInt())
	}
	return vv
}

func PeerTypeFromUserIDType32(v int32) UserIDType {
	return UserIDType(v)
}

func PeerTypeFromUserIDType(v int64) UserIDType {
	return UserIDType(v)
}

func (m UserIDType) ToInt32() int32 {
	return int32(m)
}

func (m UserIDType) ToInt() int64 {
	return int64(m)
}

func PeerTypeFromChannelIDType32(v int32) ChannelIDType {
	return ChannelIDType(v)
}

func PeerTypeFromChannelIDType(v int64) ChannelIDType {
	return ChannelIDType(v)
}

func (m ChannelIDType) ToInt32() int32 {
	return int32(m)
}

func (m ChannelIDType) ToInt() int64 {
	return int64(m)
}

func PeerTypeFromChatIDType32(v int32) ChatIDType {
	return ChatIDType(v)
}

func PeerTypeFromChatIDType(v int64) ChatIDType {
	return ChatIDType(v)
}

func (m ChatIDType) ToInt32() int32 {
	return int32(m)
}

func (m ChatIDType) ToInt() int64 {
	return int64(m)
}

func PeerTypeFromInt32(v int32) PeerType {
	return PeerType(v)
}

func (m PeerType) ToInt32() int32 {
	return int32(m)
}

func PeerNotifyTypeFromInt32(v int32) PeerNotifyType {
	return PeerNotifyType(v)
}

func (m PeerNotifyType) ToInt32() int32 {
	return int32(m)
}

func RegisterDeviceTypeFromInt32(v int32) RegisterDeviceType {
	return RegisterDeviceType(v)
}

func (m RegisterDeviceType) ToInt32() int32 {
	return int32(m)
}

func DevicePlatformTypeFromString(s string) DevicePlatformType {

	switch s {
	case DevicePlatformTypePC.ToString(), DevicePlatformTypeUnknown.ToString():
		return DevicePlatformTypePC
	case DevicePlatformTypeAndroid.ToString():
		return DevicePlatformTypeAndroid
	case DevicePlatformTypeIOS.ToString():
		return DevicePlatformTypeIOS
	case DevicePlatformTypeWeb.ToString():
		return DevicePlatformTypeWeb
	default:
		log.Errorf("DevicePlatformTypeFromString error:%s", s)
		return DevicePlatformTypeUnknown
	}
}

func (m DevicePlatformType) ToString() string {
	return string(m)
}

func MessageBoxesTypeFromInt32(v int32) MessageBoxesType {
	return MessageBoxesType(v)
}

func (m MessageBoxesType) ToInt32() int32 {
	return int32(m)
}

func UpdateTypeFromInt32(v int32) UpdateType {
	return UpdateType(v)
}

func (m UpdateType) ToInt32() int32 {
	return int32(m)
}

func MessageTypeFromInt32(v int32) MessageType {
	return MessageType(v)
}

func (m MessageType) ToInt32() int32 {
	return int32(m)
}

//
//func DialogTypeFromInt32(v int32) DialogType {
//	return DialogType(v)
//}
//
//func (m DialogType) ToInt32() int32 {
//	return int32(m)
//}

func AccessHashTypeFromString(s string) AccessHashType {
	switch s {
	//case AccessHashTypeContact.ToString():
	//return AccessHashTypeContact
	case AccessHashTypeInputUser.ToString():
		return AccessHashTypeInputUser
	case AccessHashTypeUser.ToString():
		return AccessHashTypeUser
	case AccessHashTypeInputChannel.ToString():
		return AccessHashTypeInputChannel
	case AccessHashTypeChat.ToString():
		return AccessHashTypeChat
	case AccessHashTypeSecretChat.ToString():
		return AccessHashTypeSecretChat
	case AccessHashTypePeer.ToString():
		return AccessHashTypePeer
	case AccessHashTypeWebLocation.ToString():
		return AccessHashTypeWebLocation
	case AccessHashTypeWebDocument.ToString():
		return AccessHashTypeWebLocation
	case AccessHashTypeGeoPoint.ToString():
		return AccessHashTypeGeoPoint
	case AccessHashTypeInputGame.ToString():
		return AccessHashTypeInputGame
	case AccessHashTypeGame.ToString():
		return AccessHashTypeGame
	case AccessHashTypeInputBotInlineMessageID.ToString():
		return AccessHashTypeInputBotInlineMessageID
	case AccessHashTypeInputPhoto.ToString():
		return AccessHashTypeInputPhoto
	case AccessHashTypePhoto.ToString():
		return AccessHashTypePhoto
	case AccessHashTypeInputWallPaper.ToString():
		return AccessHashTypeInputWallPaper
	case AccessHashTypeWallPaper.ToString():
		return AccessHashTypeWallPaper
	case AccessHashTypeGroupCall.ToString():
		return AccessHashTypeGroupCall
	case AccessHashTypeInputPhoneCall.ToString():
		return AccessHashTypeInputPhoneCall
	case AccessHashTypePhoneCall.ToString():
		return AccessHashTypePhoneCall
	case AccessHashTypeEncryptedFile.ToString():
		return AccessHashTypeEncryptedFile
	case AccessHashTypeInputStickerSet.ToString():
		return AccessHashTypeInputStickerSet
	case AccessHashTypeStickerSet.ToString():
		return AccessHashTypeStickerSet
	case AccessHashTypeFile.ToString():
		return AccessHashTypeFile
	case AccessHashTypeFileLocation.ToString():
		return AccessHashTypeFileLocation
	case AccessHashTypeDocument.ToString():
		return AccessHashTypeDocument
	case AccessHashTypeTheme.ToString():
		return AccessHashTypeTheme
	default:
		log.Errorf("AccessHashTypeFromString error:%s", s)
		return AccessHashTypeUnknown
	}
}

func (m AccessHashType) ToString() string {
	return string(m)
}

func InputUserTypeFromInt32(v int32) InputUserType {
	return InputUserType(v)
}

func (m InputUserType) ToInt32() int32 {
	return int32(m)
}

func UpdateGetDifferenceTypeFromInt32(v int32) UpdateGetDifferenceType {
	return UpdateGetDifferenceType(v)
}

func (m UpdateGetDifferenceType) ToInt32() int32 {
	return int32(m)
}

func InputMessagesFilterTypeFromInt32(v int32) InputMessagesFilterType {
	return InputMessagesFilterType(v)
}

func (m InputMessagesFilterType) ToInt32() int32 {
	return int32(m)
}

//if (load_type == 4) {
//req.add_offset = -count + 5;
//} else if (load_type == 3) {
//req.add_offset = -count / 2;
//} else if (load_type == 1) {
//req.add_offset = -count - 1;
//} else if (load_type == 2 && max_id != 0) {
//req.add_offset = -count + 10;
//} else {
//if (lower_part < 0 && max_id != 0) {
//TLRPC.Chat chat = getChat(-lower_part);
//if (ChatObject.isChannel(chat)) {
//req.add_offset = -1;
//req.limit += 1;
//}
//}
//}
func GenerateLoadHistoryType(bChannel bool, addOffset, limit, maxId int32) LoadHistoryType {
	if limit == 1 {
		return LoadHistoryTypeMax
	}

	if bChannel && addOffset == -1 && maxId != 0 {
		return LoadHistoryTypeBackward
	}

	if addOffset == -limit+5 {
		return LoadHistoryTypeFirstAroundDate // 4
	}

	if addOffset == -limit/2 {
		return LoadHistoryTypeFirstAroundMessage
	}

	if addOffset == -limit-1 {
		return LoadHistoryTypeForward
	}

	if addOffset == -limit+10 && maxId != 0 {
		return LoadHistoryTypeFirstUnRead
	}
	return LoadHistoryTypeBackward

	//if addOffset == 0 {
	//	return LoadHistoryTypeBackward
	//} else
	//if addOffset == -1 {
	//	return LoadHistoryTypeBackward
	//} else if addOffset == -limit+5 {
	//	return LoadHistoryTypeFirstAroundDate
	//} else if addOffset == -limit/2 {
	//	return LoadHistoryTypeFirstAroundMessage
	//} else if addOffset == -limit-1 {
	//	return LoadHistoryTypeForward
	//} else if addOffset == -limit+6 {
	//	if maxId != 0 {
	//		return LoadHistoryTypeFirstUnRead
	//	}
	//}
	//return LoadHistoryTypeBackward
}

func ToMessageType(message *mtproto.Message) MessageType {
	switch message.ClassName {
	case mtproto.ClassMessageEmpty:
		return MessageTypeEmpty
	case mtproto.ClassMessage:
		return MessageTypeMessage
	case mtproto.ClassMessageService:
		return MessageTypeMessageService
	default:
		panic(fmt.Sprintf("ToMessageType error:%s", message.ClassName))
	}
	return MessageTypeEmpty
}

func ChannelParticipantsFilterTypeFromInt32(m int32) ChannelParticipantsFilterType {
	return ChannelParticipantsFilterType(m)
}

func (m ChannelParticipantsFilterType) ToInt32() int32 {
	return int32(m)
}

func AccountStatusFromInt8(s int8) AccountStatus {
	return AccountStatus(s)
}
func (m AccountStatus) ToInt8() int8 {
	return int8(m)
}

func MessageActionFromInt32(s int32) MessageAction {
	return MessageAction(s)
}
func (m MessageAction) ToInt32() int32 {
	return int32(m)
}

func MessageMediaFromInt32(s int32) MessageMedia {
	return MessageMedia(s)
}
func (m MessageMedia) ToInt32() int32 {
	return int32(m)
}
