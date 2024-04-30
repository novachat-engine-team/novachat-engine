/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 18:17
 * @File : config.go
 */

package conf

import "novachat_engine/pkg/config"

type BizMQ struct {
	Name          string
	KafkaProducer config.MQKafkaProducerConfig
}

type BizRpc struct {
	EnableMQ     bool
	BizRpcClient config.EtcdClientConfig
	BizMQList    []BizMQ
}

type Conf struct {
	Log           config.LogConfig
	Pprof         config.PprofConfig
	RpcDiscovery  config.EtcdServerConfig
	RpcServer     config.RpcServer
	GateRpcClient config.EtcdClientConfig
	AuthRpcClient config.EtcdClientConfig
	BizRpc        BizRpc
}
