/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package generator

import (
	"errors"
	"fmt"
	"github.com/pion/turn/v2"
	"math"
	"net"
	"strconv"
	"time"
)

var (
	errNilConn            = errors.New("turn: conn cannot not be nil")
	errMaxRetriesExceeded = errors.New("turn: max retries exceeded")
)

type AllocatePacketConn interface {
	Allocate(uint16)
	Free(uint16)
}

type relayPacketConn struct {
	port               uint16
	conn               net.PacketConn
	allocatePacketConn AllocatePacketConn
}

func newRelayPacketConn(port uint16, conn net.PacketConn, packetConn AllocatePacketConn) *relayPacketConn {
	if packetConn != nil {
		packetConn.Allocate(port)
	}
	return &relayPacketConn{
		port:               port,
		conn:               conn,
		allocatePacketConn: packetConn,
	}
}

func (r *relayPacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return r.conn.ReadFrom(p)
}

func (r *relayPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return r.conn.WriteTo(p, addr)
}

func (r *relayPacketConn) Close() error {
	if r.allocatePacketConn != nil {
		r.allocatePacketConn.Free(r.port)
	}
	return r.conn.Close()
}

func (r *relayPacketConn) LocalAddr() net.Addr {
	return r.conn.LocalAddr()
}

func (r *relayPacketConn) SetDeadline(t time.Time) error {
	return r.conn.SetDeadline(t)
}

func (r *relayPacketConn) SetReadDeadline(t time.Time) error {
	return r.conn.SetReadDeadline(t)
}

func (r *relayPacketConn) SetWriteDeadline(t time.Time) error {
	return r.conn.SetWriteDeadline(t)
}

type RelayAddressGenerator struct {
	generator          turn.RelayAddressGeneratorPortRange
	allocatePacketConn AllocatePacketConn
}

func NewRelayAddressGenerator(endPoint string, minPort uint16, maxPort uint16, allocatePacketConn AllocatePacketConn) (*RelayAddressGenerator, error) {
	if minPort <= 0 {
		return nil, fmt.Errorf("minPort invalid:%d", minPort)
	}

	if minPort > math.MaxUint16 {
		return nil, fmt.Errorf("min port:%d >= max %d", minPort, math.MaxUint16)
	}

	if maxPort > math.MaxUint16 {
		return nil, fmt.Errorf("maxPort port:%d >= max %d", maxPort, math.MaxUint16)
	}

	if maxPort <= minPort {
		maxPort = minPort
	}

	return &RelayAddressGenerator{
		generator: turn.RelayAddressGeneratorPortRange{
			RelayAddress: net.ParseIP(endPoint),
			Address:      "0.0.0.0",
			MinPort:      minPort,
			MaxPort:      maxPort,
		},
		allocatePacketConn: allocatePacketConn,
	}, nil
}

func (r *RelayAddressGenerator) Validate() error {
	return r.generator.Validate()
}

func (r *RelayAddressGenerator) AllocatePacketConn(network string, requestedPort int) (net.PacketConn, net.Addr, error) {
	if requestedPort != 0 {
		conn, err := r.generator.Net.ListenPacket(network, fmt.Sprintf("%s:%d", r.generator.Address, requestedPort))
		if err != nil {
			return nil, nil, err
		}
		relayAddr, ok := conn.LocalAddr().(*net.UDPAddr)
		if !ok {
			return nil, nil, errNilConn
		}

		relayAddr.IP = r.generator.RelayAddress
		return conn, relayAddr, nil
	}

	if r.generator.MaxPort <= r.generator.MinPort {
		conn, err := r.generator.Net.ListenPacket(network, r.generator.Address+":"+strconv.Itoa(int(r.generator.MinPort)))
		if err != nil {
			return nil, nil, err
		}

		// Replace actual listening IP with the user requested one of RelayAddressGeneratorStatic
		relayAddr, ok := conn.LocalAddr().(*net.UDPAddr)
		if !ok {
			return nil, nil, errNilConn
		}

		relayAddr.IP = r.generator.RelayAddress

		return newRelayPacketConn(r.generator.MinPort, conn, r.allocatePacketConn), relayAddr, nil
	}

	for try := 0; try < r.generator.MaxRetries; try++ {
		port := r.generator.MinPort + uint16(r.generator.Rand.Intn(int((r.generator.MaxPort+1)-r.generator.MinPort)))
		conn, err := r.generator.Net.ListenPacket(network, fmt.Sprintf("%s:%d", r.generator.Address, port))
		if err != nil {
			continue
		}

		relayAddr, ok := conn.LocalAddr().(*net.UDPAddr)
		if !ok {
			return nil, nil, errNilConn
		}

		relayAddr.IP = r.generator.RelayAddress
		return newRelayPacketConn(port, conn, r.allocatePacketConn), relayAddr, nil
	}

	return nil, nil, errMaxRetriesExceeded
}

func (r *RelayAddressGenerator) AllocateConn(network string, requestedPort int) (net.Conn, net.Addr, error) {
	return r.generator.AllocateConn(network, requestedPort)
}
