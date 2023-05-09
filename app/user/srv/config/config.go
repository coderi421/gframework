package config

import (
	"github.com/CoderI421/gmicro/pkg/app"
	cliflag "github.com/CoderI421/gmicro/pkg/common/cli/flag"
	"github.com/CoderI421/gmicro/pkg/log"
)

var _ app.CliOptions = &Config{}

func New() *Config {
	return &Config{
		Log: log.NewOptions(),
	}
}

type Config struct {
	Log *log.Options `json:"log" mapstructure:"log"`
}

// Flags implements app.CliOptions interface.Add flags to the specified FlagSet object.
func (c *Config) Flags() (fss cliflag.NamedFlagSets) {
	// fss.FlagSet("logs") -> 创建一个FlagSet对象，命名为logs，做为专属的 logs 传递给 Config.Log
	c.Log.AddFlags(fss.FlagSet("logs"))
	return fss
}

// Validate 将配置中的所有校验子逻辑 注册到当前实例的校验中
func (c *Config) Validate() (errors []error) {
	// 将 Log 中的校验，注册到 user 服务的，校验逻辑中
	errors = append(errors, c.Log.Validate()...)
	return
}
