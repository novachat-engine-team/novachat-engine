/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/9 17:52
 * @File : config.go
 */

package conf

import (
	"novachat_engine/pkg/config"
)

type RangePort struct {
	MinPort uint16
	MaxPort uint16
}

type Relay struct {
	EndPoint  string
	RangePort RangePort
}

type RTC struct {
	RelayConn Relay
	TurnPort  uint16
	StunPort  uint16
	Realm     [][]string
	RealmTag  string
}

type UdpRelay struct {
	EndPoint string
	Port     int32
}

type Config struct {
	Log          config.LogConfig
	Pprof        config.PprofConfig
	RpcServer    config.RpcServer
	RpcDiscovery config.EtcdServerConfig
	RTC          RTC
	UdpRelayOnly UdpRelay
}
