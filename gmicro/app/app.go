package app

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/CoderI421/gframework/pkg/log"

	"github.com/google/uuid"

	"github.com/CoderI421/gframework/gmicro/registry"
)

type App struct {
	// the app options
	opts options
	// lock for concurrent safe
	mux sync.Mutex
	// the service instance for registry
	instance *registry.ServiceInstance
}

// New create a new app instance
func New(opts ...Option) *App {
	o := options{
		// different system has different default signals
		sigs:             []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		registrarTimeout: 10 * time.Second,
		stopTimeout:      10 * time.Second,
	}
	// generate default uuid for service instance
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	// apply options
	for _, opt := range opts {
		opt(&o)
	}
	return &App{opts: o}
}

// Run start the app instance
func (a *App) Run() error {
	// create the service instance info
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}

	// lock the instance for concurrent safe, other goroutine may access it
	a.mux.Lock()
	a.instance = instance
	a.mux.Unlock()

	go func() {
		// start the rpc server
		if a.opts.rpcServer != nil {
			err := a.opts.rpcServer.Start()
			if err != nil {
				log.Errorf("start rpc server error: %s", err)
				return
			}
		}
	}()

	// register the service instance
	if a.opts.registrar != nil {
		ctx, cancelFunc := context.WithTimeout(context.Background(), a.opts.registrarTimeout)
		defer cancelFunc()

		err := a.opts.registrar.Register(ctx, instance)
		if err != nil {
			// TODO: twice error logic
			log.Errorf("registrar service error: %s", err)
			return err
		}
	}

	// listen the exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)
	// wait the exit signal
	<-c
	return nil
}

// Stop the app instance
func (a *App) Stop() error {
	if a.opts.registrar == nil || a.instance == nil {
		return nil
	}

	a.mux.Lock()
	instance := a.instance
	a.mux.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), a.opts.stopTimeout)
	defer cancel()

	// deregister the service instance in the registry
	return a.opts.registrar.Deregister(ctx, instance)
}

// buildInstance create the service instance info
func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	// build the service instance
	i := &registry.ServiceInstance{
		ID:   a.opts.id,
		Name: a.opts.name,
	}
	endpoints := make([]string, 0, len(a.opts.endpoints)+1)
	for _, ep := range a.opts.endpoints {
		if ep != nil {
			endpoints = append(endpoints, ep.String())
		}
	}
	//从rpcserver，restserver去主动获取这些信息
	if a.opts.rpcServer != nil {
		u := &url.URL{
			Scheme: "grpc",
			Host:   a.opts.rpcServer.Address(),
		}
		endpoints = append(endpoints, u.String())
	}
	// if the registry is empty, then use the default registry
	//if a.opts.registrar == nil {
	//	a.opts.registrar = a.opts.defaultRegistrar()
	//}
	return i, nil
}
