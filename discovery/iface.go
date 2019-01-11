package discovery

import (
	"github.com/popeyeio/gobalance/instance"
)

type Discovery interface {
	Discover() ([]instance.Instance, error)
}
