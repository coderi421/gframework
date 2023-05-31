package code

//go:generate codeg -type=int
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase int = iota + 100101

	// ErrConnectDB - 500: Init db error.
	ErrConnectDB
	// ErrConnectGRPC - 500: Connect to grpc error.
	ErrConnectGRPC
)
