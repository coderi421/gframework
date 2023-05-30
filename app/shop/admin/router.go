package admin

import (
	"github.com/coderi421/gframework/app/shop/admin/controller"
	"github.com/coderi421/gframework/gmicro/server/restserver"
)

func initRouter(g *restserver.Server) {
	v1 := g.Group("/v1")
	ugroup := v1.Group("/user")
	ucontroller := controller.NewUserController()
	ugroup.GET("list", ucontroller.List)
}
