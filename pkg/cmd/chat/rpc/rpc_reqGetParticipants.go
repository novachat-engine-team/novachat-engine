package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
)

func (impl *Impl) ReqGetParticipants(ctx context.Context, request *chatService.GetParticipants) (*types.Any, error) {
	panic("ReqGetParticipants")
}
