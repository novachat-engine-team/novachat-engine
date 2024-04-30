/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *SessionImpl) PushUpdates(ctx context.Context, pushUpdatesEvent *sessionService.PushUpdatesEvent) (*types.Any, error) {

	log.Debugf("PushUpdates authKeyId:%d sessionId:%v", pushUpdatesEvent.AuthKeyId, pushUpdatesEvent.SessionId)

	var updates mtproto.Updates
	err := types.UnmarshalAny(pushUpdatesEvent.Updates, &updates)
	if err != nil {
		log.Infof("SessionPushUpdates authKeyId:%d Updates:%v ok == false", pushUpdatesEvent.AuthKeyId, pushUpdatesEvent.Updates)
		return nil, err
	}

	session := m.messageChannel.GetChannel(pushUpdatesEvent.AuthKeyId)
	ok, err := session.PushUpdates(pushUpdatesEvent.AuthKeyId, pushUpdatesEvent.SessionId, &updates)
	if err != nil {
		log.Infof("SessionPushUpdates authKeyId:%d PushUpdates Updates:%v ok == false", pushUpdatesEvent.AuthKeyId, pushUpdatesEvent.Updates)
		return nil, err
	}

	log.Debugf("PushUpdates authKeyId:%d sessionId:%d ok:%v", pushUpdatesEvent.AuthKeyId, pushUpdatesEvent.SessionId, ok)
	return types.MarshalAny(mtproto.ToMTBool(ok))
}
