package npns

import (
	"errors"
	"novachat_engine/pkg/cmd/sync/conf"
	"novachat_engine/pkg/mq"
)

type Npns struct {
	producer mq.Producer
	consumer mq.Consumer
}

func NewNpns(conf *conf.Conf) *Npns {
	npns := &Npns{}
	if conf.NpnsProducer != nil {
		npns.producer = mq.NewKafkaProducer(conf.NpnsProducer)
	}

	if conf.NpnsConsumer != nil {
		npns.consumer = mq.NewKafkaConsumer(conf.NpnsConsumer, mq.WithConsumerCB(npns))
	}

	return npns
}

func (m *Npns) SendMessage(key string, value []byte) (partition int32, offset int64, err error) {
	if m.producer == nil {
		err = errors.New("npns sendMessage enable false")
		return 0, 0, err
	}
	return m.producer.SendMessage(key, value)
}

func (m *Npns) SendMessages(key string, value [][]byte) error {
	if m.producer == nil {
		return errors.New("npns sendMessages enable false")
	}

	keys := make([]string, len(value))
	for idx := 0; idx < len(value); idx++ {
		keys = append(keys, key)
	}
	return m.producer.SendMessages(keys, value)
}

func (m *Npns) Close() {
	if m.producer != nil {
		m.producer.Close()
	}

	if m.consumer != nil {
		m.consumer.Close()
	}
}
