package admin

import (
	"github.com/CoderI421/gframework/app/shop/admin/config"
	"github.com/CoderI421/gframework/gmicro/server/restserver"
)

// NewUserHTTPServer 创建一个http server
func NewUserHTTPServer(conf *config.Config) (*restserver.Server, error) {
	uRestServer := restserver.NewServer(
		restserver.WithPort(conf.Server.HttpPort),
		restserver.WithEnableProfiling(true),
		restserver.WithMiddlewares(conf.Server.Middlewares),
	)

	//	配置好路由
	initRouter(uRestServer)
	return uRestServer, nil
}
