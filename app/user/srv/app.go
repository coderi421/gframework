package srv

import (
	"github.com/hashicorp/consul/api"

	"github.com/coderi421/gframework/app/pkg/options"
	"github.com/coderi421/gframework/app/user/srv/config"
	gapp "github.com/coderi421/gframework/gmicro/app"
	"github.com/coderi421/gframework/gmicro/registry"
	"github.com/coderi421/gframework/gmicro/registry/consul"
	"github.com/coderi421/gframework/pkg/app"
	"github.com/coderi421/gframework/pkg/log"
)

func NewApp(basename string) *app.App {
	c := config.New()

	options := []app.Option{
		app.WithOptions(c),
		app.WithRunFunc(run(c)),
		//app.WithNoConfig(),
	}

	return app.NewApp("user", "gmicro", options...)
}

// NewRegistrar 创建注册中心
func NewRegistrar(registry *options.RegistryOptions) registry.Registrar {
	c := api.DefaultConfig()
	c.Address = registry.Address
	c.Scheme = registry.Scheme
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	return consul.New(client, consul.WithHealthCheck(true))
}

func NewUserApp(cfg *config.Config) (*gapp.App, error) {
	//初始化log
	log.Init(cfg.Log)
	defer log.Flush()
	//服务注册
	register := NewRegistrar(cfg.Registry)
	//生成rpc服务
	rpcServer, err := NewUserRPCServer(cfg)
	if err != nil {
		return nil, err
	}
	return gapp.New(
		gapp.WithName(cfg.Server.Name),
		gapp.WithRPCServer(rpcServer),
		gapp.WithRegistrar(register),
	), nil

}

// 闭包，以便可以使用 config.Config
// controller(参数校验) -> service(具体的业务逻辑)-> data(数据库的接口)
func run(c *config.Config) app.RunFunc {
	return func(basename string) error {
		userApp, err := NewUserApp(c)
		if err != nil {
			return err
		}
		//启动
		if err := userApp.Run(); err != nil {
			log.Errorf("run user app error: %s", err)
		}
		return nil
	}
}
