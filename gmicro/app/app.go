package app

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/CoderI421/gframework/gmicro/server"

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

	// the context for stopping http app
	cancel func()
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

	//现在启动了两个server，一个是http，一个是rpc
	/*
		这两个server是否必须启动成功？
		如果有一个启动失败，那么我们就要停止另外一个server
		启动了多个，如果其中一个启动失败，其他应该被取消
		如果剩余的server的状态：
			1.没有开始调用start
				不进行就行或者调用stop
			2.start进行中
				调用进行中的cancel
			3.start已经完成
				调用stop
	*/

	var servers []server.Server
	if a.opts.rpcServer != nil {
		servers = append(servers, a.opts.rpcServer)
	}
	if a.opts.restServer != nil {
		servers = append(servers, a.opts.restServer)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	a.cancel = cancelFunc
	errGroup, ctx := errgroup.WithContext(ctx)
	wg := sync.WaitGroup{}

	for _, srv := range servers {

		// 启动一个goroutine 监听是否有服务在启动过程中存在失败的情况
		errGroup.Go(func() error {
			<-ctx.Done()
			timeoutCtx, srcCancelFunc := context.WithTimeout(context.Background(), a.opts.stopTimeout)
			defer srcCancelFunc()
			return srv.Stop(timeoutCtx)
		})

		wg.Add(1)
		// 启动一个goroutine 执行服务的启动
		errGroup.Go(func(wg *sync.WaitGroup) func() error {
			return func() error {
				wg.Done()
				return srv.Start(ctx)
			}
		}(&wg))
	}
	//go func() {
	//	// start the rpc server goroutine
	//	if a.opts.rpcServer != nil {
	//		err := a.opts.rpcServer.Start(context.Background())
	//		if err != nil {
	//			log.Errorf("start rpc server error: %s", err)
	//			return
	//		}
	//	}
	//}()
	// 这里需要等服务都启动完毕后，才能进行下面的注册服务
	wg.Wait()

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
	//<-c
	errGroup.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return a.Stop()
		}
	})

	return errGroup.Wait()
}

// Stop the app instance
func (a *App) Stop() error {
	log.Info("start deregister service")
	//自己生成的context生成cancel后往服务中传递，所以能通知到所有的服务下的context
	if a.cancel != nil {
		log.Infof("start cancel context")
		a.cancel()
	}

	// 注销注册中心的逻辑
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
	// 这里未能处理 restServer 的 address 到注册中心，如果有需要可以拓展
	// if the registry is empty, then use the default registry
	//if a.opts.registrar == nil {
	//	a.opts.registrar = a.opts.defaultRegistrar()
	//}
	return &registry.ServiceInstance{
		ID:        a.opts.id,
		Name:      a.opts.name,
		Endpoints: endpoints,
	}, nil
}
