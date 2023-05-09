// Package v1 contains API types that are common to all versions.
//
// The package contains two categories of types:
//   - external (serialized) types that lack their own version (e.g TypeMeta)
//   - internal (never-serialized) types that are needed by several different
//     api groups, and so live here, to avoid duplication and/or import loops
//     (e.g. LabelSelector).
//
// In the future, we will probably move these categories of objects into
// separate packages.
package v1 // import "github.com/CoderI421/gmicro/pkg/common/meta/v1"
