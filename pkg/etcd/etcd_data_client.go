/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/6 16:18
 * @File : etcd_data_client.go
 */

package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/hack"
	"novachat_engine/pkg/log"
	"strings"
	"time"
)

type DataOperator int32
const (
	DataOperatorAdd  	DataOperator = 0
	DataOperatorRemove  DataOperator = 1
)

type DataOperatorData struct {
	Operator DataOperator
	Key		 string
	Value	 []byte
}

type DataClient struct {
	ctx 	context.Context
	cancel 	context.CancelFunc
	client 	*clientv3.Client
	conf 	config.EtcdClientConfig

	closeChan 	chan none

	chanData	chan DataOperatorData
}

func NewDataClient(conf config.EtcdClientConfig, chanData chan DataOperatorData) (*DataClient, error) {

	var err error
	m := &DataClient{
		closeChan: make(chan none),
		chanData: chanData,
	}

	m.conf = conf
	if m.conf.DialTimeout == 0 {
		m.conf.DialTimeout = 3
	}

	m.client, err = clientv3.New(clientv3.Config{
		Endpoints:            strings.Split(conf.EtcdAddr, ";"),
		AutoSyncInterval:     0,
		DialTimeout:          time.Second * time.Duration(m.conf.DialTimeout),
		DialKeepAliveTime:    time.Second * time.Duration(m.conf.DialKeepAliveTime),
		DialKeepAliveTimeout: time.Duration(conf.DialKeepAliveTimeout),
		MaxCallSendMsgSize:   int(m.conf.MaxCallSendMsgSize),
		MaxCallRecvMsgSize:   int(m.conf.MaxCallRecvMsgSize),
		TLS:                  nil,
		Username:             conf.UserName,
		Password:             conf.Password,
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  true,
	})
	if err != nil {
		log.Errorf("NewDataClient error:%v", err)
	}

	return m, err
}

func (m *DataClient) GoString() string {
	return fmt.Sprintf("DataClient")
}

func (m *DataClient) String() string {
	return fmt.Sprintf("DataClient")
}

func (m *DataClient) Stop() {
	if m.ctx != nil {
		m.ctx.Done()
	}
	if m.cancel != nil {
		m.cancel()
	}
	if m.client != nil {
		m.client.Close()
	}
}

func (m *DataClient) Start() {
	maxBase := time.Duration(20)
	minBase := time.Duration(3)
	span := time.Duration(5)
	base := minBase
	go func() {
		var err error
		for {
			if err = m.watchAddrUpdates(); err != nil {
				if base >= maxBase {
					base = maxBase
				} else {
					base += span
				}
				time.Sleep(time.Second * base)
				if base >= maxBase {
					log.Fatalf("DataClient::Start::watchAddrUpdates error:%s", err.Error())
				} else {
					log.Errorf("DataClient::Start::watchAddrUpdates error:%s", err.Error())
				}
			} else {
				base = minBase
			}
		}
	}()

	go func() {
		for {
			select {
			case <- m.closeChan:
				close(m.chanData)
				return
			}
		}
	}()
}

func (m *DataClient) watchAddrUpdates() error {
	m.ctx, m.cancel = context.WithTimeout(context.Background(), time.Second * time.Duration(m.conf.DialTimeout))
	rootPath := MakeWatchKey(m.conf.Schema, m.conf.ServerName)
	resp, err := m.client.Get(m.ctx, rootPath, clientv3.WithPrefix())
	if err != nil {
		log.Errorf("watchAddrUpdates error:%s", err)
		return err
	}

	if resp.Kvs != nil {
		for i := range resp.Kvs {
			m.chanData <- DataOperatorData{
				Operator: DataOperatorAdd,
				Key:      hack.String(resp.Kvs[i].Key),
				Value:    resp.Kvs[i].Value,
			}
		}
	}

	respChan := m.client.Watch(context.Background(), rootPath, clientv3.WithPrefix())
	for resp1 := range respChan {
		for _, v := range resp1.Events {
			switch v.Type {
			case mvccpb.DELETE:
				m.chanData <- DataOperatorData{
					Operator: DataOperatorRemove,
					Key:      hack.String(v.Kv.Key),
					Value:    v.Kv.Value,
				}
			case mvccpb.PUT:
				m.chanData <- DataOperatorData{
					Operator: DataOperatorAdd,
					Key:      hack.String(v.Kv.Key),
					Value:    v.Kv.Value,
				}
			}
		}
	}
	return nil
}
