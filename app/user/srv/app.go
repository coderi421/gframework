package srv

import (
	"fmt"

	"github.com/CoderI421/gmicro/app/user/srv/config"
	"github.com/CoderI421/gmicro/pkg/app"
)

func NewApp(basename string) *app.App {
	c := config.New()

	options := []app.Option{
		app.WithOptions(c),
		app.WithRunFunc(run(c)),
		app.WithNoConfig(),
	}

	return app.NewApp("user", "gmicro", options...)
}

// 闭包，以便可以使用 config.Config
func run(c *config.Config) app.RunFunc {
	return func(basename string) error {
		fmt.Println(c.Log.Level)
		return nil
	}
}
