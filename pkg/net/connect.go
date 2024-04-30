/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : connect.go
 */

package net

type Connection interface {
	GetConnID() uint64
	IsClosed() bool
	Close() error
	Codec() Codec
	Receive() (interface{}, error)
	Send(msg interface{}) error
}

type Callback interface {
	OnNewConnection(conn *TcpConnection)
	OnConnectionDataArrived(c *TcpConnection, msg interface{}) error
	OnConnectionClosed(c *TcpConnection)
}