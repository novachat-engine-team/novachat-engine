/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/25 18:31
 * @File : etcd_client_manually.go
 */

package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/hack"
	"novachat_engine/pkg/hashing"
	"novachat_engine/pkg/log"
	"strings"
	"sync"
	"time"
)

type none struct{}
type ManualClient struct {
	ctx 	context.Context
	cancel 	context.CancelFunc
	client 	*clientv3.Client
	conf 	config.EtcdClientConfig

	bClose      *atomic.Bool
	closeChan 	chan none
	mutex 		sync.RWMutex
	addrs  		[]*resolver.Address
	operatorAdd chan *resolver.Address
	operatorRemove chan *resolver.Address

	clientConnMap map[string]*grpc.ClientConn
	ketama *hashing.Ketama
}

func NewManualClient(conf config.EtcdClientConfig) (*ManualClient, error) {

	var err error
	m := &ManualClient{
		closeChan: make(chan none),
		addrs: make([]*resolver.Address, 0, 0),
		operatorAdd: make(chan *resolver.Address, 1024),
		operatorRemove: make(chan *resolver.Address, 1024),
		bClose: atomic.NewBool(false),
		clientConnMap: make(map[string]*grpc.ClientConn),
		ketama: hashing.NewKetama(hashing.DefaultReplicas, nil),
	}

	m.conf = conf

	if m.conf.DialTimeout <= 0 {
		m.conf.DialTimeout = 3
	}
	if m.conf.DialKeepAliveTime <= 0 {
		m.conf.DialKeepAliveTime = 3
	}

	if m.conf.MaxCallSendMsgSize <= 0 {
		m.conf.MaxCallSendMsgSize = grpcMaxSendMsgSize
	}

	if m.conf.MaxCallRecvMsgSize <= 0 {
		m.conf.MaxCallRecvMsgSize = grpcMaxCallMsgSize
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
		log.Errorf("NewManualClient error:%v", err)
	}


	return m, err
}

func (m *ManualClient) GoString() string {
	s := make([]string, 0, len(m.clientConnMap))
	m.mutex.RLock()
	for k := range m.clientConnMap {
		s = append(s, k)
	}
	m.mutex.RUnlock()

	return fmt.Sprintf("ManualClient clientConnMap:%v", s)
}

func (m *ManualClient) String() string {
	s := make([]string, 0, len(m.clientConnMap))
	m.mutex.RLock()
	for k := range m.clientConnMap {
		s = append(s, k)
	}
	m.mutex.RUnlock()

	return fmt.Sprintf("ManualClient clientConnMap:%v", s)
}

func (m *ManualClient) Stop() {
	if m.ctx != nil {
		m.ctx.Done()
	}
	if m.cancel != nil {
		m.cancel()
	}
	if m.client != nil {
		m.client.Close()
	}

	if m.bClose.Load() == false {
		m.bClose.CAS(false, true)
		close(m.closeChan)
		close(m.operatorAdd)
		close(m.operatorRemove)

		m.mutex.Lock()
		for _, v := range m.clientConnMap {
			v.Close()
		}
		m.mutex.Unlock()
	}
}

func (m *ManualClient) Start() {
	go func() {
		var err error
		for {
			if err = m.watchAddrUpdates(); err != nil {
				log.Errorf("ManualClient::Start::watchAddrUpdates error:%s", err.Error())
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <- m.closeChan:
				return
			case d:= <- m.operatorAdd:
				{
					bFound := false
					m.mutex.Lock()

					for idx := range m.addrs {
						if m.addrs[idx].Addr == d.Addr {
							bFound = true
							break
						}
					}
					if bFound == false {
						addr := d.Addr
						addr = strings.TrimPrefix(addr, MakeWatchKey(m.conf.Schema, m.conf.ServerName))
						log.Debugf("NewGRPCClientConnAddr addr:%s", addr)
						ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
						conn, err := NewGRPCClientConnAddr(ctx, addr, m.conf.MaxCallRecvMsgSize, m.conf.MaxCallSendMsgSize)
						if err != nil {
							log.Errorf("NewGRPCClientConn error:%v", err.Error())
						} else {
							m.addrs = append(m.addrs, d)
							m.clientConnMap[d.Addr] = conn
							m.ketama.Add(d.Addr)
						}
						cancelFunc()
					}
					m.mutex.Unlock()
				}
			case d:= <- m.operatorRemove :
				m.mutex.Lock()
				for idx := range m.addrs {
					if m.addrs[idx].Addr == d.Addr {
						m.addrs = append(m.addrs[:idx], m.addrs[idx+1:]...)

						conn, ok := m.clientConnMap[d.Addr]
						if ok == true {
							conn.Close()
							delete(m.clientConnMap, d.Addr)
							m.ketama.Remove(d.Addr)
						}
						break
					}
				}
				m.mutex.Unlock()
			}
		}
	}()
}

func (m *ManualClient) watchAddrUpdates() error {
	m.ctx, m.cancel = context.WithTimeout(context.Background(), time.Second * time.Duration(m.conf.DialTimeout))
	rootPath := MakeWatchKey(m.conf.Schema, m.conf.ServerName)
	resp, err := m.client.Get(m.ctx, rootPath, clientv3.WithPrefix())
	if err != nil {
		log.Errorf("watchAddrUpdates error:%s", err)
		return err
	}

	if resp.Kvs != nil {
		for i := range resp.Kvs {
			m.operatorAdd <- &resolver.Address{
				Addr:       hack.String(resp.Kvs[i].Key),
				ServerName: m.conf.ServerName,
			}
		}
	}

	respChan := m.client.Watch(context.Background(), rootPath, clientv3.WithPrefix())
	for resp1 := range respChan {
		for _, v := range resp1.Events {
			switch v.Type {
			case mvccpb.DELETE:
				m.operatorRemove <- &resolver.Address{
					Addr:       hack.String(v.Kv.Key),
				}
			case mvccpb.PUT:
				m.operatorAdd <- &resolver.Address{
					Addr:       hack.String(v.Kv.Key),
				}
			}
		}
	}
	return nil
}

func (m *ManualClient) GetClientConn(addr string) *grpc.ClientConn {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	builder := strings.Builder{}
	builder.WriteString("/")
	builder.WriteString(m.conf.Schema)
	builder.WriteString("/")
	builder.WriteString(addr)
	
	conn, _ := m.clientConnMap[builder.String()]
	return conn
}

func (m *ManualClient) getClientConn(addr string) *grpc.ClientConn {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	conn, _ := m.clientConnMap[addr]
	return conn
}

func (m *ManualClient) GetClientConnByKey(key string) *grpc.ClientConn {
	addr, ok := m.ketama.Get(key)
	if ok == false {
		return nil
	}

	return m.getClientConn(addr)
}

func (m *ManualClient) GetAddList() []string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	l := make([]string, 0, len(m.addrs))
	for _, v := range m.addrs {
		l = append(l, strings.TrimPrefix(v.Addr, MakeWatchKey(m.conf.Schema, m.conf.ServerName)))
	}
	return l
}
