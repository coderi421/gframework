package app

import (
	"net/url"
	"os"
	"time"

	"github.com/CoderI421/gmicro/gmicro/registry"
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
	registrar        registry.Registrar

	// timeout for stopping the service
	stopTimeout time.Duration
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
