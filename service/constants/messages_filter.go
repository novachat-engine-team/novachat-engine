/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/7 14:23
 * @File : input_messages_filter.go
 */

package constants

import (
	"novachat_engine/mtproto"
)

type MessagesFilterType struct{}

func MakeMessagesFilterType(filter *mtproto.MessagesFilter) InputMessagesFilterType {
	switch filter.ClassName {
	case mtproto.ClassInputMessagesFilterEmpty:
		return InputMessagesFilterEmpty
	case mtproto.ClassInputMessagesFilterPhotos:
		return InputMessagesFilterPhotos
	case mtproto.ClassInputMessagesFilterVideo:
		return InputMessagesFilterVideo
	case mtproto.ClassInputMessagesFilterPhotoVideo:
		return InputMessagesFilterPhotoVideo
	case mtproto.ClassInputMessagesFilterDocument:
		return InputMessagesFilterDocument
	case mtproto.ClassInputMessagesFilterUrl:
		return InputMessagesFilterUrl
	case mtproto.ClassInputMessagesFilterGif:
		return InputMessagesFilterGif
	case mtproto.ClassInputMessagesFilterVoice:
		return InputMessagesFilterVoice
	case mtproto.ClassInputMessagesFilterMusic:
		return InputMessagesFilterMusic
	case mtproto.ClassInputMessagesFilterChatPhotos:
		return InputMessagesFilterChatPhotos
	case mtproto.ClassInputMessagesFilterPhoneCalls:
		return InputMessagesFilterPhoneCalls
	case mtproto.ClassInputMessagesFilterRoundVoice:
		return InputMessagesFilterRoundVoice
	case mtproto.ClassInputMessagesFilterRoundVideo:
		return InputMessagesFilterRoundVideo
	case mtproto.ClassInputMessagesFilterMyMentions:
		return InputMessagesFilterMyMentions
	case mtproto.ClassInputMessagesFilterGeo:
		return InputMessagesFilterGeo
	case mtproto.ClassInputMessagesFilterContacts:
		return InputMessagesFilterContacts
	case mtproto.ClassInputMessagesFilterPinned:
		return InputMessagesFilterPinned
	case mtproto.ClassInputMessagesFilterPhotoVideoDocuments:
		return InputMessagesFilterPhotoVideoDocuments
	default:
		return InputMessagesFilterUnknown
	}
}
