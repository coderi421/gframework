package restserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mws "github.com/CoderI421/gframework/gmicro/server/restserver/middlewares"
	"github.com/CoderI421/gframework/gmicro/server/restserver/pprof"
	"github.com/CoderI421/gframework/gmicro/server/restserver/validation"
	"github.com/CoderI421/gframework/pkg/errors"
	"github.com/CoderI421/gframework/pkg/log"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

type JwtInfo struct {
	//defaults to "JWT"
	Realm string
	//defaults to empty
	Key string
	//defaults to 7 days
	Timeout time.Duration
	//defaults to 7 days 刷新时长
	MaxRefresh time.Duration
}

// wrapper for gin.Engine
type Server struct {
	*gin.Engine
	//端口号 默认值 8080
	port int
	//开发模式
	mode string
	//是否开启健康检查接口，默认开启，如果开启会自动添加/health接口
	healthz bool
	//是否开启pprof接口，默认开启,如果开启会自动添加/debug/pprof接口
	enableProfiling bool
	//是否开启metrics接口，默认开启，如果开启会自动添加/metrics接口
	enableMetrics bool
	//中间件(拦截器)两种用法 1.提前写好，直接配置名称就可以，用起来方便，比rpc自定义的实现弱，2.自定义gin.HandlerFunc
	middlewares       []string
	customMiddlewares []gin.HandlerFunc
	//jwt配置信息
	jwt *JwtInfo
	//翻译器 默认：zh
	transName string
	trans     ut.Translator
	// 将 http.Server 作为作为 Server 的一个属性，为了重写 gin 的 Run 方法 实现后续的优雅推出
	server *http.Server

	//tracing 的服务名，默认为 restserver
	serviceName string
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		port:    8080,
		mode:    "debug",
		healthz: true,
		jwt: &JwtInfo{
			"JWT",
			"Gd%YCfP1agNHo5x6xm2Qs33Bf!B#Gi!o",
			7 * 24 * time.Hour,
			7 * 24 * time.Hour,
		},
		Engine:    gin.Default(),
		transName: "zh",
	}
	for _, o := range opts {
		o(srv)
	}
	// 循环遍历中间件，如果是自定义的中间件，就添加到gin.Engine中
	for _, m := range srv.middlewares {
		mw, ok := mws.Middlewares[m]
		if !ok {
			log.Warnf("can not find middleware:%s", m)
			continue
			//panic(errors.Errorf("can not find middleware:%s", m))
		}
		log.Infof("install middleware:%s", m)
		srv.Use(mw)
	}
	return srv
}

// start rest server
func (s *Server) Start(ctx context.Context) error {
	/*
		debug模式和release模式区别主要是打印的日志不同
		环境变量的模式，在docker k8s部署中很常用
		gin.SetMode(gin.ReleaseMode)
	*/
	if s.mode != gin.DebugMode && s.mode != gin.ReleaseMode && s.mode != gin.TestMode {
		return errors.New("mode must be one of debug/release/test")
	}

	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s(%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	//TODO 初始化翻译器
	err := s.initTrans(s.transName)
	if err != nil {
		log.Errorf("initTrans error: %s", err.Error())
		return err
	}

	//注册mobile验证器
	validation.RegisterMobile(s.trans)
	//根据配置初始化pprof路由
	if s.enableProfiling {
		pprof.Register(s.Engine)
	}
	//根据配置初始化metrics路由
	if s.enableMetrics {
		// get global Monitor object
		m := ginmetrics.GetMonitor()
		// +optional set metric path, default /debug/metrics
		m.SetMetricPath("/metrics")
		// +optional set slow time, default 5s
		m.SetSlowTime(10)
		// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
		// used to p95, p99
		m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
		//反向注入
		m.Use(s)
	}

	// 如果开启了健康检查接口，就添加/health接口
	if s.healthz {
		s.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "Ok",
			})
		})
	}

	log.Infof("rest server is running on port: %d", s.port)
	address := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    address,
		Handler: s.Engine,
	}
	_ = s.SetTrustedProxies(nil)
	if err = s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	//log.Infof("rest server is stopping on port: %d", s.port)
	if err := s.server.Shutdown(ctx); err != nil {
		log.Errorf("rest server is shutting down: %v", err)
		return err
	}
	log.Infof("rest server stopped on port: %d", s.port)
	return nil
}
