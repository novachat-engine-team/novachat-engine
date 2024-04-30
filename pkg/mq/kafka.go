/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/16 13:21
 * @File : kafka.go
 */

package mq

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	time "time"
)

type KafkaProducer struct {
	conf *config.MQKafkaProducerConfig
	sarama.SyncProducer
}

func NewKafkaProducer(conf *config.MQKafkaProducerConfig) Producer {
	m := &KafkaProducer{
		conf: conf,
	}

	mqConf := sarama.NewConfig()
	mqConf.Producer.RequiredAcks = sarama.WaitForAll
	mqConf.Producer.Retry.Max = 10
	mqConf.Producer.Return.Successes = true
	mqConf.Producer.Partitioner = sarama.NewRandomPartitioner
	p, err := sarama.NewSyncProducer(conf.Brokers, mqConf)
	if err != nil {
		panic(err)
	}
	m.SyncProducer = p
	return m
}

func (m *KafkaProducer) SendMessage(key string, value []byte) (int32, int64, error) {
	partition, offset, err := m.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic:     m.conf.Topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(value),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	})
	if err != nil {
		return 0, 0, err
	}

	return partition, offset, nil
}

func (m *KafkaProducer) SendMessages(key []string, value [][]byte) error {
	err := m.SyncProducer.BeginTxn()
	if err != nil {
		log.Fatalf("KafkaProducer SendMessages txn can't be started error:%s", err.Error())
		return err
	}
	if m.SyncProducer.TxnStatus()&sarama.ProducerTxnFlagInTransaction == 0 {
		log.Fatalf("KafkaProducer SendMessages transaction must be started error:%s", err.Error())
		return err
	}

	var partition int32
	var offset int64
	for idx := range key {
		partition, offset, err = m.SyncProducer.SendMessage(&sarama.ProducerMessage{
			Topic:     m.conf.Topic,
			Key:       sarama.StringEncoder(key[idx]),
			Value:     sarama.ByteEncoder(value[idx]),
			Headers:   nil,
			Metadata:  nil,
			Offset:    0,
			Partition: 0,
			Timestamp: time.Time{},
		})
		if err != nil {
			log.Errorf("KafkaProducer SendMessages message should have been produced successfully error:%s", err.Error())

			err1 := m.SyncProducer.AbortTxn()
			if err1 != nil {
				log.Errorf("KafkaProducer SendMessages message AbortTxn error:%s", err.Error())
			}
			return err
		}
		_, _ = partition, offset
	}

	err = m.SyncProducer.CommitTxn()
	if err != nil {
		log.Fatalf("KafkaProducer SendMessages txn can't be committed, got %s", err)
	}
	return err
}

func (m *KafkaProducer) Close() error {
	if m.SyncProducer != nil {
		return m.SyncProducer.Close()
	}

	return nil
}

func (m *KafkaProducer) Topic() string {
	return m.conf.Topic
}

type _consumerOption struct {
	cb ConsumerCallback
}
type Option func(option *_consumerOption)

func WithConsumerCB(cb ConsumerCallback) Option {
	return func(option *_consumerOption) {
		option.cb = cb
	}
}

type KafkaConsumer struct {
	cb   ConsumerCallback
	conf *config.MQKafkaConsumerConfig
	*cluster.Consumer
}

func NewKafkaConsumer(conf *config.MQKafkaConsumerConfig, opt ...Option) Consumer {
	m := &KafkaConsumer{
		conf: conf,
	}

	op := _consumerOption{}
	for _, v := range opt {
		v(&op)
	}

	mqConf := cluster.NewConfig()
	mqConf.Consumer.Return.Errors = true
	mqConf.Consumer.Offsets.AutoCommit.Enable = true
	mqConf.Consumer.Offsets.AutoCommit.Interval = time.Second * 3
	mqConf.Consumer.Offsets.CommitInterval = time.Second * 3
	mqConf.Group.Return.Notifications = true
	c, err := cluster.NewConsumer(m.conf.Brokers, m.conf.Group, m.conf.Topic, mqConf)
	if err != nil {
		panic(err)
	}
	m.Consumer = c
	if op.cb != nil {
		m.SetOnMessageData(op.cb)
	}
	return m
}

func (m *KafkaConsumer) Run() {
	for {
		select {
		case err := <-m.Consumer.Errors():
			log.Errorf("kafka consumer error(%v)", err)
		case notifies := <-m.Consumer.Notifications():
			log.Infof("kafka consumer Notifications(%v)", notifies)
		case msg, ok := <-m.Consumer.Messages():
			if ok == false {
				return
			}
			if m.cb == nil {
				log.Warnf("kafka SetOnMessageData need")
			} else {
				if err := m.cb.OnMessageData(string(msg.Key), msg.Value); err != nil {
					log.Errorf("Kafka OnMessageData key:%s error:%s", string(msg.Key), err.Error())
				} else {
					m.Consumer.MarkOffset(msg, "")
				}
			}
		}
	}
}

func (m *KafkaConsumer) SetOnMessageData(cb ConsumerCallback) {
	m.cb = cb
}

func (m *KafkaConsumer) Topic() []string {
	return m.conf.Topic
}

func (m *KafkaConsumer) Group() string {
	return m.conf.Group
}

func (m *KafkaConsumer) Close() {
	m.Consumer.Close()
}
