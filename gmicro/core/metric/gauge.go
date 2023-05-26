package metric

import "github.com/prometheus/client_golang/prometheus"

type (
	GaugeVecOpts VectorOpts // GaugeVecOpts 指标向量选项

	GaugeVec interface {
		Set(v float64, labels ...string) // Set 设置标签
		Inc(labels ...string)            // Inc 增加1标签
		Dec(labels ...string)            // Dec 减少1标签
		Add(v float64, labels ...string) // Add 添加任意值标签
		Sub(v float64, labels ...string) // Sub 减少任意标签
	}

	promGaugeVec struct {
		gauge *prometheus.GaugeVec
	}
)

// NewGaugeVec 封装 prometheus.NewGaugeVec 为了增加 labels 参数
func NewGaugeVec(cfg *GaugeVecOpts) GaugeVec {
	if cfg == nil {
		return nil
	}

	vec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      cfg.Name,
		Help:      cfg.Help,
	}, cfg.Labels)
	// 反向注册，将指标注册到指标注册器中
	prometheus.MustRegister(vec)
	//将注册到指标注册器中的指标包装成promGaugeVec
	return &promGaugeVec{gauge: vec}
}

func (p *promGaugeVec) Set(v float64, labels ...string) {
	p.gauge.WithLabelValues(labels...).Set(v)
}

func (p *promGaugeVec) Inc(labels ...string) {
	p.gauge.WithLabelValues(labels...).Inc()
}

func (p *promGaugeVec) Dec(labels ...string) {
	p.gauge.WithLabelValues(labels...).Dec()
}

func (p *promGaugeVec) Add(v float64, labels ...string) {
	p.gauge.WithLabelValues(labels...).Add(v)
}

func (p *promGaugeVec) Sub(v float64, labels ...string) {
	p.gauge.WithLabelValues(labels...).Sub(v)
}
