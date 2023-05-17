package rpcserver

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"

	"github.com/CoderI421/gframework/gmicro/registry"
	"github.com/CoderI421/gframework/gmicro/server/rpcserver/selector"
)

const (
	balancerName = "selector"
)

var (
	_ base.PickerBuilder = &balancerBuilder{}
	_ balancer.Picker    = &balancerPicker{}
)

// InitBuilder initializes the builder for the selector balancer.
// 初始化 grpc balancer
func InitBuilder() {
	b := base.NewBalancerBuilder(
		balancerName,
		&balancerBuilder{
			// 获取全局的 selector
			builder: selector.GlobalSelector(),
		},
		base.Config{HealthCheck: true},
	)
	balancer.Register(b)
}

type balancerBuilder struct {
	builder selector.Builder
}

// Build creates a grpc Picker.
func (b *balancerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		// Block the RPC until a new picker is available via UpdateState().
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	nodes := make([]selector.Node, 0, len(info.ReadySCs))
	for conn, info := range info.ReadySCs {
		ins, _ := info.Address.Attributes.Value("rawServiceInstance").(*registry.ServiceInstance)
		nodes = append(nodes, &grpcNode{
			Node:    selector.NewNode("grpc", info.Address.Addr, ins),
			subConn: conn,
		})
	}
	p := &balancerPicker{
		//自己传递的算法 注意先后顺序
		selector: b.builder.Build(),
	}
	p.selector.Apply(nodes)
	return p
}

// balancerPicker is a grpc picker.
type balancerPicker struct {
	selector selector.Selector
}

// Pick pick instances.
func (p *balancerPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	n, done, err := p.selector.Select(info.Ctx)
	if err != nil {
		return balancer.PickResult{}, err
	}

	return balancer.PickResult{
		SubConn: n.(*grpcNode).subConn,
		Done: func(di balancer.DoneInfo) {
			done(info.Ctx, selector.DoneInfo{
				Err:           di.Err,
				BytesSent:     di.BytesSent,
				BytesReceived: di.BytesReceived,
				ReplyMD:       Trailer(di.Trailer),
			})
		},
	}, nil
}

// Trailer is a grpc trailder MD.
type Trailer metadata.MD

// Get get a grpc trailer value.
func (t Trailer) Get(k string) string {
	v := metadata.MD(t).Get(k)
	if len(v) > 0 {
		return v[0]
	}
	return ""
}

type grpcNode struct {
	selector.Node
	subConn balancer.SubConn
}
