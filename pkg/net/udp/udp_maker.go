/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/16 11:50
 * @File : udp_maker.go
 */

package udp

import (
	"net"
	"novachat_engine/pkg/log"
)

func ListenUdp(addr2 string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", addr2)
	if err != nil {
		log.Errorf("net.ResolveUDPAddr fail - %v", err)
		return nil, err
	}
	log.Info(addr.String())
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Errorf("net.ListenUDP fail - %v", err)
		return nil, err
	}
	return conn, nil
}

func ListenPacket(addr1 string) (net.PacketConn, error) {
	netType := "udp"
	addr, err := net.ResolveUDPAddr(netType, addr1)
	if err != nil {
		log.Errorf("net.ResolveUDPAddr fail - %v", err)
		return nil, err
	}
	log.Info(addr.String())

	conn, err := net.ListenPacket(addr.Network(), addr.String())
	if err != nil {
		log.Errorf("net.ListenPacket fail - %v", err)
		return nil, err
	}
	return conn, err
}
