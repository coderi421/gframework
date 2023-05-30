package options

import (
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/spf13/pflag"
)

type TelemetryOptions struct {
	// Name is the name of the service.
	Name string `json:"name"`
	// 连接地址
	Endpoint string `json:"endpoint"`
	// SampleRate is the rate at which traces are sampled. 1.0 means all traces are sampled, 0.0 means no traces are sampled.
	// 意思是采样率，1.0表示所有的都采样，0.0表示不采样
	Sampler float64 `json:"sampler"`
	// Batcher is the type of batcher to use for sending traces to the collector.
	// Batcher是用于将跟踪发送到收集器的批处理程序类型。
	Batcher string `json:"batcher"`
}

func NewTelemetryOptions() *TelemetryOptions {
	return &TelemetryOptions{
		Name:     "shop-service",
		Endpoint: "http://127.0.0.1:14268/api/traces",
		Sampler:  1.0,
		Batcher:  "jaeger",
	}
}

func (t *TelemetryOptions) Validate() (errs []error) {
	if t.Batcher != "jaeger" && t.Batcher != "zipkin" {
		errs = append(errs, errors.New("open-telemetry batcher only supports: jaeger and zipkin"))
	}
	return
}

func (t *TelemetryOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&t.Name, "telemetry.name", t.Name, "open-telemetry name")
	fs.StringVar(&t.Endpoint, "telemetry.endpoint", t.Endpoint, "open-telemetry endpoint")
	fs.Float64Var(&t.Sampler, "telemetry.sampler", t.Sampler, "open-telemetry sampler")
	fs.StringVar(&t.Batcher, "telemetry.batcher", t.Batcher, "open-telemetry batcher,only support jaeger and zipkin")
}
