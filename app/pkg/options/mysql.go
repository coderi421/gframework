package options

import (
	"time"

	"github.com/spf13/pflag"
)

type MySQLOptions struct {
	Host     string `json:"host,omitempty" mapstructure:"host,omitempty"`
	Port     string `json:"port,omitempty" mapstructure:"port,omitempty"`
	Username string `json:"username,omitempty" mapstructure:"username,omitempty"`
	Password string `json:"password,omitempty" mapstructure:"password,omitempty"`
	Database string `json:"database" mapstructure:"database"`
	LogLevel int    `json:"log-level,omitempty" mapstructure:"log-level,omitempty"`
	//连接池会用到
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxConnectionLifetime time.Duration `json:"max-connection-lifetime,omitempty" mapstructure:"max-connection-lifetime,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections,omitempty"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "127.0.0.1",
		Port:                  "3306",
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifetime: time.Duration(10) * time.Second,
		LogLevel:              1,
	}
}

// Validate verifies flags passed to ServerOptions.
func (so *MySQLOptions) Validate() []error {
	errs := []error{}
	return errs
}

// AddFlags adds flags related to server storage for a specific APIServer to the specified FlagSet.
func (mo *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&mo.Host, "mysql.host", mo.Host, ""+
		"MySQL service host address. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&mo.Port, "mysql.port", mo.Port, ""+
		"MySQL service port")

	fs.StringVar(&mo.Username, "mysql.username", mo.Username, ""+
		"Username for access to mysql service.")

	fs.StringVar(&mo.Password, "mysql.password", mo.Password, ""+
		"Password for access to mysql, should be used pair with password.")

	fs.StringVar(&mo.Database, "mysql.database", mo.Database, ""+
		"Database name for the server to use.")

	fs.IntVar(&mo.MaxIdleConnections, "mysql.max-idle-connections", mo.MaxOpenConnections, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&mo.MaxOpenConnections, "mysql.max-open-connections", mo.MaxOpenConnections, ""+
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&mo.MaxConnectionLifetime, "mysql.max-connection-life-time", mo.MaxConnectionLifetime, ""+
		"Maximum connection life time allowed to connecto to mysql.")

	fs.IntVar(&mo.LogLevel, "mysql.log-mode", mo.LogLevel, ""+
		"Specify gorm log level.")
}
