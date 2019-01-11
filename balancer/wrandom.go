package balancer

import (
	"github.com/popeyeio/gobalance/instance"
	"github.com/valyala/fastrand"
)

type WRandomBalancer struct {
}

var _ Balancer = (*WRandomBalancer)(nil)

func NewWRandomBalancer() Balancer {
	return &WRandomBalancer{}
}

func (b *WRandomBalancer) Name() string {
	return BalancerTypeWRandom
}

func (b *WRandomBalancer) NewPicker(instances []instance.Instance) Picker {
	total := 0
	for _, instance := range instances {
		total += instance.GetWeight()
	}

	return &wrandomPicker{
		instances: instances,
		total:     total,
	}
}

type wrandomPicker struct {
	instances []instance.Instance
	total     int
}

var _ Picker = (*wrandomPicker)(nil)

func (p *wrandomPicker) Pick(...string) (instance.Instance, error) {
	if len(p.instances) > 0 {
		w, n := uint32(0), fastrand.Uint32n(uint32(p.total))+1
		for _, instance := range p.instances {
			w = uint32(instance.GetWeight())
			if n <= w {
				return instance, nil
			}
			n -= w
		}
	}
	return nil, instance.ErrNoInstance
}
