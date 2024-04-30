package balancer

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"math/rand"
	"novachat_engine/pkg/log"
	"sync"
	"time"
)

const Random = "random"

var r *rand.Rand

func newRandomBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Random, &randomPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newRandomBuilder())
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type randomPickerBuilder struct{}

func (*randomPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	log.Infof("randomPickerBuilder: newPicker called with readySCs: %v", info.ReadySCs)

	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	var scs []balancer.SubConn
	for sc := range info.ReadySCs {
		scs = append(scs, sc)
	}

	return &randomPicker{
		subConns: scs,
	}
}
//func (*randomPickerBuilder) Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker {
//	log.Infof("randomPickerBuilder: newPicker called with readySCs: %v", readySCs)
//
//	if len(readySCs) == 0 {
//		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
//	}
//
//	var scs []balancer.SubConn
//	for _, sc := range readySCs {
//		scs = append(scs, sc)
//	}
//
//	return &randomPicker{
//		subConns: scs,
//	}
//}

type randomPicker struct {
	subConns []balancer.SubConn
	mu   sync.Mutex
	next int
}

func (p *randomPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	size := len(p.subConns)
	sc := p.subConns[r.Int() % size]
	p.mu.Unlock()
	return balancer.PickResult{
		SubConn: sc,
		Done:    nil,
	}, nil
}

