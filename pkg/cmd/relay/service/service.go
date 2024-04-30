/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package service

import relayService "novachat_engine/pkg/cmd/relay/rpc_client"

type Service interface {
	CreatePhoneCall(callId int64, adminId int64, participantId int64, version string) (*relayService.CallConnections, error)
	DiscardPhoneCall(callId int64) error
	Check(callId int64) error
}
