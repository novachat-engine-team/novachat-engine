/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package turn

import (
	"fmt"
	"github.com/pion/turn/v2"
	"math"
	"net"
	"novachat_engine/pkg/cmd/relay/conf"
	"novachat_engine/pkg/log"
)

func NewPacketConnService(conf *conf.RTC, generator turn.RelayAddressGenerator) ([]turn.PacketConnConfig, error) {
	minPort := conf.TurnPort
	maxPort := conf.TurnPort
	if minPort <= 0 {
		return nil, nil
	}

	if minPort > math.MaxUint16 {
		return nil, fmt.Errorf("TurnPort minPort:%d >= max %d", minPort, math.MaxUint16)
	}

	if maxPort > math.MaxUint16 {
		return nil, fmt.Errorf("TurnPort maxPort port:%d >= max %d", maxPort, math.MaxUint16)
	}

	endpoint := conf.RelayConn.EndPoint
	if maxPort <= minPort {
		maxPort = minPort + 1
	}

	networkType := "udp"
	c := make([]turn.PacketConnConfig, 0, maxPort-minPort+1)
	for v := minPort; v < maxPort; v++ {
		//addr, err := net.ResolveIPAddr(networkType, fmt.Sprintf("%s:%d", endpoint, v))
		//if err != nil {
		//	log.Fatalf("ResolveIPAddr error: %v", err)
		//	return nil, err
		//}

		lsn, err := net.ListenPacket(networkType, fmt.Sprintf("%s:%d", endpoint, v))
		if err != nil {
			log.Fatalf("ListenPacket error: %v", err)
			return nil, err
		}

		c = append(c, turn.PacketConnConfig{
			PacketConn:            lsn,
			RelayAddressGenerator: generator,
			PermissionHandler:     nil,
		})
	}
	return c, nil
}
