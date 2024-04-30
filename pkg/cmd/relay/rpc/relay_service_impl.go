/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/9 18:36
 * @File : relay_service_impl.go
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	relayService "novachat_engine/pkg/cmd/relay/rpc_client"
	"novachat_engine/pkg/cmd/relay/server"
	"novachat_engine/pkg/log"
)

type ServiceImpl struct {
	server *server.RelayServer
}

func NewServiceImpl(server *server.RelayServer) *ServiceImpl {
	impl := new(ServiceImpl)

	impl.server = server
	return impl
}

func (m *ServiceImpl) ReqCreatePhoneCall(ctx context.Context, req *relayService.CreatePhoneCall) (*relayService.CallConnections, error) {
	_ = ctx
	log.Debugf("CreatePhoneCall req:%v", req)
	callConnections, err := m.server.CreatePhoneCall(req.Id, req.AdminId, req.ParticipantId, req.Version)
	if err != nil {
		log.Errorf("CreatePhoneCall req:%v error:%s", req, err.Error())
		return nil, err
	}

	log.Infof("CreatePhoneCall req:%v reply ok !!!!!!! call:%+v connection:%v", req, callConnections, callConnections)
	return callConnections, err
}

func (m *ServiceImpl) ReqDiscardPhoneCall(ctx context.Context, req *relayService.DiscardPhoneCall) (*types.Any, error) {
	_ = ctx
	log.Debugf("DiscardPhoneCall req:%v", req)

	err := m.server.DiscardPhoneCall(req.Id)
	if err != nil {
		log.Warnf("DiscardPhoneCall req:%v error:%s", req, err.Error())
	}

	log.Infof("DiscardPhoneCall req:%v reply ok!!!!!!!!!", req)
	return types.MarshalAny(mtproto.ToMTBool(true))
}
