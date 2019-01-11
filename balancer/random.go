package balancer

import (
	"github.com/popeyeio/gobalance/instance"
	"github.com/valyala/fastrand"
)

type RandomBalancer struct {
}

var _ Balancer = (*RandomBalancer)(nil)

func NewRandomBalancer() Balancer {
	return &RandomBalancer{}
}

func (b *RandomBalancer) Name() string {
	return BalancerTypeRandom
}

func (b *RandomBalancer) NewPicker(instances []instance.Instance) Picker {
	return &randomPicker{
		instances: instances,
		size:      uint32(len(instances)),
	}
}

type randomPicker struct {
	instances []instance.Instance
	size      uint32
}

var _ Picker = (*randomPicker)(nil)

func (p *randomPicker) Pick() (instance.Instance, error) {
	if p.size <= 0 {
		return nil, instance.ErrNoInstance
	}
	return p.instances[fastrand.Uint32n(p.size)], nil
}
