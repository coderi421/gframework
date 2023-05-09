package registry

import "context"

// ServiceInstance is the service instance for registry center.
type ServiceInstance struct {
	// the id of service in the registry center: etcd, consul, zookeeper etc.
	ID string `json:"id"`
	// service name
	Name string `json:"name"`
	// service version
	Version string `json:"version"`
	// service metadata
	Metadata map[string]interface{} `json:"metadata"`

	//http://127.0.0.1:8080
	//grpc://127.0.0.1:9000
	EndPoints []string `json:"endpoints"`
}

// Registrar is the interface that wraps the basic Register and Deregister method.
type Registrar interface {
	// Register registers the service instance to registry center.
	Register(ctx context.Context, service *ServiceInstance) error
	// Deregister deregisters the service instance from registry center.
	Deregister(ctx context.Context, service *ServiceInstance) error
}

// Discovery is the interface that wraps the basic GetService and Watch method.
type Discovery interface {
	// GetService gets the service instance by serviceName.
	//	if the args is id, then get the only service instance by id.
	//	if the args is name, then may get the many service instance by name, load balancing is variable.
	GetService(ctx context.Context, serviceName string) (*ServiceInstance, error)
	// Watch creates the Watcher instance.
	Watch(ctx context.Context, serviceName string) (Watcher, error)
}

// Watcher is the interface that wraps the basic Next and Stop method.
type Watcher interface {
	// Next gets the service instance, it will return the service instance in the following cases:
	//	1. the first time to watch, if the service instance list is not empty, then return the service instance list.
	//	2. if the service instance changed, then return the service instance list.
	//	3. if the above two cases are not satisfied, it will block until the context deadline or cancel.
	Next() ([]*ServiceInstance, error)
	// active stop watching
	Stop() error
}
