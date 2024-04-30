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

type Config struct {
	RpcServer    config.RpcServer
	RpcDiscovery config.EtcdServerConfig
	Log          config.LogConfig
	Pprof        config.PprofConfig
	Mongodb      config.MongodbConfig
	AuthClient   config.EtcdClientConfig
}
