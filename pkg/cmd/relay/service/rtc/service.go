/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rtc

import (
	"errors"
	"github.com/pion/turn/v2"
	"math/rand"
	"novachat_engine/pkg/cmd/relay/conf"
	relayService "novachat_engine/pkg/cmd/relay/rpc_client"
	"novachat_engine/pkg/cmd/relay/service/rtc/generator"
	stunService "novachat_engine/pkg/cmd/relay/service/rtc/stun"
	turnService "novachat_engine/pkg/cmd/relay/service/rtc/turn"
	"novachat_engine/pkg/log"
)

type Service struct {
	server    *turn.Server
	generator turn.RelayAddressGenerator
	conf      *conf.RTC
}

func NewTurnService(conf *conf.RTC) (*Service, error) {
	if len(conf.RelayConn.EndPoint) == 0 {
		conf.RelayConn.EndPoint = "0.0.0.0"
	}

	if len(conf.Realm) == 0 {
		return nil, errors.New("Realm is empty")
	}

	impl := &Service{}
	generator, err := generator.NewRelayAddressGenerator(conf.RelayConn.EndPoint, conf.RelayConn.RangePort.MinPort, conf.RelayConn.RangePort.MaxPort, impl)
	if err != nil {
		return nil, err
	}

	listener, err := stunService.NewListenerService(conf, generator)
	if err != nil {
		return nil, err
	}
	relayPacketConn, err := turnService.NewPacketConnService(conf, generator)
	if err != nil {
		return nil, err
	}

	if listener == nil && relayPacketConn == nil {
		return nil, nil
	}

	server, err := turn.NewServer(turn.ServerConfig{
		PacketConnConfigs: relayPacketConn,
		ListenerConfigs:   listener,
		AuthHandler:       impl.RegisterAuth(conf),
	})
	if err != nil {
		return nil, err
	}

	impl.server = server
	impl.generator = generator
	impl.conf = conf
	return impl, nil
}

func (s *Service) Check(callId int64) error {
	return nil
}

func (s *Service) DiscardPhoneCall(callId int64) error {
	log.Debugf("DiscardPhoneCall  callId:%d", callId)
	return nil
}

func (s *Service) CreatePhoneCall(callId int64, adminId int64, participantId int64, version string) (*relayService.CallConnections, error) {
	log.Debugf("CreatePhoneCall  callId:%d adminId:%d participantId:%d version:%s", callId, adminId, participantId, version)

	conn := &relayService.CallConnections{
		EndpointId:    callId,
		WebRTCConnect: make([]*relayService.CallConnections_WebRTCConnect, 0, 2),
	}

	if s.conf.StunPort >= 0 {
		conn.WebRTCConnect = append(conn.WebRTCConnect, &relayService.CallConnections_WebRTCConnect{
			Stun: true,
			Ip:   s.conf.RelayConn.EndPoint,
			Ipv6: "",
			Port: int32(s.conf.StunPort),
		})
	}

	var realm []string

	realm = s.conf.Realm[rand.Intn(len(s.conf.Realm))]

	conn.WebRTCConnect = append(conn.WebRTCConnect, &relayService.CallConnections_WebRTCConnect{
		Stun:     false,
		Ip:       s.conf.RelayConn.EndPoint,
		Ipv6:     "",
		Port:     int32(s.conf.StunPort),
		Username: realm[0],
		Password: realm[1],
	})

	return conn, nil
}

func (s *Service) Allocate(uint16) {
	//TODO:Coderxw
}

func (s *Service) Free(uint16) {
	//TODO:Coderxw
}
