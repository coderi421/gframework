package srv

import (
	"fmt"

	upbv1 "github.com/CoderI421/gframework/api/user/v1"
	srv1 "github.com/CoderI421/gframework/app/user/srv/service/v1"

	"github.com/CoderI421/gframework/app/user/srv/config"
	"github.com/CoderI421/gframework/app/user/srv/controller/user"
	"github.com/CoderI421/gframework/app/user/srv/data/v1/mock"
	"github.com/CoderI421/gframework/gmicro/server/rpcserver"
)

func NewUserRPCServer(cfg *config.Config) (*rpcserver.Server, error) {
	data := mock.NewUsers()
	srv := srv1.NewUserService(data)
	userver := user.NewUserServer(srv)

	rpcAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	urpcServer := rpcserver.NewServer(rpcserver.WithAddress(rpcAddr))

	// 注册 user 模块的 rpc 服务
	upbv1.RegisterUserServer(urpcServer.Server, userver)

	return urpcServer, nil

}
