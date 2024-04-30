package balancer

import (
	"fmt"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"novachat_engine/pkg/hashing"
	"novachat_engine/pkg/log"
	"sync"
)

const Ketama = "ketama"

func newKetamaBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Ketama, &ketamaPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newKetamaBuilder())
}

type ketamaPickerBuilder struct{}

func (*ketamaPickerBuilder)Name() string {
	return Ketama
}

func (*ketamaPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	log.Infof("ketamaPickerBuilder: newPicker called with readySCs: %v", info.ReadySCs)

	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	replicas := len(info.ReadySCs)
	if replicas < hashing.DefaultReplicas {
		replicas = hashing.DefaultReplicas
	}

	subConns := make(map[string]balancer.SubConn)
	ketamaHashing := hashing.NewKetama(replicas, nil)
	for addr, sc := range info.ReadySCs {
		ketamaHashing.Add(sc.Address.Addr)
		subConns[sc.Address.Addr] = addr
	}

	return &ketamaPicker{
		subConns: subConns,
		ketamaHashing: ketamaHashing,
	}
}

//func (*ketamaPickerBuilder) Build(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
//	log.Infof("ketamaPickerBuilder: newPicker called with readySCs: %v", readySCs)
//
//	if len(readySCs) == 0 {
//		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
//	}
//
//	replicas := len(readySCs)
//	if replicas < hashing.DefaultReplicas {
//		replicas = hashing.DefaultReplicas
//	}
//
//	subConns := make(map[string]balancer.SubConn)
//	ketamaHashing := hashing.NewKetama(replicas, nil)
//	for addr, sc := range readySCs {
//		ketamaHashing.Add(addr.Addr)
//		subConns[addr.Addr] = sc
//	}
//
//	return &ketamaPicker{
//		subConns: subConns,
//		ketamaHashing: ketamaHashing,
//	}
//}

type ketamaPicker struct {
	subConns map[string]balancer.SubConn

	ketamaHashing *hashing.Ketama
	mu   sync.Mutex
	next int
}

func (p *ketamaPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	addr, ok := p.ketamaHashing.Get(opts.FullMethodName)
	if ok == false {
		p.mu.Unlock()
		err := fmt.Errorf("ketamaPicker::Pick error len(subConns) = %d, subConns:%v", len(p.subConns), p.subConns)
		log.Warnf(err.Error())
		return balancer.PickResult{}, err
	}

	sc := p.subConns[addr]
	p.mu.Unlock()
	return balancer.PickResult{
		SubConn: sc,
		Done:    nil,
	}, nil
}
