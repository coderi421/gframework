package code

//go:generate codeg -type=int
const (
	// ErrConnectDB - 500: Init db error.
	ErrConnectDB int = iota + 100101

	// ErrConnectGRPC - 500: Connect to grpc error.
	ErrConnectGRPC
)
