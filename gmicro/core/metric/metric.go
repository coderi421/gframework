package metric

// VectorOpts 指标选项
type VectorOpts struct {
	Namespace string   // 命名空间
	Subsystem string   // 子系统
	Name      string   // 指标名称
	Help      string   // 指标帮助信息
	Labels    []string // 指标标签
}
