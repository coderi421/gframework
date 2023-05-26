package metric

import "github.com/prometheus/client_golang/prometheus"

type (
	// HistogramVecOpts 柱状图选项
	HistogramVecOpts struct {
		VectorOpts
		Buckets []float64 //柱状图指标新增桶的概念
	}

	// A HistogramVec interface represents a histogram vector.
	HistogramVec interface {
		// Observe adds observation v to labels.
		Observe(v int64, labels ...string)
	}

	promHistogramVec struct {
		histogram *prometheus.HistogramVec
	}
)

// NewHistogramVec returns a HistogramVec.
func NewHistogramVec(cfg *HistogramVecOpts) HistogramVec {
	if cfg == nil {
		return nil
	}

	vec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      cfg.Name,
		Help:      cfg.Help,
		Buckets:   cfg.Buckets,
	}, cfg.Labels)
	prometheus.MustRegister(vec)

	return &promHistogramVec{histogram: vec}
}

func (p promHistogramVec) Observe(v int64, labels ...string) {
	p.histogram.WithLabelValues(labels...).Observe(float64(v))
}
