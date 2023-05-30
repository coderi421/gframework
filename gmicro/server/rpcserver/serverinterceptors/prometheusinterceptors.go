package serverinterceptors

import (
	"context"
	"strconv"
	"time"

	"github.com/coderi421/gframework/gmicro/core/metric"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const serverNamespace = "rpc_server"

var (
	metricServerReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		VectorOpts: metric.VectorOpts{
			// 服务的名称
			Namespace: serverNamespace,
			Subsystem: "requests",
			Name:      "your_service_name_duration_microseconds",
			Help:      "rpc server requests duration in microseconds",
			Labels:    []string{"method"},
		},
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})

	metricServerReqCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "your_service_status_code_total",
		Help:      "rpc server requests status code count",
		Labels:    []string{"method", "code"},
	})
)

// UnaryPrometheusInterceptor prometheus server interceptor.
func UnaryServerPrometheusInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {

	//进入执行逻辑前
	startTime := time.Now()
	// 执行请求的逻辑
	resp, err = handler(ctx, req)

	metricServerReqDur.Observe(int64(time.Since(startTime)/time.Millisecond), info.FullMethod)

	//记录了状态码记录那个方法，获取了那个状态码
	metricServerReqCodeTotal.Inc(info.FullMethod, strconv.Itoa(int(status.Code(err))))

	return
}
