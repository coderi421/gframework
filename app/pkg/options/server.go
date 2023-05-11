package options

import (
	"github.com/CoderI421/gframework/pkg/errors"
	"github.com/spf13/pflag"
)

// ServerOptions service discovery center options console, etcd, nacos etc.
type ServerOptions struct {
	Name string `json:"name" mapstructure:"name"`
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Host: "127.0.0.1",
		Port: 8078,
		Name: "user-srv",
	}
}

func (o *ServerOptions) Validate() (errs []error) {
	if o.Host == "" || o.Port == 0 {
		errs = append(errs, errors.New("Host and Port must be"))
	}
	return
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "server.host", o.Host, "server host default is 127.0.0.1")

	fs.IntVar(&o.Port, "server.port", o.Port, "server port default is 8078")
	fs.StringVar(&o.Name, "server.name", o.Name, "server name default is mxshop-user-srv")
}
