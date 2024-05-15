package rpc

import (
	grpc "google.golang.org/grpc"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/biz_server/conf"
	account "novachat_engine/pkg/cmd/biz_server/rpc/account"
	auth "novachat_engine/pkg/cmd/biz_server/rpc/auth"
	bots "novachat_engine/pkg/cmd/biz_server/rpc/bots"
	channels "novachat_engine/pkg/cmd/biz_server/rpc/channels"
	contacts "novachat_engine/pkg/cmd/biz_server/rpc/contacts"
	folders "novachat_engine/pkg/cmd/biz_server/rpc/folders"
	help "novachat_engine/pkg/cmd/biz_server/rpc/help"
	langpack "novachat_engine/pkg/cmd/biz_server/rpc/langpack"
	messages "novachat_engine/pkg/cmd/biz_server/rpc/messages"
	payments "novachat_engine/pkg/cmd/biz_server/rpc/payments"
	phone "novachat_engine/pkg/cmd/biz_server/rpc/phone"
	photos "novachat_engine/pkg/cmd/biz_server/rpc/photos"
	stats "novachat_engine/pkg/cmd/biz_server/rpc/stats"
	stickers "novachat_engine/pkg/cmd/biz_server/rpc/stickers"
	updates "novachat_engine/pkg/cmd/biz_server/rpc/updates"
	upload "novachat_engine/pkg/cmd/biz_server/rpc/upload"
	users "novachat_engine/pkg/cmd/biz_server/rpc/users"
	chatClient "novachat_engine/pkg/cmd/chat/rpc_client"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	sessionClient "novachat_engine/pkg/cmd/session/rpc_client"
	sfsClient "novachat_engine/pkg/cmd/sfs/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/rpc/rpc"
	helpApp "novachat_engine/service/app/help"
	"novachat_engine/service/mgo"
	"sync"
)

type Impl struct {
	conf           *conf.Config
	streamMessages sync.Map
	rpcDataPool    sync.Pool
}

func NewImpl(s *grpc.Server, conf *conf.Config) *Impl {
	mgo.InstallMongodb(&conf.MongoDB)
	err := cache.InstallRedis(conf.Redis)
	if err != nil {
		panic(err.Error())
	}

	helpApp.ParseConfig(conf.ConfigPath)

	sfsClient.InstallSfsClient(conf.SfsClient)
	msgClient.InstallMsgClient(conf.MsgClient)
	syncClient.InstallSyncClient(conf.SyncClient)
	authClient.InstallAuthClient(conf.AuthClient)
	sessionClient.InstallSessionManualClient(conf.SessionClient)
	chatClient.InstallChatClient(conf.ChatClient)

	helpServiceImpl := help.NewHelpServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCHelpService_serviceDesc, helpServiceImpl)
	mtproto.RegisterRPCHelpServiceServer(s, helpServiceImpl)

	paymentsServiceImpl := payments.NewPaymentsServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCPaymentsService_serviceDesc, paymentsServiceImpl)
	mtproto.RegisterRPCPaymentsServiceServer(s, paymentsServiceImpl)

	phoneServiceImpl := phone.NewPhoneServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCPhoneService_serviceDesc, phoneServiceImpl)
	mtproto.RegisterRPCPhoneServiceServer(s, phoneServiceImpl)

	usersServiceImpl := users.NewUsersServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCUsersService_serviceDesc, usersServiceImpl)
	mtproto.RegisterRPCUsersServiceServer(s, usersServiceImpl)

	photosServiceImpl := photos.NewPhotosServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCPhotosService_serviceDesc, photosServiceImpl)
	mtproto.RegisterRPCPhotosServiceServer(s, photosServiceImpl)

	langpackServiceImpl := langpack.NewLangpackServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCLangpackService_serviceDesc, langpackServiceImpl)
	mtproto.RegisterRPCLangpackServiceServer(s, langpackServiceImpl)

	foldersServiceImpl := folders.NewFoldersServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCFoldersService_serviceDesc, foldersServiceImpl)
	mtproto.RegisterRPCFoldersServiceServer(s, foldersServiceImpl)

	statsServiceImpl := stats.NewStatsServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCStatsService_serviceDesc, statsServiceImpl)
	mtproto.RegisterRPCStatsServiceServer(s, statsServiceImpl)

	authServiceImpl := auth.NewAuthServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCAuthService_serviceDesc, authServiceImpl)
	mtproto.RegisterRPCAuthServiceServer(s, authServiceImpl)

	contactsServiceImpl := contacts.NewContactsServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCContactsService_serviceDesc, contactsServiceImpl)
	mtproto.RegisterRPCContactsServiceServer(s, contactsServiceImpl)

	messagesServiceImpl := messages.NewMessagesServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCMessagesService_serviceDesc, messagesServiceImpl)
	mtproto.RegisterRPCMessagesServiceServer(s, messagesServiceImpl)

	uploadServiceImpl := upload.NewUploadServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCUploadService_serviceDesc, uploadServiceImpl)
	mtproto.RegisterRPCUploadServiceServer(s, uploadServiceImpl)

	stickersServiceImpl := stickers.NewStickersServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCStickersService_serviceDesc, stickersServiceImpl)
	mtproto.RegisterRPCStickersServiceServer(s, stickersServiceImpl)

	accountServiceImpl := account.NewAccountServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCAccountService_serviceDesc, accountServiceImpl)
	mtproto.RegisterRPCAccountServiceServer(s, accountServiceImpl)

	updatesServiceImpl := updates.NewUpdatesServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCUpdatesService_serviceDesc, updatesServiceImpl)
	mtproto.RegisterRPCUpdatesServiceServer(s, updatesServiceImpl)

	channelsServiceImpl := channels.NewChannelsServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCChannelsService_serviceDesc, channelsServiceImpl)
	mtproto.RegisterRPCChannelsServiceServer(s, channelsServiceImpl)

	botsServiceImpl := bots.NewBotsServiceImpl(conf)
	rpc.RegisterRpcImpl(mtproto.RPCBotsService_serviceDesc, botsServiceImpl)
	mtproto.RegisterRPCBotsServiceServer(s, botsServiceImpl)

	impl := &Impl{
		conf:        conf,
		rpcDataPool: sync.Pool{New: func() interface{} { return &rpc.RpcStreamData{} }},
	}
	return impl
}

func (m *Impl) loadOrStore(address string) (chan *rpc.RpcStreamData, bool) {
	queue, ok := m.streamMessages.LoadOrStore(address, make(chan *rpc.RpcStreamData, 1<<10))
	q := queue.(chan *rpc.RpcStreamData)
	return q, ok
}

func (m *Impl) delete(address string) {
	m.streamMessages.Delete(address)
}
