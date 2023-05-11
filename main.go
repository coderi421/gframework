package main

import (
	"github.com/gin-gonic/gin"

	"github.com/CoderI421/gframework/pkg/log"
)

func main() {
	log.Init(log.NewOptions())

	log.Error("testing", log.String("key", "value"))
	log.Errorf("testing %s", "myerror")
	// init gin context for log
	ctx := &gin.Context{}
	ctx.Set(log.KeyRequestID, "111")
	//ctx.Set(log.KeyUsername, "coderI")
	log.ErrorC(ctx, "testing", log.String("key", "value"))
}
