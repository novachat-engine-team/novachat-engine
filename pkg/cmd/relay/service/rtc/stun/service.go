/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package stun

import (
	"fmt"
	"github.com/pion/turn/v2"
	"math"
	"net"
	"novachat_engine/pkg/cmd/relay/conf"
	"novachat_engine/pkg/log"
)

func NewListenerService(conf *conf.RTC, generator turn.RelayAddressGenerator) ([]turn.ListenerConfig, error) {
	minPort := conf.StunPort
	maxPort := conf.StunPort
	if minPort <= 0 {
		return nil, nil
	}

	if minPort > math.MaxUint16 {
		return nil, fmt.Errorf("StunPort minPort:%d >= max %d", minPort, math.MaxUint16)
	}

	if maxPort > math.MaxUint16 {
		return nil, fmt.Errorf("StunPort maxPort port:%d >= max %d", maxPort, math.MaxUint16)
	}

	endpoint := conf.RelayConn.EndPoint
	if maxPort <= minPort {
		maxPort = minPort + 1
	}

	networkType := "tcp"
	c := make([]turn.ListenerConfig, 0, maxPort-minPort+1)
	for v := minPort; v < maxPort; v++ {
		//addr, err := net.ResolveIPAddr(networkType, fmt.Sprintf("%s:%d", endpoint, v))
		//if err != nil {
		//	log.Fatalf("ResolveIPAddr error: %v", err)
		//	return nil, err
		//}

		lsn, err := net.Listen(networkType, fmt.Sprintf("%s:%d", endpoint, v))
		if err != nil {
			log.Fatalf("Listen error: %v", err)
			return nil, err
		}

		c = append(c, turn.ListenerConfig{
			Listener:              lsn,
			RelayAddressGenerator: generator,
			PermissionHandler:     nil,
		})
	}

	return c, nil
}
