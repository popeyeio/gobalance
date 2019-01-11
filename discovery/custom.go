package discovery

import (
	"github.com/popeyeio/gobalance/instance"
)

type CustomDiscovery struct {
	instances []instance.Instance
}

var _ Discovery = (*CustomDiscovery)(nil)

func NewCustomDiscovery(instances ...instance.Instance) Discovery {
	return &CustomDiscovery{
		instances: instances,
	}
}

func (d *CustomDiscovery) Discover() ([]instance.Instance, error) {
	return d.instances, nil
}
