package balancer

import (
	"sync"

	"github.com/popeyeio/gobalance/instance"
)

type WRRBalancer struct {
}

var _ Balancer = (*WRRBalancer)(nil)

func NewWRRBalancer() Balancer {
	return &WRRBalancer{}
}

func (b *WRRBalancer) Name() string {
	return BalancerTypeWRR
}

func (b *WRRBalancer) NewPicker(instances []instance.Instance) Picker {
	picker := &wrrPicker{
		instances: instances,
		size:      len(instances),
		weights:   make([]int, len(instances)),
	}
	for _, ins := range instances {
		picker.total += ins.GetWeight()
	}
	return picker
}

type wrrPicker struct {
	sync.Mutex
	instances []instance.Instance
	size      int
	weights   []int
	total     int
}

var _ Picker = (*wrrPicker)(nil)

func (p *wrrPicker) Pick(...string) (instance.Instance, error) {
	if p.size <= 0 {
		return nil, instance.ErrNoInstance
	}

	p.Lock()

	max := 0
	for i, ins := range p.instances {
		p.weights[i] += ins.GetWeight()
		if p.weights[i] > p.weights[max] {
			max = i
		}
	}
	p.weights[max] -= p.total

	p.Unlock()

	return p.instances[max], nil
}
