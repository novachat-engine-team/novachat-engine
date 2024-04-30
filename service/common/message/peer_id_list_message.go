/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/22 11:00
 * @File : peer_id_list_message.go
 */

package message

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
)

type PickPeerIdListType struct {
	PeerIdList    []int64
	ChatIdList    []int64
	ChannelIdList []int64
}

func wrapFindAndAppend(l []int64, id int64) []int64 {
	if id == 0 {
		return l
	}
	for _, v := range l {
		if id == v {
			return l
		}
	}
	l = append(l, id)
	return l
}

func pickPeerIdListTypeAppendPeerId(r *PickPeerIdListType, peer *mtproto.Peer) {
	switch peer.ClassName {
	case mtproto.ClassPeerUser:
		r.PeerIdList = wrapFindAndAppend(r.PeerIdList, constants.PeerTypeFromUserIDType32(peer.UserId).ToInt())
	case mtproto.ClassPeerChat:
		r.ChatIdList = wrapFindAndAppend(r.ChatIdList, constants.PeerTypeFromChannelIDType32(peer.ChatId).ToInt())
	case mtproto.ClassPeerChannel:
		r.ChannelIdList = wrapFindAndAppend(r.ChannelIdList, constants.PeerTypeFromChannelIDType32(peer.ChannelId).ToInt())
	}
}

func PickUserList(idList []int64, id int64) []int64 {
	if id == 0 {
		return idList
	}
	for _, v := range idList {
		if v == id {
			return idList
		}
	}
	idList = append(idList, id)
	return idList
}

func PickMessagePeerIdList(messageList []*mtproto.Message) PickPeerIdListType {
	ret := PickPeerIdListType{
		PeerIdList:    make([]int64, 0, 1),
		ChatIdList:    make([]int64, 0, 1),
		ChannelIdList: make([]int64, 0, 1),
	}

	for _, v := range messageList {
		if v.FromId90DDDC1171 != 0 {
			ret.PeerIdList = wrapFindAndAppend(ret.PeerIdList, constants.PeerTypeFromUserIDType32(v.FromId90DDDC1171).ToInt())
		}

		if v.FromId286FA604119 != nil {
			pickPeerIdListTypeAppendPeerId(&ret, v.FromId286FA604119)
		}

		if v.ToId != nil {
			pickPeerIdListTypeAppendPeerId(&ret, v.ToId)
		}

		if v.PeerId != nil {
			pickPeerIdListTypeAppendPeerId(&ret, v.PeerId)
		}

		if v.Action != nil {
			ret.PeerIdList = wrapFindAndAppend(ret.PeerIdList, constants.PeerTypeFromUserIDType32(v.Action.UserId).ToInt())
			ret.PeerIdList = wrapFindAndAppend(ret.PeerIdList, constants.PeerTypeFromUserIDType32(v.Action.InviterId).ToInt())

			for _, v1 := range v.Action.Users {
				ret.PeerIdList = wrapFindAndAppend(ret.PeerIdList, constants.PeerTypeFromUserIDType32(v1).ToInt())
			}

			ret.ChatIdList = wrapFindAndAppend(ret.ChatIdList, constants.PeerTypeFromChannelIDType32(v.Action.ChatId).ToInt())

			ret.ChannelIdList = wrapFindAndAppend(ret.ChannelIdList, constants.PeerTypeFromChannelIDType32(v.Action.ChannelId).ToInt())
		}
	}

	return ret
}
