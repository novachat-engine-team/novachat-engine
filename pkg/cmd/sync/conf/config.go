/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 18:17
 * @File : config.go
 */

package conf

import (
	"novachat_engine/pkg/config"
)

type ProducerMQ struct {
	Name          string
	KafkaProducer config.MQKafkaProducerConfig
}

type ConsumerMQ struct {
	Name            string
	MessageConsumer config.MQKafkaConsumerConfig
}

type Conf struct {
	Server           config.ServerConfig
	Log              config.LogConfig
	Pprof            config.PprofConfig
	RpcDiscovery     config.EtcdServerConfig
	RpcServer        config.RpcServer
	SessionRpcClient config.EtcdClientConfig
	AuthRpcClient    config.EtcdClientConfig
	Producer         []ProducerMQ
	Consumer         []ConsumerMQ
	NpnsProducer     *config.MQKafkaProducerConfig
	NpnsConsumer     *config.MQKafkaConsumerConfig
	MongoDB          *config.MongodbConfig
}
