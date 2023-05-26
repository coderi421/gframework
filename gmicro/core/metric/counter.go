package metric

import "github.com/prometheus/client_golang/prometheus"

type (
	CounterVecOpts VectorOpts // CounterVecOpts 计数选项

	// CounterVec 计数向量重写了 prometheus.CounterVec
	CounterVec interface {
		Inc(labels ...string)            // Inc 增加1标签
		Add(v float64, labels ...string) // Add 添加任意值标签
	}

	// promCounterVec 封装了 prometheus.CounterVec
	promCounterVec struct {
		counter *prometheus.CounterVec
	}
)

// NewCounterVec 封装 prometheus.NewCounterVec 为了增加 labels 参数
func NewCounterVec(cfg *CounterVecOpts) CounterVec {
	if cfg == nil {
		return nil
	}

	vec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      cfg.Name,
		Help:      cfg.Help,
	}, cfg.Labels)
	prometheus.MustRegister(vec)

	return &promCounterVec{counter: vec}
}

func (p *promCounterVec) Inc(labels ...string) {
	p.counter.WithLabelValues(labels...).Inc()
}

func (p *promCounterVec) Add(v float64, labels ...string) {
	p.counter.WithLabelValues(labels...).Add(v)
}
