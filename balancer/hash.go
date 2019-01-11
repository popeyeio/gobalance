package balancer

import (
	"gitee.com/johng/gf/g/encoding/ghash"
	"github.com/modern-go/reflect2"
	"github.com/popeyeio/gobalance/instance"
)

type HashFunc func([]byte) uint32

type HashBalancer struct {
	hash HashFunc
}

var _ Balancer = (*HashBalancer)(nil)

func NewHashBalancer(fs ...HashFunc) Balancer {
	f := ghash.BKDRHash
	if len(fs) > 0 {
		f = fs[0]
	}

	return &HashBalancer{
		hash: f,
	}
}

func (b *HashBalancer) Name() string {
	return BalancerTypeHash
}

func (b *HashBalancer) NewPicker(instances []instance.Instance) Picker {
	return &hashPicker{
		instances: instances,
		size:      uint32(len(instances)),
		hash:      b.hash,
	}
}

type hashPicker struct {
	instances []instance.Instance
	size      uint32
	hash      HashFunc
}

var _ Picker = (*hashPicker)(nil)

func (p *hashPicker) Pick(keys ...string) (instance.Instance, error) {
	if p.size <= 0 {
		return nil, instance.ErrNoInstance
	}

	var key string
	if len(keys) > 0 {
		key = keys[0]
	}
	return p.instances[p.hash(reflect2.UnsafeCastString(key))%p.size], nil
}
