package random

import (
	"context"
	"math/rand"

	"github.com/CoderI421/gframework/gmicro/server/rpcserver/selector"
	"github.com/CoderI421/gframework/gmicro/server/rpcserver/selector/node/direct"
)

const (
	// Name is random balancer name
	Name = "random"
)

var _ selector.Balancer = &Balancer{} // Name is balancer name

// Balancer is a random balancer.
type Balancer struct{}

// New an random selector.
func New() selector.Selector {
	return NewBuilder().Build()
}

// Pick is pick a weighted node.
func (p *Balancer) Pick(_ context.Context, nodes []selector.WeightedNode) (selector.WeightedNode, selector.DoneFunc, error) {
	if len(nodes) == 0 {
		return nil, nil, selector.ErrNoAvailable
	}
	cur := rand.Intn(len(nodes))
	selected := nodes[cur]
	d := selected.Pick()
	return selected, d, nil
}

// NewBuilder returns a selector builder with random balancer
func NewBuilder() selector.Builder {
	return &selector.DefaultBuilder{
		Balancer: &Builder{},
		Node:     &direct.Builder{},
	}
}

// Builder is random builder
type Builder struct{}

// Build creates Balancer
func (b *Builder) Build() selector.Balancer {
	return &Balancer{}
}
