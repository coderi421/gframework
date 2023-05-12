package rpcserver

import (
	"context"
	"time"

	"google.golang.org/grpc"
	grpcinsecure "google.golang.org/grpc/credentials/insecure"

	"github.com/CoderI421/gframework/gmicro/registry"
	"github.com/CoderI421/gframework/gmicro/server/rpcserver/clientinterceptors"
	"github.com/CoderI421/gframework/gmicro/server/rpcserver/resolver/discovery"
	"github.com/CoderI421/gframework/pkg/log"
)

type ClientOption func(o *clientOptions)
type clientOptions struct {
	// 服务端的地址
	endpoint string
	// 超时时间
	timeout time.Duration
	// 服务发现接口
	discovery registry.Discovery
	// Unary 服务的拦截器
	unaryInts []grpc.UnaryClientInterceptor
	// Stream 服务的拦截器
	streamInts []grpc.StreamClientInterceptor
	// 用户自己设置 grpc 连接的结构体,例如: grpc.WithInsecure()， grpc.WithTransportCredentials()
	rpcOpts []grpc.DialOption
	// 根据 Name 生成负载均衡的策略
	balancerName string

	// 客户端的日志
	logger log.Logger
}

// 设置服务端的地址
func WithEndpoint(endpoint string) ClientOption {
	return func(o *clientOptions) {
		o.endpoint = endpoint
	}
}

// 设置超时时间
func WithClientTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOptions) {
		o.timeout = timeout
	}
}

// 设置服务发现
func WithDiscovery(d registry.Discovery) ClientOption {
	return func(o *clientOptions) {
		o.discovery = d
	}
}

// 设置拦截器
func WithClientUnaryInterceptor(in ...grpc.UnaryClientInterceptor) ClientOption {
	return func(o *clientOptions) {
		o.unaryInts = in
	}
}

// 设置stream拦截器
func WithClientStreamInterceptor(in ...grpc.StreamClientInterceptor) ClientOption {
	return func(o *clientOptions) {
		o.streamInts = in
	}
}

// 设置grpc的dial选项
func WithClientOptions(opts ...grpc.DialOption) ClientOption {
	return func(o *clientOptions) {
		o.rpcOpts = opts
	}
}

// 设置负载均衡器
func WithBalancerName(name string) ClientOption {
	return func(o *clientOptions) {
		o.balancerName = name
	}
}

// 设置日志
func WithClientLogger(logger log.Logger) ClientOption {
	return func(o *clientOptions) {
		o.logger = logger
	}
}

// 非安全拨号
func DialInsecure(ctx context.Context, opts ...ClientOption) (*grpc.ClientConn, error) {
	return dial(ctx, true, opts...)
}

func Dial(ctx context.Context, opts ...ClientOption) (*grpc.ClientConn, error) {
	return dial(ctx, false, opts...)
}

func dial(ctx context.Context, insecure bool, opts ...ClientOption) (*grpc.ClientConn, error) {
	// 默认配置
	options := clientOptions{
		timeout:      2000 * time.Millisecond,
		balancerName: "round_robin",
	}

	for _, o := range opts {
		o(&options)
	}

	//TODO 客户端默认拦截器
	ints := []grpc.UnaryClientInterceptor{
		clientinterceptors.TimeoutInterceptor(options.timeout),
	}
	streamInts := []grpc.StreamClientInterceptor{}

	if len(options.unaryInts) > 0 {
		ints = append(ints, options.unaryInts...)
	}
	if len(options.streamInts) > 0 {
		streamInts = append(streamInts, options.streamInts...)
	}

	//可以由用户端自己传递 这些默认的
	grpcOpts := []grpc.DialOption{
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "` + options.balancerName + `"}`),
		grpc.WithChainUnaryInterceptor(ints...),
		grpc.WithChainStreamInterceptor(streamInts...),
	}

	//服务发现的选项
	if &options.discovery != nil {
		grpcOpts = append(grpcOpts, grpc.WithResolvers(
			discovery.NewBuilder(options.discovery,
				discovery.WithInsecure(insecure)),
		))
	}

	// 如果是非安全模式
	if insecure {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(grpcinsecure.NewCredentials()))
	}

	if len(options.rpcOpts) > 0 {
		grpcOpts = append(grpcOpts, options.rpcOpts...)
	}

	return grpc.DialContext(ctx, options.endpoint, grpcOpts...)
}
