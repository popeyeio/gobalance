package instance

import (
	"fmt"
)

type CustomInstance struct {
	Value   interface{}
	Weight  int
	IDC     string
	Cluster string
}

var _ Instance = (*CustomInstance)(nil)

func NewCustomInstance(value interface{}) *CustomInstance {
	return &CustomInstance{
		Value: value,
	}
}

func (i *CustomInstance) SetWeight(weight int) *CustomInstance {
	i.Weight = weight
	return i
}

func (i *CustomInstance) String() string {
	return fmt.Sprintf("value:%+v, weight:%d, idc:%s, cluster:%s",
		i.Value, i.Weight, i.IDC, i.Cluster)
}

func (i *CustomInstance) GetValue() interface{} {
	return i.Value
}

func (i *CustomInstance) GetWeight() int {
	return i.Weight
}

func (i *CustomInstance) GetIDC() string {
	return i.IDC
}

func (i *CustomInstance) GetCluster() string {
	return i.Cluster
}
