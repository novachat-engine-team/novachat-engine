/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package server

import (
	"fmt"
	"novachat_engine/pkg/cmd/relay/conf"
	relayService "novachat_engine/pkg/cmd/relay/rpc_client"
	"novachat_engine/pkg/cmd/relay/service"
	"novachat_engine/pkg/cmd/relay/service/relay_udp"
	turnService "novachat_engine/pkg/cmd/relay/service/rtc"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"sync"
	"time"
)

type RelayServer struct {
	rtc           *turnService.Service
	udpRelay      *relay_udp.Service
	callIdService sync.Map
}

func NewServer(conf *conf.Config) (*RelayServer, error) {
	rtcServer, err := turnService.NewTurnService(&conf.RTC)
	if err != nil {
		log.Errorf("TurnService error:%s", err.Error())
		return nil, err
	}

	udpRelay, err := relay_udp.NewRelayUdpService(conf.UdpRelayOnly)
	if err != nil {
		log.Errorf("RelayUdpService error:%s", err.Error())
		return nil, err
	}

	if rtcServer == nil && udpRelay == nil {
		return nil, fmt.Errorf("conf not found UdpRelayOnly or RTC conf:%v", conf)
	}

	rs := &RelayServer{
		rtc:      rtcServer,
		udpRelay: udpRelay,
	}
	go func() {
		ticker := time.NewTicker(time.Second * 10)
		var err1 error
		for {
			select {
			case <-ticker.C:
				rs.callIdService.Range(func(key, value interface{}) bool {
					serviceIns := value.(service.Service)
					err1 = serviceIns.Check(key.(int64))
					if err1 != nil {
						rs.callIdService.Delete(key)
					}
					return false
				})
			}
		}
	}()
	return rs, nil
}

func (r *RelayServer) UpdateAuth(conf *conf.Config) {
	r.rtc.RegisterAuth(&conf.RTC)
}

func (r *RelayServer) DiscardPhoneCall(callId int64) error {
	log.Debugf("DiscardPhoneCall  callId:%d", callId)
	v, ok := r.callIdService.Load(callId)
	if !ok {
		log.Warnf("DiscardPhoneCall not found callId:%d", callId)
		return nil
	}

	vv, ok := v.(service.Service)
	if !ok {
		log.Errorf("DiscardPhoneCall callId:%d is nil", callId)
		return nil
	}

	vv.DiscardPhoneCall(callId)
	r.callIdService.Delete(callId)
	return nil
}

func (r *RelayServer) CreatePhoneCall(callId int64, adminId int64, participantId int64, version string) (*relayService.CallConnections, error) {
	log.Debugf("CreatePhoneCall  callId:%d adminId:%d participantId:%d version:%s", callId, adminId, participantId, version)
	switch version {
	case constants.PhoneVersion2_4_4:
		conn, err := r.udpRelay.CreatePhoneCall(callId, adminId, participantId, version)
		if err != nil {
			log.Errorf("CreatePhoneCall udpRelay callId:%d adminId:%d participantId:%d version:%s error:%s", callId, adminId, participantId, version, err.Error())
			return nil, err
		}
		r.callIdService.Store(callId, r.udpRelay)
		return conn, err

	case constants.PhoneVersion2_5_0:
		fallthrough
	case constants.PhoneVersion2_8_8:
		fallthrough
	default:
		err := fmt.Errorf("unsupport version:%d", version)
		log.Errorf("CreatePhoneCall callId:%d adminId:%d participantId:%d error:%s", callId, adminId, participantId, err.Error())
		return nil, err

	case constants.PhoneVersion2_7_7:
		fallthrough
	case constants.PhoneVersion3_0_0:
		conn, err := r.rtc.CreatePhoneCall(callId, adminId, participantId, version)
		if err != nil {
			log.Errorf("CreatePhoneCall rtc callId:%d adminId:%d participantId:%d version:%s error:%s", callId, adminId, participantId, version, err.Error())
			return nil, err
		}
		r.callIdService.Store(callId, r.rtc)
		return conn, err
	}
}
