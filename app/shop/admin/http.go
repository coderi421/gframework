package admin

import (
	"github.com/CoderI421/gframework/app/shop/admin/config"
	"github.com/CoderI421/gframework/gmicro/server/restserver"
)

// NewUserHTTPServer 创建一个http server
func NewUserHTTPServer(conf *config.Config) (*restserver.Server, error) {
	//trace.InitAgent(trace.Options{
	//	Name:     conf.Telemetry.Name,
	//	Endpoint: conf.Telemetry.Endpoint,
	//	Sampler:  conf.Telemetry.Sampler,
	//	Batcher:  conf.Telemetry.Batcher,
	//})

	uRestServer := restserver.NewServer(
		restserver.WithPort(conf.Server.HttpPort),
		restserver.WithEnableProfiling(true),
		restserver.WithMiddlewares(conf.Server.Middlewares),
		restserver.WithEnableProfiling(true),
		restserver.WithMetrics(true),
	)
	//_ = tracerProvider()

	//	配置好路由
	initRouter(uRestServer)
	return uRestServer, nil
}

//var tp *otelsdktrace.TracerProvider
//
//// 初始化Provider
//func tracerProvider() error {
//	url := "http://127.0.0.1:14268/api/traces"
//	jexp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
//	if err != nil {
//		panic(err)
//	}
//
//	tp = otelsdktrace.NewTracerProvider(
//		otelsdktrace.WithBatcher(jexp),
//		otelsdktrace.WithResource(
//			resource.NewWithAttributes(
//				semconv.SchemaURL,
//				semconv.ServiceNameKey.String("mxshop-user"),
//				attribute.String("environment", "dev"),
//				attribute.Int("ID", 1),
//			),
//		),
//	)
//	otel.SetTracerProvider(tp)
//	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
//	return nil
//}
