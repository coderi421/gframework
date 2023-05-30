package user

import (
	v1 "github.com/coderi421/gframework/api/user/v1"
	srv1 "github.com/coderi421/gframework/app/user/srv/service/v1"
)

type userServer struct {
	v1.UnimplementedUserServer
	srv srv1.UserSrv
}

func NewUserServer(srv srv1.UserSrv) *userServer {
	return &userServer{srv: srv}
}
