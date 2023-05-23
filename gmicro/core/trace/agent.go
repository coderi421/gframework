package trace

import (
	"sync"

	"github.com/CoderI421/gframework/pkg/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
)

const (
	kindJaeger = "jaeger"
	kingZipkin = "zipkin"
)

var (
	// struct{}空结构体 不占内存，zerobase
	// 这个map的作用是，记录已经初始化过的endpoint，避免重复初始化
	agents = make(map[string]struct{})
	// 保证map的并发安全
	lock sync.Mutex
)

func InitAgent(o Options) {
	lock.Lock()
	defer lock.Unlock()

	// 如果已经存在了，就不再初始化
	_, ok := agents[o.Endpoint]
	if ok {
		return
	}

	err := startAgent(o)
	if err != nil {
		return
	}
	// 无错误后，将endpoint加入到map中
	agents[o.Endpoint] = struct{}{}
}

func startAgent(o Options) error {
	var sexp trace.SpanExporter
	var err error

	opts := []trace.TracerProviderOption{
		// 采样率
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(o.Sampler))),
		// 设置service
		trace.WithResource(resource.NewSchemaless(semconv.ServiceNameKey.String(o.Name))),
	}

	if len(o.Endpoint) > 0 {
		switch o.Batcher {
		case kindJaeger:
			sexp, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(o.Endpoint)))
			if err != nil {
				return err
			}
		case kingZipkin:
			sexp, err = zipkin.New(o.Endpoint)
			if err != nil {
				return err
			}
		default:
			return nil
		}
		opts = append(opts, trace.WithBatcher(sexp))
	}

	//	创建一个新的trace provider，将初始化的参数传入
	tp := trace.NewTracerProvider(opts...)
	//	设置全局的trace provider
	otel.SetTracerProvider(tp)
	// 	设置全局的 propagator 传播提取器
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	//	设置 open-telemetry 的错误处理函数
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		log.Errorf("[otel] error: %v", err)
	}))
	return nil
}
