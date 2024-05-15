/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/3/28 16:50
 * @File : rpc_client_option.go
 */

package rpc_util

import "google.golang.org/grpc"

type clientOption struct {
	UnaryClientRecover       UnaryClientCallback
	UnaryClient              UnaryClientCallback
	UnaryClientStreamRecover UnaryClientCallback
	retry                    int32
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

type CallOption struct {
	grpc.EmptyCallOption
	retry int32
	apply func(c *CallOption)
}

type ClientOption func(option *clientOption)

func WithRetry(retry int32) *CallOption {
	return &CallOption{
		apply: func(c *CallOption) {
			c.retry = retry
		},
	}
}

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
