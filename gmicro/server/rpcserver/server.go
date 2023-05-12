package rpcserver

import (
	"net"
	"net/url"
	"time"

	"github.com/CoderI421/gframework/pkg/log"

	apimd "github.com/CoderI421/gframework/api/metadata"
	srvints "github.com/CoderI421/gframework/gmicro/server/rpcserver/serverinterceptors"
	"github.com/CoderI421/gframework/pkg/host"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type ServerOption func(o *Server)

// Server grpc 服务端结构体
type Server struct {
	// 匿名字段，用于实现 grpc 注册 Server
	*grpc.Server

	address string
	//可传递拦截器
	unaryInts []grpc.UnaryServerInterceptor
	//stream
	streamInts []grpc.StreamServerInterceptor
	//客户自己实现ServerOption
	grpcOpts []grpc.ServerOption
	//监听器
	lis net.Listener
	//timeout
	timeout time.Duration

	//rpc的健康检查
	health *health.Server
	//一个grpc接口查看所有rpc服务 查看当前服务的所有rpc服务接口
	metadata *apimd.Server
	//设置监听的ip port
	endpoint *url.URL
}

func (s *Server) Address() string {
	return s.address
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		//自己生成端口号
		address: ":0",
		health:  health.NewServer(),
		//timeout: 1 * time.Second,
	}
	for _, o := range opts {
		o(srv)
	}
	//TODO 我们现在希望用户不设置拦截器的情况下，我们会自动默认加上一些必须的拦截器 , crash tracing
	unaryInts := []grpc.UnaryServerInterceptor{
		srvints.UnaryCrashInterceptor,
		//srvints.UnaryTimeoutInterceptor(srv.timeout),
	}
	//timeout可以交给用户设置，不设置就不用此拦截器
	if srv.timeout > 0 {
		unaryInts = append(unaryInts, srvints.UnaryTimeoutInterceptor(srv.timeout))
	}
	if len(srv.unaryInts) > 0 {
		unaryInts = append(unaryInts, srv.unaryInts...)
	}

	//把我们传入的拦截器转换成grpc的ServerOption
	grpcOpts := []grpc.ServerOption{grpc.ChainUnaryInterceptor(unaryInts...)}
	//把用户自己传入的grpc.ServerOption放在一起
	if len(srv.grpcOpts) > 0 {
		grpcOpts = append(grpcOpts, srv.grpcOpts...)
	}
	// 创建一个 grpc server，至此 当前文件的 Server 结构体已经实现了 grpc 的 Server 接口
	srv.Server = grpc.NewServer(grpcOpts...)

	//注册metadata的Server
	srv.metadata = apimd.NewServer(srv.Server)
	//自动解析address
	err := srv.listenAndEndpoint()
	if err != nil {
		return nil
	}
	//注册health
	grpc_health_v1.RegisterHealthServer(srv.Server, srv.health)
	//可以支持用户直接通过grpc的一个接口查看当前支持的所有的rpc服务
	apimd.RegisterMetadataServer(srv.Server, srv.metadata)
	reflection.Register(srv.Server)
	return srv
}

func WithAddress(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func WithLis(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}
func WithUnaryInterceptor(in ...grpc.UnaryServerInterceptor) ServerOption {
	return func(s *Server) {
		s.unaryInts = in
	}
}
func WithStreamInterceptor(in ...grpc.StreamServerInterceptor) ServerOption {
	return func(s *Server) {
		s.streamInts = in
	}
}
func WithOptions(opts ...grpc.ServerOption) ServerOption {
	return func(s *Server) {
		s.grpcOpts = opts
	}
}

// 完成ip和端口的提取
func (s *Server) listenAndEndpoint() error {
	if &s.lis == nil {
		lis, err := net.Listen("tcp", s.address)
		if err != nil {
			return err
		}
		s.lis = lis
	}
	addr, err := host.Extract(s.address, s.lis)
	if err != nil {
		_ = s.lis.Close()
		return err
	}
	s.endpoint = &url.URL{Scheme: "grpc", Host: addr}
	return nil
}

func (s *Server) Start() error {
	log.Infof("[grpc] server listening on: %s", s.lis.Addr().String())
	//改grpc核心变量 状态
	//只有.Resume()之后，请求才能进来
	//s.health.Shutdown()相反
	s.health.Resume()
	return s.Server.Serve(s.lis)

}
func (s *Server) Stop() error {
	//设置服务的状态为not_serving 防止接受新的请求
	s.health.Shutdown()
	s.Server.GracefulStop()
	log.Infof("[grpc] server stopped")
	return nil
}
