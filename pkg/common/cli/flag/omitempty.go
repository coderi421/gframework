package flag

// OmitEmpty is an interface for flags to report whether their underlying value
// is "empty." If a flag implements OmitEmpty and returns true for a call to Empty(),
// it is assumed that flag may be omitted from the command line.
type OmitEmpty interface {
	Empty() bool
}
