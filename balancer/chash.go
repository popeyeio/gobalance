package balancer

import (
	"gitee.com/johng/gf/g/encoding/ghash"
	"github.com/modern-go/reflect2"
	"github.com/popeyeio/gobalance/instance"
)

type HashFunc64 func([]byte) uint64

type CHashBalancer struct {
	hash HashFunc64
}

var _ Balancer = (*CHashBalancer)(nil)

func NewCHashBalancer(fs ...HashFunc64) Balancer {
	f := ghash.BKDRHash64
	if len(fs) > 0 {
		f = fs[0]
	}

	return &CHashBalancer{
		hash: f,
	}
}

func (b *CHashBalancer) Name() string {
	return BalancerTypeCHash
}

func (b *CHashBalancer) NewPicker(instances []instance.Instance) Picker {
	return &chashPicker{
		instances: instances,
		size:      int64(len(instances)),
		hash:      b.hash,
	}
}

type chashPicker struct {
	instances []instance.Instance
	size      int64
	hash      HashFunc64
}

var _ Picker = (*chashPicker)(nil)

func (p *chashPicker) Pick(keys ...string) (instance.Instance, error) {
	if p.size <= 0 {
		return nil, instance.ErrNoInstance
	}

	var key string
	if len(keys) > 0 {
		key = keys[0]
	}
	return p.instances[choose(p.hash(reflect2.UnsafeCastString(key)), p.size)], nil
}

// choose uses jump consistent hash.
// introduction refers to https://blog.helong.info//blog/2015/03/13/jump_consistent_hash/
// paper refers to http://arxiv.org/ftp/arxiv/papers/1406/1406.2294.pdf
func choose(key uint64, size int64) (index int64) {
	index = -1
	for i := int64(0); i < size; {
		index = i
		key = key*2862933555777941757 + 1
		i = int64(float64(index+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return
}
