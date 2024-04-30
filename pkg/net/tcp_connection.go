/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : tcp_connection.go
 */

package net

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
)

const gcCounter int32 = 128

var (
	closeCounter int32  = 0
	globalID     uint64 = 0
)

type none struct{}

type TcpConnection struct {
	conn      net.Conn
	codec     Codec
	closeFlag int32
	cb        Callback
	id        uint64
	closeChan chan none
	sendChan  chan interface{}
	recvMutex sync.Mutex
	sendMutex sync.RWMutex
	Context   interface{}
}

func NewTcpConnection(codec Codec, conn net.Conn, sendChanSize int32, cb Callback) *TcpConnection {

	ins := &TcpConnection{
		conn:      conn,
		codec:     codec,
		closeFlag: 0,
		cb:        cb,
		id:        atomic.AddUint64(&globalID, 1),
		closeChan: make(chan none),
		sendChan:  make(chan interface{}, sendChanSize),
	}

	go ins.runEventLoop()
	return ins
}

func (t *TcpConnection) GoString() string {
	if t.conn == nil {
		return fmt.Sprintf("id:%d", t.id)
	}

	return fmt.Sprintf("id:%d conn:%+v", t.id, t.conn)
}

func (t *TcpConnection) String() string {
	if t.conn == nil {
		return fmt.Sprintf("id:%d", t.id)
	}

	return fmt.Sprintf("id:%d conn:%+v", t.id, t.conn)
}

func (t *TcpConnection) RemoteAddr() net.Addr {
	return t.conn.RemoteAddr()
}

func (t *TcpConnection) runEventLoop() {
	defer t.Close()

	for {
		select {
		case msg, ok := <-t.sendChan:
			if !ok || t.codec.Send(msg) != nil {
				return
			}
		case <-t.closeChan:
			return
		}
	}
}

func (t *TcpConnection) GetConnID() uint64 {
	return t.id
}

func (t *TcpConnection) IsClosed() bool {
	return atomic.LoadInt32(&t.closeFlag) == 1
}

func (t *TcpConnection) Close() error {
	if atomic.CompareAndSwapInt32(&t.closeFlag, 0, 1) {
		if t.cb != nil {
			t.cb.OnConnectionClosed(t)
		}
		close(t.closeChan)

		t.sendMutex.Lock()
		defer t.sendMutex.Unlock()

		close(t.sendChan)
		atomic.AddInt32(&closeCounter, 1)
		err := t.codec.Close()
		return err
	}

	if atomic.LoadInt32(&closeCounter) >= gcCounter {
		atomic.StoreInt32(&closeCounter, 0)

		runtime.GC()
	}
	return nil
}

func (t *TcpConnection) Codec() Codec {
	return t.codec
}

func (t *TcpConnection) Receive() (interface{}, error) {
	t.recvMutex.Lock()
	defer t.recvMutex.Unlock()

	msg, err := t.codec.Receive()
	if err != nil {
		t.Close()
	}
	return msg, err
}

func (t *TcpConnection) Send(msg interface{}) error {
	if t.sendChan == nil {
		if t.IsClosed() {
			return errors.New("closed")
		}

		//t.sendMutex.Lock()

		err := t.codec.Send(msg)
		if err != nil {
			//t.sendMutex.Unlock()
			t.Close()
		} else {
			//t.sendMutex.Unlock()
		}
		return err
	}

	t.sendMutex.RLock()
	if t.IsClosed() {
		t.sendMutex.RUnlock()
		return errors.New("closed")
	}

	select {
	case t.sendChan <- msg:
		t.sendMutex.RUnlock()
		return nil
	default:
		t.sendMutex.RUnlock()
		t.Close()
		return errors.New("error")
	}
}
