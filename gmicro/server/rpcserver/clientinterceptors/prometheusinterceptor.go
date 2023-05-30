package clientinterceptors

import (
	"context"
	"strconv"
	"time"

	"github.com/coderi421/gframework/gmicro/core/metric"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const serverNamespace = "rpc-client"

var (
	metricServerReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		VectorOpts: metric.VectorOpts{
			Namespace: serverNamespace,
			Subsystem: "requests",
			Name:      "your_service_name_duration_microseconds",
			Help:      "rpc client requests duration in microseconds",
			Labels:    []string{"method"}},
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})

	metricServerReqCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "your_service_status_code_total",
		Help:      "rpc client requests status code count",
		Labels:    []string{"method", "code"},
	})
)

// UnaryClientPrometheusInterceptor prometheus sclient interceptor.
func UnaryClientPrometheusInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		//记录了耗时
		metricServerReqDur.Observe(int64(time.Since(startTime)/time.Millisecond), method)

		//记录了状态码
		metricServerReqCodeTotal.Inc(method, strconv.Itoa(int(status.Code(err))))
		return err
	}
}
