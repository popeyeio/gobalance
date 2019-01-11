package balancer

import (
	"sync/atomic"

	"github.com/popeyeio/gobalance/instance"
)

type RRBalancer struct {
}

var _ Balancer = (*RRBalancer)(nil)

func NewRRBalancer() Balancer {
	return &RRBalancer{}
}

func (b *RRBalancer) Name() string {
	return BalancerTypeRR
}

func (b *RRBalancer) NewPicker(instances []instance.Instance) Picker {
	return &rrPicker{
		instances: instances,
		size:      uint32(len(instances)),
	}
}

type rrPicker struct {
	instances []instance.Instance
	size      uint32
	next      uint32
}

var _ Picker = (*rrPicker)(nil)

func (p *rrPicker) Pick() (instance.Instance, error) {
	if p.size <= 0 {
		return nil, instance.ErrNoInstance
	}
	return p.instances[(atomic.AddUint32(&p.next, 1)-1)%p.size], nil
}
