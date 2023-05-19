# gmicro
Microservice framework implemented based on Golang.

## Features
微服务框架：
抽取出公共的服务，快速的启动项目，当需要某些功能的时候通过配置文件或者具体的实现接口
微服务框架抽取出来的是业务无关的功能
- 微服务的技术选型:
	http - gin
	rpc - gRPC

- 服务启动的时候做什么:
	1. rpc 服务的启动所需组件
	2. 处理 web 服务的启动所需组件、recovery 拦截器
	3. 如果希望系统可以被监控 - 需要实现 pprof 接口
	4. 自启动 metrics 的接口
	5. 可以自注册 health 接口(gRPC,http)
	6. app 与 gin 以及 grpc 解耦
	7. 完成自动服务注册(可选)
	8. 优雅退出
	9. 启动过程中的日志打印
	10. go-zero 中处理的信号 – 某些信号产生后在退出前会写入当前的进行的信息(信号量应该由使用者确定)

## Framework
```shell
├─api 存放与外部交互的接口及 proto 文件
│  ├─metadata
│  └─user
│      └─v1
├─app 具体服务相关的实现
│  ├─shop 举例 shop http 服务 同级的还可以创建其他文件夹
│  │  └─admin shop 下的 admin 侧 http 服务 同级的还可以创建其他文件夹
│  │      ├─config 配置文件
│  │      └─controller 表示层
│  ├─pkg 服务共通的包
│  │  ├─code 服务的错误码
│  │  ├─options
│  │  └─translator
│  │      └─gin
│  └─user 举例 user gRPC 服务，可以同时建立 http 服务，参考 shop
│      ├─client 用于本地测试 user rpc 服务的客户端
│      └─srv
│          ├─config 服务的配置项，Log error 等子服务的相关逻辑全部注册到配置中
│          ├─controller 表示层
│          │  └─user
│          ├─data 数据访问层：数据库，缓存，消息队列等，无业务逻辑
│          │  └─v1
│          │      ├─db 数据库相关的操作
│          │      └─mock 数据库的 mock 用于测试
│          └─service 业务逻辑层
│          │   └─v1
│          ├─app.go user 服务生成逻辑
│          └─rpc.go user 模块的 rpc 服务 初始化逻辑
├─cmd 服务启动入口
│  ├─admin
│  └─user 启动示例 user 服务
├─configs 配置文件
│  ├─admin
│  └─user
├─gmicro 微服务相关包
│  ├─app 服务启动相关的结构体
│  │  └─app.go 这个 app 是 GRPC，服务名称，注册中心等的集合
│  ├─code
│  ├─core
│  │  ├─metric
│  │  └─trace
│  ├─registry
│  │  └─consul 服务注册中心相关逻辑，参考 kratos
│  └─server
│      ├─restserver http 服务的初始化配置
│      │  ├─middlewares http 服务的中间件
│      │  │  └─auth
│      │  ├─pprof http 服务的 pprof 相关逻辑
│      │  └─validation http 服务的参数校验
│      └─rpcserver rpc 服务的初始化配置
│          ├─client.go rpc 客户端的初始化配置
│          ├─server.go rpc 服务端的初始化配置
│          ├─clientinterceptors 客户端的拦截器：超时连接器
│          ├─resolver 服务发现相关的逻辑 解析器
│          │  ├─direct 直连
│          │  └─discovery 服务发现，负载均衡
│          │      ├─builder.go 服务发现的构建器
│          │      └─resolver.go 服务发现的解析器,负载均衡的逻辑在这里实现 UpdateState 核心
│          ├─selector 重写 grpc 接口，具体服务相关的实现 
│          │  ├─node gRPC 服务节点
│          │  │  ├─direct 直连节点
│          │  │  └─ewma ewma算法节点，用于实现 p2c 负载均衡策略
│          │  ├─p2c  负载均衡 [Power of Two Random Choices] 算法
│          │  ├─random 负载均衡随机算法
│          │  └─wrr 负载均衡 加权轮询 [Weighted Round Robin] 算法
│          └─serverinterceptors 服务端的拦截器：超时，crash恢复
├─pkg
│  ├─app 包括的项目的启动服务，配置文件的读取，命令行工具，以及其他选项
│  │  ├─app.go 服务启动：命令行工具，日志，错误包，配置等
│  │  ├─config.go 读取配置文件
│  │  ├─flag.go 命令行工具
│  │  ├─cmd.go 服务启动命令
│  │  ├─options.go 服务启动选项
│  │  └─help.go 帮助信息
│  ├─common 共通相关的包
│  │  ├─auth
│  │  ├─cli
│  │  │  ├─flag
│  │  │  └─globalflag
│  │  ├─core
│  │  ├─json
│  │  ├─meta
│  │  │  └─v1
│  │  ├─runtime
│  │  ├─scheme
│  │  ├─selection
│  │  ├─term
│  │  ├─time
│  │  ├─tools
│  │  ├─util
│  │  │  ├─clock 时间相关的工具
│  │  │  ├─fileutil
│  │  │  ├─homedir
│  │  │  ├─idutil
│  │  │  ├─iputil
│  │  │  ├─jsonutil
│  │  │  ├─net
│  │  │  ├─retryutil
│  │  │  ├─runtime
│  │  │  ├─sets
│  │  │  ├─signals
│  │  │  ├─sliceutil
│  │  │  ├─stringutil
│  │  │  └─wait
│  │  ├─validation
│  │  │  └─field
│  │  └─version
│  │      └─verflag
│  ├─errors 基础的 error 包 可以单独使用
│  ├─host
│  └─log 基础的 error 包 可以单独使用
└─tools
    └─codegen errorCode 代码生成工具 
```

## 命令行工具
```shell
# 查看帮助信息
go run .\cmd\***\***.go --help
```


