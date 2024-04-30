/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:13
 * @File : config.go
 */

package conf

import (
	"fmt"
	"novachat_engine/pkg/config"
)

type Config struct {
	Log              config.LogConfig        `toml:"log"`
	Pprof            config.PprofConfig      `toml:"pprof"`
	Key              config.KeyConfig        `toml:"key"`
	Server           config.ServerConfig     `toml:"server"`
	RpcServer        config.RpcServer        `toml:"rpcServer"`
	RpcDiscovery     config.EtcdServerConfig `toml:"rpcDiscovery"`
	SessionRpcClient config.EtcdClientConfig `toml:"sessionRpcClient"`
	AuthRpcClient    config.EtcdClientConfig `toml:"authRpcClient"`
}

func (c *Config) String() string {
	return fmt.Sprintf("{"+
		"log=%+v, "+
		"PProf=%+v,"+
		"Server=%+v,"+
		"RpcServer=%+v,"+
		"RpcDiscovery=%+v,"+
		"RpcClient=%+v}",
		c.Log, c.Pprof,
		c.Server,
		c.RpcServer,
		c.RpcDiscovery,
		c.SessionRpcClient)
}

func (c *Config) GoString() string {
	return fmt.Sprintf("{"+
		"log=%+v, "+
		"PProf=%+v,"+
		"Server=%+v,"+
		"RpcServer=%+v,"+
		"RpcDiscovery=%+v,"+
		"RpcClient=%+v}",
		c.Log, c.Pprof,
		c.Server,
		c.RpcServer,
		c.RpcDiscovery,
		c.SessionRpcClient)
}

const Filename = "gate.toml"
