package instance

import (
	"errors"
	"fmt"
)

var (
	ErrNoInstance = errors.New("no instance available")
)

type Instance interface {
	fmt.Stringer
	GetValue() interface{}
	GetWeight() int
	GetIDC() string
	GetCluster() string
}
