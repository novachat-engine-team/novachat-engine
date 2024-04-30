/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package auth

import (
	"fmt"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"novachat_engine/pkg/cmd/auth/conf"
	"novachat_engine/pkg/cmd/auth/rpc"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/command"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/rpc_util"
	"os"
	"sync"
)

type Application struct {
	conf          conf.Config
	pprofListener net.Listener
	rpcServer     *grpc.Server
	wg            sync.WaitGroup
	rpcDiscovery  *etcd.EtcdNamingServer
}

func (m *Application) Initialize(fileName string) (*config.LogConfig, error) {
	fi, err := os.Stat(fileName)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", fileName, err.Error())
	}
	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("not found %s", fileName)
	}

	data, _ := ioutil.ReadFile(fileName)
	err = yaml.Unmarshal(data, &m.conf)
	if err != nil {
		return nil, fmt.Errorf("fileName:%s parse error:%s", fileName, err.Error())
	}

	logLevel := log.FormatLevelString(m.conf.Log.LogLevel)
	if !log.LevelStringCheck(logLevel) {
		return nil, fmt.Errorf("log Configure level error:`%s` %+v", m.conf.Log.LogLevel, m.conf.Log)
	}
	return &m.conf.Log, nil
}

func (m *Application) RunLoop() error {
	listener, err := net.Listen("tcp", m.conf.RpcServer.Addr)
	if err != nil {
		log.Fatalf("failed to tcp server: %s", err)
		return err
	}
	defer listener.Close()
	m.rpcServer = rpc_util.NewRpcServer()

	authService.RegisterAuthServiceServer(m.rpcServer, rpc.NewAuthImpl(&m.conf))

	m.rpcDiscovery, err = etcd.NewEtcdServer(m.conf.RpcDiscovery, m.conf.RpcServer.Addr)
	if err != nil {
		log.Fatalf("%#v", err)
		return err
	}
	defer m.rpcDiscovery.UnRegister(m.conf.RpcDiscovery.ServerName, m.conf.RpcServer.Addr)

	if err = m.rpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to rpcServer: %s", err)
		return err
	}
	return nil
}

func (m *Application) Close() {
	if m.rpcDiscovery != nil {
		m.rpcDiscovery.UnRegister(m.conf.RpcDiscovery.ServerName, m.conf.RpcServer.Addr)
	}

	if m.rpcServer != nil {
		m.rpcServer.GracefulStop()
	}

	if m.pprofListener != nil {
		m.pprofListener.Close()
	}
}

func (m *Application) OnEvent(filename string) {
	log.Warnf("Watcher OnEvent:%s", filename)

	cnf := conf.Config{}
	data, _ := ioutil.ReadFile(filename)
	err := yaml.Unmarshal(data, &cnf)
	if err != nil {
		log.Errorf("Unmarshal filename:%s error:%s", filename, err.Error())
		return
	}
	command.Watcher.Add(filename)

	logLevel := log.FormatLevelString(cnf.Log.LogLevel)
	if !log.LevelStringCheck(logLevel) {
		log.Errorf("log Configure level error:`%s` %+v", cnf.Log.LogLevel, cnf.Log)
		return
	}

	allScopes := log.Scopes()
	scopes := allScopes[log.DefaultScopeName]
	if scopes != nil {
		scopes.SetOutputLevel(log.StringToLevel(logLevel))
	}
}
