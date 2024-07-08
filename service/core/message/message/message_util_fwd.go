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
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	"novachat_engine/service/data/messages/message"
)

func messageFwdUtil(fwd *mtproto.MessageFwdHeader) *data_message.Fwd {
	if fwd == nil {
		return nil
	}

	peerType := constants.PeerTypeUser.ToInt32()
	fwdFromUserId := fwd.FromIdFADFF4AC71
	savePeerType := constants.PeerTypeEmpty.ToInt32()
	savePeerId := int32(0)
	if fwd.FromIdFADFF4AC71 != 0 {
		fwdFromUserId = fwd.FromIdFADFF4AC71
	} else if fwd.ChannelId != 0 {
		fwdFromUserId = fwd.ChannelId
		peerType = constants.PeerTypeChannel.ToInt32()
	} else {
		if fwd.FromId5F777DCE119 != nil {
			switch fwd.FromId5F777DCE119.ClassName {
			case mtproto.ClassPeerUser:
				fwdFromUserId = fwd.FromId5F777DCE119.UserId
				peerType = constants.PeerTypeUser.ToInt32()
			case mtproto.ClassPeerChat:
				fwdFromUserId = fwd.FromId5F777DCE119.ChatId
				peerType = constants.PeerTypeChat.ToInt32()
			case mtproto.ClassPeerChannel:
				fwdFromUserId = fwd.FromId5F777DCE119.ChannelId
				peerType = constants.PeerTypeChannel.ToInt32()
			}
		}
	}

	if fwd.SavedFromPeer != nil {
		switch fwd.SavedFromPeer.ClassName {
		case mtproto.ClassPeerUser:
			savePeerId = fwd.SavedFromPeer.UserId
			savePeerType = constants.PeerTypeUser.ToInt32()
		case mtproto.ClassPeerChat:
			savePeerId = fwd.SavedFromPeer.ChatId
			savePeerType = constants.PeerTypeChat.ToInt32()
		case mtproto.ClassPeerChannel:
			savePeerId = fwd.SavedFromPeer.ChannelId
			savePeerType = constants.PeerTypeChannel.ToInt32()
		}
	}

	return &data_message.Fwd{
		PeerId:       constants.PeerTypeFromUserIDType32(fwdFromUserId).ToInt(),
		PeerType:     peerType,
		MsgId:        fwd.SavedFromMsgId,
		Date:         fwd.Date,
		FromName:     fwd.FromName,
		SavePeerId:   constants.PeerTypeFromUserIDType32(savePeerId).ToInt(),
		SavePeerType: savePeerType,
		ChannelPost:  fwd.ChannelPost,
		PostAuthor:   fwd.PostAuthor,
	}
}

func toMessageFwd(fwd *data_message.Fwd) *mtproto.MessageFwdHeader {
	if fwd == nil {
		return nil
	}

	f := &mtproto.MessageFwdHeader{}

	f.SavedFromMsgId = fwd.MsgId
	f.FromName = fwd.FromName
	f.PsaType = fwd.PsaType
	f.Date = fwd.Date
	f.ChannelPost = fwd.ChannelPost
	f.PostAuthor = fwd.PostAuthor

	if fwd.SavePeerType != 0 && fwd.SavePeerId != 0 {
		switch fwd.SavePeerType {
		case constants.PeerTypeUser.ToInt32():
			f.SavedFromPeer = mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(fwd.SavePeerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChat.ToInt32():
			f.SavedFromPeer = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromUserIDType(fwd.SavePeerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChannel.ToInt32():
			f.SavedFromPeer = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromUserIDType(fwd.SavePeerId).ToInt32()}).To_Peer()
		}
	}

	if fwd.PeerType != 0 {
		switch fwd.PeerType {
		case constants.PeerTypeUser.ToInt32():
			f.FromId5F777DCE119 = mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(fwd.PeerId).ToInt32()}).To_Peer()
			f.FromIdFADFF4AC71 = constants.PeerTypeFromUserIDType(fwd.PeerId).ToInt32()
		case constants.PeerTypeChat.ToInt32():
			f.FromId5F777DCE119 = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromUserIDType(fwd.PeerId).ToInt32()}).To_Peer()
		case constants.PeerTypeChannel.ToInt32():
			f.FromId5F777DCE119 = mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromUserIDType(fwd.PeerId).ToInt32()}).To_Peer()
			f.ChannelId = constants.PeerTypeFromUserIDType(fwd.PeerId).ToInt32()
		}
	}

	return mtproto.NewTLMessageFwdHeader(f).To_MessageFwdHeader()
}
