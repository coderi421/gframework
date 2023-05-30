package rpcserver

import (
	"context"
	"time"

	"google.golang.org/grpc"
	grpcinsecure "google.golang.org/grpc/credentials/insecure"

	"github.com/coderi421/gframework/gmicro/registry"
	"github.com/coderi421/gframework/gmicro/server/rpcserver/resolver/discovery"
	"github.com/coderi421/gframework/pkg/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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

	// 是否开启链路追踪
	enableTracing bool
}

// WithEndpoint 设置服务端的地址
func WithEndpoint(endpoint string) ClientOption {
	return func(o *clientOptions) {
		o.endpoint = endpoint
	}
}

// WithClientTimeout 设置超时时间
func WithClientTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOptions) {
		o.timeout = timeout
	}
}

// WithDiscovery 设置服务发现
func WithDiscovery(d registry.Discovery) ClientOption {
	return func(o *clientOptions) {
		o.discovery = d
	}
}

// WithClientUnaryInterceptor 设置拦截器
func WithClientUnaryInterceptor(in ...grpc.UnaryClientInterceptor) ClientOption {
	return func(o *clientOptions) {
		o.unaryInts = in
	}
}

// WithClientStreamInterceptor 设置stream拦截器
func WithClientStreamInterceptor(in ...grpc.StreamClientInterceptor) ClientOption {
	return func(o *clientOptions) {
		o.streamInts = in
	}
}

// WithClientOptions 设置grpc的dial选项
func WithClientOptions(opts ...grpc.DialOption) ClientOption {
	return func(o *clientOptions) {
		o.rpcOpts = opts
	}
}

// WithBalancerName 设置负载均衡器
func WithBalancerName(name string) ClientOption {
	return func(o *clientOptions) {
		o.balancerName = name
	}
}

// WithClientLogger 设置日志
func WithClientLogger(logger log.Logger) ClientOption {
	return func(o *clientOptions) {
		o.logger = logger
	}
}

// WithClientTracing 设置链路追踪
func WithClientTracing() ClientOption {
	return func(o *clientOptions) {
		o.enableTracing = true
	}
}

// DialInsecure 非安全拨号
func DialInsecure(ctx context.Context, opts ...ClientOption) (*grpc.ClientConn, error) {
	return dial(ctx, true, opts...)
}

func Dial(ctx context.Context, opts ...ClientOption) (*grpc.ClientConn, error) {
	return dial(ctx, false, opts...)
}

func dial(ctx context.Context, insecure bool, opts ...ClientOption) (*grpc.ClientConn, error) {
	// 默认配置
	options := clientOptions{
		timeout:       2000 * time.Millisecond,
		balancerName:  "round_robin",
		enableTracing: true,
	}

	for _, o := range opts {
		o(&options)
	}

	//TODO 客户端默认拦截器
	ints := []grpc.UnaryClientInterceptor{
		//应该是闭包特性，直接调用后返回resp供grpc拦截器调用
		//otelgrpc.UnaryClientInterceptor(),
	}
	// 可给用户自己设置需不需链路追踪
	if options.enableTracing {
		ints = append(ints, otelgrpc.UnaryClientInterceptor())
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

	//服务发现的选项 这里调用 resolver 的直连模式或者是服务发现模式
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
