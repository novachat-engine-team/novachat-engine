/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/23 18:04
 * @File : config.go
 */

package conf

import "novachat_engine/pkg/config"

type Relay struct {
	PhoneCallWebRTC bool
	P2P             bool
	LibVersion      string
}

type Config struct {
	Server         config.ServerConfig
	RpcServer      config.RpcServer
	RpcDiscovery   config.EtcdServerConfig
	Log            config.LogConfig
	Pprof          config.PprofConfig
	MongoDB        config.MongodbConfig
	MySQL          config.MySQLConfig
	Redis          config.RedisConfig
	SfsClient      config.EtcdClientConfig
	MsgClient      config.EtcdClientConfig
	SyncClient     config.EtcdClientConfig
	AuthClient     config.EtcdClientConfig
	SessionClient  config.EtcdClientConfig
	ChatClient     config.EtcdClientConfig
	RelayConfig    Relay
	EnvMode        string
	ConfigPath     string
	CountryFile    string
	AppConfig      string
	InviteText     string
	QRLoginCodeKey string
}
