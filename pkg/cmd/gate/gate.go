/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package gate

import (
	"fmt"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/gate/conf"
	"novachat_engine/pkg/cmd/gate/handler"
	"novachat_engine/pkg/cmd/gate/handshake"
	"novachat_engine/pkg/cmd/gate/rpc"
	sessionClient "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/command"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
	"novachat_engine/pkg/net/mtserver"
	rpcPkg "novachat_engine/pkg/rpc/rpc"
	"novachat_engine/pkg/rpc/rpc_util"
	"os"
	"strconv"
	"sync"
)

type Application struct {
	conf          conf.Config
	pprofListener net.Listener
	rpcServer     *grpc.Server
	wg            sync.WaitGroup
	serverMain    *mtserver.MTProtoServer
	rpcDiscovery  *etcd.EtcdNamingServer
	gateHandler   *handler.GateHandler
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
	handshake.KeyFingerPrintLength = len(m.conf.Key.KeyFingerPrint)
	keyFingerprint, err := strconv.ParseUint(m.conf.Key.KeyFingerPrint, 10, 64)
	if err != nil {
		log.Fatalf("%#v", err)
		return err
	}

	hk, err := handshake.NewHandshake(m.conf.Key.KeyFile, int64(keyFingerprint))
	if err != nil {
		log.Fatalf("%#v", err)
		return err
	}

	listener, err := net.Listen("tcp", m.conf.RpcServer.Addr)
	if err != nil {
		log.Fatalf("failed to tcp server: %s", err)
		return err
	}
	defer listener.Close()
	s := rpc_util.NewRpcServer()

	rpcPkg.RegisterRpcServiceServer(s, rpc.NewGateImpl(m))

	defer s.GracefulStop()

	m.wg.Add(1)
	go func() {
		if err = s.Serve(listener); err != nil {
			log.Fatalf("failed to rpcServer: %s", err)
			panic(err)
		}
		m.wg.Done()
	}()

	m.rpcDiscovery, err = etcd.NewEtcdServer(m.conf.RpcDiscovery, m.conf.RpcServer.Addr)
	if err != nil {
		log.Fatalf("%#v", err)
		return err
	}

	sessionClient.InstallSessionManualClient(m.conf.SessionRpcClient)
	authClient.InstallAuthClient(m.conf.AuthRpcClient)

	m.gateHandler = handler.NewGateHandler(&m.conf, hk)
	m.serverMain = mtserver.NewMTProtoServer(m.conf.Server.Addr, m.gateHandler)

	return m.serverMain.Serve()
}

func (m *Application) Close() {
	if m.rpcDiscovery != nil {
		m.rpcDiscovery.UnRegister(m.conf.RpcDiscovery.ServerName, m.conf.RpcServer.Addr)
	}

	if m.serverMain != nil {
		m.serverMain.Stop()
	}

	if m.rpcServer != nil {
		m.rpcServer.GracefulStop()
	}

	if m.pprofListener != nil {
		m.pprofListener.Close()
	}

	m.wg.Wait()
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

	m.conf.Key = cnf.Key
}

func (m *Application) GetConnection(connId uint64) *net2.TcpConnection {
	return m.serverMain.GetConnection(connId)
}
