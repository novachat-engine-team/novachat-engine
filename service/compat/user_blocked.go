/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/10/26 14:58
 * @File : user_blocked.go
 */

package compat

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
)

const UserBlockLayer = int32(118)

func NewUserBlock(userId int64, blocked bool, layer int32) *mtproto.Update {

	if layer <= UserBlockLayer {
		return mtproto.NewTLUpdateUserBlocked(&mtproto.Update{
			UserId:  constants.PeerTypeFromUserIDType(userId).ToInt32(),
			Blocked: mtproto.ToMTBool(blocked),
		}).To_Update()
	} else {
		return mtproto.NewTLUpdatePeerBlocked(&mtproto.Update{
			PeerId:  mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer(),
			Blocked: mtproto.ToMTBool(blocked),
		}).To_Update()
	}
}
