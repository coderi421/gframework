package options

import (
	"github.com/spf13/pflag"
)

// ServerOptions service discovery center options console, etcd, nacos etc.
type ServerOptions struct {
	Name              string   `json:"name,omitempty" mapstructure:"name"`                 // 服务名称
	Host              string   `json:"host,omitempty" mapstructure:"host"`                 // 服务地址
	Port              int      `json:"port,omitempty" mapstructure:"port"`                 // 服务端口
	HttpPort          int      `json:"http_port,omitempty" mapstructure:"http_port"`       // http端口
	EnableProfiling   bool     `json:"profiling,omitempty" mapstructure:"profiling"`       // 是否开启性能分析
	EnableLimit       bool     `json:"limit,omitempty" mapstructure:"limit"`               // 是否开启限流
	EnableMetrics     bool     `json:"metrics,omitempty" mapstructure:"metrics"`           // 是否开启指标
	EnableHealthCheck bool     `json:"health_check,omitempty" mapstructure:"health_check"` // 是否开启健康检查
	Middlewares       []string `json:"middlewares,omitempty" mapstructure:"middlewares"`   // 中间件
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Host:            "127.0.0.1",
		Port:            8078,
		HttpPort:        8079,
		Name:            "user-srv",
		EnableProfiling: true,
		EnableLimit:     true,
		EnableMetrics:   true,
	}
}

func (so *ServerOptions) Validate() (errs []error) {
	errs = []error{}
	return
}

func (so *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&so.EnableProfiling, "server.enable-profiling", so.EnableProfiling,
		"enable-profiling, if true, will add <host>:<port>/debug/pprof/, default is true")
	fs.BoolVar(&so.EnableMetrics, "server.enable-metrics", so.EnableMetrics,
		"enable-metrics, if true, will add /metrics, default is true")
	fs.BoolVar(&so.EnableHealthCheck, "server.enable-health-check", so.EnableHealthCheck,
		"enable-health-check, if true, will add health check route, default is true")
	//fs.StringVarP 带有简写命令
	// （接收值的变量，命令名称，默认值，描述）
	fs.StringVar(&so.Host, "server.host", so.Host, "server host default is 127.0.0.1")
	fs.IntVar(&so.Port, "server.port", so.Port, "server port default is 8078")
	fs.IntVar(&so.HttpPort, "server.http-port", so.HttpPort, "server http port default is 8079")
	fs.StringVar(&so.Name, "server.name", so.Name, "server name default is user-srv")
}
