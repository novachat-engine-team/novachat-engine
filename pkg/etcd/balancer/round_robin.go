package balancer

//
//var (
//	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
//	mu sync.Mutex
//)
//
//// Intn implements rand.Intn on the grpcrand global source.
//func Intn(n int) int {
//	mu.Lock()
//	res := r.Intn(n)
//	mu.Unlock()
//	return res
//}
//
//const Name = "round_robin"
//
//// newBuilder creates a new roundrobin balancer builder.
//func newBuilder() balancer.Builder {
//	return base.NewBalancerBuilderWithConfig(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
//}
//
//func init() {
//	balancer.Register(newBuilder())
//}
//
//type rrPickerBuilder struct{}
//
//func (*rrPickerBuilder) Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker {
//	grpclog.Infof("roundrobinPicker: newPicker called with readySCs: %v", readySCs)
//	if len(readySCs) == 0 {
//		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
//	}
//	var scs []balancer.SubConn
//	for _, sc := range readySCs {
//		scs = append(scs, sc)
//	}
//	return &rrPicker{
//		subConns: scs,
//		// Start at a random index, as the same RR balancer rebuilds a new
//		// picker when SubConn states change, and we don't want to apply excess
//		// load to the first server in the list.
//		next: Intn(len(scs)),
//	}
//}
//
//type rrPicker struct {
//	// subConns is the snapshot of the roundrobin balancer when this picker was
//	// created. The slice is immutable. Each Get() will do a round robin
//	// selection from it and return the selected SubConn.
//	subConns []balancer.SubConn
//
//	mu   sync.Mutex
//	next int
//}
//
//func (p *rrPicker) Pick(ctx context.Context, opts balancer.PickOptions) (balancer.SubConn, func(balancer.DoneInfo), error) {
//	p.mu.Lock()
//	sc := p.subConns[p.next]
//	p.next = (p.next + 1) % len(p.subConns)
//	p.mu.Unlock()
//	return sc, nil, nil
//}
