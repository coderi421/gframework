package app

import (
	"net/url"
	"os"
	"time"

	"github.com/coderi421/gframework/gmicro/server/restserver"

	"github.com/coderi421/gframework/gmicro/server/rpcserver"

	"github.com/coderi421/gframework/gmicro/registry"
)

type Option func(o *options)
type options struct {
	// service id in registry center
	id   string
	name string
	//http://127.0.0.1:8080
	//grpc://127.0.0.1:9000
	endpoints []*url.URL
	// signal for stopping the service
	sigs []os.Signal
	// allow user to provide own implementation
	registrarTimeout time.Duration
	// the registry implementation of registry center
	registrar registry.Registrar

	// timeout for stopping the service
	stopTimeout time.Duration

	// rpc server instance
	rpcServer *rpcserver.Server
	// rest server instance
	restServer *restserver.Server
}

// WithRegistrar allows user to provide own implementation of registry center
func WithRegistrar(registrar registry.Registrar) Option {
	return func(o *options) {
		o.registrar = registrar
	}
}

// WithRPCServer allows user to provide own implementation of rpc server
func WithRPCServer(server *rpcserver.Server) Option {
	return func(o *options) {
		o.rpcServer = server
	}
}

// WithRestServer allows user to provide own implementation of rest server
func WithRestServer(server *restserver.Server) Option {
	return func(o *options) {
		o.restServer = server
	}
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithEndpoints(endpoints []*url.URL) Option {
	return func(o *options) {
		o.endpoints = endpoints
	}
}

func WithSigs(sigs []os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

func WithRegistrarTimeout(registrarTimeout time.Duration) Option {
	return func(o *options) {
		o.registrarTimeout = registrarTimeout
	}
}

func WithStopTimeout(stopTimeout time.Duration) Option {
	return func(o *options) {
		o.stopTimeout = stopTimeout
	}
}
