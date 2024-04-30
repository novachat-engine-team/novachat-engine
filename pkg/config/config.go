/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:15
 * @File : config.go
 */

package config

type ServerConfig struct {
	Ver       string `toml:"ver"`
	ServerId  int32  `toml:"serverId"`
	ProtoName string `toml:"protoName"`
	Addr      string `toml:"addr"`
}

type KeyConfig struct {
	KeyFile        string `toml:"keyFile"`
	KeyFingerPrint string `toml:"keyFingerPrint"`
}

type LogConfig struct {
	LogPath          string `toml:"logPath"`
	LogLevel         string `toml:"logLevel"`
	RotateMaxSize    int    `toml:"rotateMaxSize"`
	RotateMaxAge     int    `toml:"rotateMaxAge"`
	RotateMaxBackups int    `toml:"rotateMaxBackups"`
}

type PprofConfig struct {
	Port   int32 `toml:"port"`
	IsOpen bool  `toml:"isOpen"`
}

type RpcServer struct {
	Addr string `toml:"addr"`
}

type EtcdServerConfig struct {
	Schema               string
	ServerName           string
	EtcdAddr             string
	TTL                  int64
	DialTimeout          int32
	DialKeepAliveTime    int32
	DialKeepAliveTimeout int32
	MaxCallSendMsgSize   int32
	MaxCallRecvMsgSize   int32
	UserName             string
	Password             string
}

type EtcdClientConfig struct {
	Schema               string
	ServerName           string
	DialTimeout          int32
	EtcdAddr             string
	DialKeepAliveTime    int32
	DialKeepAliveTimeout int32
	MaxCallSendMsgSize   int32
	MaxCallRecvMsgSize   int32
	UserName             string
	Password             string
	Balancer             string
}

type RedisMode string

const (
	RedisModeAlone    RedisMode = "alone"
	RedisModeSentinel RedisMode = "sentinel"
	RedisModeCluster  RedisMode = "cluster"
)

type RedisConfig struct {
	Addrs  []string
	Mode   RedisMode
	Test   bool
	Idle   int32
	Active int32
	Passwd string
	DBNum  int32
}

type MySQLConfig struct {
	Name    string `toml:"name"`   // tag name
	DSN     string `toml:"dsn"`    // data source name
	Active  int32  `toml:"active"` // pool
	Idle    int32  `toml:"idle"`   // pool
	Migrate bool   `toml:"migrate"`
	Debug   bool   `toml:"debug"`
}

//
//// 短信模板设置
//type SmsTemplate struct {
//	Title string // 标题
//	Code  string // 模板
//}
//
//// 短信签名配置
//type SubmailSignConfig struct {
//	SubmailAppID             string //国内短信应用ID
//	SubmailAppKey            string //国内短信应用签名
//	SubmailHost              string //国内短信发送接口
//	InternationalAppID       string //国际短信应用ID
//	InternationalAppKey      string //国际短信应用签名
//	InternationalSubmailHost string //国际短信发送接口
//}

type SmsConfig struct {
	Title          string
	SmsPlatform    string
	ConfigFileName string
}

type MQKafkaProducerConfig struct {
	Topic   string
	Brokers []string
}

type MQKafkaConsumerConfig struct {
	Topic   []string
	Group   string
	Brokers []string
}

type MongodbConfig struct {
	Direct      bool
	Addr        []string
	ReplicaSet  string
	Database    string
	Username    string
	Password    string
	MaxPoolSize int32
	MinPoolSize int32
}

type MQRabbitConfig struct {
	AMQP         string //amqp/amqps
	Addr         string //127.0.0.1:5672
	UserName     string
	Password     string
	Kind         string // fanout/topic/direct
	URI          string
	ExchangeName string
	QueueName    string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Exclusive    bool
	Delay        bool
}

type ExternalProducerConfig struct {
	Enable   bool
	Producer MQKafkaProducerConfig
}

type Cache struct {
	Shards             int
	LifeWindow         int
	CleanWindow        int
	MaxEntriesInWindow int
	MaxEntrySize       int
	HardMaxCacheSize   int
}
