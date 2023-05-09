package flag

import (
	goflag "flag"

	"github.com/spf13/pflag"
)

// NoOp implements goflag.Value and plfag.Value,
// but has a noop Set implementation.
type NoOp struct{}

var (
	_ goflag.Value = NoOp{}
	_ pflag.Value  = NoOp{}
)

func (NoOp) String() string {
	return ""
}

func (NoOp) Set(val string) error {
	return nil
}

func (NoOp) Type() string {
	return "NoOp"
}
