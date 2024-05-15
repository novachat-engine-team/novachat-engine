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

type Config struct {
	Server       config.ServerConfig
	Log          config.LogConfig        `yaml:"log"`
	Pprof        config.PprofConfig      `yaml:"pprof"`
	RpcServer    config.RpcServer        `yaml:"rpcServer"`
	RpcDiscovery config.EtcdServerConfig `yaml:"rpcDiscovery"`

	MessageProducer config.MQKafkaProducerConfig `yaml:"messageProducer"`
	ChatProducer    config.MQKafkaProducerConfig `yaml:"chatProducer"`
	MessageConsumer config.MQKafkaConsumerConfig `yaml:"messageConsumer"`
	ChatInConsumer  config.MQKafkaConsumerConfig `yaml:"chatInConsumer"`
	SyncRpcClient   config.EtcdClientConfig      `yaml:"syncRpcClient"`
	//ChatRpcClient   config.EtcdClientConfig      `yaml:"chatRpcClient"`

	Redis   config.RedisConfig   `yaml:"redis"`
	MongoDB config.MongodbConfig `yaml:"mongodb"`
}
