package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"

	v1 "github.com/CoderI421/gframework/api/user/v1"
	"github.com/CoderI421/gframework/gmicro/registry/consul"
	"github.com/CoderI421/gframework/gmicro/server/rpcserver"
)

func main() {
	conf := api.DefaultConfig()
	conf.Address = "127.0.0.1:8500"
	conf.Scheme = "http"

	client, err := api.NewClient(conf)
	if err != nil {
		panic(err)
	}

	// 服务注册
	r := consul.New(client, consul.WithHealthCheck(true))

	conn, err := rpcserver.DialInsecure(
		context.Background(), rpcserver.WithDiscovery(r),
		/*
			第3个/是为了第二个参数是空的
			默认格式：direct://<authority>/127.0.0.1:8078
			以后使用nacos或者其他的中心 也不用改discovery 只修改conf就可以
			服务发现可以直接去kartors里面copy registry下的etcd nacos等使用
		*/
		rpcserver.WithEndpoint("discovery:///user-srv"),
		rpcserver.WithClientTimeout(50*time.Second),
	)
	//conn, err := grpc.Dial("127.0.0.1:8078", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	uc := v1.NewUserClient(conn)
	re, err := uc.GetUserList(context.Background(), &v1.PageInfo{})
	if err != nil {
		panic(err)
	}
	fmt.Println(re)
}
