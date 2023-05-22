package trace

// Options 这个配置和 业务中 app 中的 tracing 冗余，为的是将 tracing 从业务中解耦出来
type Options struct {
	Name     string  `json:"name"`     // jeager 名称
	Endpoint string  `json:"endpoint"` // jeager 地址
	Sampler  float64 `json:"sampler"`  // 采样率
	Batcher  string  `json:"batcher"`  // 批量发送
}
