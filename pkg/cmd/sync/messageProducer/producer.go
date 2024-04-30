package messageProducer

import (
	"github.com/pkg/errors"
	"novachat_engine/pkg/cmd/sync/conf"
	"novachat_engine/pkg/mq"
	"novachat_engine/pkg/util"
)

var ErrorNotFound = errors.New("producer name not found")

type ProducerType = string

const (
	UsersProducer    ProducerType = "users"
	ChatsProducer    ProducerType = "chats"
	ChannelsProducer ProducerType = "channels"
)

type UpdatesKey = string

const (
	SyncUpdatesKey UpdatesKey = "syncUpdatesKey"
	PushUpdatesKey UpdatesKey = "pushUpdatesKey"
)

type Producer struct {
	_v map[string]mq.Producer
}

func NewMessageProducer(conf *[]conf.ProducerMQ) *Producer {
	if len(*conf) == 0 {
		panic("producer not found")
	}

	impl := &Producer{
		_v: make(map[string]mq.Producer, 3),
	}

	if index := util.Index(*conf, func(idx int) bool {
		return (*conf)[idx].Name == UsersProducer
	}); index > 0 {
		impl._v[UsersProducer] = mq.NewKafkaProducer(&(*conf)[index].KafkaProducer)
	}

	var ok bool
	for _, v := range *conf {
		switch v.Name {
		case ChatsProducer:
			fallthrough
		case ChannelsProducer:
			impl._v[v.Name] = mq.NewKafkaProducer(&v.KafkaProducer)
		case UsersProducer:
		default:
		}

		_, ok = impl._v[UsersProducer]
		if !ok {
			impl._v[UsersProducer] = mq.NewKafkaProducer(&v.KafkaProducer)
		}
	}
	if _, ok = impl._v[ChatsProducer]; !ok {
		impl._v[ChatsProducer] = impl._v[UsersProducer]
	}
	if _, ok = impl._v[ChannelsProducer]; !ok {
		impl._v[ChannelsProducer] = impl._v[UsersProducer]
	}

	return impl
}

func (m *Producer) SendMessage(producerType ProducerType, key string, value []byte) (partition int32, offset int64, err error) {
	p, ok := m._v[producerType]
	if !ok {
		return 0, 0, ErrorNotFound
	}

	return p.SendMessage(key, value)
}

func (m *Producer) SendMessages(producerType ProducerType, key string, value [][]byte) error {
	p, ok := m._v[producerType]
	if !ok {
		return ErrorNotFound
	}

	keys := make([]string, len(value))
	for idx := 0; idx < len(value); idx++ {
		keys = append(keys, key)
	}
	return p.SendMessages(keys, value)
}

func (m *Producer) Close() {
	cacheProducer := make(map[mq.Producer]struct{})
	for _, v := range m._v {
		_, ok := cacheProducer[v]
		if !ok {
			v.Close()
			cacheProducer[v] = struct{}{}
		}
	}
}
