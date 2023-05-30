package main

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/coderi421/gframework/api/user/v1"
	"github.com/coderi421/gframework/gmicro/registry/consul"
	"github.com/coderi421/gframework/gmicro/server/rpcserver"
	"github.com/coderi421/gframework/gmicro/server/rpcserver/selector"
	"github.com/coderi421/gframework/gmicro/server/rpcserver/selector/random"
	"github.com/hashicorp/consul/api"
)

func main() {
	//客户端，设置全局负载均衡策略，这里选择了 random
	// 这个逻辑中，balancerName是selector,在 gmicro/server/rpcserver/balancer.go中规定的
	selector.SetGlobalSelector(random.NewBuilder())
	// selector.SetGlobalSelector(random.NewBuilder()) 设定的全局的 selector 然后这里，调用 selector 进行注册
	rpcserver.InitBuilder()

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
		rpcserver.WithBalancerName("selector"),
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

	// 测试 grpc 服务的负载均衡用例
	// 在终端中，使用 --server.port=*** --server.http-port -c configs/user/srv.yaml 命令启动启动多个，这里就会轮询调用，测试负载均衡
	for {
		re, err := uc.GetUserList(context.Background(), &v1.PageInfo{
			Pn:    1,
			PSize: 4,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(re)
		time.Sleep(time.Second * 5)
	}
}
