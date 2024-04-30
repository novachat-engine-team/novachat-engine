/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/3/28 16:50
 * @File : rpc_client_option.go
 */

package rpc_util

type clientOption struct {
	UnaryClientRecover 			UnaryClientCallback
	UnaryClient 	   			UnaryClientCallback
	UnaryClientStreamRecover 	UnaryClientCallback
}

func evaluateClientOptions(opts []ClientOption) *clientOption {
	optCopy := &clientOption{
		UnaryClientRecover:       nil,
		UnaryClient:              nil,
		UnaryClientStreamRecover: nil,
	}

	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type ClientOption func(option *clientOption)

func WithUnaryClientRecoveryHandler(f UnaryClientCallback) ClientOption {
	return func(o *clientOption) {
		o.UnaryClientRecover = f
	}
}

func WithUnaryClientHandler(f UnaryClientCallback) ClientOption {
	return func(o *clientOption) {
		o.UnaryClient = f
	}
}

func WithUnaryClientStreamHandler(f UnaryClientCallback) ClientOption {
	return func(o *clientOption) {
		o.UnaryClientStreamRecover = f
	}
}