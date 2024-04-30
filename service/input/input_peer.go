/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/10 15:47
 * @File : peer.go
 */

package input

import (
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

type InputPeer struct {
	peerId   int64
	peerType constants.PeerType
}

func MakeInputPeerIntValue(peerId int64, pt int32) *InputPeer {
	return &InputPeer{
		peerId:   peerId,
		peerType: constants.PeerType(pt),
	}
}

func MakeInputPeerValue(peerId int64, pt constants.PeerType) *InputPeer {
	return &InputPeer{
		peerId:   peerId,
		peerType: pt,
	}
}

func (m *InputPeer) UpdatePeerId(userId int64) {
	m.peerId = userId
}

func (m *InputPeer) GetPeerId() int64 {
	return m.peerId
}

func (m *InputPeer) GetPeerType() constants.PeerType {
	return m.peerType
}

func (m *InputPeer) ToPeer() *mtproto.Peer {
	switch m.peerType {
	case constants.PeerTypeChat:
		return mtproto.NewTLPeerChat(&mtproto.Peer{
			ChatId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32(),
		}).To_Peer()
	case constants.PeerTypeChannel:
		return mtproto.NewTLPeerChannel(&mtproto.Peer{
			ChannelId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32(),
		}).To_Peer()
	case constants.PeerTypeUser:
		return mtproto.NewTLPeerUser(&mtproto.Peer{
			UserId: constants.PeerTypeFromUserIDType(m.peerId).ToInt32(),
		}).To_Peer()
	case constants.PeerTypeSelf:
		return mtproto.NewTLPeerUser(&mtproto.Peer{
			UserId: constants.PeerTypeFromUserIDType(m.peerId).ToInt32(),
		}).To_Peer()
	default:
		panic(fmt.Sprintf("ToPeer %v", m.peerType))
	}
	return nil
}

func (m *InputPeer) ToInputPeer() *mtproto.InputPeer {
	switch m.peerType {
	case constants.PeerTypeEmpty:
		return mtproto.NewTLInputPeerEmpty(nil).To_InputPeer()
	case constants.PeerTypeChat:
		return mtproto.NewTLInputPeerChat(&mtproto.InputPeer{
			ChatId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32(),
		}).To_InputPeer()
	case constants.PeerTypeChannel:
		return mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
			ChannelId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32(),
		}).To_InputPeer()
	case constants.PeerTypeUser:
		return mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId: constants.PeerTypeFromUserIDType(m.peerId).ToInt32(),
		}).To_InputPeer()
	case constants.PeerTypeSelf:
		return mtproto.NewTLInputPeerSelf(&mtproto.InputPeer{
			UserId: constants.PeerTypeFromUserIDType(m.peerId).ToInt32(),
		}).To_InputPeer()
	case constants.PeerTypeUserFromMessage: // unused
		return mtproto.NewTLInputPeerUserFromMessage(&mtproto.InputPeer{
			MsgId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32(),
		}).To_InputPeer()
	case constants.PeerTypeChannelFromMessage: // unused
		return mtproto.NewTLInputPeerChannelFromMessage(&mtproto.InputPeer{
			MsgId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32(),
		}).To_InputPeer()
	default:
		log.Errorf("ToInputPeer error:%s value:%v", m.peerType, m)
		return nil
	}
}

func MakeInputPeer(inputPeer *mtproto.InputPeer) *InputPeer {
	p := &InputPeer{}
	switch inputPeer.ClassName {
	case mtproto.ClassInputPeerEmpty:
		p.peerType = constants.PeerTypeEmpty
	case mtproto.ClassInputPeerChat:
		p.peerType = constants.PeerTypeChat
		p.peerId = constants.PeerTypeFromChannelIDType32(inputPeer.ChatId).ToInt()
	case mtproto.ClassInputPeerChannel:
		p.peerType = constants.PeerTypeChannel
		p.peerId = constants.PeerTypeFromChannelIDType32(inputPeer.ChannelId).ToInt()
	case mtproto.ClassInputPeerUser:
		p.peerType = constants.PeerTypeUser
		p.peerId = constants.PeerTypeFromUserIDType32(inputPeer.UserId).ToInt()
	case mtproto.ClassInputPeerSelf:
		p.peerType = constants.PeerTypeSelf
		p.peerId = constants.PeerTypeFromUserIDType32(inputPeer.UserId).ToInt()
	case mtproto.ClassInputPeerUserFromMessage: // unused
		p.peerType = constants.PeerTypeUserFromMessage
		p.peerId = constants.PeerTypeFromUserIDType32(inputPeer.MsgId).ToInt()
	case mtproto.ClassInputPeerChannelFromMessage: // unused
		p.peerType = constants.PeerTypeChannelFromMessage
		p.peerId = constants.PeerTypeFromUserIDType32(inputPeer.MsgId).ToInt()
	default:
		log.Errorf("MakeInputPeer error:%s value:%v", inputPeer.ClassName, inputPeer)
		return nil
	}

	return p
}

func MakePeer(inputPeer *mtproto.Peer) *InputPeer {
	switch inputPeer.ClassName {
	case mtproto.ClassPeerChat:
		return &InputPeer{
			peerId:   constants.PeerTypeFromChannelIDType32(inputPeer.ChatId).ToInt(),
			peerType: constants.PeerTypeChat,
		}
	case mtproto.ClassPeerChannel:
		return &InputPeer{
			peerId:   constants.PeerTypeFromChannelIDType32(inputPeer.ChannelId).ToInt(),
			peerType: constants.PeerTypeChannel,
		}
	case mtproto.ClassPeerUser:
		return &InputPeer{
			peerId:   constants.PeerTypeFromUserIDType32(inputPeer.UserId).ToInt(),
			peerType: constants.PeerTypeUser,
		}
	default:
		log.Errorf("MakePeer error:%s value:%v", inputPeer.ClassName, inputPeer)
		return nil
	}
}
