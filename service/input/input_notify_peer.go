/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/12 15:40
 * @File : input_notify_peer.go
 */

package input

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

type PeerNotify struct {
	peerId         int64
	peerType       constants.PeerType
	peerNotifyType constants.PeerNotifyType
}

func (m *PeerNotify) GetPeerId() int64 {
	return m.peerId
}

func (m *PeerNotify) GetPeerNotifyType() constants.PeerNotifyType {
	return m.peerNotifyType
}

func (m *PeerNotify) GetPeerType() constants.PeerType {
	return m.peerType
}

func (m *PeerNotify) ToInputNotifyPeer() *mtproto.InputNotifyPeer {
	switch m.peerNotifyType {
	case constants.PeerNotifyInputNotifyAll:
		return mtproto.NewTLInputNotifyAll(nil).To_InputNotifyPeer()
	case constants.PeerNotifyInputNotifyBroadcasts:
		return mtproto.NewTLInputNotifyBroadcasts(nil).To_InputNotifyPeer()
	case constants.PeerNotifyInputNotifyChats:
		return mtproto.NewTLInputNotifyChats(nil).To_InputNotifyPeer()
	case constants.PeerNotifyInputNotifyUsers:
		return mtproto.NewTLInputNotifyUsers(nil).To_InputNotifyPeer()
	case constants.PeerNotifyInputNotifyPeer:
		var peer *mtproto.InputPeer
		switch m.peerType {
		case constants.PeerTypeUser, constants.PeerTypeSelf:
			peer = mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32()}).To_InputPeer()
		case constants.PeerTypeChat:
			peer = mtproto.NewTLInputPeerChat(&mtproto.InputPeer{ChatId: constants.PeerTypeFromChatIDType(m.peerId).ToInt32()}).To_InputPeer()
		case constants.PeerTypeChannel:
			peer = mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{ChannelId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32()}).To_InputPeer()
		default:
			panic(m.peerType)
		}
		return mtproto.NewTLInputNotifyPeer(&mtproto.InputNotifyPeer{
			Peer: peer,
		}).To_InputNotifyPeer()
	default:
		log.Errorf("PeerNotify ToInputNotifyPeer error:%v", m.peerNotifyType)
		return nil
	}
}

func (m *PeerNotify) ToNotifyPeer() *mtproto.NotifyPeer {
	switch m.peerNotifyType {
	case constants.PeerNotifyInputNotifyAll:
		return mtproto.NewTLNotifyAll(nil).To_NotifyPeer()
	case constants.PeerNotifyInputNotifyBroadcasts:
		return mtproto.NewTLNotifyBroadcasts(nil).To_NotifyPeer()
	case constants.PeerNotifyInputNotifyChats:
		return mtproto.NewTLNotifyChats(nil).To_NotifyPeer()
	case constants.PeerNotifyInputNotifyUsers:
		return mtproto.NewTLNotifyUsers(nil).To_NotifyPeer()
	case constants.PeerNotifyInputNotifyPeer:
		var peer *mtproto.Peer
		switch m.peerType {
		case constants.PeerTypeUser, constants.PeerTypeSelf:
			peer = mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChat:
			peer = mtproto.NewTLPeerChat(&mtproto.Peer{ChatId: constants.PeerTypeFromChatIDType(m.peerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChannel:
			peer = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(m.peerId).ToInt32()}).To_Peer()
		default:
			panic(m.peerType)
		}
		return mtproto.NewTLNotifyPeer(&mtproto.NotifyPeer{Peer: peer}).To_NotifyPeer()
	default:
		log.Errorf("PeerNotify ToNotifyPeer error:%v", m.peerNotifyType)
		return nil
	}
}

func MakePeerNotify(peer *mtproto.InputNotifyPeer, selfId constants.UserIDType) *PeerNotify {
	p := &PeerNotify{}
	switch peer.ClassName {
	case mtproto.ClassInputNotifyAll:
		p.peerNotifyType = constants.PeerNotifyInputNotifyAll
	case mtproto.ClassInputNotifyBroadcasts:
		p.peerNotifyType = constants.PeerNotifyInputNotifyBroadcasts
	case mtproto.ClassInputNotifyChats:
		p.peerNotifyType = constants.PeerNotifyInputNotifyChats
	case mtproto.ClassInputNotifyUsers:
		p.peerNotifyType = constants.PeerNotifyInputNotifyUsers
	case mtproto.ClassInputNotifyPeer:
		p.peerNotifyType = constants.PeerNotifyInputNotifyPeer
		if peer.Peer != nil {
			switch peer.Peer.ClassName {
			case mtproto.ClassInputPeerSelf:
				p.peerType = constants.PeerTypeSelf
				p.peerId = selfId.ToInt()
				break
			case mtproto.ClassInputPeerUser:
				p.peerType = constants.PeerTypeUser
				p.peerId = constants.PeerTypeFromUserIDType32(peer.Peer.UserId).ToInt()
				break
			case mtproto.ClassInputPeerChat:
				p.peerType = constants.PeerTypeChat
				p.peerId = constants.PeerTypeFromChannelIDType32(peer.Peer.ChatId).ToInt()
				break
			case mtproto.ClassInputPeerChannel:
				p.peerType = constants.PeerTypeChannel
				p.peerId = constants.PeerTypeFromChannelIDType32(peer.Peer.ChannelId).ToInt()
				break
			}
		}
	default:
		log.Errorf("MakePeerNotify error:%v peer:%v", peer.ClassName, p)
		return nil
	}
	return p
}
