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
│  └─user
│      └─v1
├─app 具体服务相关的实现
│  ├─pkg 服务共通的包
│  │  └─code 服务的错误码
│  └─user 举例 user 服务
│      └─srv
│          ├─config 服务的配置项，Log error 等子服务的相关逻辑全部注册到配置中
│          ├─controller  表示层
│          │  └─user
│          ├─data 数据访问层：数据库，缓存，消息队列等，无业务逻辑
│          │  └─v1
│          │      ├─db
│          │      └─mock
│          └─service 业务逻辑层
│              └─v1
├─cmd 服务启动入口
│  └─user 启动示例 user 服务
├─configs 配置文件
├─gmicro
│  ├─app 服务启动相关的结构体
│  └─registry 服务实例注册信息相关结构体
├─pkg
│  ├─app
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
│  │  │  ├─clock  时间相关的工具
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
│  └─log 基础的 error 包 可以单独使用
└─tools 工具包
	└─codegen errorCode 代码生成工具 
```



