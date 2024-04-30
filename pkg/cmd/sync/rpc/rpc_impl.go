package rpc

import (
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	sessionClient "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/cmd/sync/conf"
	"novachat_engine/pkg/cmd/sync/handler"
	"novachat_engine/pkg/cmd/sync/messageProducer"
)

type Impl struct {
	conf     *conf.Conf
	handler  *handler.Handler
	producer *messageProducer.Producer
}

func NewImpl(conf *conf.Conf, handler *handler.Handler) *Impl {
	impl := &Impl{
		conf:     conf,
		handler:  handler,
		producer: messageProducer.NewMessageProducer(&conf.Producer),
	}

	authClient.InstallAuthClient(conf.AuthRpcClient)
	sessionClient.InstallSessionManualClient(conf.SessionRpcClient)
	return impl
}

func (m *Impl) Close() {
	m.producer.Close()
	m.handler.Close()
}
