package balancer

import (
	"errors"
	"strings"

	"github.com/popeyeio/gobalance/instance"
)

const (
	BalancerTypeRR      = "RR"
	BalancerTypeWRR     = "WRR"
	BalancerTypeRandom  = "RANDOM"
	BalancerTypeWRandom = "WRANDOM"
	BalancerTypeHash    = "HASH"
	BalancerTypeCHash   = "CHASH"
)

var (
	ErrNoBalancer = errors.New("no balancer available")
)

type Balancer interface {
	Name() string
	NewPicker([]instance.Instance) Picker
}

type Picker interface {
	Pick(...string) (instance.Instance, error)
}

func CreateBalancer(name string) (Balancer, error) {
	switch strings.ToUpper(name) {
	case BalancerTypeRR:
		return NewRRBalancer(), nil
	case BalancerTypeWRR:
		return NewWRRBalancer(), nil
	case BalancerTypeRandom:
		return NewRandomBalancer(), nil
	case BalancerTypeWRandom:
		return NewWRandomBalancer(), nil
	case BalancerTypeHash:
		return NewHashBalancer(), nil
	case BalancerTypeCHash:
		return NewCHashBalancer(), nil
	}
	return nil, ErrNoBalancer
}
