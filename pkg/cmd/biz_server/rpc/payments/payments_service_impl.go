/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : payments_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cmd/biz_server/conf"
)

type PaymentsServiceImpl struct {
	conf *conf.Config
	//*payment.Payment
	//*NotifyServer
	//rabbitMQProducer *mq.RabbitMqProducer
	//rabbitMQConsumer *mq.RabbitMqConsumer
}

func NewPaymentsServiceImpl(conf *conf.Config) *PaymentsServiceImpl {

	//if conf.BizLogic.Payment == true {
	//	payment.InstallPayment(conf.PaymentRpcClient)
	//
	//}

	impl := &PaymentsServiceImpl{
		conf: conf,
		//Payment: payment.NewPayment(conf.BizLogic.Payment),
	}

	//if conf.BizLogic.Payment == true {
	//    impl.NotifyServer = NewNotifyServer(conf, impl)
	//
	//    impl.rabbitMQProducer = mq.NewRabbitMqProducer(conf.RabbitMQProducer)
	//    impl.rabbitMQConsumer = mq.NewRabbitMqConsumer(conf.RabbitMQConsumer)
	//    impl.rabbitMQConsumer.SetOnMessageData(impl)
	//    impl.rabbitMQConsumer.Run()
	//
	//    impl.Start()
	//}

	return impl
}
