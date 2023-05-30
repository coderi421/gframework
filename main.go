package main

import (
	"fmt"
	"time"
)

func main() {
	//log.Init(log.NewOptions())
	//
	//log.Error("testing", log.String("key", "value"))
	//log.Errorf("testing %s", "myerror")
	//// init gin context for log
	//ctx := &gin.Context{}
	//ctx.Set(log.KeyRequestID, "111")
	////ctx.Set(log.KeyUsername, "coderI")
	//log.ErrorC(ctx, "testing", log.String("key", "value"))

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		go func(i int) {
			fmt.Println(&i)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
