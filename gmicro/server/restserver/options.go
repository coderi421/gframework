package restserver

import "github.com/gin-gonic/gin"

type ServerOption func(*Server)

func WithEnableProfiling(profiling bool) ServerOption {
	return func(s *Server) {
		s.enableProfiling = profiling
	}
}

func WithServiceName(srvName string) ServerOption {
	return func(s *Server) {
		s.serviceName = srvName
	}
}

func WithMode(mode string) ServerOption {
	return func(s *Server) {
		s.mode = mode
	}
}
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithMiddlewares(middlewares []string) ServerOption {
	return func(s *Server) {
		s.middlewares = middlewares
	}
}

func WithHealthz(healthz bool) ServerOption {
	return func(s *Server) {
		s.healthz = healthz
	}
}

func WithJwt(jwt *JwtInfo) ServerOption {
	return func(s *Server) {
		s.jwt = jwt
	}
}
func WithTransName(transName string) ServerOption {
	return func(s *Server) {
		s.transName = transName
	}
}

func WithCustomMiddlewares(middlewares []gin.HandlerFunc) ServerOption {
	return func(s *Server) {
		s.customMiddlewares = middlewares
	}
}

func WithMetrics(metrics bool) ServerOption {
	return func(s *Server) {
		s.enableMetrics = metrics
	}
}
