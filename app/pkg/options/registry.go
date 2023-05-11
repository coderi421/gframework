package options

import (
	"github.com/CoderI421/gframework/pkg/errors"
	"github.com/spf13/pflag"
)

// RegistryOptions registry center options console, etcd, nacos etc.
type RegistryOptions struct {
	Address string `json:"address" mapstructure:"address,omitempty"`
	// 指明基于什么进行注册 如：nacos consul etcd等，因为不通的注册中心 可能使用 http tcp grpc 等不同的协议
	Scheme string `json:"scheme" mapstructure:"scheme,omitempty"`
}

func NewRegistryOptions() *RegistryOptions {
	return &RegistryOptions{
		Address: "127.0.0.1:8500",
		Scheme:  "http"}
}

// Validate 校验配置是否正确
func (o *RegistryOptions) Validate() (errs []error) {
	if o.Address == "" || o.Scheme == "" {
		errs = append(errs, errors.New("address and scheme is empty"))
	}
	return
}

// AddFlags 将配置信息添加到 FlagSet 中
func (o *RegistryOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Address, "consul.address", o.Address, ""+
		"consul address, if left , default is 127.0.0.1:8500")
	fs.StringVar(&o.Scheme, "consul.scheme", o.Scheme, ""+
		"registry scheme, if left , default is http")
}
