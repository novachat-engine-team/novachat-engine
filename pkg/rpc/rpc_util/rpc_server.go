package rpc_util

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

func NewRpcServer() *grpc.Server {
	opts := []Option{
		WithUnaryRecoveryHandler(BizUnaryRecoveryHandler),
		WithUnaryRecoveryHandler2(BizUnaryRecoveryHandler2),
		WithStreamRecoveryHandler(BizStreamRecoveryHandler),
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             time.Duration(10) * time.Second,
			PermitWithoutStream: true,
		}),
		grpc_middleware.WithUnaryServerChain(
			UnaryServerInterceptor(opts...)),
		grpc_middleware.WithStreamServerChain(
			StreamServerInterceptor(opts...)))
	return s
}
