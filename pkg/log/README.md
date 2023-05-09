# How to use it
## init
```golang
	log.Init(log.NewOptions())
	
	log.Error("testing", log.String("key", "value"))
	log.Errorf("testing %s", "myerror")
	// init gin context for log
	ctx := &gin.Context{}
	ctx.Set(log.KeyRequestID, "111")
	ctx.Set(log.KeyUsername, "coderI")
	log.ErrorC(ctx, "testing", log.String("key", "value"))
```

